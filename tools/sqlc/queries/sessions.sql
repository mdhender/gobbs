--  Copyright (c) 2024 Michael D Henderson. All rights reserved.

-- name: GetSessionData :one
SELECT *
FROM sessions
WHERE sid = ?
;

