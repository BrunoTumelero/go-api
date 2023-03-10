package routes

import (
	"log"
	"net/http"

	"github.com/BrunoTumelero/go-api/controlers"
	"github.com/BrunoTumelero/go-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controlers.Home)
	r.HandleFunc("/persons", controlers.AllPersons).Methods("Get")
	r.HandleFunc("/persons/{id}", controlers.ReturnPerson).Methods("Get")
	r.HandleFunc("/persons/{id}", controlers.DeletePerson).Methods("Delete")
	r.HandleFunc("/persons/{id}", controlers.EditPerson).Methods("Put")
	r.HandleFunc("/persons/", controlers.CreatePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
