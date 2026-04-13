# Merge Plan

## Overview

A few years ago, our MyBB forum was copied to create a second, independent
forum. Since then both forums have been running side by side — same members,
same history up to the split, but with new posts, new users, and small
changes accumulating independently on each side.

This plan describes how we bring everything back together into one forum
without losing any content from either side.

### What "merging" means in practice

Think of it like reuniting two scrapbooks that started as photocopies of the
same original. Many pages are identical. Some pages were only added to one
copy. A few pages exist in both but have small differences — someone added a
note in one copy but not the other.

Our job is to:

1. **Identify the duplicates.** For every user, thread, and post we ask:
   "Does this exist on both forums?" We do this by comparing key details —
   usernames, dates, and content — not just database row numbers, which
   diverged after the split.

2. **Keep everything unique.** Any thread or post that only exists on one
   forum gets carried over to the merged result automatically.

3. **Resolve differences.** When the same item differs between the two
   forums (for example, a user updated their profile on one side but not the
   other), we follow simple rules: the original forum's version wins for
   account details, and we preserve the most complete version of content.

4. **Reassemble the forum.** Once we know what goes in, we rebuild the
   forum database with clean, consistent numbering and verify that every
   thread, post, and attachment is intact.

5. **Load it back.** The merged data is pushed into a fresh copy of the
   original forum's database. The forum boots up with the combined content
   of both sides.

### What this does NOT do

- It does not merge private messages. Those stay in backups and can be
  retrieved on request.
- It does not carry over plugin data. Plugins can be reinstalled afterward.
- It does not change how the forum looks or works — it only combines the
  underlying data.

### Timeline and safety

The merge is done offline against snapshots, so neither forum is affected
until we're ready. We take a development snapshot to build and test against,
then a final snapshot just before cutover. The original forum goes read-only
only for the brief cutover window.

Every step produces reports we can review before moving on. Nothing is
pushed to the live database until we're confident the result is correct.

---

## Technical Detail

Everything below is implementation detail for the engineering team.

This document tracks the plan for merging two diverged MyBB forum forks back
into a single coherent dataset.

## Terminology

| Term       | Prefix     | Description                                                      |
|------------|------------|------------------------------------------------------------------|
| **Origin** | `mybb_`    | The original forum. Still live; will be the final target.        |
| **Fork**   | `PBMnet_`  | The derivative fork. Now read-only.                              |
| **Merged** | `mybb_`    | The final merged dataset, exported back into Origin's MySQL.     |

The two forks diverged around **2022-12-19**.

## Key Decisions

- **Auth authority**: Origin. Passwords, salts, and login keys come from Origin.
- **Fork A (Fork) is read-only**: no resync needed; import once.
- **Origin will have two sync points**: one snapshot for development, one for
  final cutover. After cutover, Origin goes read-only briefly while we push
  the merged dataset back.
- **Late Origin updates win**: any content updated on Origin between the dev
  snapshot and the final snapshot automatically takes precedence over Fork data.
- **Separate SQLite databases**: `fork.sqlite3` and `origin.sqlite3`, not one
  DB with two prefixes.
- **Merge workspace**: `merge.sqlite3` holds metadata, identity maps, match
  candidates, conflicts, and final ID assignments.
- **Merged output**: `merged.sqlite3` contains the final MyBB-shaped tables
  ready for MySQL export.

## Architecture

```
MySQL Origin (mybb_)          MySQL Fork (PBMnet_)
       │                              │
       ▼                              ▼
  origin.sqlite3                 fork.sqlite3
       │                              │
       └──────────┬───────────────────┘
                  ▼
            merge.sqlite3
         (identity map, matches,
          conflicts, overrides)
                  │
                  ▼
           merged.sqlite3
         (canonical UUID tables →
          final mybb_ tables with
          fresh integer IDs)
                  │
                  ▼
         Target MySQL (mybb_)
```

## Tables by Treatment

### Content tables (full merge logic)

These are the tables that carry user-generated content and need real
deduplication, matching, and conflict resolution.

