// files.go

package utl

import (
	"os"
	"time"
)

func LoadFileText(filePath string) (rawBytes []byte, err error) {
	rawBytes, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return rawBytes, nil
}

func SaveFileText(filePath string, rawBytes []byte) error {
	err := os.WriteFile(filePath, rawBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(filePath string) {
	if FileExist(filePath) {
		if err := os.Remove(filePath); err != nil {
			panic(err.Error())
		}
	}
}

func FileUsable(filePath string) (e bool) {
	// True if file EXISTS && has SOME content
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
	if err != nil {
		return 0
	}
	return f.Size()
}

func FileModTime(filePath string) int {
	// Modified time in Unix epoch
	f, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return int(f.ModTime().Unix())
}

func FileAge(filePath string) int64 {
	// Return file age in seconds
	if FileUsable(filePath) {
		fileEpoc := int64(FileModTime(filePath))
		return int64(time.Now().Unix()) - fileEpoc
	}
	return int64(0)
}
