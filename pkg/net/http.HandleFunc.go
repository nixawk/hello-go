package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	html := `
        <html>
                <head>
                        <title>Chat</title>
                </head>
                <body>
                        Let's chat !
                </body>
        </html>
        `

	w.Write([]byte(html))
}

func main() {
	log.Println("Listening on 0.0.0.0:8080")

	http.HandleFunc("/", index)

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
