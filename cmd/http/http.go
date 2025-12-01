package http_server

import (
	"fmt"
	"net"

	"github.com/binhbeng/goex/internal/proto"
	"github.com/binhbeng/goex/internal/router"
	"github.com/binhbeng/goex/wire"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Cmd = &cobra.Command{
		Use:     "server",
		Short:   "Start HTTP server",
		Example: "go run main.go server",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	host string
	port int
)

func init() {
	Cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "app host")
	Cmd.Flags().IntVarP(&port, "port", "P", 9001, "app port")
}

func run() error {
	errCh := make(chan error, 2)

	// gRPC server
	go func() {
		if err := runGrpc(); err != nil {
			errCh <- fmt.Errorf("gRPC error: %v", err)
		}
	}()

	// HTTP server
	go func() {
		if err := runHttp(); err != nil {
			errCh <- fmt.Errorf("HTTP error: %v", err)
		}
	}()

	return <-errCh
}

func runGrpc() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	w, err := wire.NewWireGrpc()
	if err != nil {
		return fmt.Errorf("failed to init gRPC deps: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, w.OrderService)
	reflection.Register(s)

	fmt.Println("gRPC server listening on :50051")
	return s.Serve(lis) 
}

func runHttp() error {
	deps, err := wire.NewWire()
	if err != nil {
		return fmt.Errorf("failed to init HTTP deps: %v", err)
	}

	engine := router.SetRouters(deps)
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Println("HTTP server listening on", addr)
	return engine.Run(addr) 
}

