package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configuration "../configuration"
	model "../model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*GoService containing router and database*/
type GoService struct {
	Router *mux.Router
	DB     *gorm.DB
	Server *http.Server
}

/*Initialize app router and configuration*/
func (application *GoService) Initialize() {
	application.Router = mux.NewRouter()

	application.Router.HandleFunc("/api/v1/event", application.PostCreate).Methods("POST")
	application.Router.HandleFunc("/api/v1/event", application.GetAll).Methods("GET")
	application.Router.HandleFunc("/api/v1/event/{id}", application.GetByID).Methods("GET")

	//load env from cloud conig server
	configuration.LoadCloudConfig()
}

/*Run app and initialize db connection and http server*/
func (application *GoService) Run() {
	//load db connection
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	//fmt.Println(dbURI)
	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}

	application.DB = conn
	application.DB.Debug().AutoMigrate(&model.Event{})

	//load http server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)

	application.Server = &http.Server{
		Handler:      application.Router,
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting http server...")
		if err := application.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}

/*WaitForShutdown application server*/
func (application *GoService) WaitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	defer application.DB.Close()
	application.Server.Shutdown(ctx)

	log.Println("Shutting down http server.")
	os.Exit(0)
}
