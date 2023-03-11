package jira

import (
	"fmt"
	"strings"
)

func BuildJQL(userName, startDt, endDt, status string) string {
	a := fmt.Sprintf("assignee=%s", userName)
	s := fmt.Sprintf("cf[10200]>%s", startDt)
	e := fmt.Sprintf("cf[10201]<%s", endDt)
	st := buildStatusCondition(status)
	return strings.Join([]string{a, s, e, st}, " and ")
}

func buildStatusCondition(status string) string {
	if status != "" {
		code := GetStatusCode(status)
		return fmt.Sprintf("status=%s", code)
	}
	res := make([]string, 0)
	for _, code := range StatusCode {
		c := fmt.Sprintf("status=%s", code)
		res = append(res, c)
	}
	return "(" + strings.Join(res, " or ") + ")"
}
