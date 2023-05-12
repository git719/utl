// strings.go

package utl

import (
	"fmt"
	"strings"
)

func SubString(large, small string) bool {
	// Case insensitive substring search
	if strings.Contains(strings.ToLower(large), strings.ToLower(small)) {
		return true
	}
	return false
}

func LastElem(s, splitter string) string {
	split := strings.Split(s, splitter) // Split the string
	return split[len(split)-1]          // Return last element
}

func FirstN(s string, n int) string {
	// Return first N characters of string n
	if len(s) <= n {
		return s
	}
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func Str(x interface{}) string {
	// Return the best printable string value for given x variable
	if x == nil {
		return ""
	}
	switch GetType(x) {
	case "bool":
		return fmt.Sprintf("%t", x)
	case "string":
		return x.(string)
	default:
		return "" // Blank for other types
	}
}

func ToStr(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func ItemInList(arg string, argList []string) bool {
	for _, value := range argList {
		if value == arg {
			return true
		}
	}
	return false
}

func PadSpaces(targetWidth, stringWidth int) string {
	// Return string of spaces for padded printing. Needed when printing terminal colors.
	// Colorize output uses % sequences that conflict with Printf's own formatting with %
	padding := targetWidth - stringWidth
	if padding > 0 {
		return fmt.Sprintf("%*s", padding, " ")
	} else {
		return ""
	}
}

func PreSpc(value interface{}, width int) string {
	// Return value as a string, padded with leading spaces, totalling width
	// Needed when printing terminal colors, because they conflict with Printf's own '%' formatting
	str := ToStr(value)
	padding := width - len(str)
	if padding > 0 {
		return fmt.Sprintf("%*s%s", padding, " ", str)
	} else {
		return str
	}
}

func PostSpc(value interface{}, width int) string {
	// Return value as a string, padded with trailing spaces, totalling width
	// Needed when printing terminal colors, because they conflict with Printf's own '%' formatting
	str := ToStr(value)
	padding := width - len(str)
	if padding > 0 {
		return fmt.Sprintf("%s%*s", str, padding, " ")
	} else {
		return str
	}
}
