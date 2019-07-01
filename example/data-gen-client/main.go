package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/the-redback/go-oneliners"

	server "github.com/kamolhasan/Way-to-Learn-Prometheus/example/api-server"
)

func main() {
	var books = &server.BookList{
		Items: []server.Book{
			{
				ID:     "1",
				Name:   "Don Quixote",
				Author: "Miguel de Cervantes",
			},
			{
				ID:     "2",
				Name:   "A Tale of Two Cities",
				Author: "Charles Dickens",
			},
			{
				ID:     "3",
				Name:   "The Lord of the Rings",
				Author: "J.R.R. Tolkien",
			},
			{
				ID:     "4",
				Name:   "The Little Prince",
				Author: "Antoine de Saint-Exuper",
			},
		},
	}
	var book = &server.Book{
		ID:     "4",
		Name:   "Harry Potter and the Philosopher’s Stone ",
		Author: "J.K. Rowling",
	}
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	flag := false
	singalChan := make(chan os.Signal, 1)
	defer close(singalChan)
	signal.Notify(singalChan, os.Interrupt)
	go func() {
		<-singalChan
		flag = true
		log.Println("Shutting down... ...")
	}()

	for {
		if flag {
			break
		}
		switch rand.Intn(6) {
		case 0:
			MakePOST(&client, books)
		case 1:
			MakeUpdate(&client, book)
		case 2:
			MakeGetBooks(&client)
		case 3:
			MakeGetBook(&client)
		case 4:
			MakeGetBook2(&client)
		default:
			MakeFalseCall(&client)

		}
		time.Sleep(5 * time.Second)
	}

}

func MakePOST(clt *http.Client, books *server.BookList) {
	reqBody, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/books", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	oneliners.PrettyJson(body)

}

func MakeUpdate(clt *http.Client, book *server.Book) {
	reqBody, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/books/4", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	oneliners.PrettyJson(body)
}

func MakeGetBooks(clt *http.Client) {
	request, err := http.NewRequest("GET", "http://localhost:8080/books", bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	oneliners.PrettyJson(body)
}

func MakeGetBook(clt *http.Client) {
	request, err := http.NewRequest("GET", "http://localhost:8080/books/2", bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(body)
	oneliners.PrettyJson(body)
}

func MakeGetBook2(clt *http.Client) {
	request, err := http.NewRequest("GET", "http://localhost:8080/books/6", bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(body)
	oneliners.PrettyJson(body)
}

func MakeFalseCall(clt *http.Client) {
	request, err := http.NewRequest("GET", "http://localhost:8080/wrongURL", bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-type", "application/json")
	request.SetBasicAuth("kamol", "hasan")

	resp, err := clt.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(body)
	oneliners.PrettyJson(body)
}
