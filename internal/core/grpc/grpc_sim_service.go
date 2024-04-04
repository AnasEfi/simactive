package grpc

import (
	"context"
	"fmt"
	pb "simactive/api/generated/github.com/fixedNick/SimHelper"
	"simactive/internal/core"
	"time"
)

type SimService interface {
	Add(ctx context.Context, s core.Sim) (err error)
	GetByID(ctx context.Context, id int) (sim core.Sim, err error)
	Remove(ctx context.Context, id int) (err error)
}

type GRPCSimService struct {
	pb.UnimplementedSimServer

	timeout    time.Duration
	simService SimService
}

func NewGRPCSimService(ss SimService, timeout time.Duration) GRPCSimService {
	return GRPCSimService{
		simService: ss,
		timeout:    timeout,
	}
}

func (gs GRPCSimService) AddSim(ctx context.Context, req *pb.AddSimRequest) (*pb.AddSimResponse, error) {
	// validate
	ctx, cancel := context.WithTimeout(ctx, gs.timeout)
	defer cancel()

	sim := core.NewSim(0, req.SimData.Number, int(req.SimData.ProviderID), req.SimData.IsActivated, req.SimData.ActivateUntil, req.SimData.IsBlocked)
	if err := gs.simService.Add(ctx, sim); err != nil {
		return nil, ErrInternal
	}
	return &pb.AddSimResponse{
		IsAdded: true,
		Message: fmt.Sprintf("sim card with number %s added", sim.Number()),
	}, nil
}

func (gs GRPCSimService) DeleteSim(ctx context.Context, req *pb.DeleteSimRequest) (*pb.DeleteSimResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, gs.timeout)
	defer cancel()
	if err := gs.simService.Remove(ctx, int(req.GetId())); err != nil {
		return nil, ErrInternal
	}
	return &pb.DeleteSimResponse{
		IsDeleted: true,
	}, nil
}
