package gotools

import (
	"fmt"
	"strings"
	"time"
)

func convertLayout(layout string) string {
	parsedLayout := layout

	switch layout {
	case "datetime":
		parsedLayout = "20060102150405"
	case "-datetime":
		parsedLayout = "2006-01-02 15:04:05"
	case "-datetime-":
		parsedLayout = "2006-01-02 15:04"
	case "-datetime--":
		parsedLayout = "2006-01-02 15"
	case "-date":
		parsedLayout = "2006-01-02"
	case "date":
		parsedLayout = "20060102"
	case "-date-":
		parsedLayout = "2006-01"
	case "-date--":
		parsedLayout = "2006"
	case "/datetime":
		parsedLayout = "2006/01/02 15:04:05"
	case "/datetime-":
		parsedLayout = "2006/01/02 15:04"
	case "/datetime--":
		parsedLayout = "2006/01/02 15"
	case "/date":
		parsedLayout = "2006/01/02"
	case "/date-":
		parsedLayout = "2006/01"
	case "/date--":
		parsedLayout = "2006"
	case ".datetime":
		parsedLayout = "2006.01.02 15:04:05"
	case ".datetime-":
		parsedLayout = "2006.01.02 15:04"
	case ".datetime--":
		parsedLayout = "2006.01.02 15"
	case ".date":
		parsedLayout = "2006.01.02"
	case ".date-":
		parsedLayout = "2006.01"
	case ".date--":
		parsedLayout = "2006"
	case "time":
		parsedLayout = "15:04:05"
	case "time-":
		parsedLayout = "15:04"
	case "time--":
		parsedLayout = "15"
	}

	return parsedLayout
}

// FmtTime format time by specified layout.
func FmtTime(time time.Time, layout string) (string, error) {
	parsedLayout := convertLayout(layout)
	if strings.TrimSpace(parsedLayout) == "" {
		return "", fmt.Errorf("invalid layout format: %s", layout)
	}

	return time.Format(parsedLayout), nil
}

// MustFmtTime format time by specified layout, will panic if layout is invalid.
func MustFmtTime(time time.Time, layout string) string {
	ti, err := FmtTime(time, layout)
	if err != nil {
		panic(fmt.Sprintf("Format time failed with layout: %s", layout))
	}

	return ti
}

// ParseLocaltime time with zone set to local.
// parsable layout can be [-datetime|-datetime-|-datetime--|-date|-date-|-date--|time|time-|time--]
// and any other layout format that can pass to time.ParseInLocation
func ParseLocaltime(toParseTime, layout string) (time.Time, error) {
	parsedLayout := convertLayout(layout)
	if strings.TrimSpace(parsedLayout) == "" {
		return time.Time{}, fmt.Errorf("invalid layout format: %s", layout)
	}

	return time.ParseInLocation(parsedLayout, toParseTime, time.Now().Location())
}

// MustParseLocaltime time with zone set to local, will panic if layout specified is invalid.
// parsable layout can be [-datetime|-datetime-|-datetime--|-date|-date-|-date--|time|time-|time--]
// and any other layout format that can pass to time.ParseInLocation
func MustParseLocaltime(toParseTime, layout string) time.Time {
	ti, err := ParseLocaltime(toParseTime, layout)
	if err != nil {
		panic(fmt.Sprintf("Parse local time failed with format :%s", layout))
	}

	return ti
}
