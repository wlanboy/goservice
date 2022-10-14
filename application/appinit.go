package application

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tkanos/gonfig"
	configuration "github.com/wlanboy/goservice/v2/configuration"
)

/*ConfigParameters for App*/
type ConfigParameters struct {
	CloudConfig string
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

	goservice.Router.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	goservice.Router.Use(loggingMiddleware)
	goservice.Router.Use(goservice.authMiddleware)

	var appconfig ConfigParameters = handleConfiguration()
	goservice.Config = &appconfig
}

func handleConfiguration() ConfigParameters {
	var appconfig ConfigParameters

	err := gonfig.GetConf("config.json", &appconfig)
	if err != nil {
		log.Fatal(".env file missing")
	}
	cloudconfigenabled := appconfig.CloudConfig

	if cloudconfigenabled == "true" {
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
