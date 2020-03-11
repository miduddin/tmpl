package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	templateFile := os.Args[1]
	inputFile := os.Args[2]

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	inputMap := map[string]interface{}{}
	yaml.Unmarshal(inputData, &inputMap)

	template.Must(template.ParseFiles(templateFile)).Execute(os.Stdout, inputMap)
}
