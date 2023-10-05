package api

import (
	"context"

	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	violations := validatieCreateTripRequestRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	// Driver must be same as the trip creator
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

func validatieCreateTripRequestRequest(req *pb.CreateTripRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidateString(req.GetContactInfo(), 1, 1024); err != nil {
		violations = append(violations, fieldViolation("contact_info", err))
	}
	// if err := validators.ValidatePassword(req.GetPassword()); err != nil {
	// 	violations = append(violations, fieldViolation("password", err))
	// }
	return violations
}
