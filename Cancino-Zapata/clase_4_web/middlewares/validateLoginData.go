package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"clase_4_web/models"
	"clase_4_web/utils"
)


type contextKey string

const LoginInputKey contextKey = "loginInput"

func ValidateLoginData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("SOY MIDDLEWARE")
		var input models.UserLogin
		
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil || input.Correo == "" || input.Password == "" {
			utils.SendResponse(res, http.StatusBadRequest, false, "ERROR: Todos los campos son obligatorios.", nil, "Error.")
			return
		}
	
		ctx := context.WithValue(req.Context(), LoginInputKey, input)
		req = req.WithContext(ctx)
		
		fmt.Println("PASE EL MIDDLEWARE")
		next.ServeHTTP(res, req)
	})
}
