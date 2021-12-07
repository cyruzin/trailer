package cmd

import (
	"time"
)

func parseDate(inDate string) string {
	if inDate == "" {
		return "Date unavailable"
	}
	layout := "2006-01-02"
	t, err := time.Parse(layout, inDate)
	if err != nil {
		return "Date unavailable"
	}
	return t.Format("2006")
}
