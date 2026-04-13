# AGENTS.md

This file gives future agents the project context we want to avoid repeating in every session.

## Mission

`gobbs` exists to preserve the PlayByMail forums as a read-only archive.

Primary goals:

- faithfully preserve forum structure, thread content, authorship, timestamps, quotes, and attachments from the MyBB source data
- generate a static archive that is easy to host and browse long-term
- keep the rendered archive recognizable as the original forum, while improving clarity when helpful
- prefer preservation and correctness over cleverness

## Current Project Shape

- `cmd/mysql2sqlite` converts the MySQL schema for SQLite
- `cmd/mysql2sqlite-data` imports MyBB data into `mybb.sqlite3`
- `cmd/mysql2sqlite-verify` compares MySQL and SQLite row counts
- `cmd/gobbs-serve` runs the local preview server
- `cmd/gobbs-static` exports the static site to `public/`
- `internal/forumsite` contains the shared forum loader, renderer, templates, and CSS
- `uploads/` contains archived attachment files that must ship with the generated site

## Working Assumptions

- The forum being archived is `https://forums.playbymail.dev/`.
- The generated site is read-only; we are not rebuilding interactive forum behavior.
- Preview and static export should stay visually and structurally aligned.
- Attachment paths matter: deploy `public/` and `uploads/` together.
- When in doubt, choose the option that best preserves the original content and navigability.
- Treat the project as effectively greenfield when needed: prefer a cleaner long-term archive architecture over preserving weak existing abstractions.
- Larger refactors are welcome when they improve maintainability, fidelity, or layout accuracy.

## Preferred Workflow

1. Inspect the existing renderer, templates, and data flow before changing behavior.
2. Use preview mode for iteration: `make serve`, `air`, or `make live`.
3. When templates or styling change, verify the result in preview before exporting.
4. When output looks right, regenerate the static site with `make build`.
5. If a data import or schema change is involved, use the conversion/import/verify tools to confirm archive integrity.

## Change Priorities

Prioritize work in this order:

1. correctness of archived content
2. attachment preservation and link integrity
3. forum/thread navigation and discoverability
4. visual polish and historical feel
5. implementation neatness

## Guardrails

- Do not remove or rewrite archived content unless there is clear evidence it is incorrect.
- Avoid introducing features that make the archive feel like a live forum.
- Keep template and export behavior consistent, but do not preserve incidental implementation details just for backward compatibility.
- Preserve user-made changes already present in the worktree unless explicitly asked to revert them.
- Document meaningful project-level decisions here when they become stable expectations.
