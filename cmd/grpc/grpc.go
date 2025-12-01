package grpc_server

import (
	"fmt"
	"log"
	"net"

	// "github.com/binhbeng/goex/internal/handler"
	// "github.com/binhbeng/goex/internal/proto"
	"github.com/binhbeng/goex/internal/proto"
	"github.com/binhbeng/goex/wire"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Cmd = &cobra.Command{
		Use:     "grpc",
		Short:   "Start grpc server",
		Example: "go run main.go grpc",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	w, _ := wire.NewWireGrpc()
	s := grpc.NewServer()

	proto.RegisterOrderServiceServer(s, w.OrderService)

	reflection.Register(s)

	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
