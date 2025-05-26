package routes

import (
	_ "fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func Home(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/home.html")

	if err != nil {
		panic(err)
	} else {
		template.Execute(res, nil)
	}
}

func Aboutus(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/aboutus.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(res, nil)
	}
}

func Params(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/params.html")
	if err != nil {
		panic(err)
	}

	params := mux.Vars(req)
	dataParams := map[string]string{
		"id":   params["id"],
		"slug": params["slug"],
	}

	template.Execute(res, dataParams)
}

func ParamsQueryString(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/paramsquerystring.html")
	if err != nil {
		panic(err)
	}

	query := req.URL.Query()
	dataParams := map[string]string{
		"id":   query.Get("id"),
		"slug": query.Get("slug"),
	}
	template.Execute(res, dataParams)
}

/* func Home(res http.ResponseWriter, req *http.Request) {
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
*/
