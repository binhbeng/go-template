package cron

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "cron",
		Short:   "Starting a scheduled task",
		Example: "go run main.go cron",
		PreRun: func(cmd *cobra.Command, args []string) {
			// data.InitData()
		},
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
)

func Start() {
	myLog := myLogger{}
	crontab := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(myLog)))
	job := cron.NewChain(cron.SkipIfStillRunning(myLog), cron.Recover(myLog)).Then(cron.FuncJob(func() {
		log.Printf("%s:%s\n", time.Now().Format("2006-01-02 15:04:05"), " cron demo")
	}))
	_, err := crontab.AddJob("@every 5s", job)

	if err != nil {
		panic("Error adding job:" + err.Error())
	}
	crontab.Start()
	select {}
}

type myLogger struct {
}

func (ml myLogger) Info(msg string, keysAndValues ...any) {
	log.Printf(msg, keysAndValues...)
}

func (ml myLogger) Error(err error, msg string, keysAndValues ...any) {
	log.Printf(msg, keysAndValues...)
}
