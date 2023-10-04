-- name: CreateTrip :one
INSERT INTO trips (
    contact_info,
    driver_id,
    max_passenger,
    price,
    able_pick_up,
    resort_id,
    departure_at,
    accept_payment_type,
    currency
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9
)RETURNING *;

-- name: CountTrip :one
SELECT count(*) FROM trips
WHERE resort_id = $1;

-- name: GetTrip :one   
SELECT * FROM trips
WHERE id = $1 LIMIT 1;

-- name: ListAllTrips :many
SELECT * FROM trips
WHERE resort_id = $1 AND departure_at>(now())
ORDER BY departure_at
LIMIT $2 OFFSET $3;

-- name: UpdateTrip :one
UPDATE trips
SET
     contact_info = COALESCE(sqlc.narg(contact_info), contact_info),
     max_passenger = COALESCE(sqlc.narg(max_passenger), max_passenger),
     able_pick_up = COALESCE(sqlc.narg(able_pick_up),able_pick_up)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTrip :exec
DELETE FROM trips
WHERE id = $1;