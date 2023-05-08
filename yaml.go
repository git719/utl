// yaml.go

package utl

import (
	"bytes"
	"fmt"
	goyaml "github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/token"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

func SaveFileYaml(yamlObject interface{}, filePath string) {
	// Save given YAML object to given filePath
	yamlData, err := yaml.Marshal(&yamlObject)
	if err != nil {
		panic(err.Error())
	}
	err = ioutil.WriteFile(filePath, yamlData, 0600)
	if err != nil {
		panic(err.Error())
	}
}

func LoadFileYaml(filePath string) (yamlObject interface{}, err error) {
	// Read/load/decode given filePath as some YAML object
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileContent, &yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlObject, nil
}

func PrintYaml(yamlObject interface{}) {
	yamlBytes, err := YamlObjToBytes(yamlObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(yamlBytes))
}

func LoadFileYamlBytes(filePath string) (yamlBytes []byte, err error) {
	// Load YAML file, including comments
	yamlBytes, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Check YAML formatting compliancy using "github.com/goccy/go-yaml"
	// which provides errors with line numbers
	var yamlObject interface{}
	err = goyaml.Unmarshal(yamlBytes, &yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlBytes, nil // We only care about returning the byte slice
}

func YamlObjToBytes(yamlObject interface{}) (yamlBytes []byte, err error) {
	// Convert YAML interface object to byte slice
	buffer := &bytes.Buffer{}
	encoder := yaml.NewEncoder(buffer)
	indent := 2
	encoder.SetIndent(indent)
	err = encoder.Encode(yamlObject)
	if err != nil {
		return nil, err
	}
	yamlBytes = buffer.Bytes()
	return yamlBytes, nil
}

func colorizeString(tk *token.Token, src string) string {
	str := Whi(src)
	switch tk.Type {
	case token.MappingKeyType:
		str = Blu(src)
	case token.StringType, token.SingleQuoteType, token.DoubleQuoteType:
		prev := tk.PreviousType()
		next := tk.NextType()
		if next == token.MappingValueType {
			str = Blu(src)
		} else if prev == token.AnchorType || prev == token.AliasType {
			str = Yel(src)
		} else {
			str = Gre(src)
		}
	case token.IntegerType, token.FloatType, token.BoolType:
		str = Mag(src)
	case token.AnchorType, token.AliasType:
		str = Yel(src)
	case token.CommentType:
		str = Gra(src)
	}
	return str
}

func PrintYamlColor(yamlObject interface{}) {
	// Colorized printout of proper YAML object, that don't usually include comments
	yamlBytes, err := YamlObjToBytes(yamlObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	PrintYamlBytesColor(yamlBytes)
}

func PrintYamlBytesColor(yamlBytes []byte) {
	//func PrintYamlBytesColor(yamlBytes []byte, allowLineNumber bool) {
	// Colorized printout of YAML byte slice which may include comments.
	// Caller must ensure yamlBytes is proper YAML.
	tokens := lexer.Tokenize(string(yamlBytes))
	if len(tokens) == 0 {
		return
	}
	printOut := []string{}
	//lineNumber := tokens[0].Position.Line
	for _, tk := range tokens {
		lines := strings.Split(tk.Origin, "\n")
		header := ""
		// if allowLineNumber {
		// 	header = fmt.Sprintf("%2d  ", lineNumber)
		// }
		if len(lines) == 1 {
			line := colorizeString(tk, lines[0])
			if len(printOut) == 0 {
				printOut = append(printOut, header+line)
				//lineNumber++
			} else {
				text := printOut[len(printOut)-1]
				printOut[len(printOut)-1] = text + line
			}
		} else {
			header := ""
			for idx, src := range lines {
				// if allowLineNumber {
				// 	header = fmt.Sprintf("%2d  ", lineNumber)
				// }
				line := colorizeString(tk, src)
				if idx == 0 {
					if len(printOut) == 0 {
						printOut = append(printOut, header+line)
						//lineNumber++
					} else {
						text := printOut[len(printOut)-1]
						printOut[len(printOut)-1] = text + line
					}
				} else {
					printOut = append(printOut, fmt.Sprintf("%s%s", header, line))
					//lineNumber++
				}
			}
		}
	}
	fmt.Println(strings.Join(printOut, "\n"))
}