| Kind         | Table(s)                          | Primary key | Match strategy                                      |
|--------------|-----------------------------------|-------------|-----------------------------------------------------|
| user         | `_users`, `_userfields`           | uid         | `lower(username) + regdate`; fallback `lower(email) + regdate` |
| forum        | `_forums`                         | fid         | `parent_uuid + normalized(name) + type`             |
| thread       | `_threads`                        | tid         | `forum_uuid + starter_uuid + dateline + normalized(subject) + root_post_hash` |
| post         | `_posts`                          | pid         | `thread_uuid + author_uuid + dateline + message_hash` |
| attachment   | `_attachments`                    | aid         | `post_uuid + sha256(file) + filename + filesize`    |
| poll         | `_polls`, `_pollvotes`            | pid (polls) | `thread_uuid + question_hash + options_hash`        |
| announcement | `_announcements`                  | aid         | `forum_uuid + author_uuid + startdate + subject`    |
| reputation   | `_reputation`                     | rid         | `target_uuid + from_uuid + dateline + comment_hash` |
| warning      | `_warnings`, `_warningtypes`      | wid         | `target_uuid + mod_uuid + dateline + title_hash`    |
| ~~pm~~       | ~~`_privatemessages`~~            | ~~pmid~~    | Excluded from merge; retrieve from backups on request. |

### Relationship / state tables (FK remap + union)

These don't carry independent content. After parent entities are matched,
remap foreign keys and deduplicate by the resulting tuples.

- `_forumpermissions` — `(forum_uuid, group_uuid)`
- `_moderators` — `(forum_uuid, user_uuid)`
- `_forumsread` — `(forum_uuid, user_uuid)`
- `_forumsubscriptions` — `(forum_uuid, user_uuid)`
- `_threadsread` — `(thread_uuid, user_uuid)`
- `_threadsubscriptions` — `(thread_uuid, user_uuid)`
- `_threadratings` — `(thread_uuid, user_uuid)`
- `_buddyrequests` — `(user_uuid, target_user_uuid)`
- `_banned` — `(user_uuid)`
- `_joinrequests` — `(user_uuid, group)`
- `_groupleaders` — `(group, user_uuid)`
- `_reportedcontent` — `(post_uuid or thread_uuid, reporter_uuid)`

### Reference / config tables (baseline from Origin)

Use Origin's version as the baseline. Map Fork's groups/types to Origin's
where needed.

- `_usergroups` — map by `normalized(title) + permission_hash`
- `_usertitles`
- `_attachtypes`
- `_profilefields`
- `_threadprefixes`
- `_mycode`
- `_smilies`
- `_icons`
- `_badwords`
- `_calendars`, `_calendarpermissions`
- `_helpdocs`, `_helpsections`
- `_promotions`, `_promotionlogs`
- `_reportreasons`
- `_warninglevels`
- `_questions`, `_questionsessions`
- `_spiders`
- `_bam`

### Excluded tables

- `_privatemessages` — excluded from merge; retrieve from backups on request
- `_bam` — plugin table; Charles can reinstall plugins post-cutover

### Operational tables (ignore / regenerate)

Do not merge. MyBB regenerates these at runtime.

- `_sessions`, `_adminsessions`
- `_captcha`
- `_searchlog`
- `_mailqueue`, `_maillogs`, `_mailerrors`
- `_adminlog`, `_moderatorlog`, `_tasklog`
- `_spamlog`
- `_stats`
- `_threadviews`
- `_datacache`
- `_upgrade_data`
- `_delayedmoderation`

### Deployment tables (from target instance)

These belong to the running MyBB instance, not the content archive.
Leave them for the target to manage.

- `_settings`, `_settinggroups`
- `_templates`, `_templategroups`, `_templatesets`
- `_themes`, `_themestylesheets`
- `_tasks`
- `_adminoptions`, `_adminviews`
- `_banfilters`
- `_awaitingactivation`
- `_events`

---

## Merge Workspace Schema

These tables live in `merge.sqlite3`.

