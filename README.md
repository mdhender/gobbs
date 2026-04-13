# gobbs

Tools for converting a MyBB forum into SQLite and rendering it as a read-only static forum archive.

## What Is Here

- `cmd/mysql2sqlite` converts the MySQL schema to SQLite-friendly SQL.
- `cmd/mysql2sqlite-data` imports MyBB data into `mybb.sqlite3`.
- `cmd/mysql2sqlite-verify` compares row counts between MySQL and SQLite.
- `cmd/gobbs-serve` runs a local preview server for the archive.
- `cmd/gobbs-static` writes the static site to `public/`.
- `internal/forumsite` contains the shared loader, renderer, templates, and styles used by both preview and export.

## Development Loop

Preview mode is intended for day-to-day work. It reads the SQLite archive directly and uses live templates from disk, so changes to Go code, HTML templates, and CSS show up on refresh.

Run the preview server directly:

```bash
env GOCACHE=$(pwd)/.gocache go run ./cmd/gobbs-serve
```

Or use:

```bash
make serve
```

Or use `air` for automatic rebuilds:

```bash
air
```

Or:

```bash
make air
```

For automatic browser refresh while you work, run:

```bash
make live
```

That starts `air` plus a BrowserSync proxy. Open `http://localhost:3000` instead of `:8080` and the page will refresh automatically when templates, CSS, or the rebuilt preview binary change.

If you prefer a named local domain such as `http://gobbs.test:3000`, add an `/etc/hosts` entry like `127.0.0.1 gobbs.test` and start BrowserSync with `BROWSER_SYNC_HOST=gobbs.test`.

If you prefer a single multiplexed terminal, use `overmind` with the included [Procfile.dev](/Users/wraith/Software/mdhender/gobbs/Procfile.dev:1):

```bash
overmind start -f Procfile.dev
```

That runs the Go preview server watcher and the BrowserSync proxy together. Open `http://localhost:3000` for the auto-refreshing preview.

The included [`.air.toml`](/Users/wraith/Software/mdhender/gobbs/.air.toml:1) is configured to:

- build `cmd/gobbs-serve`
- use a repo-local `.gocache`
- watch `cmd/`, `internal/`, `.go`, `.html`, and `.css` files

`make live` uses `npx --yes browser-sync`, so it expects `npm`/`npx` to be available on your machine.

Example:

```bash
BROWSER_SYNC_HOST=gobbs.test overmind start -f Procfile.dev
```

## Static Export

When the preview looks right, generate the static site:

```bash
env GOCACHE=$(pwd)/.gocache go run ./cmd/gobbs-static -out public
```

Or:

```bash
make build
```

That writes:

- `public/index.html`
- `public/f/<fid>/index.html`
- `public/t/<tid>/index.html`
- `public/assets/site.css`
- `public/uploads/...`

By default, the static builder copies the configured `uploads/` tree into `public/uploads/` so attachments and avatars ship with the exported archive.

## Templates

Live preview reads template assets from:

- [internal/forumsite/templates/site.html](/Users/wraith/Software/mdhender/gobbs/internal/forumsite/templates/site.html:1)
- [internal/forumsite/templates/site.css](/Users/wraith/Software/mdhender/gobbs/internal/forumsite/templates/site.css:1)

Static export uses embedded copies of those same files, so the preview and export paths stay aligned.

## Notes

- `.env`, `setup.json`, and `*.sqlite3` are ignored by Git.
- `GOCACHE` only controls where Go stores build cache artifacts; it does not change program behavior.
- `gobbs-static` now copies the configured `uploads/` tree into `public/uploads/` so exported attachments and avatars travel with the generated site by default.
- The current renderer handles common MyBB BBCode and attachments, but there is still room to improve HTML-heavy forum descriptions and some quote edge cases.
