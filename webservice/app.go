package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ah := authorHandler{newHardCodedAuthorService()}

	r := mux.NewRouter()
	r.HandleFunc("/authors/{uuid}", ah.getAuthor).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type authorHandler struct {
	as authorService
}

func (ah *authorHandler) getAuthor(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid := vars["uuid"]

	requestedAuthor, found := ah.as.getAuthorByUUID(uuid)
	if !found {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(writer)
	//should have error handling here
	enc.Encode(requestedAuthor)
}
