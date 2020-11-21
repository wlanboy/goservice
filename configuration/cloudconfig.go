package configuration

/*CloudConfig struct*/
type CloudConfig struct {
	Name            string            `json:"name"`
	PropertySources []PropertySources `json:"propertySources"`
}

/*PropertySources struct*/
type PropertySources struct {
	Name   string `json:"name"`
	Source Source `json:"source"`
}

/*Source struct*/
type Source struct {
	DbName      string `json:"db_name"`
	DbUser      string `json:"db_user"`
	DbPass      string `json:"db_pass"`
	DbType      string `json:"db_type"`
	DbHost      string `json:"db_host"`
	DbPort      string `json:"db_port"`
	Realm       string `json:"realm"`
	RealmUser   string `json:"realm_user"`
	RealmSecret string `json:"realm_secret"`
	HTTPPort    string `json:"http_port"`
}
