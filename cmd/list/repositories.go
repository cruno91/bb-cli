package list

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var (
	projectKey string
)

// CmdListRepositories represents the list workspace command
var CmdListRepositories = &cobra.Command{
	Use:   "repositories",
	Short: "List repositories from a Bitbucket project.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()
		listRepositories(bb, workspaceSlug, projectKey)
	},
}

func init() {
	CmdList.AddCommand(CmdListRepositories)
	CmdListRepositories.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdListRepositories.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdListRepositories.Flags().StringVarP(&projectKey, "project", "n", "", "Bitbucket project (Example: \"PROJ\"")
	if err := CmdListRepositories.MarkFlagRequired("project"); err != nil {
		fmt.Println(err)
	}
}

func listRepositories(bb *bitbucket.Client, workspace string, project string) {
	options := &bitbucket.RepositoriesOptions{
		Owner: workspace,
	}

	repositories, err := bb.Repositories.ListForAccount(options)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	PrintRepositoriesTable(repositories.Items)
}

func PrintRepositoriesTable(repositories []bitbucket.Repository) {
	// Initialize table.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the header row.
	_, err := fmt.Fprintln(w, "Slug\tName")
	if err != nil {
		return
	}

	// Print each workspace in a row.
	for _, repository := range repositories {
		_, err := fmt.Fprintf(w, "%s\t%s\n", repository.Slug, repository.Name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Ensure the output is flushed to standard output.
	err = w.Flush()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
