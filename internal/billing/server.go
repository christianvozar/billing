package billing

import (
	"context"

	pb "github.com/christianvozar/billing/pkg/rpc"
)

var PerUnitCost float64 = 3.50

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) AddIOMetrics(ctx context.Context, dio *pb.IOMetricsReq) (res *pb.CustomerBilledResp, err error) {
	units := dio.ReadBytes + dio.WriteBytes
	c := ToUSD(PerUnitCost)

	return &pb.CustomerBilledResp{
		Name:                dio.Name,
		UnitsBilled:         units,
		PerUnitCost:         PerUnitCost,
		TotalMonthlyCharged: c.Multiply(float64(units)).Float64(),
	}, nil
}
