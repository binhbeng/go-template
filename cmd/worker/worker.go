package worker

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/binhbeng/goex/config"
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/pkg/kafka"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "worker",
		Short:   "Start worker",
		Example: "go run main.go worker",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {

}

type Worker struct {
	consumer kafka.KafkaConsumer
}

func NewWorker() *Worker {
	return &Worker{
		consumer: kafka.NewKafkaConsumer(
			[]string{"localhost:9092"},
			"test",
			"test",
			nil,
		),
	}
}

func (w *Worker) Start(ctx context.Context) error {
	return nil
}

func (w *Worker) Shutdown(ctx context.Context) error {
	if err := w.consumer.Close(); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Kafka connection closed successfully")

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Shutdown timeout exceeded")
			return ctx.Err()
		}
	default:
	}

	fmt.Println("Worker shutdown completed")

	return nil
}

func run() {
	config.Load()
	data.InitData()
	worker := NewWorker()

	if worker == nil {
		fmt.Println("worker is nil")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	var wg sync.WaitGroup

	wg.Go(func() {
		if err := worker.Start(ctx); err != nil && err != context.Canceled {
			fmt.Println("start worker failed", err)
		}
	})

	<-ctx.Done()
	fmt.Println("receive signal to stop worker")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := worker.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Shutdown failed")
	}

	fmt.Println("Main process terminated")
	wg.Wait()
}
