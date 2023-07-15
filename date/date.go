package date

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func GetExpirationDate(date string, turnaroundTimeNumber int) timestamp.Timestamp {

	// Convert the timestamp to a time.Time
	t, err := time.Parse("2006-01-02T15:04:05.000Z", date)
	if err != nil {
		fmt.Println(err)
	}

	// Add the turnaround time to the date
	t = t.Add(time.Duration(turnaroundTimeNumber) * time.Hour)

	// Loop until the date is on a workday between 9-17 hours
	for {
		// Get the weekday and the hour of the date
		weekday := t.Weekday()
		hour := t.Hour()

		// If the date is on a weekend, add one day
		if weekday == time.Saturday || weekday == time.Sunday {
			t = t.Add(24 * time.Hour)
			continue
		}

		// If the date is before 9 am, set it to 9 am
		if hour < 9 {
			t = time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location())
			continue
		}

		// If the date is after 5 pm, set it to 9 am of the next day
		if hour >= 17 {
			t = time.Date(t.Year(), t.Month(), t.Day()+1, 9, 0, 0, 0, t.Location())
			continue
		}

		// If the date is valid, break the loop
		break
	}

	// Convert the time.Time back to a timestamp.Timestamp
	return timestamp.Timestamp{Seconds: t.Unix()}
}
