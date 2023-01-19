// utils.go

package utl

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"runtime"
	"sort"
)

func Die(format string, args ...interface{}) {
	fmt.Printf(format, args...) // Same as print function but does not return
	os.Exit(1)                  // Always exit with return code 1
}

func Trace() string {
	// Return string showing current "File_path [line number] function_name"
	// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
	progCounter, fp, ln, ok := runtime.Caller(1)
	if !ok {
		return fmt.Sprintf("%s\n    %s:%d\n", "?", "?", 0)
	}
	funcPointer := runtime.FuncForPC(progCounter)
	if funcPointer == nil {
		return fmt.Sprintf("%s\n    %s:%d\n", "?", fp, ln)
	}
	return fmt.Sprintf("%s\n    %s:%d\n", funcPointer.Name(), fp, ln)
}

func ValidUuid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func SameType(a, b interface{}) bool {
	// Check if two variables are of the same type
	a_type := fmt.Sprintf("%T", a)
	b_type := fmt.Sprintf("%T", b)
	return a_type == b_type
}

func GetType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func IsHexDigit(c rune) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F') {
		return true
	}
	return false
}

func SortStringMapByKeys(inMap map[string]string) (sortedMap map[string]string) {
	keys := make([]string, 0, len(inMap))
	for k := range inMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedMap = make(map[string]string, len(inMap))
	for _, k := range keys {
		sortedMap[k] = inMap[k]
	}
	return sortedMap
}
