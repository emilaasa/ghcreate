package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type githubKey struct {
	Key string `json:"key"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Needs one argument, ghcreate reponame")
	}
	repoName := os.Args[1]
	var credentials githubKey
	credsFile, err := ioutil.ReadFile("./key.json")
	json.Unmarshal(credsFile, &credentials)

	url := createRepoURL()
	body := []byte("{ \"name\": \"" + repoName + "\" }")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic("write better error handling")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+credentials.Key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("res status ", res.Status)

}

func createRepoURL() string {
	githubAPI := "https://api.github.com/"
	githubUser := "user"
	reposURL := "/repos"
	fullURL := githubAPI + githubUser + reposURL
	return fullURL
}
