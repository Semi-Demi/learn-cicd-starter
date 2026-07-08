package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headerKey     string
		headerValue   string
		expected      string
		expectedError error
	}{
		{"", "", "", ErrNoAuthHeaderIncluded},
		{"Authorization", "", "", ErrNoAuthHeaderIncluded},
		{"Authorization", "token key", "", errors.New("malformed authorization header")},
		{"Authorization", "ApiKey token", "token", nil},
	}

	for _, test := range tests {
		header := http.Header{}
		header.Add(test.headerKey, test.headerValue)
		value, err := GetAPIKey(header)

		if value != test.expected || (err != nil && err.Error() != test.expectedError.Error()) || (err == nil && test.expectedError != nil) {
			t.Errorf("GetAPIKey(%v) = %v, %v; want %v, %v", header, value, err, test.expected, test.expectedError)
		}
	}

}
