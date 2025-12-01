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
		Example: "go-layout crawl -c config.yml",
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
	Cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "监听服务器地址")
	Cmd.Flags().IntVarP(&port, "port", "P", 9001, "监听服务器端口")
	Cmd.Flags().BoolVarP(&setRoute, "set-route", "R", false, "设置数据库数据库API路由表")
}

func run() {
	fmt.Println("start crawl")
	select {}
}
