
package main

import (
	"flag"
  "fmt"
	"os"
  "io/ioutil"
  "encoding/json"
  "text/template"
)

var tmplFile string
var jsonFile string
var jsonMap interface{}
var jsonData interface{}
var envParam string

func init() {
  flag.StringVar(&tmplFile, "tmpl", "", "Template File")
  flag.StringVar(&jsonFile, "json", "", "JSON Data File")
	flag.StringVar(&envParam, "env", "", "Environment Parameter")
  flag.Parse()
}

func main() {
	file, err := ioutil.ReadFile(jsonFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading JSON file: %v\n", err)
    os.Exit(1)
  }

  err = json.Unmarshal(file, &jsonMap)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error parsing JSON file: %v\n", err)
    os.Exit(1)
  }
	jsonData := jsonMap.(map[string]interface{})

	if envParam != "" {
		jsonData = jsonData[envParam].(map[string]interface{})
	}

  tmpl, err := template.ParseFiles(tmplFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
    os.Exit(1)
  }
  err = tmpl.Execute(os.Stdout, jsonData)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
    os.Exit(1)
  }
}
