package add

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	pipelineVariable string
	variableLabel    string
	secureFlag       bool
)

// CmdAddPipelineVariable represents the get project command
var CmdAddPipelineVariable = &cobra.Command{
	Use:   "variable",
	Short: "Add a pipeline variable for a Bitbucket repository.",
	Long: `Pass in a workspace, repository, variable value, and the variable label to add it to a Bitbucket repository.
			Add the "s" flag to make it a secure variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()
		addPipelineVariable(bb, workspaceSlug, repositorySlug, pipelineVariable, variableLabel, secureFlag)
	},
}

func init() {
	CmdAdd.AddCommand(CmdAddPipelineVariable)
	CmdAddPipelineVariable.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdAddPipelineVariable.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdAddPipelineVariable.Flags().StringVarP(&repositorySlug, "repository", "r", "", "Bitbucket project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdAddPipelineVariable.MarkFlagRequired("repository"); err != nil {
		fmt.Println(err)
	}
	CmdAddPipelineVariable.Flags().StringVarP(&pipelineVariable, "pipeline-variable", "v", "", "Pipeline variable")
	if err := CmdAddPipelineVariable.MarkFlagRequired("pipeline-variable"); err != nil {
		fmt.Println(err)
	}
	CmdAddPipelineVariable.Flags().StringVarP(&variableLabel, "variable-label", "l", "", "Pipeline variable label")
	if err := CmdAddPipelineVariable.MarkFlagRequired("variable-label"); err != nil {
		fmt.Println(err)
	}
	CmdAddPipelineVariable.Flags().BoolVarP(&secureFlag, "secure", "s", true, "Make the variable secure")

}

func addPipelineVariable(bb *bitbucket.Client, workspace string, repository string, value string, label string, secure bool) {
	variableOpts := &bitbucket.RepositoryPipelineVariableOptions{
		Owner:    workspace,
		RepoSlug: repository,
		Key:      label,
		Value:    value,
		Secured:  secure,
	}

	_, err := bb.Repositories.Repository.AddPipelineVariable(variableOpts)

	if err != nil {
		fmt.Println("Error adding pipeline variable")
		fmt.Println(err)
		return
	}

	fmt.Println("Pipeline variable added")
}
