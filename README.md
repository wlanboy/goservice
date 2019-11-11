# goservice
Golang Rest service with Postgresql backend

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
- "github.com/joho/godotenv"

# build
go get -d -v
go clean
go build

# run
go run main.go

# dockerize
docker build -t goservice .
docker run -p 8000:8000 goservice

# call
curl -X POST http://127.0.0.1:8000/api/v1/event -H 'Content-Type: application/json' -d '{"name": "test", "type": "info"}'
curl -X GET http://127.0.0.1:8000/api/v1/event 
