-- name: CreatePassenger :one
INSERT INTO passengers (
    passenger_id,
    trip_id
)VALUES(
    $1,$2
)RETURNING *;

-- name: ListPassengers :many
SELECT * FROM passengers
WHERE trip_id = $1
ORDER BY created_at;

-- name: DeletePassenger :exec
DELETE FROM passengers
WHERE id = $1;