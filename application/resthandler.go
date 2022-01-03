package application

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/wlanboy/goservice/v2/model"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

/*PostCreate POST method*/
func (goservice *GoService) PostCreate(w http.ResponseWriter, r *http.Request) {

	event := model.Event{}

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println("Cannot parse JSON")
		log.Println(err)
		WriteJSONErrorResponse(w, "Cannot parse JSON", http.StatusBadRequest)
	} else {
		errdb, resp := model.SaveEvent(event, goservice.DB)
		if errdb != "" {
			log.Println("Model error")
			log.Println(errdb)
			WriteJSONErrorResponse(w, errdb, http.StatusInternalServerError)
		} else {
			WriteJSONResponse(w, resp, http.StatusCreated)
		}
	}
}

/*GetByID GET method*/
func (goservice *GoService) GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	uuid, uuiderr := uuid.FromString(id)

	if uuiderr != nil {
		WriteJSONErrorResponse(w, "Cannot parse UUID", http.StatusBadRequest)
	} else {
		errdb, resp := model.GetEventByID(uuid, goservice.DB)
		if errdb != "" {
			log.Println("Model error")
			log.Println(errdb)
			WriteJSONErrorResponse(w, errdb, http.StatusNotFound)
		} else {
			WriteJSONResponse(w, resp, http.StatusOK)
		}
	}
}

/*GetAll GET method*/
func (goservice *GoService) GetAll(w http.ResponseWriter, r *http.Request) {

	errdb, resp := model.GetAllEvents(goservice.DB)
	if errdb != "" {
		log.Println("Model error")
		log.Println(errdb)
		WriteJSONErrorResponse(w, errdb, http.StatusNotFound)
	} else {
		WriteJSONArrayResponse(w, resp, http.StatusOK)
	}
}
