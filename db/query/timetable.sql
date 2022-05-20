-- name: Createtimetable :one
INSERT INTO timetable(
    route, 
    arrival,
    departure,
    distance,
    price
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: Gettimetable :one
SELECT * FROM timetable
WHERE route = $1
LIMIT 1;

-- name: Listtimetables :many
SELECT * FROM timetable
ORDER BY route
LIMIT $1;

-- name: Deletetimetable :exec
DELETE FROM timetable
WHERE route = $1;