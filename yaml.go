// yaml.go

package utl

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

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

func PrintYaml(yamlObject interface{}) {
	pretty, err := yaml.Marshal(&yamlObject)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf(string(pretty))
	}
}
