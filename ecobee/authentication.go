package ecobee

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type PinRequest struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
}


func Get_ecobee_pin(ecobeeApiKey string) (error) {
	var Scopes = []string{"smartRead", "smartWrite"}

	uv := url.Values{
		"response_type": {"ecobeePin"},
		"client_id":     {ecobeeApiKey},
		"scope":         {strings.Join(Scopes, ",")},
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "api.ecobee.com",
		Path:     "authorize",
		RawQuery: uv.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("error retrieving response: %s", err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("ecobeePIN.json", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	return err


}
