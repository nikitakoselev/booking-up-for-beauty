package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	tv, _ := time.Parse("January 2, 2006 15:04:05", date)
	return tv.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	tv, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return inTheAfternoon(tv)
}

func inTheAfternoon(tv time.Time) bool {
	return tv.Hour() >= 12 && tv.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	tv, _ := time.Parse("1/2/2006 15:04:05", date)
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %.2d:%.2d.", tv.Weekday(), tv.Month(), tv.Day(), tv.Year(), tv.Hour(), tv.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return t
}
