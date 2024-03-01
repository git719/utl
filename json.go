// json.go

package utl

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func LoadFileJson(filePath string) (jsonObject interface{}, err error) {
	// Read/load/decode text JSON object file
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

func LoadFileJsonGzip(filePath string) (jsonObject interface{}, err error) {
	// Read/load/decode gzipped JSON object file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	byteValue, err := ioutil.ReadAll(gzipReader)
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
	// Save given JSON object as text file
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		panic(err.Error())
	}
}
func SaveFileJsonGzip(jsonObject interface{}, filePath string) {
	// Save given JSON object as gzipped file
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		panic(err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write(jsonData)
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

func JsonBytesToJsonObj(jsonBytes []byte) (jsonObject interface{}, err error) {
	// Convert JSON byte slice to JSON interface object, with default 2-space indentation
	err = json.Unmarshal(jsonBytes, &jsonObject)
	return jsonObject, err
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

func StringInJson(jsonObject interface{}, filter string) bool {
	// Recursive function returns True if filter string value is anywhere within jsonObject
	switch value := jsonObject.(type) {
	case string:
		return SubString(value, filter)
	case []interface{}:
		for _, v := range value {
			if StringInJson(v, filter) {
				return true
			}
		}
	case map[string]interface{}:
		for _, v := range value {
			if StringInJson(v, filter) {
				return true
			}
		}
	}
	return false
}
