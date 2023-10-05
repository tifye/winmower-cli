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

	viper.Set("ApiKey", "224e94f74e9841c388f255f10ae60c4f")
	viper.Set("AccessToken", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkZXZlbG9wZXIiOnRydWUsImdsb2JhbF9hZG1pbiI6ZmFsc2UsIm5hbWUiOiJKb3NodWEgRGUgTWF0YXMgKGV4dGVybmFsKSIsInJvbGVzIjp7ImF1dG9jaGVjayI6IndyaXRlIiwiYnVuZGxlcyI6InJlYWQiLCJleHQ6YnVuZGxlcyI6InJlYWQiLCJleHQ6Y2ljLmFwcC5sdWNpZGNoYXJ0IjoicmVhZCIsImV4dDpncm91cC5hcHBsLmxlYXJuaW5nIjoicmVhZCIsImV4dDpncm91cC5zeXN0LnZwbi5hY2Nlc3Muc2UuY2l0cml4LmludGVybmFsIjoicmVhZCIsImV4dDpncm91cC5zeXN0LnZwbi5lbXBsb3llLmFsbC51c2VycyI6InJlYWQiLCJleHQ6dGRtIjoicmVhZCIsImZvdGEiOiJyZWFkIiwibW93ZXJjeWNsZSI6InJlYWQiLCJwcm9kdWN0Y2F0YWxvZyI6InJlYWQiLCJwcm94eWluY2xvdWQiOiJyZWFkIiwicmR0IjoicmVhZCIsInJkdHVzZXJtYW5hZ2VtZW50Ijoid3JpdGUiLCJyb2JvdGljcyI6InJlYWQiLCJyb2xsb3V0bWFuYWdlbWVudCI6IndyaXRlIiwic2NyaXB0c3RvcmFnZSI6InJlYWQiLCJ0ZG0iOiJ3cml0ZSJ9LCJwcm9maWxlIjp7ImZpcnN0bmFtZSI6Ikpvc2h1YSIsImxhc3RuYW1lIjoiRGUgTWF0YXMiLCJlbWFpbCI6Ikpvc2h1YS5EZS5NYXRhc0BodXNxdmFybmFncm91cC5jb20iLCJsb2NhdGlvbiI6bnVsbCwiZnVsbG5hbWUiOiJKb3NodWEgRGUgTWF0YXMiLCJodW1hbiI6dHJ1ZX0sImF1dGgiOnsic3NvX2lkIjoiMTZmNzU3MDgtMzMwNi00YjFlLTljODYtZjc2NDdkYWU2OTkzIiwiY2xpZW50X2tleSI6Ikpvc2h1YS5EZS5NYXRhc0BodXNxdmFybmFncm91cC5jb20iLCJjbGllbnRfc2VjcmV0IjoiWCJ9LCJpc3N1ZWQiOjE2OTY0NDkwNjQ2NDYsImV4cGlyZXMiOjE2OTcwNTM4NjQ2NDYsImFwaV9rZXkiOiIyMjRlOTRmNzRlOTg0MWMzODhmMjU1ZjEwYWU2MGM0ZiIsImlkIjoiMDMwYmVkODktZDlhMS00Y2ViLTlmMzAtNmRmNmM0YzU0ZDhjIiwic3lzdGVtX3Rlc3QiOmZhbHNlLCJiZXRhX3Rlc3QiOmZhbHNlLCJpbnRlcm5hbF90ZXN0IjpmYWxzZX0.wGXOZXg2Fc3XcMZXwLXsM04mB8ZVuJQ-OGkjjV_78b8")
	viper.Set("x-api-key", "fruit-pie")
	// encryptedProfile := tViper.GetString("Profile")
	// const key = "{7f8d534a-bf20-4e69-bbf8-54f4a9378f23}"
}
