// run "go mod init api" to create a new module
package api

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
	todoCh := make(chan []todo)
	errCh := make(chan error)
	// auto-execute the goroutine with ()
	go func() {
		todos, err := fetchTodos()
		if err != nil {
			errCh <- err
			return
		}
		todoCh <- todos
	}()

	fmt.Printf("Fetching todos... \n")
	defer close(todoCh)
	defer close(errCh)
	// you need to listen to the channel to get the data
	select {
	case todos := <-todoCh:
		fmt.Println("Todos: ", todos)
	case err := <-errCh:
		fmt.Println("Error: ", err)
	}
}
