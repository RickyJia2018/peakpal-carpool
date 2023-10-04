// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: trip_requests.sql

package db

import (
	"context"
)

const createTripRequest = `-- name: CreateTripRequest :one
INSERT INTO trip_requests (
    trip_id,
    passenger_id,
    boarding_station,
    payment_type,
    currency,
    contact_info
)VALUES(
    $1,$2,$3,$4,$5,$6
)RETURNING id, trip_id, passenger_id, boarding_station, payment_type, currency, contact_info, approved, created_at
`

type CreateTripRequestParams struct {
	TripID          int64  `json:"trip_id"`
	PassengerID     int64  `json:"passenger_id"`
	BoardingStation int64  `json:"boarding_station"`
	PaymentType     string `json:"payment_type"`
	Currency        string `json:"currency"`
	ContactInfo     string `json:"contact_info"`
}

func (q *Queries) CreateTripRequest(ctx context.Context, arg CreateTripRequestParams) (TripRequest, error) {
	row := q.db.QueryRow(ctx, createTripRequest,
		arg.TripID,
		arg.PassengerID,
		arg.BoardingStation,
		arg.PaymentType,
		arg.Currency,
		arg.ContactInfo,
	)
	var i TripRequest
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.PassengerID,
		&i.BoardingStation,
		&i.PaymentType,
		&i.Currency,
		&i.ContactInfo,
		&i.Approved,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTripRequest = `-- name: DeleteTripRequest :exec
DELETE FROM trip_requests
WHERE id = $1
`

func (q *Queries) DeleteTripRequest(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTripRequest, id)
	return err
}

const getTripRequest = `-- name: GetTripRequest :one
SELECT id, trip_id, passenger_id, boarding_station, payment_type, currency, contact_info, approved, created_at FROM trip_requests
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTripRequest(ctx context.Context, id int64) (TripRequest, error) {
	row := q.db.QueryRow(ctx, getTripRequest, id)
	var i TripRequest
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.PassengerID,
		&i.BoardingStation,
		&i.PaymentType,
		&i.Currency,
		&i.ContactInfo,
		&i.Approved,
		&i.CreatedAt,
	)
	return i, err
}

const listTripRequests = `-- name: ListTripRequests :many
SELECT id, trip_id, passenger_id, boarding_station, payment_type, currency, contact_info, approved, created_at FROM trip_requests
WHERE trip_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3
`

type ListTripRequestsParams struct {
	TripID int64 `json:"trip_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTripRequests(ctx context.Context, arg ListTripRequestsParams) ([]TripRequest, error) {
	rows, err := q.db.Query(ctx, listTripRequests, arg.TripID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TripRequest{}
	for rows.Next() {
		var i TripRequest
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.PassengerID,
			&i.BoardingStation,
			&i.PaymentType,
			&i.Currency,
			&i.ContactInfo,
			&i.Approved,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}