```sql
-- Registry of imported sources.
CREATE TABLE gobbs_sources (
  source       TEXT PRIMARY KEY,            -- 'origin' or 'fork'
  sqlite_path  TEXT NOT NULL,               -- relative path to source SQLite DB
  table_prefix TEXT NOT NULL,               -- 'mybb_' or 'PBMnet_'
  attachment_root TEXT NOT NULL,            -- 'uploads/origin/' or 'uploads/fork/'
  imported_at  TEXT NOT NULL                -- ISO 8601 timestamp
);

-- Inventory of every source row we care about.
CREATE TABLE gobbs_source_records (
  source     TEXT NOT NULL,
  kind       TEXT NOT NULL,                 -- logical kind: 'user', 'forum', etc.
  legacy_id  TEXT NOT NULL,                 -- original PK (TEXT for composite keys)
  row_hash   TEXT,                          -- SHA-256 of canonical field values
  match_key  TEXT,                          -- conservative cross-fork matching key
  row_json   TEXT,                          -- full row as JSON for conflict inspection
  PRIMARY KEY (source, kind, legacy_id)
);

-- The core identity map: (source, kind, id) → uuid.
CREATE TABLE gobbs_identity_map (
  source     TEXT NOT NULL,
  kind       TEXT NOT NULL,
  legacy_id  TEXT NOT NULL,
  uuid       TEXT,                          -- assigned UUID v4 (NULL until matched)
  confidence INTEGER NOT NULL DEFAULT 0,    -- 0=unmatched, 100=exact hash, 50=logical key, etc.
  basis      TEXT NOT NULL DEFAULT '',       -- human-readable reason for match
  status     TEXT NOT NULL DEFAULT 'pending', -- pending/auto/manual/ignored
  PRIMARY KEY (source, kind, legacy_id)
);
CREATE INDEX idx_identity_uuid ON gobbs_identity_map (kind, uuid);

-- Candidate matches awaiting confirmation.
CREATE TABLE gobbs_match_candidates (
  kind          TEXT NOT NULL,
  origin_id     TEXT NOT NULL,              -- legacy_id in origin
  fork_id       TEXT NOT NULL,              -- legacy_id in fork
  score         REAL NOT NULL,              -- 0.0–1.0
  basis         TEXT NOT NULL,              -- 'exact_hash', 'username+regdate', etc.
  status        TEXT NOT NULL DEFAULT 'pending', -- pending/accepted/rejected
  PRIMARY KEY (kind, origin_id, fork_id)
);

-- Field-level conflicts for matched entities.
CREATE TABLE gobbs_conflicts (
  kind           TEXT NOT NULL,
  uuid           TEXT NOT NULL,
  field_name     TEXT NOT NULL,
  origin_value   TEXT,
  fork_value     TEXT,
  severity       TEXT NOT NULL,             -- 'info', 'warn', 'blocker'
  resolution     TEXT NOT NULL DEFAULT 'unresolved',  -- unresolved/origin/fork/manual/recompute
  resolved_value TEXT,
  PRIMARY KEY (kind, uuid, field_name)
);

-- Final integer ID assignments for the merged dataset.
CREATE TABLE gobbs_final_id_map (
  kind       TEXT NOT NULL,
  uuid       TEXT NOT NULL,
  final_id   INTEGER NOT NULL,
  PRIMARY KEY (kind, uuid),
  UNIQUE (kind, final_id)
);
```

## Conflict Resolution Rules

| Field class                | Resolution strategy                        |
|----------------------------|--------------------------------------------|
| Immutable (regdate, etc.)  | Mismatch = blocker, manual review          |
| Auth (password, salt, etc.)| Always Origin                              |
| Profile (email, sig, etc.) | Origin wins (live data is more current)    |
| Body content (post message)| Origin wins if edited after fork date; else Fork wins if edited after fork date; else identical = no conflict |
| Visibility / moderation    | Preserve content (prefer visible) for archive; flag when one side deleted |
| Counters (postnum, views)  | Recompute from merged data, never trust source values |
| Sets (additionalgroups)    | Union after FK remap                       |

Special rule for final cutover resync: if Origin's row changed between dev
snapshot and final snapshot, Origin's version wins unconditionally.

---

## Tasks

Tasks are ordered by dependency. Each task should be completable in a single
working session. Tasks marked with `[gate]` block all subsequent tasks in
their phase.

### Phase 1: Import Origin

- [ ] **1.1** Create `origin.setup.json` with Origin's MySQL connection details
      and `mybb_` prefix.
- [ ] **1.2** Obtain and store Origin's MySQL schema dump.
      Import `backups/mybb_origin_20260413.sql` or dump from live instance.
- [ ] **1.3** Generate `origin-schema.sql` (SQLite) using `cmd/mysql2sqlite`
      with Origin's schema as input.
- [ ] **1.4** Diff the Origin and Fork SQLite schemas. Document any tables or
      columns present in one but not the other. This is essential context for
      all matching logic. `[gate]`
