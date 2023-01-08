package cmd

import (
	"fmt"
	"os"

	"github.com/codingpierogi/gurl/pkg/curl"
	"github.com/codingpierogi/gurl/pkg/fetch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	output    []string
	verbose   bool
	userAgent string
)

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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "K", "", "Read config from a file")
	rootCmd.PersistentFlags().StringSliceVarP(&output, "output", "o", []string{""}, "Write to file instead of stdout")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Make the operation more talkative")
	rootCmd.PersistentFlags().StringVarP(&userAgent, "user-agent", "A", "curl/7.68.0", "Send User-Agent <name> to server")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("user-agent", rootCmd.PersistentFlags().Lookup("user-agent"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".gurlrc")
	}

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if !rootCmd.PersistentFlags().Lookup("verbose").Changed && viper.IsSet("verbose") {
		rootCmd.PersistentFlags().Set("verbose", viper.GetString("verbose"))
	}

	if !rootCmd.PersistentFlags().Lookup("user-agent").Changed && viper.IsSet("user_agent") {
		rootCmd.PersistentFlags().Set("user-agent", viper.GetString("user_agent"))
	}
}
