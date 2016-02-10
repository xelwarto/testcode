// http://blog.charmes.net/2015/07/reverse-proxy-in-go.html
// https://golang.org/pkg/net/http/httputil/

package main

import (
        "net/http"
        "net/http/httputil"
        "net/url"
)

func main() {
        proxy := httputil.NewSingleHostReverseProxy(&url.URL{
                Scheme: "http",
                Host:   "localhost:9091",
        })
        http.ListenAndServe(":9090", proxy)
}
