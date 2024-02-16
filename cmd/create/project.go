package create

import (
	"bb-cli/cmd/auth"
	"bb-cli/cmd/get"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	workspaceName      string
	projectName        string
	projectKey         string
	projectDescription string
	privateFlag        bool
)

// CmdCreateProject represents the create project command
var CmdCreateProject = &cobra.Command{
	Use:   "project",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		project, err := get.FetchProject(bb, workspaceName, projectName)

		if err != nil {
			return
		} else if project != nil {
			fmt.Println("Project already exists.")
			return
		}

		createProject(bb, workspaceName, projectName, projectKey, privateFlag)
	},
}

func init() {
	CmdCreate.AddCommand(CmdCreateProject)
	CmdCreateProject.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Bitbucket workspace")
	if err := CmdCreateProject.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdCreateProject.Flags().StringVarP(&projectName, "name", "n", "", "Project name (wrap in quotes if it contains spaces)")
	if err := CmdCreateProject.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	CmdCreateProject.Flags().StringVarP(&projectKey, "key", "k", "", "Project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdCreateProject.MarkFlagRequired("key"); err != nil {
		fmt.Println(err)
	}
	CmdCreateProject.Flags().StringVarP(&projectDescription, "description", "d", "", "Project description.")
	CmdCreateProject.Flags().BoolVarP(&privateFlag, "private", "p", true, "Make the project private")
}

func createProject(bb *bitbucket.Client, workspace string, name string, key string, private bool) {
	//en := url.QueryEscape(name)
	projectOpt := &bitbucket.ProjectOptions{
		Owner:     workspace,
		Name:      name,
		Key:       key,
		IsPrivate: private,
	}

	_, err := bb.Workspaces.CreateProject(projectOpt)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
