/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tifye/wim/cmd/download"
	"github.com/tifye/wim/cmd/info"
	"github.com/tifye/wim/cmd/list"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wim",
	Short: "CLI tool for winmower operations.",
	Long:  `Wim is a CLI tool for winmower operations.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(download.DownloadCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(list.ListCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.winmower-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".winmower-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wim")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	initTifAppConfig()
}

func initTifAppConfig() {
	tViper := viper.New()

	roaming, err := os.UserConfigDir()
	cobra.CheckErr(err)

	roaming = roaming + "/Tif App"

	tViper.AddConfigPath(roaming)
	tViper.SetConfigType("json")
	tViper.SetConfigName("Settings")
	if err := tViper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.Set("ApiKey", "")
	viper.Set("AccessToken", "")
	viper.Set("x-api-key", "fruit-pie")
	// encryptedProfile := tViper.GetString("Profile")
	// const key = "{7f8d534a-bf20-4e69-bbf8-54f4a9378f23}"
}
