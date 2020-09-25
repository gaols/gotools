package gotools

import (
	"time"
)

// BeginningOfDate returns the beginning date of a time.
func BeginningOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// BeginningOfThisWeek returns the beginning date of this week.
func BeginningOfThisWeek() time.Time {
	return BeginningOfWeek(time.Now())
}

// BeginningOfThisMonth returns the first date of this month.
func BeginningOfThisMonth() time.Time {
	return BeginningOfMonth(time.Now())
}

// LastWeek returns monday of last week.
func LastWeek() time.Time {
	return BeginningOfThisWeek().AddDate(-7, 0, 0)
}

// LastMonth returns the first date of last month.
func LastMonth() time.Time {
	return BeginningOfThisMonth().AddDate(0, -1, 0)
}

// LastYear returns the first date of last year.
func LastYear() time.Time {
	return BeginningOfThisYear().AddDate(0, -1, 0)
}

// BeginningOfThisYear returns the first date of this year.
func BeginningOfThisYear() time.Time {
	return BeginningOfYear(time.Now())
}

// BeginningOfMonth returns the first date of a month where t belongs to.
func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// BeginningOfWeek return the first date of the week where t belongs to.
func BeginningOfWeek(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	d := t.AddDate(0, 0, -wd+1)
	return BeginningOfDate(d)
}

// BeginningOfYear returns the first date of a year where t belongs to.
func BeginningOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndingOfDate returns the ending nano of a date where t belongs to.
// Deprecated: Use EndOfDate instead.
func EndingOfDate(t time.Time) time.Time {
	return EndOfDate(t)
}

// EndOfDate returns the ending nano of a date where t belongs to.
func EndOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// EndOfThisMonth returns the last date of this month.
func EndOfThisMonth() time.Time {
	return BeginningOfThisMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfThisYear returns the last date of this year.
func EndOfThisYear() time.Time {
	return BeginningOfThisYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Yesterday returns the beginning date of yesterday.
func Yesterday() time.Time {
	today := Today()
	return today.AddDate(0, 0, -1)
}

// Today returns the beginning date of today.
func Today() time.Time {
	now := time.Now()
	return BeginningOfDate(now)
}

// Tomorrow returns the beginning date of tomorrow.
func Tomorrow() time.Time {
	today := Today()
	return today.AddDate(0, 0, 1)
}

// DaysAgo returns the time subtract specified days.
func DaysAgo(t time.Time, days int) time.Time {
	if days < 0 {
		panic("days should greater than or equals 0")
	}

	return t.AddDate(0, 0, -days)
}

// DaysAfter returns the time add specified days.
func DaysAfter(t time.Time, days int) time.Time {
	if days < 0 {
		panic("days should greater than or equals 0")
	}

	return t.AddDate(0, 0, days)
}

// DaysBetween calc days between two days.
// goutils.DaysBetween(goutils.Yesterday(), goutils.Today()) = 1
// goutils.DaysBetween(goutils.Today(), goutils.Yesterday()) = -1
func DaysBetween(startTime, endTime time.Time) int {
	diff := BeginningOfDate(endTime).Sub(BeginningOfDate(startTime))
	return int(diff.Hours() / 24)
}
