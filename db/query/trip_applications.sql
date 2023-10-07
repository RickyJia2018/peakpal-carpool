-- name: CreateTripApplication :one
INSERT INTO trip_applications (
    trip_id,
    passenger_id,
    boarding_station,
    payment_type,
    currency,
    total_passenger,
    contact_info
)VALUES(
    $1,$2,$3,$4,$5,$6,$7
)RETURNING *;


-- name: GetTripApplication :one   
SELECT
    ta.*,
    t.driver_id
FROM
    trip_applications ta
LEFT JOIN trips t ON ta.trip_id = t.id
-- LEFT JOIN stations s ON ta.boarding_station = s.id
WHERE
    ta.id = $1;

-- name: ListTripApplications :many
SELECT * FROM trip_applications
WHERE trip_id = sqlc.narg(trip_id) OR passenger_id = sqlc.narg(passenger_id) 
ORDER BY created_at
LIMIT $1 OFFSET $2;

-- name: DeleteTripApplication :exec
DELETE FROM trip_applications
WHERE id = $1;