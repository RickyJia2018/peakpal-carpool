-- name: CreateStation :one
INSERT INTO stations (
    trip_id,
    station_name,
    arrival_minutes
)VALUES(
    $1,$2,$3
)RETURNING *;

-- name: GetStation :one
SELECT 
    s.id AS station_id,
    s.station_name,
    s.arrival_minutes,
    t.driver_id,
    p.passenger_id
FROM
    stations s
LEFT JOIN trips t ON s.tip_id = t.id
LEFT JOIN trip_applications p ON s.tip_id = p.tip_id
WHERE
    s.id = $1;

-- name: ListStations :many
SELECT * FROM stations
WHERE trip_id = $1
ORDER BY arrival_minutes;

-- name: DeleteStation :exec
DELETE FROM stations
WHERE id = $1;