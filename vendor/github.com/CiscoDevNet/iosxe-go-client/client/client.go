package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	gabs "github.com/CiscoDevNet/iosxe-go-client/container"
)

const (
	headerAccept            = "Accept"
	headerContentType       = "Content-Type"
	mimeTypeApplicationJSON = "application/yang-data+json"
)

// V2 client for IOSXE REST APIs
type V2 struct {
	host           string
	deviceUsername string
	devicePassword string
	httpClient     *http.Client
}

// NewV2 Constructor for V2 Client
//
// e.g. NewV2("https://sandbox@cisco.com:8890", "admin", "Cisco123", 30, true, "", "", "")
func NewV2(host, deviceUsername, devicePassword string, defaultTimeout int, insecure bool, caFile, proxyURL, proxyCreds string) (*V2, error) {
	if !strings.HasPrefix(host, "http") {
		return nil, errors.New("hostURL scheme must be 'http(s)'")
	}

	// hostURL should end with port number and not with trailing slash
	// if trailing slash present, remove it
	host = strings.TrimSuffix(host, "/")
	baseURL := fmt.Sprintf("%s/restconf", host)
	transport := http.DefaultTransport.(*http.Transport)

	// Add proxy if required
	if proxyURL != "" {
		pUrl, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(pUrl)
		if proxyCreds != "" {
			basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxyCreds))
			transport.ProxyConnectHeader = http.Header{}
			transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
		}
	}
	// Set the insecure boolean
	tlsConfig := &tls.Config{InsecureSkipVerify: insecure}

	// Add CA certificate if required
	if len(caFile) > 0 {
		caCert, err := ioutil.ReadFile(filepath.Clean(caFile))
		if err != nil {
			return nil, err
		}
		caCertPool, _ := x509.SystemCertPool()
		if caCertPool == nil {
			caCertPool = x509.NewCertPool()
		}
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("could not append CA file to CA Cert pool")
		}
		tlsConfig.RootCAs = caCertPool
	}

	transport.TLSClientConfig = tlsConfig
	httpClient := &http.Client{
		Timeout:   time.Second * time.Duration(defaultTimeout),
		Transport: transport,
	}

	return &V2{
		host:           baseURL,
		deviceUsername: deviceUsername,
		devicePassword: devicePassword,
		httpClient:     httpClient,
	}, nil
}

// Do function performs HTTP API Call
func (c *V2) Do(req *http.Request) (*http.Response, error) {
	log.Printf("[DEBUG] Begining DO method %s", req.URL.String())
	var resp *http.Response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, checkForErrors(resp)
}

// PrepareRequest Creates *http.Request with required headers
func (c *V2) PrepareRequest(method string, endpoint string, body *gabs.Container, queryParams *map[string]string) (*http.Request, error) {
	urlString := fmt.Sprintf("%s%s", c.host, endpoint)
	// validate url
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	if method == http.MethodGet || method == http.MethodDelete {
		req, err = http.NewRequest(method, urlString, nil)
	} else if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req, err = http.NewRequest(method, urlString, bytes.NewBuffer(body.Bytes()))
	} else {
		return nil, errors.New("invalid method")
	}
	if err != nil {
		return nil, err
	}
	// deviceUsername, devicePassword works like basic auth
	req.SetBasicAuth(c.deviceUsername, c.devicePassword)

	req.Header.Set(headerAccept, mimeTypeApplicationJSON)
	req.Header.Set(headerContentType, mimeTypeApplicationJSON)

	if queryParams != nil {
		q := req.URL.Query()
		for qk, qv := range *queryParams {
			q.Add(qk, qv)
		}
		req.URL.RawQuery = q.Encode()
	}
	return req, nil
}

// checkForErrors - Checks for http status code based errors
func checkForErrors(resp *http.Response) error {
	resourcePath := strings.Replace(resp.Request.URL.Path, "/restconf", "", 1)
	method := resp.Request.Method

	switch resp.StatusCode {
	// success checks
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return nil

	// client config/parameter/payload related error checks
	case http.StatusNotFound:
		return fmt.Errorf("not-found: %s", resourcePath)
	case http.StatusUnauthorized:
		return errors.New("unauthorized: please check your credentials")
	case http.StatusForbidden:
		return errors.New("forbidden: you do not have permissions")
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("not-allowed: %s is not allowed on %s", method, resourcePath)
	case http.StatusNotAcceptable:
		container, err := GetContainer(resp)
		if err != nil {
			return fmt.Errorf("not-acceptable: %v", err)
		}

		return fmt.Errorf("not-acceptable: %v", formatNotAcceptableErrors(container))

	// server side errors
	case http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusBadGateway:
		return fmt.Errorf("server-error: %s", resp.Status)
	default:
		log.Printf("[DEBUG] Unable to identify HTTP error")
	}

	// fallback success check
	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return nil
	}
	// Handling of unknown errors
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] HTTP REQUEST FAILED [%s] %s %d - response: %s",
		method, resourcePath, resp.StatusCode, bodyString)
	return fmt.Errorf("failed: status code: %d - error: %v", resp.StatusCode, bodyString)
}

func formatNotAcceptableErrors(data *gabs.Container) string {
	var errString string

	if data.Exists("token") {
		errString += fmt.Sprintf("Token: %s: %s", data.S("token").Data(), data.S("message").Data())
	} else {
		for _, child := range data.Children() {
			errString += fmt.Sprintf("\nToken: %s\nMessage: %s\n", child.S("token").Data(), child.S("message").Data())
			for k, v := range child.ChildrenMap() {
				if k != "message" && k != "token" {
					errString += fmt.Sprintf("%s: %s\n", k, v.String())
				}
			}
		}
	}

	if errString == "" {
		return data.String()
	} else {
		return errString
	}
}
