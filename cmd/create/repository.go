package create

import (
	"bb-cli/cmd/auth"
	"bb-cli/cmd/list"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	repositorySlug string
)

// CmdCreateRepository represents the create repository command
var CmdCreateRepository = &cobra.Command{
	Use:   "repository",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		createRepository(bb, workspaceName, projectName, repositorySlug, privateFlag)
	},
}

func init() {
	CmdCreate.AddCommand(CmdCreateRepository)
	CmdCreateRepository.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Bitbucket workspace")
	if err := CmdCreateRepository.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdCreateRepository.Flags().StringVarP(&projectName, "project", "n", "", "Project name (wrap in quotes if it contains spaces)")
	if err := CmdCreateRepository.MarkFlagRequired("project"); err != nil {
		fmt.Println(err)
	}
	CmdCreateRepository.Flags().StringVarP(&repositorySlug, "repository", "r", "", "Repository slug")
	if err := CmdCreateRepository.MarkFlagRequired("repository"); err != nil {
		fmt.Println(err)
	}
	CmdCreateRepository.Flags().BoolVarP(&privateFlag, "private", "p", true, "Make the project private")
}

func createRepository(bb *bitbucket.Client, workspace string, project string, name string, private bool) {
	isPrivate := "true"

	if !private {
		isPrivate = "false"
	}

	repoOpts := &bitbucket.RepositoryOptions{
		Owner:     workspace,
		Project:   project,
		RepoSlug:  name,
		IsPrivate: isPrivate,
	}

	res, err := bb.Repositories.Repository.Create(repoOpts)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Repository created...")
	list.PrintRepositoriesTable([]bitbucket.Repository{*res})
}
