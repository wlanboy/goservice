package application

import (
	configuration "../configuration"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*Initialize app router and configuration*/
func (goservice *GoService) Initialize() {
	goservice.Router = mux.NewRouter()

	goservice.Router.HandleFunc("/health", goservice.healthCheckHandler)

	goservice.Router.HandleFunc("/api/v1/event", goservice.PostCreate).Methods("POST")
	goservice.Router.HandleFunc("/api/v1/event", goservice.GetAll).Methods("GET")
	goservice.Router.HandleFunc("/api/v1/event/{id}", goservice.GetByID).Methods("GET")

	goservice.Router.Use(loggingMiddleware)

	var cloudconfig configuration.CloudConfig
	cloudconfig = configuration.LoadCloudConfig()

	var appconfig ConfigParameters

	appconfig.DbUser = cloudconfig.PropertySources[0].Source.DbUser
	appconfig.DbPass = cloudconfig.PropertySources[0].Source.DbPass
	appconfig.DbName = cloudconfig.PropertySources[0].Source.DbName
	appconfig.DbHost = cloudconfig.PropertySources[0].Source.DbHost
	appconfig.DbPort = cloudconfig.PropertySources[0].Source.DbPort
	appconfig.DbType = cloudconfig.PropertySources[0].Source.DbType

	goservice.Config = &appconfig
}
