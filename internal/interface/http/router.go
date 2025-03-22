package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func PublicRouter(namesHandler *NamesHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	r.HandleFunc("/generate/names", namesHandler.ServeHTTP).Methods("POST")
	return r
}

func AuthRouter(authHandler *AuthHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/signup", authHandler.Signup).Methods("POST")

	return r
}
