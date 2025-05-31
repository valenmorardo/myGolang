package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
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

type Datos struct {
	Nombre      string
	Apellido    string
	Edad        int
	Habilidades []Habilidad
}
type Habilidad struct {
	Nombre string
}

func Estructuras(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/estructuras.html")
	if err != nil {
		panic(err)
	}
	Habilidad1 := Habilidad{"Videojuegos"}
	Habilidad2 := Habilidad{"Programacion"}

	valentin := Datos{
		Nombre:      "Valentin",
		Apellido:    "Morardo",
		Edad:        2,
		Habilidades: []Habilidad{Habilidad1, Habilidad2},
	}

	template.Execute(res, valentin)
}

//////////////////////////////////////////////////////////////////////
// ruta para req a otro endpoint
// endpoint: www.api.tamila.cl/
/*
www.api.tamila.cl/api/login
{
"correo": "info@tamila.cl",
"password": "p2gHNiENUw"
}
*/
// token de inicio de sesion. Reemplazarlo las veces que haga falta loggeando con los datos de arriba
var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTc0ODY1NjgzNiwiZXhwIjoxNzUxMjQ4ODM2fQ.PUYUmxdB9VoAsF9GBmiRDqEiM3DgyGWIKw32_5Ezies"

type Categoria struct {
	Id int
	Nombre string
	Slug string
	
}
type Categorias []Categoria

func ClienteHttp_GetCategories(res http.ResponseWriter, req *http.Request) {
	cliente := &http.Client{}
	newReq, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias", nil)
	if err != nil {
		panic(err)
	}

	newReq.Header.Set("Authorization", Token)
	response, err := cliente.Do(newReq)
	if err != nil {
		panic(err)
	}
	
	defer response.Body.Close()
	fmt.Println(response.Status)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	datos := Categorias{}
	errJson := json.Unmarshal(responseBody, &datos)
	if errJson != nil {
		panic(errJson)
	}
	for _, v := range datos {
		fmt.Println(v)
	}
}
