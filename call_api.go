package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	url := "https://google.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	var todo Todo 
	err = json.Unmarshal(body, &todo)
	


}
