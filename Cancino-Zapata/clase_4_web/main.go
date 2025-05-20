package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"clase_4_web/routes"

	"clase_4_web/utils"

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
	mux.HandleFunc("/about-us", routes.Aboutus)
	mux.HandleFunc("/params/{id:[0-9]+}/{slug:[a-zA-Z0-9-]+}", routes.Params)
	mux.HandleFunc("/params-querystring", routes.ParamsQueryString)

	envData := utils.GetEnvData()

	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", envData.SERVER, envData.SERVER_PORT),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server up --> localhost:3001")
	log.Fatal(server.ListenAndServe())
}
