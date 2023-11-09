package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const baseurl = "https://jsonplaceholder.typicode.com/todos/%d"

func main() {
	t := time.Now()
	ch1 := make(chan *ToDo)

	defer func() {
		fmt.Println("Total tim spent ", time.Since(t))
	}()
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	mapRes := map[int]*ToDo{}

	for _, v := range n {
		//spawn different go routine for all api
		//wg.Add(1)
		go func(id int) {
			todo, e := CallTodoApi(id)
			if e != nil {

			}
			ch1 <- todo
			//wg.Done()
		}(v)

	}
	//wg.Wait()

	fmt.Println("All routines fired ")
	for x := 0; x < len(n); x++ {
		todo := <-ch1
		mapRes[todo.ID] = todo
	}
	fmt.Println("Parsing map ")

	for i, v := range mapRes {
		fmt.Println(i, v)
	}

}

func CallTodoApi(id int) (*ToDo, error) {

	urlst := fmt.Sprintf(baseurl, id)
	fmt.Println("URL called ", urlst)
	r, e := http.Get(urlst)
	if e != nil {
		log.Println("the error errored ", e)
		return nil, e
	}
	// response we got
	//fmt.Println(urlst)
	todo1 := ToDo{}
	json.NewDecoder(r.Body).Decode(&todo1)
	fmt.Println(todo1)

	return &todo1, nil
}
