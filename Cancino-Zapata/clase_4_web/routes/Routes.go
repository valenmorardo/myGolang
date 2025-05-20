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

func ParamsQueryString(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Println(req.URL.RawQuery)
	fmt.Println(req.URL.Query())

	id := req.URL.Query().Get("id")
	fmt.Fprintf(res, "ID: %v\n", id)
	slug := req.URL.Query().Get("slug")
	fmt.Fprintf(res, "Slug: %v\n", slug)
}
