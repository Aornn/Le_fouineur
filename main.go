package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	t "rise_perso/tree"

	"github.com/gorilla/mux"
)

func computeResearch(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	q := r.URL.Query().Get("query")
	fmt.Println("Research : " + q)
	res := tree.SearchWord(q)
	if res != nil {
		w.WriteHeader(http.StatusOK)
		for i := 0; i < len(res); i++ {
			fmt.Fprintf(w, res[i]+"\n")
		}
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Mot Introuvable")
	}
}

func indexDico(tree t.Tree) {

	file, err := os.Open("liste_francais.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tree.AddWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println("Indexing Done")
}

func main() {
	tree := t.NewTree()
	indexDico(tree)
	r := mux.NewRouter()
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		computeResearch(w, r, tree)
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
