package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	mux.HandleFunc("/params", routes.Params)
	mux.HandleFunc("/params/{id:[0-9]+}/{slug:[a-zA-Z0-9-]+}", routes.Params)
	mux.HandleFunc("/params-querystring", routes.ParamsQueryString)
	mux.HandleFunc("/estructuras", routes.Estructuras)


	// ruta para hace request a otro endpoint
	mux.HandleFunc("/cliente-http/get-categories", routes.ClienteHttp_GetCategories).Methods("GET")
	mux.HandleFunc("/cliente-http/post-categorie", routes.ClienteHttp_PostCategorie)
	mux.HandleFunc("/cliente-http/edit-categorie/{id:[0-9]+}", routes.ClienteHttp_EditCategorie)
	mux.HandleFunc("/cliente-http/delete-categorie/{id:[0-9]+}", routes.ClienteHttp_DeleteCategorie)

	envData := utils.GetEnvData()

	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", envData.SERVER, envData.SERVER_PORT),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Ejecutamos el servidor en una goroutine
	go func() {
		log.Printf("Server up --> http://%v:%v\n", envData.SERVER, envData.SERVER_PORT)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperamos Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // Queda "parado" acá hasta que presiones Ctrl+C

	log.Println("Apagando el servidor...")

	// Contexto con timeout para shutdown elegante
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error al apagar el servidor: %v", err)
	}

	log.Println("Servidor apagado correctamente")
}
