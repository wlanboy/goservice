package application

import (
	"net/http"
	"time"
)

/*GetLong GET method*/
func (goservice *GoService) GetLong(w http.ResponseWriter, r *http.Request) {

	resultChan := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		resultChan <- "operation done"
	}()
	result := <-resultChan
	WriteJSONInfoResponse(w, result, http.StatusOK)

}
