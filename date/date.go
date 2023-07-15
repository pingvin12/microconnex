package date

import (
	"fmt"
	"time"
)

func calculateExpirationDate(t time.Time, turnaroundTimeNumber int) time.Time {
	t = t.Add(time.Duration(turnaroundTimeNumber) * time.Hour)

	for {
		weekday := t.Weekday()
		hour := t.Hour()

		if weekday == time.Saturday || weekday == time.Sunday {
			t = t.Add(24 * time.Hour)
			continue
		}

		if hour < 9 {
			t = time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location())
			continue
		}

		if hour >= 17 {
			t = time.Date(t.Year(), t.Month(), t.Day()+1, 9, 0, 0, 0, t.Location())
			continue
		}

		break
	}

	return t
}

func GetExpirationDate(date string, turnaroundTimeNumber int) (string, error) {
	t, err := parseTimestamp(date)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	t = calculateExpirationDate(t, turnaroundTimeNumber)

	return t.Format("2006-01-02T15:04:05.000Z"), nil
}
func parseTimestamp(date string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", date)
}
