package cmd

import "time"

func parseDate(inDate string) string {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, inDate)
	if err != nil {
		return inDate
	}
	return t.Format("02-Jan-2006")
}
