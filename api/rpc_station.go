package api

import (
	"context"
	"errors"

	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateStation(ctx context.Context, req *pb.CreateStationRequest) (*pb.CreateStationResponse, error) {
	violations := validateCreateStationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trip, err := server.store.GetTrip(ctx, int64(req.GetTripeId()))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "Trip not found")
		}
		return nil, status.Errorf(codes.Internal, "Fail to find Trip with ID: %d", req.GetTripeId())
	}
	if trip.DriverID != authPayload.UserId {
		return nil, status.Errorf(codes.Unauthenticated, "You are not the driver of this trip")
	}
	station, err := server.store.CreateStation(ctx, db.CreateStationParams{
		TripID:         req.GetTripeId(),
		StationName:    req.GetName(),
		ArrivalMinutes: int32(req.GetArrivalMinutes()),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to create Station: %s", err.Error())
	}
	rsp := &pb.CreateStationResponse{
		Station: convertStation(station),
	}
	return rsp, nil
}

func (server *Server) ListStations(ctx context.Context, req *pb.ListStationsRequest) (*pb.ListStationsResponse, error) {
	violations := validateListStationsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	stations, err := server.store.ListStations(ctx, int64(req.GetTripId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Stations with ID: %d", req.GetTripId())
	}
	rsp := &pb.ListStationsResponse{
		Stations: convertStations(stations),
	}
	return rsp, nil
}

func (server *Server) DeleteStations(ctx context.Context, req *pb.DeleteStationRequest) (*pb.DeleteStationResponse, error) {
	violations := validateDeleteStationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	stationData, err := server.store.GetStation(ctx, int64(req.GetID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Station with ID: %d", req.GetID())
	}

	if stationData.DriverID.Int64 != authPayload.UserId {
		if db.ErrRecordNotFound == err {
			return nil, status.Errorf(codes.Internal, "Data not found %s", err.Error())
		}
		return nil, status.Errorf(codes.PermissionDenied, "No permission to delete this station")
	}
	if err := server.store.DeleteStation(ctx, int64(req.GetID())); err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to delete Station with ID: %d", req.GetID())
	}
	rsp := &pb.DeleteStationResponse{
		Success: true,
	}
	return rsp, nil
}

func validateCreateStationRequest(req *pb.CreateStationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetTripeId()); err != nil {
		violations = append(violations, fieldViolation("trip_id", err))
	}
	return violations
}
func validateListStationsRequest(req *pb.ListStationsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetTripId()); err != nil {
		violations = append(violations, fieldViolation("trip_id", err))
	}
	return violations
}
func validateDeleteStationRequest(req *pb.DeleteStationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetID()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
