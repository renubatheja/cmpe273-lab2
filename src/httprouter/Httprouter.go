package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"    
)

/*
* Request and Response structs
*/

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Greeting string `json:"greeting"` 
}


/*
* hello function - Function to service GET REST request
*/

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

}

/*
* createGreeting function - Function to service POST REST request
*/

func createGreeting(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    // Stub a request to be populated from the body
    request := Request{}
    response := Response{}

    // Populate the request data
    json.NewDecoder(req.Body).Decode(&request)

	fmt.Println("The name passed in request json : ",request.Name)
	
	response.Greeting = "Hello, " + request.Name + "!"
	
	// Marshal provided interface into JSON structure
	responseJson, _ := json.Marshal(response)
	
	// Write content-type, statuscode, payload
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", responseJson)
}


/*
* Main function - Setting up of httprouter and REST handlers
*/

func main() {
    mux := httprouter.New()
    
    //handler to service GET request
    mux.GET("/hello/:name", hello)
    
    //handler to service POST request
    mux.POST("/hello", createGreeting)
    
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}