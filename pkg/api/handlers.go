package api

import (
	"fmt"
	"net/http"
)

func GetTopics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetTopics handler")
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreateTopic handler")
}
