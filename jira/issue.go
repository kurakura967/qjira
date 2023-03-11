package jira

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Issue struct {
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary   string `json:"summary"`
	Status    Status `json:"status"`
	StartDate string `json:"customfield_10200"`
	EndDate   string `json:"customfield_10201"`
}

type Status struct {
	Name string `json:"name"`
}

type searchResult struct {
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Issues     []Issue `json:"issues"`
}

type IssueService struct {
	client *Client
}

func NewIssueService(client *Client) *IssueService {
	return &IssueService{
		client: client,
	}
}

func (i *IssueService) Search(ctx context.Context, jql string, maxResults int, fields []string) ([]Issue, error) {
	url := i.client.baseURL + "/search"

	req, err := i.client.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	if jql != "" {
		q.Add("jql", jql)
	}
	if maxResults != 0 {
		q.Add("maxResults", strconv.Itoa(maxResults))
	}

	if strings.Join(fields, ",") != "" {
		q.Add("fields", strings.Join(fields, ","))
	}

	req.URL.RawQuery = q.Encode()

	res, err := i.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var s searchResult
	if err := json.Unmarshal(body, &s); err != nil {
		log.Println(err)
		return nil, err
	}

	return s.Issues, nil
}
