package cmd

import (
	"fmt"
	"os"

	fetch "github.com/codingpierogi/gurl/pkg"
	"github.com/spf13/cobra"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "Golang curl",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := fetch.Body(args[0])

		if err != nil {
			fmt.Printf("Error making request")
		}

		//fmt.Print(body)
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
}
