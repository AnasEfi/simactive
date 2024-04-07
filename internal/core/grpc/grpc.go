package grpc

import (
	"fmt"
	"log"
	"net"
	pb "simactive/api/generated/github.com/fixedNick/SimHelper"
	"simactive/internal/config"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// errors
var (
	ErrInternal = status.Error(codes.Internal, "Internal Server Error")
)

type GRPCServer struct {
	port    int
	timeout time.Duration

	// gRPC services
	server *grpc.Server
}

func NewGRPCServer(cfg *config.Config) *GRPCServer {
	return &GRPCServer{
		port:    cfg.GRPC.Port,
		timeout: cfg.GRPC.Timeout,
	}
}

// MustRun runs the GRPCServer.
//
// It takes a SimService, a ServiceService, a ProviderService, and a UsedService as arguments.
// It truly panics if the gRPC server fails to start.
func (s *GRPCServer) MustRun(sim SimService, ss ServiceService, ps ProviderService, us UsedService) {
	addr := fmt.Sprintf("127.0.0.1:%d", s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}

	gs := grpc.NewServer()
	s.server = gs

	pb.RegisterSimServer(gs, NewGRPCSimService(sim, s.timeout))
	pb.RegisterServiceServer(gs, NewGRPCServiceService(ss, s.timeout))
	pb.RegisterProviderServer(gs, NewGRPCProviderService(ps, s.timeout))
	pb.RegisterUsedServer(gs, NewGRPCUsedService(us, s.timeout))

	log.Println("Starting gRPC server...")
	if err = gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC requests: %v", err)
	}
}

func (gs *GRPCServer) Stop() {
	gs.server.GracefulStop()
}
