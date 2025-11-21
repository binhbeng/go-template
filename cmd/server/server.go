package server

import (
	"fmt"
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/routers"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "go run main.go server",
		PreRun: func(cmd *cobra.Command, args []string) {
			data.InitData()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	host     string
	port     int
)

func init() {
	Cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "app host")
	Cmd.Flags().IntVarP(&port, "port", "P", 9001, "app port")
}

func run() error {
	engine := routers.SetRouters()
	err := engine.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	return nil
}
