package jira

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

const (
	testJiraInstanceURL = "https://issues.apache.org/jira/"
)

func TestNewRequestClient(t *testing.T) {
	httpClient := http.DefaultClient
	c := NewClient(httpClient, testJiraInstanceURL, "")
	url := c.baseURL + "rest/api/2/issue/"

	input := Issue{
		Key: "Key",
		Fields: Fields{
			"Summary",
			Status{
				"Status",
			},
			"StartDt",
			"EndDt",
		},
	}
	want := `{"key":"Key","fields":{"summary":"Summary","status":{"name":"Status"},"customfield_10200":"StartDt","customfield_10201":"EndDt"}}` + "\n"

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(input)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := c.NewRequestWithContext(context.Background(), http.MethodGet, url, &b)
	body, _ := io.ReadAll(req.Body)
	got := string(body)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
