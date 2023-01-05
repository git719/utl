// utils.go

package chores

import (
	"os"
	"fmt"
	"runtime"

	"github.com/google/uuid"
)

func exit(code int) {
	os.Exit(code)       // Syntactic sugar. Easier to type
}

func print(format string, args ...interface{}) (n int, err error) {
	return fmt.Printf(format, args...) // More syntactic sugar
}

func die(format string, args ...interface{}) {
	fmt.Printf(format, args...) // Same as print function but does not return
	os.Exit(1)                  // Always exit with return code 1
}

func sprint(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)	// More syntactic sugar
}

func trace() (string) {
	// Return string showing current "File_path [line number] function_name"
	// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
    progCounter, fp, ln, ok := runtime.Caller(1)
    if !ok { return sprint("%s\n    %s:%d\n", "?", "?", 0) }
    funcPointer := runtime.FuncForPC(progCounter)
    if funcPointer == nil { return sprint("%s\n    %s:%d\n", "?", fp, ln) }
	return sprint("%s\n    %s:%d\n", funcPointer.Name(), fp, ln)
}

func ValidUuid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func SameType(a, b interface{}) bool {
	// Check if two variables are of the same type
	a_type := sprint("%T", a)
	b_type := sprint("%T", b)
	return a_type == b_type
}

func VarType(v interface{}) string {
	return sprint("%T", v)
}

