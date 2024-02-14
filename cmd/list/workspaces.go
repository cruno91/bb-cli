package list

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// CmdListWorkspaces represents the list workspace command
var CmdListWorkspaces = &cobra.Command{
	Use:   "workspaces",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()

		listWorkspaces(bb)
	},
}

func init() {
	CmdList.AddCommand(CmdListWorkspaces)
}

func listWorkspaces(bb *bitbucket.Client) {
	workspaceList, err := bb.Workspaces.List()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Initialize table.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the header row.
	_, err = fmt.Fprintln(w, "Name\tUUID\tSlug")
	if err != nil {
		return
	}

	// Print each workspace in a row.
	for _, workspace := range workspaceList.Workspaces {
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\n", workspace.Name, workspace.UUID, workspace.Slug)
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