- [ ] **1.5** Import Origin data into `origin.sqlite3` using
      `cmd/mysql2sqlite-data --setup-file origin.setup.json --sqlite-path origin.sqlite3 --all-tables`.
- [ ] **1.6** Verify row counts with `cmd/mysql2sqlite-verify`.
- [ ] **1.7** Rename existing `mybb.sqlite3` to `fork.sqlite3` for clarity.
      Update `setup.json` / Makefile / tooling references accordingly.
- [ ] **1.8** Inventory Origin's attachment files. Copy or symlink into
      `uploads/origin/`. Verify file existence against `mybb_attachments` rows.
- [ ] **1.9** Inventory Fork's attachment files under `uploads/fork/` (or
      current `uploads/`). Verify file existence against `PBMnet_attachments`.

### Phase 2: Merge Infrastructure

- [ ] **2.1** Create `internal/merge` package with merge DB schema creation,
      source registration, and ATTACH helpers.
- [ ] **2.2** Create `cmd/gobbs-merge` as the single entry point for merge
      operations (subcommands: `init`, `inventory`, `match`, `report`,
      `resolve`, `build`, `export`).
- [ ] **2.3** Implement `gobbs-merge init`: create `merge.sqlite3`, register
      both sources, detect prefixes, create all metadata tables. `[gate]`
- [ ] **2.4** Implement `gobbs-merge inventory`: iterate both source DBs,
      populate `gobbs_source_records` for all content-table rows. Compute
      `row_hash` and `match_key` per kind. Store `row_json`. `[gate]`
- [ ] **2.5** Write table-classification config: which tables are content,
      relationship, reference, operational, or deployment. This drives which
      rows get inventoried and which are skipped.
- [ ] **2.6** Add schema-diff reporting to `gobbs-merge inventory` output.
      Flag missing tables, extra columns, type mismatches.

### Phase 3: Entity Matching

Matching proceeds top-down. Parent entities must be matched before children
so that child match keys can reference parent UUIDs.

- [ ] **3.1** Implement user matching. `[gate]`
  - Exact hash match first (all non-auth, non-counter fields).
  - Logical key: `lower(username) + regdate`.
  - Fallback: `lower(email) + regdate`.
  - Assign UUID to matched pairs; assign unique UUIDs to unmatched rows.
  - Report: matched count, origin-only, fork-only, ambiguous.
- [ ] **3.2** Implement usergroup matching.
  - Logical key: `lower(title)`.
  - Build Origin→Fork group ID mapping for FK remapping in later phases.
- [ ] **3.3** Implement forum matching. `[gate]`
  - Must run after user matching (forum metadata references UIDs).
  - Logical key: `parent_uuid + lower(name) + type`.
  - Build forum tree from both sources, match level by level.
- [ ] **3.4** Implement thread matching. `[gate]`
  - Logical key: `forum_uuid + starter_uuid + dateline + normalized(subject)`.
  - Use root post content hash as tiebreaker for near-matches.
- [ ] **3.5** Implement post matching. `[gate]`
  - Within matched threads: `author_uuid + dateline + message_hash`.
  - Guest posts: match on `username_text + dateline + message_hash`.
  - Flag posts present on only one fork (new replies after divergence).
- [ ] **3.6** Implement attachment matching.
  - `post_uuid + sha256(file) + original_filename + filesize`.
  - If file is missing on one side, match on metadata only and flag.
- [ ] **3.7** Implement poll matching.
  - `thread_uuid + question_hash + options_hash`.
  - Poll votes: FK remap after poll + user matching.
- [ ] ~~**3.8** PM matching~~ — excluded from merge.
- [ ] **3.9** Implement announcement matching.
  - `forum_uuid + author_uuid + startdate + subject_hash`.
- [ ] **3.10** Implement reputation + warning matching.
- [ ] **3.11** Match coverage report: summary statistics per kind, list of
      all unmatched and ambiguous entities.

### Phase 4: Conflict Detection and Resolution

- [ ] **4.1** Implement field-level conflict detection for matched entities.
      For every matched UUID, compare field values between Origin and Fork.
      Apply severity rules from the conflict resolution table above.
      Populate `gobbs_conflicts`. `[gate]`
- [ ] **4.2** Implement `gobbs-merge report`: generate conflict reports in
      Markdown. Group by kind, severity, and resolution status.
