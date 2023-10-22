package handler

import (
	"fmt"
	"net/http"
)

func MiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler middle start")
		handler.ServeHTTP(w, r)
		fmt.Println("Mdidleware Finished!")

	})
}
