package main

import (
	//"TestZD/internal/Person"
	"TestZD/internal/app/service"
	"TestZD/internal/pkg/app/postgresql"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type App struct {
	echo *echo.Echo
}

func main() {

	//db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	router := mux.NewRouter()
	//a.echo.GET("/people", People)
	router.HandleFunc("/people", People).Methods("GET")
	//router.HandleFunc("/people/{id}", repository.OnePerson(context.TODO(), )).Methods("GET")
	//router.HandleFunc("/people", env.createPerson).Methods("POST")
	//router.HandleFunc("/people/{id}", env.updatePerson).Methods("PUT")
	//router.HandleFunc("/people/{id}", env.deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func People(w http.ResponseWriter, r *http.Request) {
	postgresqlClient, err := postgresql.NewClient(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	repository := service.NewRepository(postgresqlClient)
	p, _ := repository.People(context.TODO())
	fmt.Print("s")
	log.Fatal("%s", p)

}

/*
func (env *Env) getPeople(w http.ResponseWriter, r *http.Request) {
	// Get people from the database and return as JSON
	rows, err := env.db.Query("SELECT * FROM person")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	people := []Person{}
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}
	json.NewEncoder(w).Encode(people)
}
func (env *Env) getPerson(w http.ResponseWriter, r *http.Request) {
	// Get person by ID from the database and return as JSON
	params := mux.Vars(r)
	id := params["id"]
	person := Person{}
	err := env.db.QueryRow("SELECT * FROM people WHERE id=$1", id).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(person)
}
func (env *Env) createPerson(w http.ResponseWriter, r *http.Request) {
	// Create a new person in the database
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	// Enrich the person data
	person.Age = getAge(person.Name)
	person.Gender = getGender(person.Name)
	person.Nationality = getNationality(person.Name)
	// Insert the person into the database
	_, err := env.db.Exec("INSERT INTO people
								(name, surname, patronymic, age, gender, nationality)
							VALUES ($1, $2, $3, $4, $5, $6)
							",person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(person)
}
func (env *Env) updatePerson(w http.ResponseWriter, r *http.Request) {
	// Update an existing person in the database
	params := mux.Vars(r)
	id := params["id"]
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	// Enrich the person data
	person.Age = getAge(person.Name)
	person.Gender = getGender(person.Name)
	person.Nationality = getNationality(person.Name)
	// Update the person in the database
	_, err := env.db.Exec("UPDATE people SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7", person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(person)
}
func (env *Env) deletePerson(w http.ResponseWriter, r *http.Request) {
	// Delete a person from the database
	params := mux.Vars(r)
	id := params["id"]
	// Delete the person from the database
	_, err := env.db.Exec("DELETE FROM people WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func getAge(name string) int {
	// Logic to determine age based on name
	return 30
}
func getGender(name string) string {
	// Logic to determine gender based on name
	return "Male"
}
func getNationality(name string) string {
	// Logic to determine nationality based on name
	return "American"
}
*/
