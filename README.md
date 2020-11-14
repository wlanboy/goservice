![Go](https://github.com/wlanboy/goservice/workflows/Go/badge.svg?branch=master)

# goservice
Golang Rest service based on gorilla, gorm using Spring cloud config and PostgreSql

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
* go get -d -v
* go clean
* go build

# build windows
* install https://jmeubank.github.io/tdm-gcc/
* go get github.com/mattn/go-sqlite3 in windows shell

# depends on
* PostgreSQL instance: https://github.com/wlanboy/Dockerfiles/tree/master/Postgres

# run
* go run main.go

# debug
* go get -u github.com/go-delve/delve/cmd/dlv
* dlv debug ./goservice

# dockerize (docker image size is 9.89MB)
* GOOS=linux GOARCH=386 go build (386 needed for busybox)
* GOOS=linux GOARCH=arm GOARM=6 go build (Raspberry Pi build)
* GOOS=linux GOARCH=arm64 go build (Odroid C2 build)
* docker build -t goservice .

## Docker publish to github registry
- docker tag goservice:latest docker.pkg.github.com/wlanboy/goservice/goservice:latest
- docker push docker.pkg.github.com/wlanboy/goservice/goservice:latest

## Docker Registry repro
- https://github.com/wlanboy/goservice/packages/278503

# run docker container
*docker run -d -p 8000:8000 goservice

# create event
* curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'
# get all events
* curl -X GET http://127.0.0.1:8000/api/v1/event 
# get prometheus metrics
* curl -X GET http://127.0.0.1:8000/metrics
 
