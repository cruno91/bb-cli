package list

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdList represents the list command
var CmdList = &cobra.Command{
	Use:   "list",
	Short: "List something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {}
