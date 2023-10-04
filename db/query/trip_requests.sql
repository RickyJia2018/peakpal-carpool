-- name: CreateTripRequest :one
INSERT INTO trip_requests (
    trip_id,
    passenger_id,
    boarding_station,
    payment_type,
    currency,
    contact_info
)VALUES(
    $1,$2,$3,$4,$5,$6
)RETURNING *;


-- name: GetTripRequest :one   
SELECT * FROM trip_requests
WHERE id = $1 LIMIT 1;

-- name: ListTripRequests :many
SELECT * FROM trip_requests
WHERE trip_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: DeleteTripRequest :exec
DELETE FROM trip_requests
WHERE id = $1;