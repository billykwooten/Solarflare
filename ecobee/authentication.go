package ecobee

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"encoding/json"
	"github.com/jmoiron/jsonq"
)

func Get_ecobee_pin(ecobeeApiKey string) (string, error) {
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
		return "", fmt.Errorf("error retrieving response: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %s", err)
	}
	resp.Body.Close()

	var r struct {
		EcobeePin string `json:"ecobeePin"`
		Code      string `json:"code"`
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %s", err)
	}

	fmt.Printf("Pin is %q\nPress <enter> after authorizing it on https://www.ecobee.com/consumerportal in the menu under 'My Apps'\n", r.EcobeePin)
	var input string
	fmt.Scanln(&input)

	return r.Code, err
}

func GetToken(Code string, ecobeeApiKey string) error {
	uv := url.Values{
		"grant_type": {"ecobeePin"},
		"code":     {Code},
		"client_id":     {ecobeeApiKey},
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "api.ecobee.com",
		Path:     "token",
		RawQuery: uv.Encode(),
	}

	resp, err := http.PostForm(u.String(), nil)
	if err != nil {
		return fmt.Errorf("error POSTing request: %s", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	err = ioutil.WriteFile("cache/Ecobee_Token.json", body, 0644)
	if err != nil {
		panic(err)
	}

	return err
}


func RefreshToken(ecobeeApiKey string) error {
	jsonFile, _ := ioutil.ReadFile("cache/Ecobee_Token.json")
	data := map[string]interface{}{}
	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		panic(err)
	}

	jq := jsonq.NewQuery(data)
	RefreshToken, err := jq.String("refresh_token")
	if err != nil {
		panic(err)
	}

	uv := url.Values{
		"grant_type": {"refresh_token"},
		"refresh_token":     {RefreshToken},
		"client_id":     {ecobeeApiKey},
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "api.ecobee.com",
		Path:     "token",
		RawQuery: uv.Encode(),
	}

	resp, err := http.PostForm(u.String(), nil)
	if err != nil {
		return fmt.Errorf("error POSTing request: %s", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	err = ioutil.WriteFile("cache/Ecobee_Token.json", body, 0644)
	if err != nil {
		panic(err)
	}

	return err
}