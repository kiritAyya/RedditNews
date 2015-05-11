package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type News struct {
	Data struct {
		Children []struct {
			Data_ struct {
				Score  float64 `json:"score"`
				Author string  `json:"author"`
				Link   string  `json:"permalink"`
				Title  string  `json:"title"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	resp, err := http.Get("http://www.reddit.com/r/golang.json")
	check(err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	body, errq := ioutil.ReadAll(resp.Body)
	check(errq)

	b := &News{}

	er := json.Unmarshal(body, &b)
	check(er)

	for i := 0; i < len(b.Data.Children); i++ {
		fmt.Printf("Title: %v\n", b.Data.Children[i].Data_.Title)
		fmt.Printf("Author: %v\n", b.Data.Children[i].Data_.Author)
		fmt.Printf("Score: %v\n", b.Data.Children[i].Data_.Score)
		fmt.Printf("Link: %v\n", b.Data.Children[i].Data_.Link)
		fmt.Printf("\n\n")
	}
}
