package cmd

import (
	"os"

	"github.com/codingpierogi/gurl/pkg/curl"
	"github.com/codingpierogi/gurl/pkg/fetch"
	"github.com/spf13/cobra"
)

var output []string
var verbose bool
var userAgent string

var rootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "Golang curl",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := fetch.Options{
			Verbose:   verbose,
			UserAgent: userAgent,
		}
		curl.Run(args, output, options)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&output, "output", "o", []string{""}, "Write to file instead of stdout")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Make the operation more talkative")
	rootCmd.PersistentFlags().StringVarP(&userAgent, "user-agent", "A", "curl/7.68.0", "Send User-Agent <name> to server")
}
