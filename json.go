// json.go

package utl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadFileJson(filePath string) (jsonObject interface{}, err error) {
	// Read/load/decode given filePath as some JSON object
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(byteValue), &jsonObject)
	if err != nil {
		return nil, err
	}
	return jsonObject, nil
}

func SaveFileJson(jsonObject interface{}, filePath string) {
	// Save given JSON object to given filePath
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		panic(err.Error())
	}
}

func PrintJson(jsonObject interface{}) {
	pretty, err := Prettify(jsonObject)
	if err != nil {
		fmt.Printf("Prettify() error\n")
	} else {
		fmt.Printf(pretty)
	}
	fmt.Printf("\n")
	os.Stdout.Sync() // Flush the output buffer
}

func JsonBytesReindent(jsonBytes []byte, indent int) (jsonBytes2 []byte, err error) {
	var prettyJson bytes.Buffer
	indentStr := strings.Repeat(" ", indent)
	err = json.Indent(&prettyJson, jsonBytes, "", indentStr)
	if err != nil {
		return nil, err
	}
	jsonBytes2 = prettyJson.Bytes()
	return jsonBytes2, nil
}

func JsonToBytesIndent(jsonObject interface{}, indent int) (jsonBytes []byte, err error) {
	// Convert JSON interface object to byte slice, with option indent spacing
	indentStr := strings.Repeat(" ", indent)
	jsonBytes, err = json.MarshalIndent(jsonObject, "", indentStr)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func JsonToBytes(jsonObject interface{}) (jsonBytes []byte, err error) {
	// Convert JSON interface object to byte slice, with default 2-space indentation
	indent := 2 // With default 2 space indent
	jsonBytes, err = JsonToBytesIndent(jsonObject, indent)
	return jsonBytes, err
}

func Prettify(jsonObject interface{}) (pretty string, err error) {
	// NOTE: To be replaced by JsonToBytes()
	j, err := json.MarshalIndent(jsonObject, "", "  ")
	return string(j), err
}

func PrintJsonColor(jsonObject interface{}) {
	// Print JSON object in color
	jsonBytes, err := JsonToBytes(jsonObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	PrintJsonBytesColor(jsonBytes)
}

func PrintJsonBytesColor(jsonBytes []byte) {
	// Prints JSON byte slice in color. Just an alias of yaml.go:PrintYamlBytesColor().
	PrintYamlBytesColor(jsonBytes)
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
