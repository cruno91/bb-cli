package add

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdAdd represents the add command
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {}
