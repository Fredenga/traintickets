-- name: Createuser :one
INSERT INTO users(
    email, 
    password
) VALUES (
    $1, $2
) RETURNING *;

-- name: Getuser :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: Listusers :many
SELECT * FROM users
ORDER BY email
LIMIT $1;

-- name: Deleteuser :exec
DELETE FROM users
WHERE email = $1;