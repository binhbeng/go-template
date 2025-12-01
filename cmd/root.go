package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/binhbeng/goex/cmd/crawl"
	"github.com/binhbeng/goex/cmd/cron"
	"github.com/binhbeng/goex/cmd/grpc"
	"github.com/binhbeng/goex/cmd/http"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.AddCommand(http_server.Cmd)
	rootCmd.AddCommand(crawl.Cmd)
	rootCmd.AddCommand(cron.Cmd)
	rootCmd.AddCommand(grpc_server.Cmd)

	RunWire()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func RunWire() error {
	log.Println("⏳ ...Running wire to generate wire_gen.go...")

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	wireDir := filepath.Join(rootDir, "wire")

	if _, err := os.Stat(wireDir); os.IsNotExist(err) {
		return fmt.Errorf("❌ folder %s does not exist", wireDir)
	}

	wireGo := filepath.Join(wireDir, "wire.go")
	if _, err := os.Stat(wireGo); os.IsNotExist(err) {
		return fmt.Errorf("❌ wire.go not found in: %s", wireDir)
	}

	cmd := exec.Command("wire")
	cmd.Dir = wireDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("❌ wire generate failed: %w", err)
	}

	log.Println("✅ Wire_go generated successfully")
	return nil
}