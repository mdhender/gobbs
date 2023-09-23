
-- name: GetUserFields :one
SELECT u.*, f.*
FROM PBMnet_users u
         LEFT JOIN PBMnet_userfields f ON (f.ufid = u.uid)
WHERE u.uid = ?
LIMIT 1
;