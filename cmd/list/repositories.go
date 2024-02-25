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
	CmdListRpeositories.Flags().StringVarP(&projectKey, "project", "n", "", "Bitbucket project (Example: \"PROJ\"")
	if err := CmdListRpeositories.MarkFlagRequired("project"); err != nil {
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
