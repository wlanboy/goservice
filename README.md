# goservice
Golang Rest service with Postgresql backend and Spring Cloud Config Server support

# using
- "os"
- "fmt"
- "time"
- "net/http"
- "encoding/json"
- "github.com/gorilla/mux"
- "github.com/jinzhu/gorm"
- "github.com/jinzhu/gorm/dialects/postgres"
- "github.com/satori/go.uuid"

# build
go get -d -v

go clean

go build

# run
go run main.go

# dockerize (dokcer image size is 9.89MB)
GOOS=linux GOARCH=386 go build
docker build -t goservice .

# run docker container
docker run -d -p 8000:8000 goservice

# call
curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'

curl -X GET http://127.0.0.1:8000/api/v1/event 
