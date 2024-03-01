package add

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdAdd represents the add command
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Add an access key, webhook, or pipeline variable to a bitbucket repository.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Use the add command with access-key, variable, or webhook.")
		}
	},
}

func init() {}
