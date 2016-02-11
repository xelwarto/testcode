// https://github.com/thbar/golang-playground/blob/master/download-files.go


package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
  fileName := "epel-release-latest-7.noarch.rpm"
  output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

  res, err := http.Get("https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm")
  if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
  defer res.Body.Close()

  _, err = io.Copy(output, res.Body)
	if err != nil {
		fmt.Println("Error while downloading", err)
		return
	}

  fmt.Println(res.Status)
  fmt.Println(res.StatusCode)
}
