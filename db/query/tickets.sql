-- name: Createticket :one
INSERT INTO tickets(
    route,
    train_number,
    coach_number,
    seat_number,
    booking_date,
    trip_date,
    fare,
    email
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: Getticket :one
SELECT * FROM tickets
WHERE ticket_id = $1
LIMIT 1;

-- name: Listtickets :many
SELECT * FROM tickets
ORDER BY ticket_id
LIMIT $1;

-- name: Deleteticket :exec
DELETE FROM tickets
WHERE ticket_id = $1;