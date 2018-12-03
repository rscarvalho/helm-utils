package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "helm-utils",
	Short: "tests and deploys a helm chart",
	Long: `This application is used to perform routing helm commands to test/preflight and install a Helm Chart.
More info at https://www.helm.sh/`,
}

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"Path to the configuration file. Defaults to $HOME/.config/helm-utils")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, ".config")

		if stat, err := os.Stat(configPath); os.IsNotExist(err) || !stat.IsDir() {
			if err := os.Mkdir(configPath, 0755); err != nil {
				panic(err)
			}
		}

		// Search config in home directory with name "helm-utils" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigName("helm-utils")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
