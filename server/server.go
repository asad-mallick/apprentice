package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func nextQuote(w http.ResponseWriter, req *http.Request) {
	q := quote.Go()
	fmt.Println("Quotation server reply is: ", q)
	w.Write([]byte(q))
}

func main() {
	http.HandleFunc("/quote", nextQuote)
	fmt.Println("Quotation server started on port 5555")
	http.ListenAndServe(":5555", nil)
}
