-- name: GetThreadsForForum :one
SELECT threads, unapprovedthreads, deletedthreads
FROM forums
WHERE fid = ?
;

-- name: GetThreadUsers :many
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
;

-- name: GetThread :one
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
WHERE tid = ?;

-- name: FetchThreads :many
select tid,
       username,
       dateline,
       subject
from threads
where fid = ?
order by dateline
;