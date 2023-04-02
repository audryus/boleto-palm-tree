package grpcserver

import (
	"fmt"
	"net"
	"time"

	grpcrouter "github.com/audryus/boleto-palm-tree/internal/controller/grpc"
	"github.com/audryus/boleto-palm-tree/pkg/grpcserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(port string, router *grpcrouter.CityRouter) *grpc.Server {
	fmt.Println("Starting server ...")

	start := time.Now()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterCityServiceServer(srv, router)
	reflection.Register(srv)

	go srv.Serve(lis)

	fmt.Printf("Server started in %d ms\n", time.Since(start).Milliseconds())

	return srv
}
