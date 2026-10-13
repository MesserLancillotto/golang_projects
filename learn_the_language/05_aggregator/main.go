package main;

import (
	"net/http"
	"encoding/json"
	"sync"
)

var (
	channel chan IncomingRawData
)

func postHandler(writer http.ResponseWriter, request * http.Request, channel chan<-IncomingRawData) {
	var payload IncomingRawData
	if 	err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		http.Error(writer, "Malformed JSON", http.StatusBadRequest)
		return
	}
	channel<-payload
	writer.WriteHeader(http.StatusCreated)
}

func HandleIncomingRequests(writer http.ResponseWriter, request * http.Request) {
	defer request.Body.Close()
	if request.Method != http.MethodPost {
		http.Error(writer, "Method should be POST: found instead PUT, PATCH, GET or DELETE", http.StatusMethodNotAllowed)
		return
	}
	postHandler(writer, request)
}
