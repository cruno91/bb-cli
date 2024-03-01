package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdCreate represents the create command
var CmdCreate = &cobra.Command{
	Use:   "create",
	Short: "Get a project or repository in a Bitbucket workspace.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Use the create command with project or repository.")
		}
	},
}

func init() {}
