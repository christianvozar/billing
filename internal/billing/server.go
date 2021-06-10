package billing

import (
	"context"

	pb "github.com/christianvozar/billing/pkg/rpc"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) AddIOMetrics(ctx context.Context, dio *pb.IOMetricsReq) (res *pb.CustomerBilledResp, err error) {
	return &pb.CustomerBilledResp{
		Id:                  dio.Id,
		UnitsBilled:         0,
		PerUnitCost:         0,
		TotalMonthlyCharged: 0,
	}, nil
}
