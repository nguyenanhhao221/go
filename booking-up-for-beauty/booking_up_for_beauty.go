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
	layout := "January 2, 2006 15:04:05"
	today := time.Now()
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
		panic("Error" + err.Error())
	}
	return today.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// panic("Please implement the IsAfternoonAppointment function")
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
		panic("Error" + err.Error())
	}
	return t.Hour() >= 12 && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// panic("Please implement the Description function")
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
		panic("Error" + err.Error())
	}
	return "You have an appointment on " + t.Format("Monday") + ", " + t.Format("January 2, 2006") + ", at " + t.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// panic("Please implement the AnniversaryDate function")
	layout := "2006-01-02 15:04:05"
	currentYear := time.Now().Format("2006")
	date := currentYear + "-09-15 00:00:00"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
		panic("Error" + err.Error())
	}
	return t
}
