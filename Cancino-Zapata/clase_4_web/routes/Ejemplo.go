package routes

import (
	"fmt"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hola mundoooooooo")
}


func Aboutus(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "About Us")
}
