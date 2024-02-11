package list

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdListWorkspaces represents the list workspace command
var CmdListWorkspaces = &cobra.Command{
	Use:   "workspaces",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspaces called")
	},
}

func init() {
	CmdList.AddCommand(CmdListWorkspaces)
}
