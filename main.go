package main

import (
	t "catcher/tree"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func computeResearch(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	q := r.URL.Query().Get("query")
	if len(q) > 0 {
		fmt.Println("Research : " + q)
		res := tree.SearchWord(q)
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

	file, err := ioutil.ReadFile("./to_index.json")
	if err != nil {
		log.Fatal(err)
	}
	var out []t.Student
	json.Unmarshal(file, &out)
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d/%d\n", i, len(out)-1)
		tree.AddWord(out[i])
	}
	println("Indexing Done")
}

func main() {
	originsOk := handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	tree := t.NewTree()
	indexDico(tree)
	r := mux.NewRouter()
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		computeResearch(w, r, tree)
	})

	err := http.ListenAndServe(port, handlers.CORS(originsOk)(r))
	if err != nil {
		panic(err)
	}
}
