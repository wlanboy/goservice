package controller

import (
	"encoding/json"
	"net/http"

	app "../application"
	model "../model"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

/*PostCreate POST method*/
var PostCreate = func(w http.ResponseWriter, r *http.Request) {

	event := &model.Event{}

	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		app.WriteJSONErrorResponse(w, "Cannot parse JSON", http.StatusBadRequest)
	} else {
		err, resp := model.SaveEvent(*event)
		if err != "" {
			app.WriteJSONErrorResponse(w, err, http.StatusInternalServerError)
		} else {
			app.WriteJSONResponse(w, resp, http.StatusCreated)
		}
	}
}

/*GetByID GET method*/
var GetByID = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	uuid, uuiderr := uuid.FromString(id)

	if uuiderr != nil {
		app.WriteJSONErrorResponse(w, "Cannot parse UUID", http.StatusBadRequest)
	} else {
		err, resp := model.GetEventByID(uuid)
		if err != "" {
			app.WriteJSONErrorResponse(w, err, http.StatusNotFound)
		} else {
			app.WriteJSONResponse(w, resp, http.StatusOK)
		}
	}
}

/*GetAll GET method*/
var GetAll = func(w http.ResponseWriter, r *http.Request) {

	err, resp := model.GetAllEvents()
	if err != "" {
		app.WriteJSONErrorResponse(w, err, http.StatusNotFound)
	} else {
		app.WriteJSONArrayResponse(w, resp, http.StatusOK)
	}
}
