package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Compo struct {
	ID          bson.ObjectId `bson:"_id" json:"id"` // ID of component
 	Name        string        `bson:"name" json:"name"` // Name of component
	CoverImage  string        `bson:"cover_image" json:"cover_image"` // Image of component
	Description string        `bson:"description" json:"description"` // Description of component
}

type ComponentDAQ {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "components"
)

func (m *ComponentDAQ) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func AllComponent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindComponent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func AddComponent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateComponent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteComponent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
