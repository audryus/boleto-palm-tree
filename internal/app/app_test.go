package app_test

import (
	"context"
	"testing"

	config "github.com/audryus/boleto-palm-tree/configs"
	"github.com/audryus/boleto-palm-tree/internal/app"
	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestReadAll(t *testing.T) {
	cfg, _ := config.NewConfig()
	app.Run(cfg)

	conn, err := grpc.Dial("localhost:"+cfg.GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCityServiceClient(conn)
	cities, _ := client.ReadAll(context.Background(), &proto.CityRead{})
	require.EqualValues(t, 0, len(cities.Cities))
}
