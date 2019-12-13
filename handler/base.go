package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func parseRequest(r *http.Request, body interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(&body)
	}

	return errors.New("no supported type")
}

func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	content, _ := json.Marshal(data)
	_, _ = w.Write(content)
}
