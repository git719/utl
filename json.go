// json.go

package chores

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadFileJson(filePath string) (jsonObject interface{}, err error) {
	// Read/load/decode given filePath as some JSON object
	f, err := os.Open(filePath)
	if err != nil { return nil, err }
	defer f.Close()
	byteValue, err := ioutil.ReadAll(f)
	if err != nil { return nil, err }
	err = json.Unmarshal([]byte(byteValue), &jsonObject)
	if err != nil { return nil, err }
	return jsonObject, nil
}

func SaveFileJson(jsonObject interface{}, filePath string) {
	// Save given JSON object to given filePath
	jsonData, err := json.Marshal(jsonObject)
	if err != nil { panic(err.Error()) }
	err = ioutil.WriteFile(filePath, jsonData, 0600)
	if err != nil { panic(err.Error()) }
}

func PrintJson(jsonObject interface{}) {
	pretty, err := Prettify(jsonObject)
	if err != nil {
		print("Prettify() error\n")
	} else {
		print(pretty)
	}
	print("\n")
}

func Prettify(jsonObject interface{}) (pretty string, err error) {
	j, err := json.MarshalIndent(jsonObject, "", "  ")
	return string(j), err
}

func MergeMaps(m1, m2 map[string]string) (result map[string]string) {
	result = map[string]string{}
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

func MergeObjects(x, y map[string]interface{}) (obj map[string]interface{}) {
	// Merge JSON object y into x
	// NOTES:
	// 1. Non-recursive, only works attributes at first level
	// 2. If attribute exists in y, it is overwritten
	obj = x
	for k, v := range x { // Update existing x values with updated y values
		obj[k] = v
		if y[k] != nil {
			obj[k] = y[k]
		}
	}
	for k, _ := range y { // Add new y values to x
		if x[k] == nil {
			obj[k] = y[k]
		}
	}
	return obj
}
