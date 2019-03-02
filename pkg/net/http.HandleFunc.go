package main

import "io"
import "fmt"
import "net/http"

func hello(res http.ResponseWriter, req *http.Request) {
	msg := `<doctype html>
        <html>
                <head><title>Hello World</title></head>
                <body>Hello World</body>
        </html>`
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, msg)

}

func main() {
	fmt.Println("[*] http server is listening on :9999")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
