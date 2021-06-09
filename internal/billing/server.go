package billing

import (
	"context"

	pb "github.com/christianvozar/billing/pkg/rpc"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) AddDiskIOMetrics(ctx context.Context, dio *pb.DiskIOMetricsReq) (res *pb.CustomerBilledResp, err error) {
	return &pb.CustomerBilledResp{
		Customer:            dio.Customer,
		UnitsBilled:         0,
		PerUnitCost:         0,
		TotalMonthlyCharged: 0,
	}, nil
}

func (s *Server) AddNetIOMetrics(ctx context.Context, nio *pb.NetIOMetricsReq) (hat *pb.CustomerBilledResp, err error) {
	return &pb.CustomerBilledResp{
		Customer:            nio.Customer,
		UnitsBilled:         0,
		PerUnitCost:         0,
		TotalMonthlyCharged: 0,
	}, nil
}
