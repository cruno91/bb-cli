package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdGetWorkspace represents the get workspace command
var CmdGetWorkspace = &cobra.Command{
	Use:   "workspace",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspace called")
	},
}

func init() {
	CmdGet.AddCommand(CmdGetWorkspace)
}
