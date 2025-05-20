package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hola mundoooooooo")
}


func Aboutus(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "About Us")
}


func Params(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	fmt.Fprintf(res, "ID: %v | SLUG: %v", params["id"], params["slug"])
}
