package api

import (
	"context"

	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/validators"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTripApplication(ctx context.Context, req *pb.CreateTripApplicationRequest) (*pb.CreateTripApplicationResponse, error) {
	violations := validateCreateTripApplicationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	if authPayload.UserId != req.GetPassengerId() {
		return nil, status.Errorf(codes.PermissionDenied, "No permission to apply this trip")
	}

	tripApplication, err := server.store.CreateTripApplication(ctx, db.CreateTripApplicationParams{
		PassengerID:     req.GetPassengerId(),
		TripID:          req.GetTripId(),
		BoardingStation: req.GetBoardingStation(),
		PaymentType:     req.GetPaymentType(),
		Currency:        req.GetCurrency(),
		TotalPassenger:  int32(req.GetTotalPassenger()),
		ContactInfo:     req.GetContactInfo(),
	})
	if err != nil {
		if db.ForeignKeyViolation == db.ErrorCode(err) {
			return nil, status.Errorf(codes.FailedPrecondition, "Foreign key violation: %s", err.Error())
		}
		return nil, status.Errorf(codes.Internal, "Fail to create Trip Application: %s", err.Error())
	}
	rsp := &pb.CreateTripApplicationResponse{
		TripApplication: convertTripApplication(tripApplication),
	}

	return rsp, nil
}
func (server *Server) ListTripApplications(ctx context.Context, req *pb.ListTripApplicationsRequest) (*pb.ListTripApplicationsResponse, error) {
	violations := validateListTripApplicationsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	_, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	arg := db.ListTripApplicationsParams{
		Limit:  1,
		Offset: 1,
	}
	if req.PassengerId != nil {
		arg.PassengerID = pgtype.Int8{
			Int64: int64(*req.PassengerId),
			Valid: req.PassengerId != nil,
		}
	}
	if req.TripId != nil {
		arg.TripID = pgtype.Int8{
			Int64: int64(*req.TripId),
			Valid: req.TripId != nil,
		}
	}
	tripApplications, err := server.store.ListTripApplications(ctx, arg)
	rsp := &pb.ListTripApplicationsResponse{
		TripApplications: convertTripApplications(tripApplications),
	}
	return rsp, nil
}
func (server *Server) DeleteTripApplications(ctx context.Context, req *pb.DeleteTripApplicationRequest) (*pb.DeleteTripApplicationResponse, error) {
	violations := validateDeleteTripApplicationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}
	tripApplication, err := server.store.GetTripApplication(ctx, req.GetID())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot find Trip Application")
	}
	if authPayload.UserId != tripApplication.PassengerID && authPayload.UserId != tripApplication.DriverID.Int64 {
		return nil, status.Errorf(codes.PermissionDenied, "No permission to delete this trip application")
	}
	err = server.store.DeleteTripApplication(ctx, req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to delete Trip Application: %s", err.Error())
	}
	rsp := &pb.DeleteTripApplicationResponse{
		Success: true,
	}
	return rsp, nil
}

func validateCreateTripApplicationRequest(req *pb.CreateTripApplicationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetTripId()); err != nil {
		violations = append(violations, fieldViolation("trip_id", err))
	}
	if err := validators.ValidateID(req.GetPassengerId()); err != nil {
		violations = append(violations, fieldViolation("passenger_id", err))
	}
	if err := validators.ValidateID(req.GetBoardingStation()); err != nil {
		violations = append(violations, fieldViolation("boarding_station", err))
	}
	if err := validators.ValidateString(req.GetPaymentType(), 3, 10); err != nil {
		violations = append(violations, fieldViolation("payment_type", err))
	}
	if err := validators.ValidateString(req.GetCurrency(), 1, 10); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}
	if err := validators.ValidateNumber(req.GetTotalPassenger(), 1, 20); err != nil {
		violations = append(violations, fieldViolation("total_passenger", err))
	}
	if err := validators.ValidateString(req.GetContactInfo(), 3, 1024); err != nil {
		violations = append(violations, fieldViolation("contact_info", err))
	}
	return violations
}

func validateListTripApplicationsRequest(req *pb.ListTripApplicationsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.PassengerId != nil {
		if err := validators.ValidateID(req.GetPassengerId()); err != nil {
			violations = append(violations, fieldViolation("passenger_id", err))
		}

	}
	if req.TripId != nil {
		if err := validators.ValidateID(req.GetTripId()); err != nil {
			violations = append(violations, fieldViolation("trip_id", err))
		}

	}
	return violations
}

func validateDeleteTripApplicationRequest(req *pb.DeleteTripApplicationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateID(req.GetID()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
