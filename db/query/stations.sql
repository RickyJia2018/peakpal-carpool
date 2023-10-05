-- name: CreateStation :one
INSERT INTO stations (
    trip_id,
    station_name,
    arrival_minutes
)VALUES(
    $1,$2,$3
)RETURNING *;

-- name: ListStations :many
SELECT * FROM stations
WHERE trip_id = $1
ORDER BY arrival_minutes;

-- name: DeleteStation :exec
DELETE FROM stations
WHERE id = $1;