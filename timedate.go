// timedate.go

package utl

import (
	"fmt"
	"strconv"
	"time"
)

func IntAbs(x int) int {
	// Return absolute value of int value
	if x < 0 {
		return -x
	}
	return x
}

func Int64Abs(x int64) int64 {
	// Return absolute value of int64 value
	if x < 0 {
		return -x
	}
	return x
}

func StringToInt64(s string) (int64, error) {
	// Convert string number to int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func Int64ToString(i int64) string {
	// Convert int64 number to string
	return strconv.FormatInt(i, 10)
}

func ValidDate(dateString, expectedFormat string) bool {
	// Check if string is a valid date in expectedFormat
	// Reference: https://pkg.go.dev/time
	_, err := time.Parse(expectedFormat, dateString)
	if err != nil {
		return false
	}
	return true
}

func EpocInt64ToTime(epocInt int64) time.Time {
	// Convert Unix epoc seconds string to Time type
	return time.Unix(epocInt, 0)
}

func EpocStringToTime(epocString string) (time.Time, error) {
	// Convert Unix epoc seconds string to Time type
	epocInt64, err := StringToInt64(epocString)
	return time.Unix(epocInt64, 0), err
}

func ConvertDateFormat(dateString, srcFormat, dstFormat string) (string, error) {
	// Converts dateString from srcFormat to dstFormat
	t, err := time.Parse(srcFormat, dateString)
	if err != nil {
		return "", err
	}
	return t.Format(dstFormat), nil
}

func DateStringToEpocInt64(dateString, dateFormat string) (int64, error) {
	// Convert dateString, given in dateFormat, to Unix Epoc seconds int64
	t, err := time.Parse(dateFormat, dateString) // First, convert string to Time
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil // Finally, convert Time type to Unix epoc seconds
}

func GetDateInDays(days string) time.Time {
	// Print yyyy-mm-dd date for given number of +/- days in future or past
	now := time.Now().Unix()
	daysInt64, err := StringToInt64(days)
	if err != nil {
		panic(err.Error())
	}
	now += (daysInt64 * 86400) // 86400 seconds in a day
	return EpocInt64ToTime(now)
}

func IsLeapYear(year int64) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func GetDaysSinceOrTo(date1 string) int64 {
	// Calculate and return number of +/- days from NOW to date given
	// Note: Calculations are all in UTC time
	// Takes leap year into account.
	start, err := time.Parse("2006-01-02", date1)
	if err != nil {
		panic(err.Error())
	}

	end := time.Now().UTC()

	var days int64 = 0
	var sign int64 = -1

	if start.After(end) {
		start, end = end, start
		sign = 1
	}

	for start.Year() < end.Year() || (start.Year() == end.Year() && start.YearDay() < end.YearDay()) {
		days++
		start = start.AddDate(0, 0, 1)

		// Adjust for leap years
		if start.Month() == time.February && start.Day() == 28 && IsLeapYear(int64(start.Year())) {
			days++
			start = start.AddDate(0, 0, 1)
		}
	}

	return sign * days
}

func PrintDays(days int64) {
	// Print number of days, also in years and days
	days_abs := Int64Abs(days)
	var years int64 = 0

	for days_abs >= 365 {
		leap := int64(0)
		if IsLeapYear(years) {
			leap = 1
		}
		if days_abs >= (365 + leap) {
			days_abs -= (365 + leap)
			years++
		} else {
			break
		}
	}

	if years > 0 {
		fmt.Printf("%d (%d years + %d days)\n", days, years, days_abs)
	} else {
		fmt.Println(days)
	}
}

func GetDaysBetween(date1, date2 string) int64 {
	// Return number of days between 2 dates
	epoc1, err := DateStringToEpocInt64(date1, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}
	epoc2, err := DateStringToEpocInt64(date2, "2006-01-02")
	if err != nil {
		panic(err.Error())
	}

	return (Int64Abs(epoc1-epoc2) / 86400)
}
