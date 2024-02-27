package add

import (
	"bb-cli/cmd/auth"
	"fmt"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/spf13/cobra"
)

var (
	webhookUrl         string
	webhookDescription string
)

// CmdAddWebhook represents the add webhook command
var CmdAddWebhook = &cobra.Command{
	Use:   "webhook",
	Short: "Get something from Bitbucket.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bb := auth.Auth()
		addWebhook(bb, workspaceSlug, projectKey, webhookUrl, webhookDescription, true)

	},
}

func init() {
	CmdAdd.AddCommand(CmdAddWebhook)
	CmdAddWebhook.Flags().StringVarP(&workspaceSlug, "workspace", "w", "", "Bitbucket workspace (Example: \"my-workspace\"")
	if err := CmdAddWebhook.MarkFlagRequired("workspace"); err != nil {
		fmt.Println(err)
	}
	CmdAddWebhook.Flags().StringVarP(&projectKey, "project", "n", "", "Bitbucket project key (Example: For a project named \"My Project\" the key could be \"MP\"")
	if err := CmdAddWebhook.MarkFlagRequired("project"); err != nil {
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

}

func addWebhook(bb *bitbucket.Client, workspace string, project string, url string, label string, active bool, events []string) {
	webhookOpts := &bitbucket.WebhooksOptions{
		Owner:       workspace,
		RepoSlug:    project,
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
