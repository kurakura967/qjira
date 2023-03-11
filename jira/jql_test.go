package jira

import (
	"testing"
)

func TestBuildJQL(t *testing.T) {
	testCase := []struct {
		userName string
		from     string
		to       string
		status   string
		want     string
	}{
		{
			"your_username",
			"2023-01-01",
			"2023-01-02",
			"Done",
			"assignee=your_username and cf[10200]>2023-01-01 and cf[10201]<2023-01-02 and status=10001",
		},
	}

	for _, tt := range testCase {
		got := BuildJQL(tt.userName, tt.from, tt.to, tt.status)
		if got != tt.want {
			t.Errorf("unexpected jql: want %s but %s\n", tt.want, got)
		}
	}
}
