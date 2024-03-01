package add

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	webhookUrl         string
	repositorySlug     string
	webhookDescription string
	events             []string
)

// CmdAddWebhook represents the add webhook command
var CmdAddWebhook = &cobra.Command{
	Use:   "webhook",
	Short: "Get something from Bitbucket.",
	Long: `Pass in a workspace, repository, webhook url, the webhook label, and a series of events for the webhook 
			to listen to to add it to a Bitbucket repository like "repo:push" or "repo:update". If you do not add
			any events, those two defaults will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(events) == 0 {
			events = []string{"repo:push", "repo:update"}
		}

		bb := auth.Auth()
		addWebhook(bb, workspaceSlug, repositorySlug, webhookUrl, webhookDescription, true, events)
	},
}

func init() {
	CmdAdd.AddCommand(CmdAddWebhook)
	CmdAddWebhook.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdAddWebhook.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdAddWebhook.Flags().StringVarP(&repositorySlug, "repository", "r", "", "Bitbucket project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdAddWebhook.MarkFlagRequired("repository"); err != nil {
		fmt.Println(err)
	}
	CmdAddWebhook.Flags().StringVarP(&webhookUrl, "url", "u", "", "Webhook Url")
	if err := CmdAddWebhook.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	CmdAddWebhook.Flags().StringVarP(&webhookDescription, "description", "d", "", "Webhook description")
	if err := CmdAddWebhook.MarkFlagRequired("description"); err != nil {
		fmt.Println(err)
	}
	CmdAddWebhook.Flags().StringSliceVarP(&events, "events", "e", []string{}, "Webhook events (e.g., repo:push, repo:update)")

}

func addWebhook(bb *bitbucket.Client, workspace string, repository string, url string, label string, active bool, events []string) {
	webhookOpts := &bitbucket.WebhooksOptions{
		Owner:       workspace,
		RepoSlug:    repository,
		Description: label,
		Url:         url,
		Active:      active,
		Events:      events,
	}

	_, err := bb.Repositories.Webhooks.Create(webhookOpts)

	if err != nil {
		fmt.Println("Error creating webhook")
		fmt.Println(err)
		return
	}

	fmt.Println("Webhook created")
}
