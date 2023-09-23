-- name: GetSessionData :one
SELECT *
FROM PBMnet_sessions
WHERE sid = ?
;

