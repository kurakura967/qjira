package jira

var StatusCode = map[string]string{
	"Done":    "10001",
	"Pending": "10116",
	"Todo":    "10118",
	"Wip":     "10119",
}

func GetStatusCode(s string) string {
	return StatusCode[s]
}
