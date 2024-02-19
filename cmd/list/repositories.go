package list

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	projectKey string
)

// CmdListRpeositories represents the list workspace command
var CmdListRpeositories = &cobra.Command{
	Use:   "repositories",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		listRepositories(bb, workspaceSlug, projectKey)
	},
}

func init() {
	CmdList.AddCommand(CmdListRpeositories)
	CmdListRpeositories.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdListRpeositories.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdListRpeositories.Flags().StringVarP(&projectKey, "project", "w", "", "Bitbucket project (Example: \"PROJ\"")
	if err := CmdListRpeositories.MarkFlagRequired("project"); err != nil {
		fmt.Println(err)
	}
}

func listRepositories(bb *bitbucket.Client, workspace string, project string) {}
