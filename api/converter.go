package api

import (
	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertTrip(trip db.Trip) *pb.Trip {
	return &pb.Trip{
		ID:                trip.ID,
		ContactInfo:       trip.ContactInfo,
		DriverId:          trip.DriverID,
		MaxPassenger:      uint32(trip.MaxPassenger),
		Price:             uint32(trip.Price),
		AblePickUp:        trip.AblePickUp,
		ResortId:          uint64(trip.ResortID),
		DepartureAt:       timestamppb.New(trip.DepartureAt),
		ReturnAt:          timestamppb.New(trip.ReturnAt),
		RoundTrip:         trip.RoundTrip,
		AcceptPaymentType: trip.AcceptPaymentType,
		Currency:          trip.Currency,
		CreatedAt:         timestamppb.New(trip.CreatedAt),
	}
}
