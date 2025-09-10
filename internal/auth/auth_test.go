package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	type getApiKeyResponse struct {
		apiKey    string
		errString string
	}

	type test struct {
		name    string
		headers http.Header
		want    getApiKeyResponse
	}

	tests := []test{
		{name: "Empty http.Header", headers: http.Header{}, want: getApiKeyResponse{apiKey: "", errString: "no authorization header included"}},
		{name: "Invalid Authorization Header", headers: http.Header{"Authorization": []string{"invalid"}}, want: getApiKeyResponse{apiKey: "", errString: "malformed authorization header"}},
		{name: "Malformed Authorization Header", headers: http.Header{"Authorization": []string{"ApiKey"}}, want: getApiKeyResponse{apiKey: "", errString: "malformed authorization header"}},
		{name: "Valid Authorization Header", headers: http.Header{"Authorization": []string{"ApiKey myFakeApiKey"}}, want: getApiKeyResponse{apiKey: "myFakeApiKey", errString: "failing on purpose"}},
	}

	for _, tc := range tests {
		apiKey, err := GetAPIKey(tc.headers)
		errorMessage := ""
		if err != nil {
			errorMessage = err.Error()
		}
		response := getApiKeyResponse{apiKey, errorMessage}
		if !reflect.DeepEqual(tc.want, response) {
			t.Fatalf("TestCase: %v. Expected: %v, got: %v", tc.name, tc.want, response)
		}
	}

}
