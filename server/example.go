// run "go mod init github.com/ndrean/apicall" to create a new module
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// need to Uppercase to export the fields
// otherwise, it will be private and we can't print them
type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func fetchTodos() ([]todo, error) {
	url := "https://jsonplaceholder.typicode.com/todos?_limit=5"
	res, err := http.Get(url)
	if err != nil {
		// log.Fatalf("Error: %s", err)
		return nil, fmt.Errorf("error: %s", err.Error())

	}
	defer res.Body.Close()

	todos := []todo{}
	err = json.NewDecoder(res.Body).Decode(&todos)

	if err != nil {
		return nil, fmt.Errorf("error when parsing server response: %s", err.Error())
	}
	fmt.Println("Todos: ", todos)
	if len(todos) > 0 {
		// range returns the index and the value
		for id, todo := range todos {
			fmt.Printf("Todo: %v %v \n", id, todo)
		}
	}

	return todos, nil
}

// you need to create a channel to cmunicate with goroutines
func Run() {
	// server is a handler, run in a goroutine
	http.HandleFunc("/", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func server(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	todos, err := fetchTodos()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// write the response
	res.Write([]byte(fmt.Sprintf("Todos: %v", todos)))
	// res.Header().Set("Content-Type", "application/json")

}
