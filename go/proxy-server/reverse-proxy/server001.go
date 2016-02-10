package main

import (
        "net/http"
        "io"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
                println("--->", req.RemoteAddr)
                io.WriteString(w, "hello, world!\n")
        })
        http.ListenAndServe(":9091", nil)
}
