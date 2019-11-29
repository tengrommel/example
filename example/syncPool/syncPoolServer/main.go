package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/src-d/go-github/github"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

var prPool = sync.Pool{
	New: func() interface{} { return new(github.PullRequestEvent) },
}

func handle(writer http.ResponseWriter, r *http.Request) {
	data := prPool.Get().(*github.PullRequestEvent)
	defer prPool.Put(data)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logrus.Errorf("could not decode request: %v", err)
		http.Error(writer, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(writer, "pull request id: %d", *data.PullRequest.ID)
}
