package cmd

import (
	"os"

	"github.com/codingpierogi/gurl/pkg/curl"
	"github.com/codingpierogi/gurl/pkg/fetch"
	"github.com/spf13/cobra"
)

var Output []string
var Verbose bool
var UserAgent string

var rootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "Golang curl",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := fetch.Options{
			Verbose:   Verbose,
			UserAgent: UserAgent,
		}
		curl.Run(args, Output, options)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&Output, "output", "o", []string{""}, "Write to file instead of stdout")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Make the operation more talkative")
	rootCmd.PersistentFlags().StringVarP(&UserAgent, "user-agent", "A", "curl/7.68.0", "Send User-Agent <name> to server")
}
