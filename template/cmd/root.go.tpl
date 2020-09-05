package cmd

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"os"
	"tele.live/server/backbone/services/__SVC_NAME__/cmd/run"
)

var (
	rootCmd = &cobra.Command{
		Use:   "__SVC_NAME__",
		Short: "__SVC_NAME__ for reference",
		Long:  "A __SVC_NAME__ for reference in golang-monorepo.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			goflag.Parse()
		},
	}
)

func init() {
	rootCmd.AddCommand(run.RunCmd())
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
