package main

import "fmt"

type Event struct {
	Id          int    `json:"Id"`
	Title       string `json:"title"`
	Descriptiom string `json:"description"`
}

type allEvents []Event

var events = allEvents{}

func main() {
	fmt.Println("Starting API")
}
