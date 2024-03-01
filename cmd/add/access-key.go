package add

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	workspaceSlug string
	accessKey     string
	keyLabel      string
)

// CmdAddAccessKey represents the get project command
var CmdAddAccessKey = &cobra.Command{
	Use:   "access-key",
	Short: "Add an access key to a Bitbucket repository.",
	Long:  `Pass in a workspace, repository, access key value, and the key label to add it to a Bitbucket repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()
		addAccessKey(bb, workspaceSlug, repositorySlug, accessKey, keyLabel)
	},
}

func init() {
	CmdAdd.AddCommand(CmdAddAccessKey)
	CmdAddAccessKey.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdAddAccessKey.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdAddAccessKey.Flags().StringVarP(&repositorySlug, "repository", "r", "", "Bitbucket project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdAddAccessKey.MarkFlagRequired("repository"); err != nil {
		fmt.Println(err)
	}
	CmdAddAccessKey.Flags().StringVarP(&accessKey, "access-key", "k", "", "Access key")
	if err := CmdAddAccessKey.MarkFlagRequired("access-key"); err != nil {
		fmt.Println(err)
	}
	CmdAddAccessKey.Flags().StringVarP(&keyLabel, "key-label", "l", "", "Access key label")
	if err := CmdAddAccessKey.MarkFlagRequired("key-label"); err != nil {
		fmt.Println(err)
	}

}

func addAccessKey(bb *bitbucket.Client, workspace string, repository string, key string, label string) {
	accessKeyOpts := &bitbucket.DeployKeyOptions{
		Owner:    workspace,
		RepoSlug: repository,
		Label:    label,
		Key:      key,
	}

	_, err := bb.Repositories.DeployKeys.Create(accessKeyOpts)

	if err != nil {
		fmt.Println("Error adding access key")
		fmt.Println(err)
		return
	}

	fmt.Println("Access key added")
}
