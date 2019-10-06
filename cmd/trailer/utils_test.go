package cmd

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		dateIn       string
		dateExpected string
	}{
		// Incorrect input
		{"", ""},
		{"not a date", "not a date"},
		{"01-02-03", "01-02-03"},
		{"2019-10-05", "2019-10-05"},
		{"2019-10-05 12:24", "2019-10-05 12:24"},

		// Proper input
		{"2016-09-02T15:04:05.000Z", "02-Sep-2016"},
	}

	for _, test := range tests {
		dateOut := parseDate(test.dateIn)

		if dateOut != test.dateExpected {
			t.Errorf("expected '%s' and got '%s'", test.dateExpected, dateOut)
		}
	}
}
