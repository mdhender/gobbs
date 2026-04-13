# AGENTS.md

This file gives future agents the project context we want to avoid repeating in every session.

## Mission

`gobbs` merges two MyBB forum forks that diverged a few years ago back into a single coherent dataset. The forums are both still live, so resyncing from either source is possible.

Along the way we built tooling to import MyBB MySQL data into a local SQLite database and a local preview server so we can inspect the data without running a MyBB instance.

## Merge Strategy

The merge works by assigning UUIDs to every record:

1. Create a mapping table with tuples `(source, kind, id, uuid)`.
2. Iterate over users, forums, threads, posts, attachments, etc. from both forks.
3. Identical records across the two sources receive the **same UUID**.
4. After assignment, verify there are no conflicting rows where `(source, kind, id)` maps to different UUIDs.
5. Export the deduplicated, UUID-keyed records into an empty MyBB instance to complete the merge.

Because we preserve the original `(source, kind, id)` triples, we can re-import and resync from either live fork at any time.

## Current Project Shape

- `cmd/mysql2sqlite` — converts the MySQL schema for SQLite
- `cmd/mysql2sqlite-data` — imports MyBB data into `mybb.sqlite3`
- `cmd/mysql2sqlite-verify` — compares MySQL and SQLite row counts
- `cmd/gobbs-serve` — runs the local preview server
- `cmd/gobbs-static` — exports the static site to `public/`
- `internal/forumsite` — shared forum loader, renderer, templates, and CSS
- `internal/setupjson` — setup/configuration loader
- `uploads/` — archived attachment files that must ship with the generated site

## Working Assumptions

- The two source forums are both live MyBB instances that forked from a common ancestor.
- `mybb.sqlite3` is the local working copy of imported forum data.
- The preview server and static export exist to let us inspect data visually without running MyBB.
- Preview and static export should stay visually and structurally aligned.
- Attachment paths matter: deploy `public/` and `uploads/` together.
- When in doubt, choose the option that best preserves the original content and navigability.
- Treat the project as effectively greenfield when needed: prefer a cleaner long-term architecture over preserving weak existing abstractions.
- Larger refactors are welcome when they improve maintainability, fidelity, or merge accuracy.

## Preferred Workflow

1. Inspect the existing renderer, templates, and data flow before changing behavior.
2. Use preview mode for iteration: `make serve`, `air`, or `make live`.
3. When templates or styling change, verify the result in preview before exporting.
4. When output looks right, regenerate the static site with `make build`.
5. If a data import or schema change is involved, use the conversion/import/verify tools to confirm data integrity.

## Change Priorities

Prioritize work in this order:

1. correctness of merge logic and data integrity
2. faithfulness to original content (authorship, timestamps, structure)
3. attachment preservation and link integrity
4. forum/thread navigation and discoverability
5. visual polish and historical feel
6. implementation neatness

## Guardrails

- Do not remove or rewrite archived content unless there is clear evidence it is incorrect.
- Avoid introducing features that make the archive feel like a live forum.
- Keep template and export behavior consistent, but do not preserve incidental implementation details just for backward compatibility.
- Preserve user-made changes already present in the worktree unless explicitly asked to revert them.
- The `(source, kind, id)` triple is the canonical identity for every record; never discard it.
- Document meaningful project-level decisions here when they become stable expectations.
