package list

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdList represents the list command
var CmdList = &cobra.Command{
	Use:   "list",
	Short: "List projects, workspaces, or repositories from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Use the list command with projects, repositories, or workspaces.")
		}
	},
}

func init() {}
