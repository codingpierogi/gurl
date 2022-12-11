package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gurl",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(args[0])
		if err != nil {
			fmt.Printf("Error making request to ")
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		fmt.Print(string(body))
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
