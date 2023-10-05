package api

import (
	"context"
	"fmt"

	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	outCtx := metadata.NewOutgoingContext(ctx, md)

	resort, err := server.peakPalClient.GetResort(outCtx, &pb.GetResortRequest{ID: int64(req.GetResortId())})

	if err != nil {

	}
	fmt.Println("GetResort resort: ", resort.Resort.Name)
	accessToken, err := util.GetToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid access token type")
	}
	authPayload, err := server.peakPalClient.AuthorizeUser(ctx, &pb.AuthorizeUserRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid access token")
	}

	if authPayload.UserId != req.GetDriverId() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Driver ID")
	}
	//TODO: validate req params
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
