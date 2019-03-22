package main

import (
	t "catcher/tree"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func computeResearch(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	q := r.URL.Query().Get("query")
	q = strings.ToLower(q)
	if len(q) > 1 {
		fmt.Println("Research : " + q)
		res := tree.SearchWord(q)
		if len(res) != 0 {
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
		// fmt.Printf("%d/%d\n", i, len(out)-1)
		// fmt.Printf("%s", strings.Split(out[i].DisplayName, " ")[0])
		if out[i].Level != "0.0" {
			q := strings.Split(out[i].DisplayName, " ")[0]
			q = strings.ToLower(q)
			// q = strings.ReplaceAll(q, " ", "-")
			tree.AddWord(out[i], q)
			tree.AddWord(out[i], out[i].Login)
		}

	}
	println("Indexing Done")
}

func main() {
	originsOk := handlers.AllowedOrigins([]string{"*"})
	// port := os.Getenv("PORT")
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
