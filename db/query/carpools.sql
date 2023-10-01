-- name: CreateCarpool :one
INSERT INTO carpools (
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