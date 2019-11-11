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