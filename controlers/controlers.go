package controlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BrunoTumelero/go-api/database"
	"github.com/BrunoTumelero/go-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	var p []models.Pesonalites
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(models.Persons)
}

func ReturnPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Pesonalites
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var personality models.Pesonalites
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Pesonalites
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Pesonalites
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
