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
func convertStation(station db.Station) *pb.Station {
	return &pb.Station{
		ID:             station.ID,
		Name:           station.StationName,
		ArrivalMinutes: int64(station.ArrivalMinutes),
		CreatedAt:      timestamppb.New(station.CreatedAt),
	}
}
func convertPassenger(passenger db.Passenger) *pb.Passenger {
	return &pb.Passenger{
		ID:          passenger.ID,
		TripeId:     passenger.TripID,
		PassengerId: passenger.PassengerID,
		CreatedAt:   timestamppb.New(passenger.CreatedAt),
	}
}
func convertTripApplication(tripApplication db.TripApplication) *pb.TripApplication {
	return &pb.TripApplication{
		ID:              tripApplication.ID,
		TripId:          tripApplication.TripID,
		PassengerId:     tripApplication.PassengerID,
		BoardingStation: tripApplication.BoardingStation,
		PaymentType:     tripApplication.PaymentType,
		Currency:        tripApplication.Currency,
		ContactInfo:     tripApplication.ContactInfo,
		Approved:        tripApplication.Approved,
		CreatedAt:       timestamppb.New(tripApplication.CreatedAt),
	}
}
