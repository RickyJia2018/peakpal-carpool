package api

import (
	"context"
	"math"
	"time"

	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/validators"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	violations := validateCreateTripRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	// Driver(User) must be same as the trip creator
	if authPayload.UserId != req.GetDriverId() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Driver ID")
	}
	//validate if the resort exists
	if _, err := server.peakPalClient.GetResort(ctx, &pb.GetResortRequest{ID: int64(req.GetResortId())}); err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Resort with Resort ID: %d", req.GetResortId())
	}

	trip, err := server.store.CreateTrip(ctx, db.CreateTripParams{
		ContactInfo:       req.GetContactInfo(),
		DriverID:          req.GetDriverId(),
		MaxPassenger:      int32(req.GetMaxPassenger()),
		Price:             int32(req.GetPrice()),
		AblePickUp:        req.GetAblePickUp(),
		ResortID:          int64(req.GetResortId()),
		DepartureAt:       req.GetDepartureAt().AsTime(),
		ReturnAt:          req.GetReturnAt().AsTime(),
		RoundTrip:         req.GetRoundTrip(),
		AcceptPaymentType: req.GetAcceptPaymentType(),
		Currency:          req.GetCurrency(),
	})
	rsp := &pb.CreateTripResponse{
		Trip: convertTrip(trip),
	}
	return rsp, nil
}

