package intefaces

import (
	"fmt"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s \n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "Price of %s is %f \n", item, price)
	case "/update":
		item := req.URL.Query().Get("item")
		price, _ := strconv.Atoi(req.URL.Query().Get("price"))
		fmt.Printf("parsed price %d \n", price)
		db[item] = dollars(float64(price))
		fmt.Fprintf(w, "Updated item %s", item)
	}
}

func Server() {
	db := database{"socks": 10, "shoes": 100}
	err := http.ListenAndServe("localhost:8080", db)
	if err != nil {
		fmt.Println("Error is")
		fmt.Println(err)
	}
}

type Errno uintptr
