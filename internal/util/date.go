// Package util is a collection of tools for handling date-times and so forth
package util

import (
	"fmt"
	"time"
)

// ParseDateOrToday - default today
func ParseDateOrToday(s string) (time.Time, error) {
	if s == "" {
		return time.Now().Truncate(24 * time.Hour), nil
	}
	return ParseDate(s)
}

// ParseDate - in regular date format
func ParseDate(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date %q (want YYYY-MM-DD)", s)
	}
	return t, nil
}

// ThisWeek returns Monday 00:00:00 to Sunday 23:59:59 in local time.
func ThisWeek() (time.Time, time.Time) {
	now := time.Now()
	offset := (int(now.Weekday()) + 6) % 7 // Monday=0
	monday := time.Date(now.Year(), now.Month(), now.Day()-offset, 0, 0, 0, 0, now.Location())
	sunday := monday.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	return monday, sunday
}

// ISOWeekBounds parses "YYYY-Www" and returns week [Mon..Sun].
func ISOWeekBounds(iso string) (time.Time, time.Time, error) {
	var y int
	var w int
	_, err := fmt.Sscanf(iso, "%d-W%d", &y, &w)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid --week, expected YYYY-Www: %w", err)
	}
	// Find the Monday of the ISO week.
	// Start from Jan 4th which is always in week 1, then adjust.
	jan4 := time.Date(y, time.January, 4, 0, 0, 0, 0, time.Local)
	_, w1 := jan4.ISOWeek()
	deltaWeeks := w - w1
	monday := jan4.AddDate(0, 0, -((int(jan4.Weekday())+6)%7)+deltaWeeks*7)
	sunday := monday.AddDate(0, 0, 6).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	return monday, sunday, nil
}
