// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: trip_applications.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTripApplication = `-- name: CreateTripApplication :one
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
)RETURNING id, trip_id, passenger_id, boarding_station, payment_type, currency, contact_info, total_passenger, approved, created_at
`

type CreateTripApplicationParams struct {
	TripID          int64  `json:"trip_id"`
	PassengerID     int64  `json:"passenger_id"`
	BoardingStation int64  `json:"boarding_station"`
	PaymentType     string `json:"payment_type"`
	Currency        string `json:"currency"`
	TotalPassenger  int32  `json:"total_passenger"`
	ContactInfo     string `json:"contact_info"`
}

func (q *Queries) CreateTripApplication(ctx context.Context, arg CreateTripApplicationParams) (TripApplication, error) {
	row := q.db.QueryRow(ctx, createTripApplication,
		arg.TripID,
		arg.PassengerID,
		arg.BoardingStation,
		arg.PaymentType,
		arg.Currency,
		arg.TotalPassenger,
		arg.ContactInfo,
	)
	var i TripApplication
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.PassengerID,
		&i.BoardingStation,
		&i.PaymentType,
		&i.Currency,
		&i.ContactInfo,
		&i.TotalPassenger,
		&i.Approved,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTripApplication = `-- name: DeleteTripApplication :exec
DELETE FROM trip_applications
WHERE id = $1
`

func (q *Queries) DeleteTripApplication(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTripApplication, id)
	return err
}

const getTripApplication = `-- name: GetTripApplication :one
SELECT
    ta.id, ta.trip_id, ta.passenger_id, ta.boarding_station, ta.payment_type, ta.currency, ta.contact_info, ta.total_passenger, ta.approved, ta.created_at,
    t.driver_id
FROM
    trip_applications ta
LEFT JOIN trips t ON ta.trip_id = t.id
WHERE
    ta.id = $1
`

type GetTripApplicationRow struct {
	ID              int64       `json:"id"`
	TripID          int64       `json:"trip_id"`
	PassengerID     int64       `json:"passenger_id"`
	BoardingStation int64       `json:"boarding_station"`
	PaymentType     string      `json:"payment_type"`
	Currency        string      `json:"currency"`
	ContactInfo     string      `json:"contact_info"`
	TotalPassenger  int32       `json:"total_passenger"`
	Approved        bool        `json:"approved"`
	CreatedAt       time.Time   `json:"created_at"`
	DriverID        pgtype.Int8 `json:"driver_id"`
}

// LEFT JOIN stations s ON ta.boarding_station = s.id
func (q *Queries) GetTripApplication(ctx context.Context, id int64) (GetTripApplicationRow, error) {
	row := q.db.QueryRow(ctx, getTripApplication, id)
	var i GetTripApplicationRow
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.PassengerID,
		&i.BoardingStation,
		&i.PaymentType,
		&i.Currency,
		&i.ContactInfo,
		&i.TotalPassenger,
		&i.Approved,
		&i.CreatedAt,
		&i.DriverID,
	)
	return i, err
}

const listTripApplications = `-- name: ListTripApplications :many
SELECT id, trip_id, passenger_id, boarding_station, payment_type, currency, contact_info, total_passenger, approved, created_at FROM trip_applications
WHERE trip_id = $3 OR passenger_id = $4 
ORDER BY created_at
LIMIT $1 OFFSET $2
`

type ListTripApplicationsParams struct {
	Limit       int32       `json:"limit"`
	Offset      int32       `json:"offset"`
	TripID      pgtype.Int8 `json:"trip_id"`
	PassengerID pgtype.Int8 `json:"passenger_id"`
}

func (q *Queries) ListTripApplications(ctx context.Context, arg ListTripApplicationsParams) ([]TripApplication, error) {
	rows, err := q.db.Query(ctx, listTripApplications,
		arg.Limit,
		arg.Offset,
		arg.TripID,
		arg.PassengerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TripApplication{}
	for rows.Next() {
		var i TripApplication
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.PassengerID,
			&i.BoardingStation,
			&i.PaymentType,
			&i.Currency,
			&i.ContactInfo,
			&i.TotalPassenger,
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
