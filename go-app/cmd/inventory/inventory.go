package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func GetItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	items := []Item{{ID: "1", Name: "item1", Price: "100"}, {ID: "2", Name: "item2", Price: "200"}, {ID: "3", Name: "item3", Price: "300"}}
	itemJson, _ := json.Marshal(items)
	fmt.Fprint(w, string(itemJson))
}

func main() {
	router := httprouter.New()
	router.GET("/items", GetItems)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
