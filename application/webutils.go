package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/wlanboy/goservice/v2/model"
)

/*WriteJSONErrorResponse with content type and status code*/
func WriteJSONErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{ "error": "%s" }`, message)
}

/*WriteJSONOkResponse with content type and status code*/
func WriteJSONInfoResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{ "Info": "%s" }`, message)
}

/*WriteJSONResponse with content type and status code*/
func WriteJSONResponse(w http.ResponseWriter, event *model.Event, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(event)
}

/*WriteJSONArrayResponse with content type and status code*/
func WriteJSONArrayResponse(w http.ResponseWriter, event []*model.Event, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(event)
}
