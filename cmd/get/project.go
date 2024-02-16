package get

import (
	"bb-cli/cmd/auth"
	"bb-cli/cmd/list"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	workspaceSlug string
	projectKey    string
)

// CmdGetProject represents the get project command
var CmdGetProject = &cobra.Command{
	Use:   "project",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		project, err := FetchProject(bb, workspaceSlug, projectKey)
		if err != nil {
			fmt.Println("Error getting project:", err)
		} else if project != nil {
			list.PrintProjectsTable([]bitbucket.Project{*project})
		} else {
			fmt.Println("Project not found.")
		}
	},
}

func init() {
	CmdGet.AddCommand(CmdGetProject)
	CmdGetProject.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdGetProject.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdGetProject.Flags().StringVarP(&projectKey, "project", "p", "", "Bitbucket project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdGetProject.MarkFlagRequired("project"); err != nil {
		fmt.Println(err)
	}
}

func FetchProject(bb *bitbucket.Client, workspace string, project string) (*bitbucket.Project, error) {
	projectOpt := &bitbucket.ProjectOptions{
		Owner: workspace,
		Key:   project,
	}

	res, err := bb.Workspaces.GetProject(projectOpt)

	if err != nil {

		if (err.Error()) == "404 Not Found" {
			return nil, nil
		} else {
			fmt.Println("Error:", err)
			return nil, err
		}
	}

	return res, nil
}
