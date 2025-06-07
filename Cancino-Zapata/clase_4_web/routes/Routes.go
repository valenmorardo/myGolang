package routes

import (
	"bytes"
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
var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTc0ODc0ODc2MiwiZXhwIjoxNzUxMzQwNzYyfQ.QSr4ZDpxSrD9AoaDEIuAAPaxy39T9VIi7oMfxHY4rgY"

type Categoria struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Slug   string `json:"slug"`
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

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	datos := Categorias{}

	errJson := json.Unmarshal(responseBody, &datos)
	if errJson != nil {
		panic(errJson)
	}

	fmt.Fprintln(res, "LISTA DE CATEGORIAS:")
	for _, v := range datos {
		fmt.Fprintln(res, v)
	}
}

func ClienteHttp_PostCategorie(res http.ResponseWriter, req *http.Request) {
	// Creo mi categoria
	miCategoria := Categoria{Nombre: "Voy a ser el mejor programador de Golang! Hi Gopher"}

	// a la categoria que esta en un tipo struct lo paso a json para usarlo como body para la req
	jsonData, err := json.Marshal(miCategoria)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes.NewBuffer(jsonData))
	// Creo la request POST con el json creado antes y seteo los headers
	reqPost, err := http.NewRequest("POST", "https://www.api.tamila.cl/api/categorias", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	reqPost.Header.Set("Authorization", Token)
	reqPost.Header.Set("Content-Type", "application/json")

	// Creo el cliente para que haga la req
	cliente := &http.Client{}
	response, err := cliente.Do(reqPost)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyResponse, _ := io.ReadAll(response.Body)
	fmt.Fprintln(res, "Categoría creada correctamente:")
	fmt.Fprintln(res, string(bodyResponse))
}

func ClienteHttp_EditCategorie(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	idToEdit := params["id"]
	newCategorie := Categoria{Nombre: "Mi categoria Valentin modificada 312"}

	jsonData, err := json.Marshal(newCategorie)
	if err != nil {
		panic(err)
	}
	fmt.Println(idToEdit)
	fmt.Println(bytes.NewBuffer(jsonData))
	reqPut, err := http.NewRequest("PUT", "https://www.api.tamila.cl/api/categorias/"+idToEdit, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	reqPut.Header.Set("Authorization", Token)
	reqPut.Header.Set("Content-Type", "application/json")

	cliente := &http.Client{}
	response, err := cliente.Do(reqPut)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	fmt.Fprintln(res, "Categoria modifcada exitosamente!")
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Respuesta del servidor:", string(responseBody))
}

func ClienteHttp_DeleteCategorie(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idToDelete := params["id"]

	reqDel, err := http.NewRequest("DELETE", "https://www.api.tamila.cl/api/categorias/"+idToDelete, nil)
	if err != nil {
		panic(err)
	}
	reqDel.Header.Set("Authorization", Token)
	reqDel.Header.Set("Content-Type", "application/json")

	cliente := &http.Client{}
	response, err := cliente.Do(reqDel)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))
	fmt.Fprintf(res, "Categoria con ID: %v eliminada correctamente\n", idToDelete)
	res.Write(responseBody)
}
