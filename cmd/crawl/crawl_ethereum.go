package crawl

import (
	"fmt"

	"github.com/binhbeng/goex/data"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "crawl",
		Short:   "Start crawler",
		Example: "go run main.go crawl -c config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			data.InitData()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	host     string
	port     int
	setRoute bool
)

func init() {
	Cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "app host")
	Cmd.Flags().IntVarP(&port, "port", "P", 9001, "app port")
	Cmd.Flags().BoolVarP(&setRoute, "set-route", "R", false, "set route")
}

func run() {
	fmt.Println("start crawl")
	select {}
}
