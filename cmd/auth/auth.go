package auth

import (
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var (
	oauthToken string
)

// CmdAuth represents the auth command
var CmdAuth = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with a Bitbucket account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if oauthToken != "" {
			// Save the token in the viper configuration
			viper.Set("oauth.token", oauthToken)

			// Attempt to save the configuration to file
			if err := viper.WriteConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					// If the config file was not found, try to write it as a new file.
					if err := viper.SafeWriteConfig(); err != nil {
						fmt.Printf("Error saving configuration: %s\n", err)
					}
				} else {
					fmt.Printf("Error updating configuration: %s\n", err)
				}
			} else {
				fmt.Println("Authentication token updated successfully.")
			}
		}
		fmt.Println("auth called with token:", oauthToken)
	},
}

func init() {
	CmdAuth.Flags().StringVarP(&oauthToken, "token", "t", "", "Bitbucket token (ex: 1234abc:0987654321zyxwvutsrqponmlkjihgfedcba)")
	if err := CmdAuth.MarkFlagRequired("token"); err != nil {
		fmt.Println(err)
	}
}

func auth(token string) (c *bitbucket.Client) {
	bitbucketClientId, bitbucketClientSecret, tokenError := splitToken(token)

	if tokenError != nil {
		fmt.Println(tokenError)
		return
	}

	return bitbucket.NewOAuthClientCredentials(bitbucketClientId, bitbucketClientSecret)
}

func splitToken(input string) (string, string, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid token format")
	}
	return parts[0], parts[1], nil
}
