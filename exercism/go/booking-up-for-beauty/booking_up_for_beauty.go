package booking

import (
	"fmt"
	"time"
)

// reference datetime:
// Monday 01/02 03:04:05PM '06 -0700

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dateT, _ := time.Parse(layout, date)

	nowT := time.Now()
	hasPassed := dateT.Before(nowT)

	return hasPassed
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	hours, minutes, _ := t.Clock()

	timeString := fmt.Sprintf("%d%02d", hours, minutes)

	return "1200" <= timeString && timeString < "1800"

}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	tNow := time.Now()
	t := time.Date(tNow.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return t
}
