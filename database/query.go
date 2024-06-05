package database

import (
	"strings"
	"text/template"
	"time"

	pkgerrors "github.com/pkg/errors"
)

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
