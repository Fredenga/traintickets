-- name: Createtrain :one
INSERT INTO trains( 
    type, 
    class, 
    max_passenger_no,
    max_speed,
    route
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: Gettrain :one
SELECT * FROM trains
WHERE train_number = $1
LIMIT 1;

-- name: Listtrains :many
SELECT * FROM trains
ORDER BY train_number
LIMIT $1;

-- name: Listtrainsbyroutes :many
SELECT * FROM trains
WHERE route = $1
ORDER BY train_number
LIMIT $2;

-- name: Deletetrain :exec
DELETE FROM trains
WHERE train_number = $1;