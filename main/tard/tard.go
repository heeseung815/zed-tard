package main

import (
	"fmt"
	"os"
	"tard/mods"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string
var verbose bool

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if configPath != "" {
			viper.SetConfigFile(configPath)
		}
		viper.SetEnvPrefix("TARD")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				fmt.Fprintln(os.Stderr, "config file not found, apply default settings")
			default:
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	})
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Specifies the config file")
	rootCmd.AddCommand(startCmd, stopCmd)

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

var rootCmd = &cobra.Command{
	Use:     "tard",
	Short:   "tard to arrange tlog",
	Long:    "tard dtag application to arrange tlog",
	Version: mods.VersionDescription(),
}

func verbosePrint(format string, args ...interface{}) {
	if verbose {
		fmt.Printf("[VERBOSE] "+format+"\n", args...)
	}
}
