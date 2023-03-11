/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/kurakura967/qjira/config"
	"github.com/kurakura967/qjira/jira"
	"github.com/kurakura967/qjira/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get and display Issues to you are assigned",
	Long:  "Get and display Issues to you are assigned",

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		from, err := cmd.Flags().GetString("from")
		if err != nil {
			return err
		}
		end, err := cmd.Flags().GetString("to")
		if err != nil {
			return err
		}
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			return err
		}
		if err := validateStatus(status); err != nil {
			return err
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		client := jira.NewClient(new(http.Client), cfg.BaseURL, cfg.Token)
		ser := jira.NewIssueService(client)

		jql := jira.BuildJQL(cfg.Username, from, end, status)
		fields := []string{"summary", "status", "customfield_10200", "customfield_10201"}
		issues, err := ser.Search(ctx, jql, 10, fields)
		if err != nil {
			return err
		}
		utils.OutputTableFormat(issues, cfg.BrowseURL)
		return nil
	},
}

func validateStatus(status string) error {
	if s := jira.GetStatusCode(status); s == "" {
		return fmt.Errorf("not exists status: %s", status)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("from", "f", "", "(required) Issue start date 'YYYY-MM-DD'")
	getCmd.Flags().StringP("to", "t", "", "(required) Issue end date 'YYYY-MM-DD'")
	getCmd.Flags().StringP("status", "s", "Done", "Issue Status ('Done', 'Pending', 'Todo', 'Wip'). Default is 'Done'")
}
