package web_server

import (
	"fmt"
	"log"
	"net/http"
)

func Handle() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s, %s, %s \n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "header: %s -> %s \n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form: %s -> %s \n", k, v)
	}
}
