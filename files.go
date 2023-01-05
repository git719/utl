// files.go

package chores

import (
	"os"
)

func RemoveFile(filePath string) {
	if FileExist(filePath) {
		if err := os.Remove(filePath); err != nil {
			panic(err.Error())
		}
	}
}

func FileUsable(filePath string) (e bool) {
	// True if file EXISTS && has SOME content && is less than 30 days old
	if FileExist(filePath) && FileSize(filePath) > 0 {
		return true
	}
	return false
}

func FileExist(filePath string) (e bool) {
	if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
		return true
	}
	return false
}

func FileNotExist(filePath string) (e bool) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return true
	}
	return false
}

func FileSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if err != nil { return 0 }
	return f.Size()
}

func FileModTime(filePath string) int {
	// Modified time in Unix epoch
	f, err := os.Stat(filePath)
	if err != nil { return 0 }
	return int(f.ModTime().Unix())
}
