// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Passenger struct {
	ID          int64            `json:"id"`
	PassengerID int64            `json:"passenger_id"`
	TripID      int64            `json:"trip_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type Station struct {
	ID          int64  `json:"id"`
	TripID      int64  `json:"trip_id"`
	StationName string `json:"station_name"`
	// How long it take from first station
	ArrivalMinutes int32            `json:"arrival_minutes"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
}

type Trip struct {
	ID           int64            `json:"id"`
	ContactInfo  string           `json:"contact_info"`
	DriverID     int64            `json:"driver_id"`
	MaxPassenger int32            `json:"max_passenger"`
	Price        int32            `json:"price"`
	AblePickUp   bool             `json:"able_pick_up"`
	ResortID     int64            `json:"resort_id"`
	DepartureAt  pgtype.Timestamp `json:"departure_at"`
	ReturnAt     pgtype.Timestamp `json:"return_at"`
	RoundTrip    bool             `json:"round_trip"`
	// provider write
	AcceptPaymentType string           `json:"accept_payment_type"`
	Currency          string           `json:"currency"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
}

type TripRequest struct {
	ID              int64            `json:"id"`
	TripID          int64            `json:"trip_id"`
	PassengerID     int64            `json:"passenger_id"`
	BoardingStation int64            `json:"boarding_station"`
	PaymentType     string           `json:"payment_type"`
	Currency        string           `json:"currency"`
	ContactInfo     string           `json:"contact_info"`
	Approved        bool             `json:"approved"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}