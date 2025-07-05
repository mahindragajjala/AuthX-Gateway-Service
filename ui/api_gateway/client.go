package apigateway

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
)

/*
	 func API_Gateway_Json_Data_Request_Signup(jsonData []byte) []byte {
		// HTTP client with insecure TLS (for local testing)
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		// Create request
		req, err := http.NewRequest("POST", "https://172.20.78.91:443/signup", bytes.NewBuffer(jsonData))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send request
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// Log the pretty JSON for debug
		var jsonDataPretty map[string]interface{}
		if err := json.Unmarshal(body, &jsonDataPretty); err == nil {
			prettyJSON, _ := json.MarshalIndent(jsonDataPretty, "", "  ")
			fmt.Println("Response:\n", string(prettyJSON))
		}

		// ✅ Return response body (not the whole http.Response)
		return body
	}
*/
func API_Gateway_Json_Data_Request_Signup(jsonData []byte) ([]byte, int, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("POST", "https://172.20.78.91:443/signup", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func API_Gateway_Json_Data_Request_Login(jsonData []byte) (*http.Response, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ⚠️ Not safe in prod
		},
	}

	req, err := http.NewRequest("POST", "https://172.20.78.91:443/login", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil // Don't close here, caller will handle it
}
