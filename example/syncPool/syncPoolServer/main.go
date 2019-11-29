package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/src-d/go-github/github"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(writer http.ResponseWriter, r *http.Request) {
	var data github.PullRequestEvent
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logrus.Errorf("could not decode request: %v", err)
		http.Error(writer, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(writer, "pull request id: %d", *data.PullRequest.ID)
}
