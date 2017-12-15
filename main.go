package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:id,omitempty`  
	Firstname string   `json:firstname,omitempty`    
	Lastname  string   `json:lastname,omitempty`    
	Address   *Address `json:address,omitempty`  
}

type Address struct {
	City  string  `json:city,omitempty`  
	State string `json:state,omitempty`  
}

var people []Person

func initData(){
	people = append(people, Person{ID:"1", Firstname:"jaruwit", Lastname:"suriyo", Address:&Address{City:"LA", State:"Calofornia"}})
	people = append(people, Person{ID:"2", Firstname:"wanatchapong", Lastname:"manothamsatit"})
	people = append(people, Person{ID:"3", Firstname:"anusit", Lastname:"maneerat"})
}

func createPeople(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var person  Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func getPeoples(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

func getPeople(w http.ResponseWriter, req *http.Request){
   params := mux.Vars(req)
   for _, item := range people {
	   if item.ID == params["id"] {
		   json.NewEncoder(w).Encode(item)
		   return
	   }
   }
   json.NewEncoder(w).Encode(&Person{})
}

func deletePeople(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
		  people = append(people[:index], people[index+1:]...)
		  return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main()  {
	initData()
	router := mux.NewRouter();
	router.HandleFunc("/peoples", createPeople).Methods("POST")
	router.HandleFunc("/peoples", getPeoples).Methods("GET")
	router.HandleFunc("/peoples/{id}", getPeople).Methods("GET")
	router.HandleFunc("/peoples/{id}", deletePeople).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}



