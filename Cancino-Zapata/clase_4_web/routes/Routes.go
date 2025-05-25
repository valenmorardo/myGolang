package routes

import (
	_ "fmt"
	"html/template"
	"net/http"

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
	} else {
		template.Execute(res, nil)
	}
}

func ParamsQueryString(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/paramsquerystring.html") 
	if err != nil {
		panic(err)
	} else {
		template.Execute(res, nil)
	}
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