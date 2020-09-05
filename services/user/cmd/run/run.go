package run

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"liemlhd.com/golang/monorepo-example/pkg/log"
	"time"

	"github.com/spf13/cobra"
)

func run(quit context.Context, done chan error) {
	// sample use of pkg
	log.WithFields(map[string]interface{}{
		"key": "value",
	}).Errorf("test error")

	for {
		select {
		case <-quit.Done():
			done <- nil
			return
		}
	}
}

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run user",
		Long:  "Run user as a long-running service.",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func(begin time.Time) {
				log.Infof("stop user after %v", time.Since(begin))
			}(time.Now())

			log.Infof("start user on %v", time.Now())

			quit, cancel := context.WithCancel(context.TODO())
			done := make(chan error)
			go run(quit, done)

			go func() {
				sigch := make(chan os.Signal)
				signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
				log.Infof("%+v", <-sigch)
				cancel()
			}()
			return <-done
		},
	}

	return cmd

}
