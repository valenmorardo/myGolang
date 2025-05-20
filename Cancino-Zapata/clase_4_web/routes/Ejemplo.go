package routes

import (
	"fmt"
	"net/http"
)



func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hola mundi")
}