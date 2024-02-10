package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CmdGet represents the get command
var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {}
