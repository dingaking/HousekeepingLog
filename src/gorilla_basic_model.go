////////////////////////////////////////////////////////////////////////
// A Basic Model
// Now that we have routes in place, it's time to create a basic Todo model that we can send and retrieve data with. In Go, a struct will typically serve as your model. Many other languages use classes for this purpose.
////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

////////////////////////////////////////////////////////////////////////
// Note that in the last line we create another type, called Todos, which is a slice(anordered collection) of Todo. You will see where this becomes useful shortly.
////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////
// Send Back Some JSON
// Now that we have a basic model, we can simulate a real response and mock out the TodoIndex with static data.
////////////////////////////////////////////////////////////////////////

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

////////////////////////////////////////////////////////////////////////
// For now, we are just creating a static slice of Todos to send back to the client. Now if you request http://localhost:8080/todos, you should get the following rmresponse
////////////////////////////////////////////////////////////////////////

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8082", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcom!")
}

/*func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}*/

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show : ", todoId)
}