func (server *Server) GetTrip(ctx context.Context, req *pb.GetTripRequest) (*pb.GetTripResponse, error) {
	violations := validateGetTripRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trip, err := server.store.GetTrip(ctx, int64(req.GetID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Trip with ID: %d", req.GetID())
	}
	rsp := &pb.GetTripResponse{
		Trip: convertTrip(trip),
	}
	return rsp, nil
}

func (server *Server) UpdateTrip(ctx context.Context, req *pb.UpdateTripRequest) (*pb.GetTripResponse, error) {
	violations := validateUpdateTripRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trip, err := server.store.GetTrip(ctx, int64(req.GetID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Trip with ID: %d", req.GetID())
	}
	if authPayload.UserId != trip.DriverID {
		return nil, status.Errorf(codes.PermissionDenied, "No permission to modify this trip")
	}
	newTrip, err := server.store.UpdateTrip(ctx, db.UpdateTripParams{
		ID: int64(req.GetID()),
		ContactInfo: pgtype.Text{
			String: req.GetContactInfo(),
			Valid:  req.ContactInfo != nil,
		},
		MaxPassenger: pgtype.Int4{
			Int32: req.GetMaxPassenger(),
			Valid: req.MaxPassenger != nil,
		},
		AblePickUp: pgtype.Bool{
			Bool:  req.GetAblePickUp(),
			Valid: req.AblePickUp != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to update Trip: %s", err.Error())
	}
	rsp := &pb.GetTripResponse{
		Trip: convertTrip(newTrip),
	}
	return rsp, nil
}

func (server *Server) ListFutureTrips(ctx context.Context, req *pb.ListFutureTripsRequest) (*pb.ListTripsResponse, error) {
	violations := validateListFutureTripsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trips, err := server.store.ListFutureTrips(ctx, db.ListFutureTripsParams{
		ResortID: int64(req.GetResortId()),
		Limit:    int32(req.GetLimit()),
		Offset:   int32(req.GetOffset()),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Trips with Resort: %s", err.Error())
	}
	totalTrips, err := server.store.CountFutureTrips(ctx, int64(req.GetResortId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to count Trips")
	}
	rsp := &pb.ListTripsResponse{
		Trips:  convertTrips(trips),
		Total:  totalTrips,
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	return rsp, nil
}

func (server *Server) ListDriverTrips(ctx context.Context, req *pb.ListDriverTripsRequest) (*pb.ListTripsResponse, error) {
	violations := validateListDriverTripsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trips, err := server.store.ListDriverTrips(ctx, db.ListDriverTripsParams{
		DriverID: int64(req.GetDriverId()),
		Limit:    int32(req.GetLimit()),
		Offset:   int32(req.GetOffset()),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Trips with Driver: %s", err.Error())
	}
	totalTrips, err := server.store.CountDriverTrips(ctx, int64(req.GetDriverId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to count Trips")
	}
	rsp := &pb.ListTripsResponse{
		Trips:  convertTrips(trips),
		Total:  totalTrips,
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	return rsp, nil
}

// Should I delete all trip related data? such as passenger, stations, etc.
// Currently, I think user can only delete trip with no passenger.
func (server *Server) DeleteTrip(ctx context.Context, req *pb.DeleteTripRequest) (*pb.DeleteTripResponse, error) {
	violations := validateDeleteTripRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	trip, err := server.store.GetTrip(ctx, int64(req.GetID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Trip with ID: %d", req.GetID())
	}
	if authPayload.UserId != trip.DriverID {
		return nil, status.Errorf(codes.PermissionDenied, "No permission to delete this trip")
	}

	//TODO: change to trasactions delete realted stations, passengers and send email notification to each passenger
	stations, err := server.store.ListStations(ctx, int64(req.GetID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to find Stations with Trip: %s", err.Error())
	}
	for _, station := range stations {
		if err = server.store.DeleteStation(ctx, station.ID); err != nil {
			return nil, status.Errorf(codes.Internal, "Fail to delete Station: %s", err.Error())
		}
	}
	// Then delete the trip
	if err = server.store.DeleteTrip(ctx, int64(req.GetID())); err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to delete Trip: %s", err.Error())
	}

	return &pb.DeleteTripResponse{
		Success: true,
	}, nil
}

func validateCreateTripRequest(req *pb.CreateTripRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateString(req.GetContactInfo(), 1, 1024); err != nil {
		violations = append(violations, fieldViolation("contact_info", err))
	}
	if err := validators.ValidateID(req.GetDriverId()); err != nil {
		violations = append(violations, fieldViolation("driver_id", err))
	}
	if err := validators.ValidateNumber(int64(req.GetMaxPassenger()), 1, 20); err != nil {
		violations = append(violations, fieldViolation("max_passenger", err))
	}
	if err := validators.ValidateNumber(int64(req.GetPrice()), 1, math.MaxInt64); err != nil {
		violations = append(violations, fieldViolation("price", err))
	}
	if err := validators.ValidateBool(req.GetAblePickUp()); err != nil {
		violations = append(violations, fieldViolation("able_pickup", err))
	}
	if err := validators.ValidateID(int64(req.GetResortId())); err != nil {
		violations = append(violations, fieldViolation("resort_id", err))
	}
	if err := validators.ValidateTime(req.GetDepartureAt().AsTime(), time.Now(), time.Now().AddDate(0, 2, 0)); err != nil {
		violations = append(violations, fieldViolation("departure_at", err))
	}
	if err := validators.ValidateTime(req.GetReturnAt().AsTime(), time.Now(), time.Now().AddDate(0, 2, 0)); err != nil {
		violations = append(violations, fieldViolation("return_at", err))
	}
	if err := validators.ValidateBool(req.GetRoundTrip()); err != nil {
		violations = append(violations, fieldViolation("round_trip", err))
	}
	if err := validators.ValidateString(req.GetAcceptPaymentType(), 1, 1024); err != nil {
		violations = append(violations, fieldViolation("accept_payment_type", err))
	}
	if err := validators.ValidateString(req.GetCurrency(), 1, 10); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	return violations
}
func validateGetTripRequest(req *pb.GetTripRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetID()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
func validateListFutureTripsRequest(req *pb.ListFutureTripsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetResortId()); err != nil {
		violations = append(violations, fieldViolation("resort_id", err))
	}
	return violations
}
func validateListDriverTripsRequest(req *pb.ListDriverTripsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetDriverId()); err != nil {
		violations = append(violations, fieldViolation("driver_id", err))
	}
	return violations
}
func validateUpdateTripRequest(req *pb.UpdateTripRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetID()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	if req.ContactInfo != nil {
		if err := validators.ValidateString(req.GetContactInfo(), 1, 1024); err != nil {
			violations = append(violations, fieldViolation("contact_info", err))
		}
	}
	if req.MaxPassenger != nil {
		if err := validators.ValidateNumber(int64(req.GetMaxPassenger()), 1, 20); err != nil {
			violations = append(violations, fieldViolation("max_passenger", err))
		}
	}
	if req.AblePickUp != nil {
		if err := validators.ValidateBool(req.GetAblePickUp()); err != nil {
			violations = append(violations, fieldViolation("able_pickup", err))
		}
	}
	return violations
}

func validateDeleteTripRequest(req *pb.DeleteTripRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetID()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
