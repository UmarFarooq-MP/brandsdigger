package godaddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type status struct {
	Available  bool   `json:"available"`
	Definitive bool   `json:"definitive"`
	Domain     string `json:"domain"`
}

type Response struct {
	Domains []status `json:"domains"`
}

type DomainValidator struct {
	apiKey    string
	apiSecret string
	baseURl   string
	client    *http.Client
}

func (gd DomainValidator) prepareRequest(domains []string) (*http.Request, error) {
	fullURl := fmt.Sprintf("%s/v1/domains/available?checkType=FAST&forTransfer=false", gd.baseURl)

	jsonData, err := json.Marshal(domains)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fullURl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	authorization := fmt.Sprintf("sso-key %s:%s", gd.apiKey, gd.apiSecret)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	return req, nil
}

func (gd DomainValidator) ValidateDomain(domains []string) (map[string]bool, error) {

	req, err := gd.prepareRequest(domains)
	if err != nil {
		return nil, err
	}

	resp, err := gd.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//TODO:: 203 is 200 for dev change it to 200 when on prod
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	m := make(map[string]bool)

	for _, domain := range response.Domains {
		m[domain.Domain] = domain.Available
	}
	return m, nil
}

func New(apikey, apiSecret, baseurl string) *DomainValidator {
	return &DomainValidator{
		apiKey:    apikey,
		apiSecret: apiSecret,
		baseURl:   baseurl,
		client:    &http.Client{},
	}
}
