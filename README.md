# goservice
Golang Rest service with Postgresql backend and Spring Cloud Config Server support

# using
- "context"
- "os"
- "os/signal"
- "syscall"
- "fmt"
- "time"
- "net/http"
- "net/http/httputil"
- "log"
- "encoding/json"
- "github.com/gorilla/mux"
- "github.com/gorilla/handlers"
- "github.com/jinzhu/gorm"
- "github.com/jinzhu/gorm/dialects/postgres"
- "github.com/satori/go.uuid"

# build
go get -d -v

go clean

go build

# run
go run main.go

# debug
go get -u github.com/go-delve/delve/cmd/dlv
dlv debug ./goservice

# dockerize (docker image size is 9.89MB)
GOOS=linux GOARCH=386 go build

docker build -t goservice .

# run docker container
docker run -d -p 8000:8000 goservice

# call
curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'

curl -X GET http://127.0.0.1:8000/api/v1/event 
