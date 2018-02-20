package ecobee

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func Get_ecobee_pin(ecobeeApiKey string) (error) {
	url := "https://api.ecobee.com/authorize"

	var jsonData = map[string]string{"response_type": "ecobeePin", "client_id": ecobeeApiKey, "scope": "smartWrite"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP POST request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		err = ioutil.WriteFile("ecobeePIN.json", data, 0644)
		if err != nil {
			panic(err)
		}
	}
	return err
}
