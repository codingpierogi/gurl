package cmd

import (
	"fmt"
	"os"

	fetch "github.com/codingpierogi/gurl/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gurl",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := fetch.Body(args[0])

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
