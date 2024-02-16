// Package cmd
package cmd

import (
	"bb-cli/cmd/auth"
	"bb-cli/cmd/create"
	"bb-cli/cmd/get"
	"bb-cli/cmd/list"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bb",
	Short: "Create and scaffold Bitbucket projects and repositories.",
	Long:  `Create and scaffold Bitbucket projects and repositories.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(auth.CmdAuth)
	rootCmd.AddCommand(get.CmdGet)
	rootCmd.AddCommand(list.CmdList)
	rootCmd.AddCommand(create.CmdCreate)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bb-cli.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
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

		cfgFilePath := home + "/.bb-cli.json"
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".bb-cli")

		// Check if the configuration file already exists
		if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
			// Create the file with initial configuration
			file, err := os.Create(cfgFilePath)
			if err != nil {
				fmt.Println("Unable to create configuration file:", err)
				os.Exit(1)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					fmt.Println("Unable to close configuration file:", err)
					os.Exit(1)
				}
			}(file)

			initialConfig := map[string]string{
				"oauth": "token",
			}

			err = viper.WriteConfigAs(cfgFilePath)
			if err != nil {
				fmt.Println("Unable to write initial configuration to file:", err)
				return
			}
			viper.Set("oauth", initialConfig["oauth"])
		}
	}

	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err != nil {
			fmt.Println("Unable to write to stderr:", err)
			return
		}
	} else {
		_, err = fmt.Fprintln(os.Stderr, "Unable to read config file:", err)
		if err != nil {
			fmt.Println("Unable to write to stderr:", err)
			return
		}
	}
}
