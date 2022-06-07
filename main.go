package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Word struct {
	ID string "json:id"
	Description string "json:description"
	CreateDate string "json:createDate"
}

type EnglishDictioary struct {
	ID string "json:id"
	Description string "json:id"
}

var words [] Word

func main(){
	fmt.Println("Merhaba")

	r:= mux.NewRouter()

	r.HandleFunc("/words", getWords).Methods("GET")
	r.HandleFunc("/words/{id}", getSpecificWord).Methods("GET")
	r.HandleFunc("/words", createWord).Methods("GET")
	r.HandleFunc("/words/{id}", updateWord).Methods("PUT")
	r.HandleFunc("/words/{id}", deleteWord).Methods("DELETE")

	log.Fatal(http.ListenAndServe(" :8000", r))
}

func getWords(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func deleteWord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range words {

		if item.ID == params["id"]{
			words = append(words[:index], words[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(words)
}

func getSpecificWord(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		for _, item := range words {
			if item.ID == params["id"]{
				json.NewEncoder(w).Encode(item)
				return
			}
		}
}

func createWord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var word Word 
	_ = json.NewDecoder(r.Body).Decode(&word)

	word.ID = strconv.Itoa((rand.Intn(10000000)))

	words = append(words, word)
	json.NewEncoder(w).Encode(word)

}

func updateWord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range words {
		if item.ID == params["id"]{
			words = append(words[:index], words[index+1:]...)
			var word Word 
			_ = json.NewDecoder(r.Body).Decode(&word)
			word.ID = params["id"]
			words = append(words, word)
			json.NewEncoder(w).Encode(word)
			
			return
		}
	}

}

func appendData(){
	words =append(words, Word{ID: "1", Description:"Apple", CreateDate:"20220609"})
	words =append(words, Word{ID: "1", Description:"Cucumber", CreateDate:"20220609"})
}

