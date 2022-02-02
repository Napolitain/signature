package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{} // use default options

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get querystring value (iFrame src)
	querystring := r.URL.Query().Get("URL")
	if querystring == "" {
		log.Fatal("No query string provided")
		return
	}

	// Inject iFrame with querystring as src
	//body := template.New("body").Parse("")
	_, _ = fmt.Fprintf(w, "<iframe src=\"https://docs.google.com/presentation/d/e/"+querystring+"/embed?start=false\" frameborder=\"0\" width=\"960\" height=\"569\" allowfullscreen=\"true\" mozallowfullscreen=\"true\" webkitallowfullscreen=\"true\"></iframe>")
	/*
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer connection.Close()
		for {
			mt, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = connection.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	*/
}
