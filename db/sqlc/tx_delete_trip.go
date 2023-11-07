package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type DeleteTripTxParams struct {
	TripID int64
}

func (store *SQLStore) DeleteTripTx(ctx context.Context, arg DeleteTripTxParams) error {

	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		if err := queries.DeleteTrip(ctx, arg.TripID); err != nil {
			return err
		}
		stations, err := queries.ListStations(ctx, arg.TripID)
		if err != nil {
			return err
		}
		for _, station := range stations {
			if err := queries.DeleteStation(ctx, station.ID); err != nil {
				return err
			}
		}
		tripApplications, err := queries.ListTripApplications(ctx, ListTripApplicationsParams{
			TripID: pgtype.Int8{
				Int64: arg.TripID,
				Valid: arg.TripID > 0,
			},
		})
		if err != nil {
			return err
		}
		for _, tripApplication := range tripApplications {
			// TODO: Sent message (Update to MessageDB)/Email to passenger/driver
			if err := queries.DeleteTripApplication(ctx, tripApplication.ID); err != nil {
				return err
			}
		}

		return err
	})
	return err
}
