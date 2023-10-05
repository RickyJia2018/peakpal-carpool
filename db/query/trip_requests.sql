-- name: CreateTripApplication :one
INSERT INTO trip_applications (
    trip_id,
    passenger_id,
    boarding_station,
    payment_type,
    currency,
    contact_info
)VALUES(
    $1,$2,$3,$4,$5,$6
)RETURNING *;


-- name: GetTripApplication :one   
SELECT * FROM trip_applications
WHERE id = $1 LIMIT 1;

-- name: ListTripApplications :many
SELECT * FROM trip_applications
WHERE trip_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: DeleteTripApplication :exec
DELETE FROM trip_applications
WHERE id = $1;