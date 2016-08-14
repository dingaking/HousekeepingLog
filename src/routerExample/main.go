////////////////////////////////////////////////////////////////////////
// OK, We Need to Split This Up!
// At this point, the project needs a little refactoring. We have too much going on in just a few files.
// We are going to now create the following files and move the code around accordingly:
// - main.go
// - hadlers.go
// - routes.go
// - todo.go
// http://thenewstack.io/make-a-restful-json-api-go/
// https://github.com/corylanou/tns-restful-json-api
////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8082", router))
}
