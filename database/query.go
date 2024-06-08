package database

import (
	"strings"
	"text/template"
	"time"

	pkgerrors "github.com/pkg/errors"
)

// compileQuery compiles a SQL query template with predefined date parameters.
// It supports placeholders for various date ranges like today, yesterday,
// this week, last week, this month, and last month.
//
// The supported placeholders in the query template are:
// - {{.Today}}: The current date in "YYYY-MM-DD" format
// - {{.Yesterday}}: The date of the previous day
// - {{.ThisWeekStart}}: The start date of the current week (Monday)
// - {{.ThisWeekEnd}}: The end date of the current week (Sunday)
// - {{.LastWeekStart}}: The start date of the previous week (Monday)
// - {{.LastWeekEnd}}: The end date of the previous week (Sunday)
// - {{.ThisMonthStart}}: The first day of the current month
// - {{.ThisMonthEnd}}: The last day of the current month
// - {{.LastMonthStart}}: The first day of the previous month
// - {{.LastMonthEnd}}: The last day of the previous month
func (d *Database) compileQuery(query string) (string, error) {
	tmpl, err := template.New("query").Parse(query)
	if err != nil {
		return "", pkgerrors.Wrap(err, "Parse")
	}

	now := time.Now()
	format := "2006-01-02"

	thisWeekStart := now.AddDate(0, 0, -int(now.Weekday())+1)
	thisWeekEnd := thisWeekStart.AddDate(0, 0, 6)
	lastWeekStart := thisWeekStart.AddDate(0, 0, -7)
	lastWeekEnd := lastWeekStart.AddDate(0, 0, 6)

	thisMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	thisMonthEnd := thisMonthStart.AddDate(0, 1, -1)
	lastMonthStart := thisMonthStart.AddDate(0, -1, 0)
	lastMonthEnd := thisMonthStart.AddDate(0, 0, -1)

	params := map[string]string{
		"Today":          now.Format(format),
		"Yesterday":      now.AddDate(0, 0, -1).Format(format),
		"ThisWeekStart":  thisWeekStart.Format(format),
		"ThisWeekEnd":    thisWeekEnd.Format(format),
		"LastWeekStart":  lastWeekStart.Format(format),
		"LastWeekEnd":    lastWeekEnd.Format(format),
		"ThisMonthStart": thisMonthStart.Format(format),
		"ThisMonthEnd":   thisMonthEnd.Format(format),
		"LastMonthStart": lastMonthStart.Format(format),
		"LastMonthEnd":   lastMonthEnd.Format(format),
	}

	var builder strings.Builder
	if err := tmpl.Execute(&builder, params); err != nil {
		return "", pkgerrors.Wrap(err, "tmpl.Execute")
	}

	return builder.String(), nil
}

// RunQuery runs a SQL query after compiling it with date placeholders.
// It executes the query and returns the results as a slice of maps, where
// each map represents a row with column names as keys and column values as values.
// Parameters:
// - query: The SQL query template containing placeholders.
func (d *Database) RunQuery(query string) ([]map[string]interface{}, error) {
	compiledQuery, err := d.compileQuery(query)
	if err != nil {
		return nil, err
	}

	rows, err := d.db.Raw(compiledQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		cm := make([]interface{}, len(cols))
		cmp := make([]interface{}, len(cols))
		for i := range cm {
			cmp[i] = &cm[i]
		}

		if err := rows.Scan(cmp...); err != nil {
			return nil, err
		}

		rm := make(map[string]interface{})
		for i, colName := range cols {
			val := cmp[i].(*interface{})
			rm[colName] = *val
		}

		results = append(results, rm)
	}

	return results, nil
}
