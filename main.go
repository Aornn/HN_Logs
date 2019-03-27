package main

import (
	t "catcher/tree"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func computeResearchTop(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	var newQ []int

	params := mux.Vars(r)
	q := params["date"]
	size, _ := strconv.Atoi(params["size"])
	q = strings.ReplaceAll(q, "-", " ")
	q = strings.ReplaceAll(q, ":", " ")
	tempQuery := strings.Split(q, " ")
	for i := 0; i < len(tempQuery); i++ {
		val, _ := strconv.Atoi(tempQuery[i])
		newQ = append(newQ, val)
	}
	res := tree.CountDataTop(newQ, size)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func computeResearch(w http.ResponseWriter, r *http.Request, tree t.Tree) {
	var newQ []int

	params := mux.Vars(r)
	q := params["date"]
	size, _ := strconv.Atoi(params["size"])
	q = strings.ReplaceAll(q, "-", " ")
	q = strings.ReplaceAll(q, ":", " ")
	tempQuery := strings.Split(q, " ")
	for i := 0; i < len(tempQuery); i++ {
		val, _ := strconv.Atoi(tempQuery[i])
		newQ = append(newQ, val)
	}
	res := tree.CountData(newQ, size)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func main() {
	originsOk := handlers.AllowedOrigins([]string{"*"})
	tree := t.NewTree()
	t.IndexFile(tree, "./hn_logs.tsv")
	r := mux.NewRouter()
	r.HandleFunc("/top/{date}/{size}", func(w http.ResponseWriter, r *http.Request) {
		computeResearchTop(w, r, tree)
	})
	r.HandleFunc("/count/{date}", func(w http.ResponseWriter, r *http.Request) {
		computeResearch(w, r, tree)
	})

	err := http.ListenAndServe(":8080", handlers.CORS(originsOk)(r))
	if err != nil {
		panic(err)
	}
}
