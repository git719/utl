// utl.go

package utl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"

	"github.com/google/uuid"
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

func IsAlpha(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}

func IsDigit(c rune) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func IsHexDigit(c rune) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F') {
		return true
	}
	return false
}

// TODO: Combine below two func with interfaces
func SortMapStringKeys(obj map[string]string) (sortedKeys []string) {
	// Return the map string object's keys sorted
	sortedKeys = make([]string, 0, len(obj))
	for k := range obj {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}
func SortObjStringKeys(obj map[string]interface{}) (sortedKeys []string) {
	// Return the object's keys sorted
	sortedKeys = make([]string, 0, len(obj))
	for k := range obj {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func PromptMsg(msg string) rune {
	// Print prompt message and return single rune character input
	fmt.Print(Yel(msg))
	reader := bufio.NewReader(os.Stdin)
	confirm, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return confirm
}
