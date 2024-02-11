package list

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

// CmdListWorkspaces represents the list workspace command
var CmdListWorkspaces = &cobra.Command{
	Use:   "workspaces",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspaces called")

		bb := auth.Auth()

		listWorkspaces(bb)
	},
}

func init() {
	CmdList.AddCommand(CmdListWorkspaces)
}

func listWorkspaces(bb *bitbucket.Client) {
	workspaces, err := bb.Workspaces.List()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Workspaces:", workspaces)
}
