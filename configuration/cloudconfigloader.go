package configuration

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*LoadCloudConfig from http://cloudconfig:8888/goservice/default*/
func LoadCloudConfig() CloudConfig {
	var cloudconfig CloudConfig

	url := os.Getenv("cloudconfig")
	if url == "" {
		url = "http://nuc:8888/goservice/default"
	}

	client := &http.Client{}

	req, errhttp := http.NewRequest("GET", url, nil)
	if errhttp != nil {
		log.Println("Cloud Config http request error")
		log.Fatal(errhttp)
	} else {
		req.Header.Add("Accept", "application/json")

		resp, errhttpclient := client.Do(req)

		if errhttpclient != nil {
			log.Println("Cloud Config http client error")
			log.Fatal(errhttpclient)
		} else {
			cloudconfig = readAndCheckHTTPResponse(resp)
		}
	}
	return cloudconfig
}

/*readAndCheckHTTPResponse to byte array*/
func readAndCheckHTTPResponse(resp *http.Response) CloudConfig {
	var cloudconfig CloudConfig

	if resp.StatusCode == 404 {
		log.Fatal("Cloud Config not found")
	}

	if resp.StatusCode != 200 {
		log.Fatal("Cloud Config http error")
	}
	defer resp.Body.Close()

	body, errio := ioutil.ReadAll(resp.Body)
	if errio != nil {
		log.Println("Cloud Config read error")
		log.Fatal(errio)
	} else {
		cloudconfig = parseAndSetConfigValues(body)
	}
	return cloudconfig
}
