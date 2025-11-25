package server

import (
	"fmt"

	"github.com/binhbeng/goex/internal/router"
	"github.com/binhbeng/goex/wire"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
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
	deps, err := wire.NewWire()
	if err != nil {
		return err
	}

	engine := router.SetRouters(deps)
	err = engine.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}

	return nil
}
