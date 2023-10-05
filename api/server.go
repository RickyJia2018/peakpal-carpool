package api

import (
	db "github.com/RickyJia2018/peakpal-carpool/db/sqlc"
	"github.com/RickyJia2018/peakpal-carpool/pb"
	"github.com/RickyJia2018/peakpal-carpool/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedCarpoolServer
	config        util.Config
	store         db.Store
	peakPalClient pb.PeekPalClient
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, peakPalClient pb.PeekPalClient) (*Server, error) {
	server := &Server{
		config:        config,
		store:         store,
		peakPalClient: peakPalClient,
	}

	return server, nil
}
