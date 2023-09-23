# SQLC Generator

This folder contains the files used by the `sqlc` command to generate the model and queries.

## Usage

Run the `./generate.sh` script from this folder.
As far as I can tell, it creates files if they're missing, overwrites them if they already exists.
It never seems to delete files, so you may need to clean things up if you remove or rename any of the SQL files in this folder.

The `schema/schema.sql` file creates the `internal/mybb/models.go` file.

The SQL files in the `queries` folder create the ".sql.go" files in the `internal/mybb` package.
For example, `queries/site.sql` creates `internal/mybb/site.sql.go`.

