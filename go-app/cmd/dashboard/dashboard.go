package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/julienschmidt/httprouter"
)

var (
	inventoryHost = os.Getenv("INVENTORY_HOST")
	salesHost     = os.Getenv("SALES_HOST")
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	client := resty.New()

	respInventory, err := client.R().
		Get(fmt.Sprintf("%s/items", inventoryHost))

	if err != nil {
		log.Fatal(err)
	}
	respSales, err := client.R().
		Get(fmt.Sprintf("%s/invoices", salesHost))

	if err != nil {
		log.Fatal(err)
	}

	respMap := make(map[string]interface{})
	respMap["inventory"] = string(respInventory.Body())
	respMap["sales"] = string(respSales.Body())
	jsonResp, _ := json.Marshal(respMap)
	fmt.Fprint(w, string(jsonResp))

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/detail", GetDetail)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
