package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	testMux := http.NewServeMux()
	testServer := httptest.NewServer(testMux)
	defer testServer.Close()

	testMux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"maxResults":1,"issues":[{"key":"PV-1111","fields":{"summary":"test1","status":{"name":"Done"},"customfield_10200":"2023-01-01","customfield_10201":"2023-01-02"}}]}`)
	})
	testClient := NewClient(new(http.Client), testServer.URL, "")
	ser := NewIssueService(testClient)
	res, err := ser.Search(context.Background(), "", 100, []string{""})
	if err != nil {
		t.Errorf("faild to search: %s", err)
	}
	want := []Issue{
		{
			"PV-1111",
			Fields{
				"test1",
				Status{
					"Done",
				},
				"2023-01-01",
				"2023-01-02",
			},
		},
	}

	if diff := cmp.Diff(want, res); diff != "" {
		t.Errorf("Search Result is mismatch (-want +res):\n%s", diff)
	}
}
