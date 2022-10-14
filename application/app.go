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

	model "github.com/wlanboy/goservice/v2/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "net/http/pprof"
)

/*GoService containing router and database*/
type GoService struct {
	Router *mux.Router
	DB     *gorm.DB
	Server *http.Server
	Config *ConfigParameters
}

/*Run app and initialize db connection and http server*/
func (goservice *GoService) Run() {
	//load db connection
	username := goservice.Config.DbUser
	password := goservice.Config.DbPass
	dbName := goservice.Config.DbName
	dbHost := goservice.Config.DbHost
	dbPort := goservice.Config.DbPort
	dbType := goservice.Config.DbType
	httpport := goservice.Config.HTTPPort

	var conn *gorm.DB
	var err error

	if dbType == "sqlite3" {
		conn, err = gorm.Open(dbType, dbName)
	} else {
		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
		conn, err = gorm.Open(dbType, dbURI)
	}
	//fmt.Println(dbURI)
	if err != nil {
		fmt.Print(err)
	}

	goservice.DB = conn
	goservice.DB.Debug().AutoMigrate(&model.Event{})

	//load http server
	if httpport == "" {
		httpport = "8000"
	}
	fmt.Println(httpport)

	goservice.Server = &http.Server{
		Handler:      goservice.Router,
		Addr:         fmt.Sprintf(":%s", httpport),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting http server...")
		if err := goservice.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}

/*WaitForShutdown application server*/
func (goservice *GoService) WaitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	defer goservice.DB.Close()
	goservice.Server.Shutdown(ctx)

	log.Println("Shutting down http server.")
	os.Exit(0)
}
