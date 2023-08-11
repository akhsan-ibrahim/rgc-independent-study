package main

import (
	"bytes"
	"encoding/json"
	// "regexp"
	// "strings"

	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://animechan.vercel.app/api/quotes/anime?title=naruto", nil)
	res, _ := client.Do(req)
	datas, _ := ioutil.ReadAll(res.Body)

	animeQuotes := []Animechan{}
	err := json.Unmarshal(datas, &animeQuotes)
	if err != nil {
		// fmt.Println(err)
		return []Animechan{}, err
	}
	return animeQuotes, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, _ := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	var pos Postman
	err := json.Unmarshal(body, &pos)
	if err != nil {
		return Postman{}, err
	}

	// Hit API https://postman-echo.com/post with method POST:
	return pos, nil // TODO: replace this
}

func main() {
	// get, _ := ClientGet()
	// fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
