package cmd

import (
	"os"

	"github.com/binhbeng/goex/cmd/crawl"
	"github.com/binhbeng/goex/cmd/cron"
	"github.com/binhbeng/goex/cmd/server"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.AddCommand(server.Cmd)
	rootCmd.AddCommand(crawl.Cmd)
	rootCmd.AddCommand(cron.Cmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
