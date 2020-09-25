package gotools

import (
	"testing"
	"time"
)

func TestYesterday(t *testing.T) {
	now := Today()
	yesterday := now.AddDate(0, 0, -1)
	if !Yesterday().Equal(yesterday) {
		t.FailNow()
	}
}

func TestTomorrow(t *testing.T) {
	now := Today()
	tomorrow := now.AddDate(0, 0, 1)
	if !Tomorrow().Equal(tomorrow) {
		t.FailNow()
	}
}

func TestDaysAgo(t *testing.T) {
	now := time.Now()
	if !BeginningOfDate(DaysAgo(now, 1)).Equal(Yesterday()) {
		t.FailNow()
	}
}

func TestDaysAfter(t *testing.T) {
	now := time.Now()
	if !BeginningOfDate(DaysAfter(now, 1)).Equal(Tomorrow()) {
		t.FailNow()
	}
}

func TestEndingOfDate(t *testing.T) {
	now := time.Now()
	end := EndingOfDate(now)
	if end.Hour() != 23 || end.Minute() != 59 || end.Second() != 59 || end.Year() != now.Year() || end.Month() != now.Month() || end.Day() != now.Day() {
		t.FailNow()
	}
}

func TestDaysBetween(t *testing.T) {
	if DaysBetween(Today(), Yesterday()) != -1 {
		t.FailNow()
	}
	if DaysBetween(time.Now(), Tomorrow()) != 1 {
		t.FailNow()
	}
	if DaysBetween(Yesterday(), Today()) != 1 {
		t.FailNow()
	}
	if DaysBetween(Yesterday().Add(time.Minute), Today()) != 1 {
		t.FailNow()
	}
	if DaysBetween(Today().Add(time.Second), Yesterday()) != -1 {
		t.FailNow()
	}
	if DaysBetween(Today(), Today()) != 0 {
		t.FailNow()
	}
	if DaysBetween(Today(), Today().Add(time.Millisecond)) != 0 {
		t.FailNow()
	}
}

func Test_BeginningOfWeek(t *testing.T) {
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-22", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-23", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-24", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-25", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-26", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-27", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-22" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-28", "-date")), "-date") {
		t.FailNow()
	}
	if "2019-07-15" != MustFmtTime(BeginningOfWeek(MustParseLocaltime("2019-07-21", "-date")), "-date") {
		t.FailNow()
	}
}
