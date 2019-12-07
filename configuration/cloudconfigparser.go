package configuration

import (
	"encoding/json"
	"fmt"
	"log"
)

func parseAndSetConfigValues(body []byte) CloudConfig {
	var cloudconfig CloudConfig
	errjson := json.Unmarshal(body, &cloudconfig)

	if errjson != nil {
		fmt.Println(errjson)
		log.Fatal("Cloud Config json unmarshal error")
	}
	return cloudconfig
}
