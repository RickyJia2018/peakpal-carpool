// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: trips.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const countDriverTrips = `-- name: CountDriverTrips :one
SELECT count(*) FROM trips
WHERE driver_id = $1
`

func (q *Queries) CountDriverTrips(ctx context.Context, driverID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countDriverTrips, driverID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countFutureTrips = `-- name: CountFutureTrips :one
SELECT count(*) FROM trips
WHERE resort_id = $1 AND departure_at>(now())
`

func (q *Queries) CountFutureTrips(ctx context.Context, resortID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countFutureTrips, resortID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countTrip = `-- name: CountTrip :one
SELECT count(*) FROM trips
WHERE resort_id = $1
`

func (q *Queries) CountTrip(ctx context.Context, resortID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countTrip, resortID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTrip = `-- name: CreateTrip :one
INSERT INTO trips (
    contact_info,
    driver_id,
    max_passenger,
    price,
    able_pick_up,
    resort_id,
    round_trip,
    departure_at,
    return_at,
    accept_payment_type,
    currency
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
)RETURNING id, contact_info, driver_id, max_passenger, price, able_pick_up, resort_id, departure_at, return_at, round_trip, accept_payment_type, currency, created_at
`

type CreateTripParams struct {
	ContactInfo       string    `json:"contact_info"`
	DriverID          int64     `json:"driver_id"`
	MaxPassenger      int32     `json:"max_passenger"`
	Price             int32     `json:"price"`
	AblePickUp        bool      `json:"able_pick_up"`
	ResortID          int64     `json:"resort_id"`
	RoundTrip         bool      `json:"round_trip"`
	DepartureAt       time.Time `json:"departure_at"`
	ReturnAt          time.Time `json:"return_at"`
	AcceptPaymentType string    `json:"accept_payment_type"`
	Currency          string    `json:"currency"`
}

func (q *Queries) CreateTrip(ctx context.Context, arg CreateTripParams) (Trip, error) {
	row := q.db.QueryRow(ctx, createTrip,
		arg.ContactInfo,
		arg.DriverID,
		arg.MaxPassenger,
		arg.Price,
		arg.AblePickUp,
		arg.ResortID,
		arg.RoundTrip,
		arg.DepartureAt,
		arg.ReturnAt,
		arg.AcceptPaymentType,
		arg.Currency,
	)
	var i Trip
	err := row.Scan(
		&i.ID,
		&i.ContactInfo,
		&i.DriverID,
		&i.MaxPassenger,
		&i.Price,
		&i.AblePickUp,
		&i.ResortID,
		&i.DepartureAt,
		&i.ReturnAt,
		&i.RoundTrip,
		&i.AcceptPaymentType,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTrip = `-- name: DeleteTrip :exec
DELETE FROM trips
WHERE id = $1
`

func (q *Queries) DeleteTrip(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTrip, id)
	return err
}

const getTrip = `-- name: GetTrip :one
SELECT id, contact_info, driver_id, max_passenger, price, able_pick_up, resort_id, departure_at, return_at, round_trip, accept_payment_type, currency, created_at FROM trips
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTrip(ctx context.Context, id int64) (Trip, error) {
	row := q.db.QueryRow(ctx, getTrip, id)
	var i Trip
	err := row.Scan(
		&i.ID,
		&i.ContactInfo,
		&i.DriverID,
		&i.MaxPassenger,
		&i.Price,
		&i.AblePickUp,
		&i.ResortID,
		&i.DepartureAt,
		&i.ReturnAt,
		&i.RoundTrip,
		&i.AcceptPaymentType,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listDriverTrips = `-- name: ListDriverTrips :many
SELECT id, contact_info, driver_id, max_passenger, price, able_pick_up, resort_id, departure_at, return_at, round_trip, accept_payment_type, currency, created_at FROM trips
WHERE driver_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListDriverTripsParams struct {
	DriverID int64 `json:"driver_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListDriverTrips(ctx context.Context, arg ListDriverTripsParams) ([]Trip, error) {
	rows, err := q.db.Query(ctx, listDriverTrips, arg.DriverID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Trip{}
	for rows.Next() {
		var i Trip
		if err := rows.Scan(
			&i.ID,
			&i.ContactInfo,
			&i.DriverID,
			&i.MaxPassenger,
			&i.Price,
			&i.AblePickUp,
			&i.ResortID,
			&i.DepartureAt,
			&i.ReturnAt,
			&i.RoundTrip,
			&i.AcceptPaymentType,
			&i.Currency,
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

const listFutureTrips = `-- name: ListFutureTrips :many
SELECT id, contact_info, driver_id, max_passenger, price, able_pick_up, resort_id, departure_at, return_at, round_trip, accept_payment_type, currency, created_at FROM trips
WHERE resort_id = $1 AND departure_at>(now())
ORDER BY departure_at
LIMIT $2 OFFSET $3
`

type ListFutureTripsParams struct {
	ResortID int64 `json:"resort_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListFutureTrips(ctx context.Context, arg ListFutureTripsParams) ([]Trip, error) {
	rows, err := q.db.Query(ctx, listFutureTrips, arg.ResortID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Trip{}
	for rows.Next() {
		var i Trip
		if err := rows.Scan(
			&i.ID,
			&i.ContactInfo,
			&i.DriverID,
			&i.MaxPassenger,
			&i.Price,
			&i.AblePickUp,
			&i.ResortID,
			&i.DepartureAt,
			&i.ReturnAt,
			&i.RoundTrip,
			&i.AcceptPaymentType,
			&i.Currency,
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

const updateTrip = `-- name: UpdateTrip :one
UPDATE trips
SET
     contact_info = COALESCE($1, contact_info),
     max_passenger = COALESCE($2, max_passenger),
     able_pick_up = COALESCE($3,able_pick_up)
WHERE id = $4
RETURNING id, contact_info, driver_id, max_passenger, price, able_pick_up, resort_id, departure_at, return_at, round_trip, accept_payment_type, currency, created_at
`

type UpdateTripParams struct {
	ContactInfo  pgtype.Text `json:"contact_info"`
	MaxPassenger pgtype.Int4 `json:"max_passenger"`
	AblePickUp   pgtype.Bool `json:"able_pick_up"`
	ID           int64       `json:"id"`
}

func (q *Queries) UpdateTrip(ctx context.Context, arg UpdateTripParams) (Trip, error) {
	row := q.db.QueryRow(ctx, updateTrip,
		arg.ContactInfo,
		arg.MaxPassenger,
		arg.AblePickUp,
		arg.ID,
	)
	var i Trip
	err := row.Scan(
		&i.ID,
		&i.ContactInfo,
		&i.DriverID,
		&i.MaxPassenger,
		&i.Price,
		&i.AblePickUp,
		&i.ResortID,
		&i.DepartureAt,
		&i.ReturnAt,
		&i.RoundTrip,
		&i.AcceptPaymentType,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
