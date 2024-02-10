package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdGetProject represents the get project command
var CmdGetProject = &cobra.Command{
	Use:   "project",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

func init() {
	CmdGet.AddCommand(CmdGetProject)
}
