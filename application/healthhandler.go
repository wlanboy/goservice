package application

import (
	"fmt"
	"log"
	"net/http"
)

/*healthCheckHandler*/
func (goservice *GoService) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := "up"
	currentdatabase := ""
	if err := goservice.DB.Exec("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%';").Error; err != nil {
		message = "down"
		log.Print(err)
	}
	log.Println(currentdatabase)
	fmt.Fprintf(w, `{ "health": "%s" }`, message)
}
