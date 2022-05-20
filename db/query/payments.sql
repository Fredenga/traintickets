-- name: Createpayment :one
INSERT INTO payments(
    ticket_id,
    amount,
    credit_card_number
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: Getpayment :one
SELECT * FROM payments
WHERE ticket_id = $1
LIMIT 1;

-- name: Listpayments :many
SELECT * FROM payments
LIMIT $1;

