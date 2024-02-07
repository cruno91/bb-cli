package auth

import (
	"fmt"
	"strings"

	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
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
		fmt.Println("auth called")
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
