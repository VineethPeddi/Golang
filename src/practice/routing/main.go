package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	// "github.com/gorilla/mux"
)

type P struct{}

func (p P) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print("in serve http")
}

func main() {
	mux := http.NewServeMux()
	fmt.Println("server about to listening on 8k")
	mux.Handle("/", P{})
	mux.HandleFunc()
	http.ListenAndServe(":8000", mux)
	fmt.Println("server listening on 8k")
	//performGetRequest()
	// performPostRequest("http://localhost:8000/post", "application/json")
	// fmt.Println("\n\ndone with first post\n\n ")
	// performPostRequest("http://localhost:8000/postform", "application/x-www-form-urlencoded")
}

func performGetRequest() {
	const url = "http://localhost:8000/get"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("resp: %+v\n\n", *resp)
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body: ", string(data))
}

type Student struct {
	Name  string
	Marks int
}

func performPostRequest(reqUrl string, contentType string) {

	student := &Student{
		Name:  "Vineeth",
		Marks: 99,
	}

	var body *strings.Reader
	if contentType == "application/json" {
		data, _ := json.Marshal(student)
		body = strings.NewReader(string(data))

	} else {
		data := url.Values{}
		data.Set("Name", "Deva")
		data.Set("Type", "form-encoded")
		body = strings.NewReader(data.Encode())
	}

	resp, err := http.Post(reqUrl, contentType, body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("resp: %+v\n\n", *resp)
	resData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body: ", string(resData))
}
