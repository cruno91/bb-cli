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
	workspaceSlug string
)

// CmdListProjects represents the list workspace command
var CmdListProjects = &cobra.Command{
	Use:   "projects",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		listProjects(bb, workspaceSlug)
	},
}

func init() {
	CmdList.AddCommand(CmdListProjects)
	CmdListProjects.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdListProjects.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
}

func listProjects(bb *bitbucket.Client, workspace string) {
	projectList, err := bb.Workspaces.Projects(workspace)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Initialize table.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the header row.
	_, err = fmt.Fprintln(w, "Name\tUUID\tKey\tPrivate\tDescription")
	if err != nil {
		return
	}

	// Print each workspace in a row.
	for _, project := range projectList.Items {
		private := "no"
		if project.Is_private == true {
			private = "yes"
		}
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", project.Name, project.Uuid, project.Key, private, project.Description)
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
