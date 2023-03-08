# qjira
qjira is a CLI tool to retrieve and display  [Atlassian Jira](https://www.atlassian.com/software/jira) Issues.


## How to use

### Set your username and access token
Define your username and token in `config/setting.yaml`
```yaml
username: your_username
toke: your_personal_access_token
baseURL: https://your_jira_domain/jira/rest/api
browseURL: https://your_jira_domain/jira/browse
```
### Build & Run
```bash
$ go build -o qjira
$ ./qjira get -f 2023-01-01 -t 2023-03-31 -s Done
+------+---------+------------+------------+--------+--------------------------------------------+
| KEY  | SUMMARY |  STARTDT   |   ENDDT    | STATUS |                  URL                       |
+------+---------+------------+------------+--------+--------------------------------------------+
| 1111 | xxxxxxx | 2023-01-01 | 2023-01-02 | 完了    | https://your_jira_domain/jira/browse/1111  |
| 1112 | yyyyyyy | 2023-01-02 | 2023-01-03 | 完了    | https://your_jira_domain/jira/browse/1112  |
+------+---------+------------+------------+------- +--------------------------------------------+

```
