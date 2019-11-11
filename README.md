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
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/satori/go.uuid
go get github.com/joho/godotenv
go clean
go build

# run
go run main.go