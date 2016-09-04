package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
)

const (
	serverUrl = "http://jsonplaceholder.typicode.com/todos/1"
)

type todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t todo

	client := http.Client{}

	req, err := http.NewRequest("GET", serverUrl, nil)

	if err != nil {
		panic(err)
	}

	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Println("got conn")
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
		WroteRequest: func(wr httptrace.WroteRequestInfo) {
			fmt.Println("Wrote request", wr)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	bResp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	json.Unmarshal(bResp, &t)

	fmt.Println(`id:`, t.ID)
	fmt.Println(`userId:`, t.UserID)
	fmt.Println(`title:`, t.Title)
	fmt.Println(`completed:`, t.Completed)
}
