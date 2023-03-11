package utils

import (
	"bytes"
	"os"
	"testing"

	"github.com/kurakura967/qjira/jira"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	return buf.String()
}

func TestOutputTableFormat(t *testing.T) {
	input := []jira.Issue{
		{
			Key: "key1",
			Fields: jira.Fields{
				Summary: "summary1",
				Status: jira.Status{
					Name: "Done",
				},
				StartDate: "2023-01-01",
				EndDate:   "2023-01-02",
			},
		},
	}

	want := `+------+----------+------------+------------+--------+-------------------+
| KEY  | SUMMARY  |  STARTDT   |   ENDDT    | STATUS |        URL        |
+------+----------+------------+------------+--------+-------------------+
| key1 | summary1 | 2023-01-01 | 2023-01-02 | Done   | https://test/key1 |
+------+----------+------------+------------+--------+-------------------+
`

	got := captureStdout(t, func() {
		OutputTableFormat(input, "https://test")
	})
	if got != want {
		t.Errorf("unexpected table: want %s \n but \n%s", want, got)
	}
}
