// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: sessions.sql

package mybb

import (
	"context"
)

const getSessionData = `-- name: GetSessionData :one
SELECT sid, uid, ip, time, location, useragent, anonymous, nopermission, location1, location2
FROM PBMnet_sessions
WHERE sid = ?
`

func (q *Queries) GetSessionData(ctx context.Context, sid string) (PbmnetSession, error) {
	row := q.db.QueryRowContext(ctx, getSessionData, sid)
	var i PbmnetSession
	err := row.Scan(
		&i.Sid,
		&i.Uid,
		&i.Ip,
		&i.Time,
		&i.Location,
		&i.Useragent,
		&i.Anonymous,
		&i.Nopermission,
		&i.Location1,
		&i.Location2,
	)
	return i, err
}
