package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	DbName string `json:"db_name"`
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbType string `json:"db_type"`
	DbHost string `json:"db_host"`
	DbPort string `json:"db_port"`
}

/*LoadCloudConfig from http://cloudconfig:8888/goservice/default*/
func LoadCloudConfig() {

	url := os.Getenv("cloudconfig")
	if url == "" {
		url = "http://nuc:8888/goservice/default"
	}

	client := &http.Client{}

	req, errhttp := http.NewRequest("GET", url, nil)
	if errhttp != nil {
		log.Fatal("Cloud Config http request error")
		log.Fatal(errhttp)
	} else {
		req.Header.Add("Accept", "application/json")

		resp, errhttpclient := client.Do(req)

		if errhttpclient != nil {
			log.Fatal("Cloud Config http client error")
			log.Fatal(errhttpclient)
		} else {
			defer resp.Body.Close()

			if resp.StatusCode == 404 {
				log.Fatal("Cloud Config not found")
			}

			if resp.StatusCode != 200 {
				log.Fatal("Cloud Config http error")
			}

			body, errio := ioutil.ReadAll(resp.Body)
			if errio != nil {
				log.Fatal("Cloud Config read error")
				log.Fatal(errio)
			} else {
				//stringproperties := string(body)
				//fmt.Println(stringproperties)

				var cloudconfig CloudConfig
				errjson := json.Unmarshal(body, &cloudconfig)

				if errjson != nil {
					fmt.Println(errjson)
					log.Fatal("Cloud Config json unmarshal error")
					os.Exit(1)
				} else {
					os.Setenv("db_user", cloudconfig.PropertySources[0].Source.DbUser)
					os.Setenv("db_pass", cloudconfig.PropertySources[0].Source.DbPass)
					os.Setenv("db_name", cloudconfig.PropertySources[0].Source.DbName)
					os.Setenv("db_host", cloudconfig.PropertySources[0].Source.DbHost)
					os.Setenv("db_port", cloudconfig.PropertySources[0].Source.DbPort)
					os.Setenv("db_type", cloudconfig.PropertySources[0].Source.DbType)
				}
			}
		}
	}
}
