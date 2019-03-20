package main

import (
	"encoding/json"
    "io/ioutil"
	"fmt"
	"log"
	"net/http"
	t "rise_perso/tree"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func computeResearch(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	q := r.URL.Query().Get("query")
	if len(q) > 0 {
		fmt.Println("Research : " + q)
		res := tree.SearchWord(q)
		fmt.Println(len(res))
		if len(res) != 0 {
			fmt.Println("in")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
		}
	}
}

func indexDico(tree t.Tree) {

	file, err := ioutil.ReadFile("./good.json")
	if err != nil {
		log.Fatal(err)
	}
	var out []t.Student
	json.Unmarshal(file, &out)
	for i := 0; i < len(out); i++{
		fmt.Println(out[i])
		tree.AddWord(out[i])
	} 
	println("Indexing Done")
}

func main() {
	originsOk := handlers.AllowedOrigins([]string{"*"})

	tree := t.NewTree()
	indexDico(tree)
	r := mux.NewRouter()
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		computeResearch(w, r, tree)
	})

	err := http.ListenAndServe(":8080", handlers.CORS(originsOk)(r))
	if err != nil {
		panic(err)
	}
}
