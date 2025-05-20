package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"clase_4_web/routes"

	"github.com/gorilla/mux"
)

/* func main() {


	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hola mundo")
	})

	fmt.Println("Server up --> localhost:3001")
	log.Fatal(http.ListenAndServe("localhost:3001", nil))
} */

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.Home)

	server := &http.Server{
		Addr:         "localhost:3001",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server up --> localhost:3001")
	log.Fatal(server.ListenAndServe())
}