- [ ] **4.3** Implement `gobbs-merge resolve`: apply a YAML/JSON overrides
      file to resolve conflicts manually. Overrides file is version-controlled.
- [ ] **4.4** Implement auto-resolution for non-blocker conflicts per the
      resolution rules (Origin wins for auth/profile, recompute for counters,
      etc.).
- [ ] **4.5** Review and resolve all blocker conflicts manually. `[gate]`
- [ ] **4.6** Implement conflict re-check: re-run detection after resolutions
      to confirm zero unresolved blockers.

### Phase 5: Build Merged Dataset

- [ ] **5.1** Implement final ID assignment in `gobbs_final_id_map`. `[gate]`
  - Users: ordered by `regdate, uuid`.
  - Forums: ordered by tree position / display order.
  - Threads: ordered by `dateline, uuid`.
  - Posts: ordered by `thread_final_id, dateline, uuid`.
  - Attachments: ordered by `dateuploaded, uuid`.
- [ ] **5.2** Build the FK rewrite map. For every `(kind, legacy_id)` in
      both sources, know the `final_id`. This is the lookup table for
      rewriting every foreign key column. `[gate]`
- [ ] **5.3** Implement `gobbs-merge build`: create `merged.sqlite3` with
      Origin's schema (using `mybb_` prefix). `[gate]`
- [ ] **5.4** Write merged user rows. Apply conflict resolutions. Rewrite
      `usergroup`, `additionalgroups`, `displaygroup` using group ID map.
- [ ] **5.5** Write merged forum rows. Rewrite `pid` (parent forum ID),
      `parentlist`. Rebuild tree-derived fields.
- [ ] **5.6** Write merged thread rows. Rewrite `fid`, `uid`, `firstpost`,
      `lastpost`, `lastposteruid`. Leave counters for recomputation.
- [ ] **5.7** Write merged post rows. Rewrite `tid`, `fid`, `uid`,
      `replyto`. Apply content conflict resolutions.
- [ ] **5.8** Write merged attachment rows. Rewrite `pid`, `uid`. Copy
      attachment files into `uploads/merged/` with canonical names.
- [ ] **5.9** Write merged poll rows + poll votes. Rewrite `tid`, `uid`.
- [ ] ~~**5.10** PMs~~ — excluded from merge.
- [ ] **5.11** Write merged announcement rows. Rewrite `fid`, `uid`.
- [ ] **5.12** Write merged relationship/state tables. Remap all FKs,
       deduplicate by resulting tuples.
- [ ] **5.13** Copy reference/config tables from Origin baseline.
       Apply any group/type remappings needed.
- [ ] **5.14** Recompute all denormalized counters and derived fields:
  - `users.postnum`, `users.threadnum`
  - `threads.replies`, `threads.attachmentcount`, `threads.lastpost`, etc.
  - `forums.threads`, `forums.posts`, `forums.lastpost`, etc.
- [ ] **5.15** Integrity checks on `merged.sqlite3`: `[gate]`
  - Every post's `tid`/`fid`/`uid` exists.
  - Every thread's `firstpost`/`lastpost` exists.
  - Every attachment's `pid`/`uid` exists.
  - No duplicate final IDs per kind.
  - No orphan FK references.
  - Merged counts are sensible (≥ max(Origin, Fork), ≤ Origin + Fork).

### Phase 6: Review

- [ ] **6.1** Update `cmd/gobbs-serve` to accept `--sqlite-path` flag so it
      can preview `merged.sqlite3`.
- [ ] **6.2** Add merge provenance pages to preview server:
  - `/__merge/stats` — per-kind merge statistics.
  - `/__merge/conflicts` — remaining warnings, resolved conflicts.
  - `/__merge/provenance/:kind/:id` — show source lineage for any entity.
- [ ] **6.3** Spot-check review:
  - 10 most active users (profile, post count, authorship).
  - 10 highest-reply threads (correct post ordering, no duplicates).
  - 5 threads known to have diverged (replies on both forks).
  - All forums (tree structure, names, descriptions).
  - Random sample of 20 posts with attachments.
  - All posts by guest users.
- [ ] **6.4** Attachment audit: verify every merged attachment row has a
      corresponding file in `uploads/merged/`. Flag missing files.
