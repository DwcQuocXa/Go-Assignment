package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	PersonalCode uuid.UUID `json:"personalCode"`
}

var personList []Person

//Get all persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(personList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Get a person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	// find one with the code from the params
	for _, item := range personList {
		codeParams, _ := uuid.FromString(params["personal-code"])
		if item.PersonalCode == codeParams {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

//Create a person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPerson Person

	newPerson.PersonalCode = uuid.NewV4()
	personList = append(personList, newPerson)

	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for index, item := range personList {
		codeParams, _ := uuid.FromString(params["personal-code"])
		if item.PersonalCode == codeParams {
			personList = append(personList[:index], personList[index+1:]...)
			var person Person

			person.PersonalCode = codeParams
			personList = append(personList, person)

			if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for index, item := range personList {
		codeParams, _ := uuid.FromString(params["personal-code"])
		if item.PersonalCode == codeParams {
			personList = append(personList[:index], personList[index+1:]...)
			break
		}
	}
}

func main() {
	//Init router
	r := mux.NewRouter()

	//First data
	personList = append(personList, Person{
		FirstName:    "John",
		LastName:     "Doe",
		PersonalCode: uuid.NewV4()})

	//Routes
	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons/{personal-code}", getPerson).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/{personal-code}", updatePerson).Methods("GET")
	r.HandleFunc("/persons/{personal-code}", deletePerson).Methods("GET")

	//Server and port
	log.Fatal(http.ListenAndServe(":8080", r))
}
