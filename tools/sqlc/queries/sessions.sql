-- name: GetSessionData :one
SELECT *
FROM sessions
WHERE sid = ?
;