- [ ] **6.5** Sign-off checkpoint. All blockers resolved, spot-checks pass,
      attachment audit clean. `[gate]`

### Phase 7: Export to MySQL

- [ ] **7.1** Create `cmd/gobbs-export-mysql` (or `cmd/sqlite2mysql-data`).
      Read from `merged.sqlite3`, write to target MySQL in dependency order.
- [ ] **7.2** Implement batched inserts with configurable batch size.
      Disable/restore indexes around bulk load for performance.
- [ ] **7.3** Include sidecar tables in export: `gobbs_sources`,
      `gobbs_identity_map`, `gobbs_final_id_map`. These enable future
      traceability and potential resync.
- [ ] **7.4** Export attachment files to target upload directory with
      canonical naming matching `mybb_attachments.attachname`.
- [ ] **7.5** Post-export verification:
  - Row counts in MySQL match `merged.sqlite3`.
  - FK integrity checks in MySQL.
  - MyBB boots and renders the forum index.
  - Thread/post/user profile pages load correctly.
  - Attachment downloads work.
- [ ] **7.6** Final cutover procedure:
  1. Put Origin read-only.
  2. Re-import Origin into `origin.sqlite3` (final snapshot).
  3. Re-run merge with "Origin wins if changed since dev snapshot" rule.
  4. Re-run build, review, export.
  5. Push to target MySQL.
  6. Verify.
  7. Go live.

---

## Resolved Questions

- **Plugins**: Don't merge plugin tables (e.g. `_bam`). Charles can
  reinstall plugins on the target instance after cutover.
- **Custom profile fields**: None beyond `fid1`/`fid2`/`fid3`.
- **Data volume**: <2,000 active users, <10,000 posts. Small enough
  that performance is not a concern; correctness is the only priority.
- **Target MySQL**: Drop and rebuild fresh. Use the same CREATE TABLE
  statements and prefix (`mybb_`) as Origin's current schema.
- **User exclusions**: Only flag users that are banned on one fork but
  not the other. No other exclusions needed.
- **PMs**: Exclude from merge entirely. PMs can be retrieved from
  backups and emailed to users on request.

## Appendix: FK Columns Requiring Rewrite

A non-exhaustive inventory of columns that contain entity IDs and must be
rewritten during the build phase. Each column references a specific kind.

| Table               | Column              | References kind |
|---------------------|---------------------|-----------------|
| `_posts`            | `tid`               | thread          |
| `_posts`            | `fid`               | forum           |
| `_posts`            | `uid`               | user            |
| `_posts`            | `replyto`           | post            |
| `_threads`          | `fid`               | forum           |
| `_threads`          | `uid`               | user            |
| `_threads`          | `firstpost`         | post            |
| `_threads`          | `lastpost`          | post            |
| `_threads`          | `lastposteruid`     | user            |
| `_attachments`      | `pid`               | post            |
| `_attachments`      | `uid`               | user            |
| `_polls`            | `tid`               | thread          |
| `_pollvotes`        | `pid`               | poll            |
| `_pollvotes`        | `uid`               | user            |
| `_forums`           | `pid` (parent)      | forum           |
| `_forums`           | `lastposteruid`     | user            |
| `_forums`           | `lastposttid`       | thread          |
| `_announcements`    | `fid`               | forum           |
| `_announcements`    | `uid`               | user            |
| `_reputation`       | `uid`               | user            |
| `_reputation`       | `adduid`            | user            |
| `_reputation`       | `pid`               | post            |
| `_warnings`         | `uid`               | user            |
| `_warnings`         | `issuedby`          | user            |
| `_warnings`         | `tid`               | thread          |
| `_warnings`         | `pid`               | post            |
| `_banned`           | `uid`               | user            |
| `_banned`           | `admin`             | user            |
| `_users`            | `usergroup`         | usergroup       |
| `_users`            | `displaygroup`      | usergroup       |
| `_users`            | `referrer`          | user            |
| `_users`            | `additionalgroups`  | usergroup (CSV) |
| `_banned`           | `oldadditionalgroups` | usergroup (CSV) |
| `_forums`           | `parentlist`        | forum (CSV)     |
| `_threadprefixes`   | `forums`            | forum (CSV)     |
| `_attachtypes`      | `groups`            | usergroup (CSV) |
| `_attachtypes`      | `forums`            | forum (CSV)     |
