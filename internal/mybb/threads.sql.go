// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: threads.sql

package mybb

import (
	"context"
	"database/sql"
)

const fetchThreads = `-- name: FetchThreads :many
select tid,
       username,
       dateline,
       subject
from threads
where fid = ?
order by dateline
`

type FetchThreadsRow struct {
	Tid      int64
	Username string
	Dateline int64
	Subject  string
}

func (q *Queries) FetchThreads(ctx context.Context, fid int64) ([]FetchThreadsRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchThreads, fid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchThreadsRow
	for rows.Next() {
		var i FetchThreadsRow
		if err := rows.Scan(
			&i.Tid,
			&i.Username,
			&i.Dateline,
			&i.Subject,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getThread = `-- name: GetThread :one
;

SELECT tid,
       fid,
       subject,
       prefix,
       icon,
       poll,
       uid,
       username,
       dateline,
       firstpost,
       lastpost,
       lastposter,
       lastposteruid,
       views,
       replies,
       closed,
       sticky,
       numratings,
       totalratings,
       notes,
       visible,
       unapprovedposts,
       deletedposts,
       attachmentcount,
       deletetime
FROM threads
WHERE tid = ?
`

func (q *Queries) GetThread(ctx context.Context, tid int64) (Thread, error) {
	row := q.db.QueryRowContext(ctx, getThread, tid)
	var i Thread
	err := row.Scan(
		&i.Tid,
		&i.Fid,
		&i.Subject,
		&i.Prefix,
		&i.Icon,
		&i.Poll,
		&i.Uid,
		&i.Username,
		&i.Dateline,
		&i.Firstpost,
		&i.Lastpost,
		&i.Lastposter,
		&i.Lastposteruid,
		&i.Views,
		&i.Replies,
		&i.Closed,
		&i.Sticky,
		&i.Numratings,
		&i.Totalratings,
		&i.Notes,
		&i.Visible,
		&i.Unapprovedposts,
		&i.Deletedposts,
		&i.Attachmentcount,
		&i.Deletetime,
	)
	return i, err
}

const getThreadUsers = `-- name: GetThreadUsers :many
;

SELECT t.tid,
       t.fid,
       t.subject,
       t.prefix,
       t.icon,
       t.poll,
       t.uid,
       t.username,
       t.dateline,
       t.firstpost,
       t.lastpost,
       t.lastposter,
       t.lastposteruid,
       t.views,
       t.replies,
       t.closed,
       t.sticky,
       t.numratings,
       t.totalratings,
       t.notes,
       t.visible,
       t.unapprovedposts,
       t.deletedposts,
       t.attachmentcount,
       t.deletetime,
       (t.totalratings / t.numratings) AS averagerating,
       t.username                      AS threadusername,
       u.username
FROM threads t
         LEFT JOIN users u ON (u.uid = t.uid)
WHERE t.fid = ?
  AND (t.visible IN (1, -1, 0))
ORDER BY t.sticky DESC, t.lastpost desc
`

type GetThreadUsersRow struct {
	Tid             int64
	Fid             int64
	Subject         string
	Prefix          int64
	Icon            int64
	Poll            int64
	Uid             int64
	Username        string
	Dateline        int64
	Firstpost       int64
	Lastpost        int64
	Lastposter      string
	Lastposteruid   int64
	Views           int64
	Replies         int64
	Closed          string
	Sticky          int64
	Numratings      int64
	Totalratings    int64
	Notes           string
	Visible         int8
	Unapprovedposts int64
	Deletedposts    int64
	Attachmentcount int64
	Deletetime      int64
	Averagerating   interface{}
	Threadusername  string
	Username_2      sql.NullString
}

func (q *Queries) GetThreadUsers(ctx context.Context, fid int64) ([]GetThreadUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getThreadUsers, fid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetThreadUsersRow
	for rows.Next() {
		var i GetThreadUsersRow
		if err := rows.Scan(
			&i.Tid,
			&i.Fid,
			&i.Subject,
			&i.Prefix,
			&i.Icon,
			&i.Poll,
			&i.Uid,
			&i.Username,
			&i.Dateline,
			&i.Firstpost,
			&i.Lastpost,
			&i.Lastposter,
			&i.Lastposteruid,
			&i.Views,
			&i.Replies,
			&i.Closed,
			&i.Sticky,
			&i.Numratings,
			&i.Totalratings,
			&i.Notes,
			&i.Visible,
			&i.Unapprovedposts,
			&i.Deletedposts,
			&i.Attachmentcount,
			&i.Deletetime,
			&i.Averagerating,
			&i.Threadusername,
			&i.Username_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getThreadsForForum = `-- name: GetThreadsForForum :one
SELECT threads, unapprovedthreads, deletedthreads
FROM forums
WHERE fid = ?
`

type GetThreadsForForumRow struct {
	Threads           int64
	Unapprovedthreads int64
	Deletedthreads    int64
}

func (q *Queries) GetThreadsForForum(ctx context.Context, fid int64) (GetThreadsForForumRow, error) {
	row := q.db.QueryRowContext(ctx, getThreadsForForum, fid)
	var i GetThreadsForForumRow
	err := row.Scan(&i.Threads, &i.Unapprovedthreads, &i.Deletedthreads)
	return i, err
}
