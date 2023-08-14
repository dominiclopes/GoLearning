package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var respData Response
	if err := json.Unmarshal(bodyBytes, &respData); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", respData)
	fmt.Printf("%#v", respData)
}
