package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// export GOOGLE_APPLICATION_CREDENTIALS="/Users/tengzhou/Desktop/teng-review-firebase-adminsdk-cuh7g-e94895a080.json"
func main() {
	var router = mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Up and running...")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
