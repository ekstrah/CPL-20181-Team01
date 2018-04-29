package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server Struct (Model)
type dummyData struct {
	Temperature int `json: "temp"`
	NumOfPpl    int `json: "ppl"`
	Humidity    int `json: "humid"`
}

//init serverPC
var dataBuff []dummyData

func extractPacket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Server dummyData
	_ = json.NewDecoder(r.Body).Decode(&Server)
	fmt.Println(Server.Temperature)
	fmt.Println(Server.NumOfPpl)
	fmt.Println(Server.Humidity)

	for index, item := range dataBuff {
		if Server.NumOfPpl == item.NumOfPpl {
			dataBuff = append(dataBuff[:index], dataBuff[index+1:]...)
			break
		}
	}
	dataBuff = append(dataBuff, Server)

	///Save to File
	fmt.Println("Success")
}

//Get All ServerList
func getServerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataBuff)
}

func main() {
	r := mux.NewRouter()

	//Mocking the data
	r.HandleFunc("/local", getServerList).Methods("GET")
	r.HandleFunc("/local", extractPacket).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
