--  Copyright (c) 2024 Michael D Henderson. All rights reserved.

-- name: FetchPosts :many
SELECT pid,
       username,
       dateline,
       subject,
       replyto,
       message
FROM posts
WHERE tid=?
ORDER BY dateline;
