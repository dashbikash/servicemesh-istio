package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Invoice struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Date   int    `json:"date"`
}

func GetInvoices(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	invoices := []Invoice{{ID: "1", Name: "invoice1", Amount: 100, Date: 20210101}, {ID: "2", Name: "invoice2", Amount: 200, Date: 20210102}, {ID: "3", Name: "invoice3", Amount: 300, Date: 20210103}, {ID: "4", Name: "invoice4", Amount: 400, Date: 20210104}}
	invoiceJson, _ := json.Marshal(invoices)
	fmt.Fprint(w, string(invoiceJson))
}

func main() {
	router := httprouter.New()
	router.GET("/invoices", GetInvoices)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
