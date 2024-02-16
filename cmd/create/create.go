package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdCreate represents the create command
var CmdCreate = &cobra.Command{
	Use:   "create",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {}
