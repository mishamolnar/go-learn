package encoding

import (
	"log"
	"os"
	"text/template"
	"time"
)

const temp = `{{.TotalCount}} issues
{{range .Items}}------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()) / 24
}

var report = template.Must(template.New(temp).Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp))

func Display(res *IssuesSearchResult) {
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
