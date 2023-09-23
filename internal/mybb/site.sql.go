// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: site.sql

package mybb

import (
	"context"
)

const getTitleCache = `-- name: GetTitleCache :one
SELECT title, cache
FROM PBMnet_datacache
`

func (q *Queries) GetTitleCache(ctx context.Context) (PbmnetDatacache, error) {
	row := q.db.QueryRowContext(ctx, getTitleCache)
	var i PbmnetDatacache
	err := row.Scan(&i.Title, &i.Cache)
	return i, err
}