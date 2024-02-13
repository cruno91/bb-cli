package get

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	workspaceName string
	projectName   string
)

// CmdGetProject represents the get project command
var CmdGetProject = &cobra.Command{
	Use:   "project",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")

		bb := auth.Auth()

		getProject(bb, workspaceName, projectName)
	},
}

func init() {
	CmdGet.AddCommand(CmdGetProject)
	CmdGetProject.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Bitbucket workspace")
	if err := CmdGetProject.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdGetProject.Flags().StringVarP(&projectName, "project", "p", "", "Bitbucket project")
	if err := CmdGetProject.MarkFlagRequired("project"); err != nil {
		fmt.Println(err)
	}
}

func getProject(bb *bitbucket.Client, workspace string, project string) {
	projectOpt := &bitbucket.ProjectOptions{
		Owner: workspace,
		Key:   project,
	}

	res, err := bb.Workspaces.GetProject(projectOpt)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(res)
}
