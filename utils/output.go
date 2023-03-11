package utils

import (
	"os"

	"github.com/kurakura967/qjira/jira"
	"github.com/olekukonko/tablewriter"
)

func OutputTableFormat(issues []jira.Issue, url string) {
	data := make([][]string, 0)
	for _, issue := range issues {
		row := []string{
			issue.Key,
			issue.Fields.Summary,
			issue.Fields.StartDate,
			issue.Fields.EndDate,
			issue.Fields.Status.Name,
			url + "/" + issue.Key,
		}

		data = append(data, row)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Summary", "StartDt", "EndDt", "Status", "URL"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
