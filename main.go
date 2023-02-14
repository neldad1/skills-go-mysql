package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"portfolio.com/skills"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/listSkill", skills.ListSkills).Methods("GET")
	router.HandleFunc("/addSkill", skills.AddSkill).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
