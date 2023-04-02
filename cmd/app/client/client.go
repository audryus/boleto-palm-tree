package main

import (
	"context"
	"fmt"

	config "github.com/audryus/boleto-palm-tree/configs"
	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg, _ := config.NewConfig()
	conn, err := grpc.Dial("localhost:"+cfg.GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCityServiceClient(conn)
	ctx := context.Background()
	read(client, ctx)
	c, _ := client.Save(ctx, &proto.City{
		Name: "Lero",
	})
	read(client, ctx)
	c.Name = "New name"
	client.Save(ctx, c)
	read(client, ctx)
}
func read(client proto.CityServiceClient, ctx context.Context) {
	fmt.Printf("\n >>>> \n")
	cities, _ := client.ReadAll(ctx, &proto.CityRead{})
	for _, city := range cities.Cities {
		fmt.Printf("%s : %s\n", city.Id, city.Name)
	}
}
