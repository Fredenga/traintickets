// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: timetable.sql

package db

import (
	"context"
	"time"
)

const createtimetable = `-- name: Createtimetable :one
INSERT INTO timetable(
    route, 
    arrival,
    departure,
    distance,
    price
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING route, arrival, departure, distance, price
`

type CreatetimetableParams struct {
	Route     string    `json:"route"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
	Distance  int32     `json:"distance"`
	Price     int32     `json:"price"`
}

func (q *Queries) Createtimetable(ctx context.Context, arg CreatetimetableParams) (Timetable, error) {
	row := q.db.QueryRowContext(ctx, createtimetable,
		arg.Route,
		arg.Arrival,
		arg.Departure,
		arg.Distance,
		arg.Price,
	)
	var i Timetable
	err := row.Scan(
		&i.Route,
		&i.Arrival,
		&i.Departure,
		&i.Distance,
		&i.Price,
	)
	return i, err
}

const deletetimetable = `-- name: Deletetimetable :exec
DELETE FROM timetable
WHERE route = $1
`

func (q *Queries) Deletetimetable(ctx context.Context, route string) error {
	_, err := q.db.ExecContext(ctx, deletetimetable, route)
	return err
}

const gettimetable = `-- name: Gettimetable :one
SELECT route, arrival, departure, distance, price FROM timetable
WHERE route = $1
LIMIT 1
`

func (q *Queries) Gettimetable(ctx context.Context, route string) (Timetable, error) {
	row := q.db.QueryRowContext(ctx, gettimetable, route)
	var i Timetable
	err := row.Scan(
		&i.Route,
		&i.Arrival,
		&i.Departure,
		&i.Distance,
		&i.Price,
	)
	return i, err
}

const listtimetables = `-- name: Listtimetables :many
SELECT route, arrival, departure, distance, price FROM timetable
ORDER BY route
LIMIT $1
`

func (q *Queries) Listtimetables(ctx context.Context, limit int32) ([]Timetable, error) {
	rows, err := q.db.QueryContext(ctx, listtimetables, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Timetable
	for rows.Next() {
		var i Timetable
		if err := rows.Scan(
			&i.Route,
			&i.Arrival,
			&i.Departure,
			&i.Distance,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
