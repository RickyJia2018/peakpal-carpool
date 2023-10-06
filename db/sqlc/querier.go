// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CountDriverTrips(ctx context.Context, driverID int64) (int64, error)
	CountFutureTrips(ctx context.Context, resortID int64) (int64, error)
	CountTrip(ctx context.Context, resortID int64) (int64, error)
	CreatePassenger(ctx context.Context, arg CreatePassengerParams) (Passenger, error)
	CreateStation(ctx context.Context, arg CreateStationParams) (Station, error)
	CreateTrip(ctx context.Context, arg CreateTripParams) (Trip, error)
	CreateTripApplication(ctx context.Context, arg CreateTripApplicationParams) (TripApplication, error)
	DeletePassenger(ctx context.Context, id int64) error
	DeleteStation(ctx context.Context, id int64) error
	DeleteTrip(ctx context.Context, id int64) error
	DeleteTripApplication(ctx context.Context, id int64) error
	GetTrip(ctx context.Context, id int64) (Trip, error)
	GetTripApplication(ctx context.Context, id int64) (TripApplication, error)
	ListDriverTrips(ctx context.Context, arg ListDriverTripsParams) ([]Trip, error)
	ListFutureTrips(ctx context.Context, arg ListFutureTripsParams) ([]Trip, error)
	ListPassengers(ctx context.Context, tripID int64) ([]Passenger, error)
	ListStations(ctx context.Context, tripID int64) ([]Station, error)
	ListTripApplications(ctx context.Context, arg ListTripApplicationsParams) ([]TripApplication, error)
	UpdateTrip(ctx context.Context, arg UpdateTripParams) (Trip, error)
}

var _ Querier = (*Queries)(nil)
