package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	gabs "github.com/CiscoDevNet/iosxe-go-client/container"

	"github.com/CiscoDevNet/iosxe-go-client/models"
)

// Errors for V2 API client
var (
	ErrPayloadCreation = errors.New("error in creating payload")
)

// Get Performs HTTP(S) GET on endpoint with queryParams
//
// Cisco IOS XE GET APIs return json response on success, available as *gabs.Container
func (c *V2) Get(endpoint string, queryParams *map[string]string) (*http.Response, *gabs.Container, error) {
	req, err := c.PrepareRequest(http.MethodGet, endpoint, nil, queryParams)
	if err != nil {
		return nil, nil, err
	}
	response, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	container, err := GetContainer(response)
	if err != nil {
		return response, nil, err
	}
	return response, container, nil
}

// Create Performs HTTP(S) POST on endpoint with model
//
// Cisco IOS XE POST APIs return json response on success, available as *gabs.Container
func (c *V2) Create(endpoint string, model models.Model) (*http.Response, *gabs.Container, error) {
	jsonPayload, err := c.PrepareModel(model)
	if err != nil {
		return nil, nil, err
	}

	log.Println("[DEBUG] CREATE Payload: ", jsonPayload.String())

	req, err := c.PrepareRequest(http.MethodPost, endpoint, jsonPayload, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := c.Do(req)
	if err != nil {
		return response, nil, err
	}

	container, err := GetContainer(response)
	if err != nil {
		return response, nil, err
	}
	return response, container, nil
}

// Patch Performs HTTP(S) PATCH on endpoint with model
//
// Cisco IOS XE PATCH APIs return json response on success, available as *gabs.Container
func (c *V2) Patch(endpoint string, model models.Model) (*http.Response, *gabs.Container, error) {
	jsonPayload, err := c.PrepareModel(model)
	if err != nil {
		return nil, nil, err
	}

	log.Println("[DEBUG] PATCH Payload: ", jsonPayload.String())

	req, err := c.PrepareRequest(http.MethodPatch, endpoint, jsonPayload, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := c.Do(req)
	if err != nil {
		return response, nil, err
	}

	container, err := GetContainer(response)
	if err != nil {
		return response, nil, err
	}
	return response, container, nil
}

// Update Performs HTTP(S) UPDATE on endpoint
//
// Cisco IOS XE UPDATE APIs does not return any response
func (c *V2) Update(endpoint string, model models.Model) (*http.Response, error) {
	jsonPayload, err := c.PrepareModel(model)
	if err != nil {
		return nil, err
	}

	log.Println("[DEBUG] UPDATE Payload: ", jsonPayload.String())

	req, err := c.PrepareRequest(http.MethodPut, endpoint, jsonPayload, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Delete Performs HTTP(S) DELETE on endpoint
//
// Cisco IOS XE DELETE APIs does not return any response
func (c *V2) Delete(endpoint string) (*http.Response, error) {
	req, err := c.PrepareRequest(http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// GetContainer parses HTTP(S) Response and returns *gabs.Container
func GetContainer(resp *http.Response) (*gabs.Container, error) {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > 0 {
		return gabs.ParseJSON(bodyBytes)
	}
	if len(bodyBytes) == 0 {
		return nil, nil
	}
	return nil, errors.New("no body found")
}

// PrepareModel Creates container from model
func (c *V2) PrepareModel(model models.Model) (*gabs.Container, error) {
	modelMap, err := model.ToMap()
	if err != nil {
		return nil, err
	}

	payload := gabs.New()

	for key, value := range modelMap {
		_, err = payload.Set(value, key)
		if err != nil {
			return nil, ErrPayloadCreation
		}
	}

	json_payload := payload.S("json_payload")
	if err != nil {
		return nil, err
	}

	json_payload_unquoted, err := strconv.Unquote(string(json_payload.Bytes()))
	if err != nil {
		return nil, err
	}

	parsed_json_payload, err := gabs.ParseJSON([]byte(json_payload_unquoted))
	if err != nil {
		return nil, err
	}

	return parsed_json_payload, nil
}
