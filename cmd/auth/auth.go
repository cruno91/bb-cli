package auth

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AuthCmd represents the diskUsage command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with a Bitbucket account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func init() {}
