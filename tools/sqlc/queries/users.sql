
--  Copyright (c) 2024 Michael D Henderson. All rights reserved.

-- name: GetUserFields :one
SELECT u.*, f.*
FROM users u
         LEFT JOIN userfields f ON (f.ufid = u.uid)
WHERE u.uid = ?
LIMIT 1
;