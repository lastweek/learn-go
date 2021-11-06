/*
 * I have to say Go's packages are so nice to use.
 * Those packages.. just great!
 *
 * I love the fact that it is still close to C.
 * But not as close as to Scale type of lanugages.
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var port = 8080

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
	dump, err := httputil.DumpRequest(r, false)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%q", dump)
}

func main() {
	http.HandleFunc("/", testHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
