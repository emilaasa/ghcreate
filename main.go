package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/* name	string	Required. The name of the repository.
	// POST /user/repos
	*/
	// Token: 5d69a732133c3c3fceeedc4c6848f9b0fbe5b0ce

	type GithubKey struct {
		Key string `json:"key"`
	}

	var credentials GithubKey

	tokenFile, err := ioutil.ReadFile("./key.json")
	json.Unmarshal(tokenFile, &credentials)

	githubApi := "https://api.github.com/"
	githubUser := "user"
	reposUrl := "/repos"
	fullUrl := githubApi + githubUser + reposUrl
	fmt.Println(fullUrl)
	exampleBody := []byte(`{ "name": "ghcreate" }`)
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(exampleBody))
	if err != nil {
		panic("write better error handling")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+credentials.Key)

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	fmt.Println("res status ", res.Status)

}
