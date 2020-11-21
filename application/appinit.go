package application

import (
	"log"
	"os"

	configuration "../configuration"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*ConfigParameters for App*/
type ConfigParameters struct {
	DbName      string
	DbUser      string
	DbPass      string
	DbType      string
	DbHost      string
	DbPort      string
	Realm       string
	RealmUser   string
	RealmSecret string
	HTTPPort    string
}

/*Initialize app router and configuration*/
func (goservice *GoService) Initialize() {
	goservice.Router = mux.NewRouter()

	goservice.Router.HandleFunc("/health", goservice.healthCheckHandler)
	goservice.Router.Handle("/metrics", promhttp.Handler())

	goservice.Router.HandleFunc("/api/v1/event", goservice.PostCreate).Methods("POST")
	goservice.Router.HandleFunc("/api/v1/event", goservice.GetAll).Methods("GET")
	goservice.Router.HandleFunc("/api/v1/event/{id}", goservice.GetByID).Methods("GET")

	goservice.Router.Use(loggingMiddleware)
	goservice.Router.Use(goservice.authMiddleware)

	var appconfig ConfigParameters = handleConfiguration()
	goservice.Config = &appconfig
}

func handleConfiguration() ConfigParameters {
	var appconfig ConfigParameters

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file missing")
	}
	cloudconfigenabled := os.Getenv("cloudconfigenabled")

	if cloudconfigenabled == "false" {
		appconfig.DbUser = os.Getenv("db_user")
		appconfig.DbPass = os.Getenv("db_pass")
		appconfig.DbName = os.Getenv("db_name")
		appconfig.DbHost = os.Getenv("db_host")
		appconfig.DbPort = os.Getenv("db_port")
		appconfig.DbType = os.Getenv("db_type")
		appconfig.Realm = os.Getenv("realm")
		appconfig.RealmUser = os.Getenv("realm_user")
		appconfig.RealmSecret = os.Getenv("realm_secret")
		appconfig.HTTPPort = os.Getenv("http_port")
	} else {
		var cloudconfig configuration.CloudConfig
		cloudconfig = configuration.LoadCloudConfig()

		appconfig.DbUser = cloudconfig.PropertySources[0].Source.DbUser
		appconfig.DbPass = cloudconfig.PropertySources[0].Source.DbPass
		appconfig.DbName = cloudconfig.PropertySources[0].Source.DbName
		appconfig.DbHost = cloudconfig.PropertySources[0].Source.DbHost
		appconfig.DbPort = cloudconfig.PropertySources[0].Source.DbPort
		appconfig.DbType = cloudconfig.PropertySources[0].Source.DbType
		appconfig.Realm = cloudconfig.PropertySources[0].Source.Realm
		appconfig.RealmUser = cloudconfig.PropertySources[0].Source.RealmUser
		appconfig.RealmSecret = cloudconfig.PropertySources[0].Source.RealmSecret
		appconfig.HTTPPort = cloudconfig.PropertySources[0].Source.HTTPPort
	}

	return appconfig
}
