package apigateway

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

func API_Gateway_Json_Data_Request(jsonData []byte) {

	// Skip cert verification ONLY for testing (not recommended in prod)
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("POST", "https://172.20.78.91:443/signup", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("✅ Server response status:", resp.Status)
}
