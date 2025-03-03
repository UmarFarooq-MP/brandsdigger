package daddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type status struct {
	Available  bool   `json:"available"`
	Definitive bool   `json:"definitive"`
	Domain     string `json:"domain"`
}

type Response struct {
	Domains []status `json:"domains"`
}

type GodaddyDomainValidator struct {
	apiKey    string
	apiSecret string
	baseURl   string
}

func (gd GodaddyDomainValidator) ValidateDomain(domain []string) (bool, error) {

	params := url.Values{}
	params.Set("checkType", "FAST")
	params.Set("forTransfer", "false")

	fullURl := fmt.Sprintf("%s/v1/domains/available?checkType=FAST&forTransfer=false", gd.baseURl)

	jsonData, err := json.Marshal(domain)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", fullURl, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	authorization := fmt.Sprintf("sso-key %s:%s", gd.apiKey, gd.apiSecret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, errors.New(resp.Status)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	fmt.Println("Response Status: ", response)

	return true, nil
}

func New(apikey, apiSecret, baseurl string) *GodaddyDomainValidator {
	return &GodaddyDomainValidator{
		apiKey:    apikey,
		apiSecret: apiSecret,
		baseURl:   baseurl,
	}
}
