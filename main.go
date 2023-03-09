package main

import (
	"log"
	"net/http"

	"github.com/YeisonMolano/primer_web_go/internal/user"
	"github.com/YeisonMolano/primer_web_go/pkg/bootstrap"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()
	loger := bootstrap.InitLogger()

	// Conexion a la base de datos
	db, err := bootstrap.DbConnection()
	if err != nil {
		loger.Fatal(err)
	}

	userRepo := user.NewRepo(loger, db)
	userService := user.NewService(loger, userRepo)
	userEnd := user.MakeEndpoints(userService)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Updte).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
