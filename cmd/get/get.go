package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdGet represents the get command
var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get a project from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Use the get command with project.")
		}
	},
}

func init() {}
