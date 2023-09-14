package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ApiResponse struct {
	Result bool   `json:"Result"`
	Data   string `json:"Data"`
}

// handler functions business logic k lye use honge. hm ne 2 hander bnae hue hen get set k lye
// yourHandlers will need a writer object (to write a response to the client)
// and a request object (to obtain information about the incoming request)
func YourHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//you can use the Encode method of the Encoder struct returned by NewEncoder to write
	//a struct that will be encoded as JSON to the client.
	json.NewEncoder(w).Encode(ApiResponse{Result: true, Data: "myname"})

}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	t := r.URL.Query().Get("t")
	t = t + " " + time.Now().String()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//you can use the Encode method of the Encoder struct returned by NewEncoder to write
	//a struct that will be encoded as JSON to the client.
	json.NewEncoder(w).Encode(ApiResponse{Result: true, Data: t})
	//w.Write([]byte("Gorilla!\n"))
}

func main() {
	// create router with NewRouter method and assign it to instance r
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	//after declaring a new router instance, you can use the HandleFunc method of your router
	//instance to assign routes to handler functions along with the request type that the
	//handler function handles.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/test", TestHandler)

	//http.ListenAndServe() function to start the server and tell it to listen for
	//new HTTP requests and
	//then serve them using the handler functions you set up
	//You can set up a server using the ListenAndServe method of the http package.
	//The ListenAndServe method takes as arguments the port you want the server to run on
	//and a router instance
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
