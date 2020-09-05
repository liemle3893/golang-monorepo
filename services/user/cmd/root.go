package cmd

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"os"
	"liemlhd.com/golang/monorepo-example/services/user/cmd/run"
)

var (
	rootCmd = &cobra.Command{
		Use:   "user",
		Short: "user for reference",
		Long:  "A user for reference in golang-monorepo.",
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
