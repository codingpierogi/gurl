package cmd

import (
	"fmt"
	"os"

	"github.com/codingpierogi/gurl/pkg/fetch"
	"github.com/spf13/cobra"
)

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
		body, err := fetch.Body(args[0], options)

		if err != nil {
			fmt.Printf("Error making request")
		}

		fmt.Print(body)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Make the operation more talkative")
	rootCmd.PersistentFlags().StringVarP(&UserAgent, "user-agent", "A", "curl/7.68.0", "Send User-Agent <name> to server")
}
