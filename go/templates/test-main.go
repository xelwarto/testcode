
package main

import (
	"flag"
  "fmt"
	"os"
  "io/ioutil"
  "encoding/json"
  "text/template"
)

var fileName string
var dataFile string

func init() {
  flag.StringVar(&fileName, "file", "test", "Input file")
  flag.StringVar(&dataFile, "data", "test2", "Data file")
  flag.Parse()
}

func main() {
  file, err := ioutil.ReadFile(dataFile)
  if err != nil {
    fmt.Printf("File error: %v\n", err)
    os.Exit(1)
  }

  var f interface{}
  json.Unmarshal(file, &f)

	m := f.(map[string]interface{})

  tmpl, err := template.ParseFiles(fileName)
  if err != nil { panic(err) }
  err = tmpl.Execute(os.Stdout, m["gce"])
  if err != nil { panic(err) }
}
