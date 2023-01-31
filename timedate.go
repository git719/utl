// timedate.go

package utl

import (
	"fmt"
	"strconv"
	"time"
)

func ValidDate(dateString, expectedFormat string) bool {
	// Check if string is a valid date in expectedFormat
	// References:
	// - https://pkg.go.dev/time
	_, err := time.Parse(expectedFormat, dateString)
	if err != nil {
		return false
	}
	return true
}

func IntAbs(val int64) int64 {
	// Return absolute value of int64 value
	if val < 0 {
		return -val
	}
	return val
}

func StringToInt64(s string) (int64, error) {
	// Convert string number to int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func EpocIntToTime(epocInt int64) time.Time {
	// Convert Unix epoc seconds string to Time type
	return time.Unix(epocInt, 0)
}

func EpocStringToTime(epocString string) (time.Time, error) {
	// Convert Unix epoc seconds string to Time type
	epocInt64, err := StringToInt64(epocString)
	return time.Unix(epocInt64, 0), err
}

func ConvertDateFormat(dateString, fromFormat, toFormat string) (string, error) {
	// Converts dateString from fromFormat to toFormat
	t, err := time.Parse(fromFormat, dateString)
	if err != nil {
		return "", err
	}
	return t.Format(toFormat), nil
}

func DateStringToEpocInt64(dateString, dateFormat string) (int64, error) {
	// Convert dateString, given in dateFormat, to Unix Epoc seconds int64
	t, err := time.Parse(dateFormat, dateString) // First, convert string to Time
	if err != nil {
		panic(err.Error())
	}
	return t.Unix(), nil // Finally, convert Time type to Unix epoc seconds
}

func PrintDateInDays(days string) {
	// Print yyyy-mm-dd date for given number of +/- days in future or past
	now := time.Now().Unix()
	daysInt64, err := StringToInt64(days)
	if err != nil {
		panic(err.Error())
	}
	now += (daysInt64 * 86400) // 86400 seconds in a day
	date1 := EpocIntToTime(now)
	fmt.Println(date1.Format("2006-01-02"))
}

func PrintDays(days int64) {
	days_abs := IntAbs(days)
	if days_abs > 365 {
		years := days_abs / 365
		modulus := days_abs % 365
		fmt.Printf("%d (%d years + %d days)\n", days, years, modulus)
	} else {
		fmt.Println(days)
	}
}

func PrintDaysSinceOrTo(date1 string) {
	// Calculate and print number of +/- days from NOW to date given
	// Note: Calculations are all in UTC time
	epoc1, err := DateStringToEpocInt64(date1, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}

	now := time.Now()
	now_secs := now.Unix()

	// Get today's epoc time at midnight UTC
	today := now.Format("2006-01-02")
	t, err := time.Parse("2006-01-02", today) // First, convert string to Time
	if err != nil {
		panic(err.Error())
	}
	midnight_secs := t.Unix()

	since_midnight := now_secs - midnight_secs
	epoc1 = epoc1 + since_midnight
	PrintDays((epoc1 - now_secs) / 86400)
}

func PrintDaysBetween(date1, date2 string) {
	// Calculate and print number of days between 2 dates
	epoc1, err := DateStringToEpocInt64(date1, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}
	epoc2, err := DateStringToEpocInt64(date2, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}
	if epoc1 > epoc2 {
		PrintDays((epoc1 - epoc2) / 86400)
	} else if epoc2 > epoc1 {
		PrintDays((epoc2 - epoc1) / 86400)
	} else {
		fmt.Println(0)
	}
}
