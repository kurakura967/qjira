package jira

import "testing"

func TestGetStatusCode(t *testing.T) {
	testCase := []struct {
		status string
		want   string
	}{
		{
			"Done",
			"10001",
		},
		{
			"Pending",
			"10116",
		},
		{
			"Todo",
			"10118",
		},
		{
			"Wip",
			"10119",
		},
		{
			"Hoge",
			"",
		},
	}

	for _, tt := range testCase {
		got := GetStatusCode(tt.status)
		if got != tt.want {
			t.Errorf("unexpected status: want %s but %s\n", tt.status, got)
		}
	}
}
