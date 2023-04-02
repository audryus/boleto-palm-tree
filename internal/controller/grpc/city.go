package grpcrouter

import (
	"context"

	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
)

type CityRouter struct {
	uc CityUseCase
	proto.UnimplementedCityServiceServer
}

type CityUseCase interface {
	ReadAll(ctx context.Context) (*proto.Cities, error)
	Save(ctx context.Context, request *proto.City) (*proto.City, error)
}

func NewCityRouter(uc CityUseCase) *CityRouter {
	return &CityRouter{
		uc: uc,
	}
}

func (r *CityRouter) Save(ctx context.Context, request *proto.City) (*proto.City, error) {
	return r.uc.Save(ctx, request)
}

func (r *CityRouter) ReadAll(ctx context.Context, request *proto.CityRead) (*proto.Cities, error) {
	return r.uc.ReadAll(ctx)
}
