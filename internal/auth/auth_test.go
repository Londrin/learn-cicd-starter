package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := http.Header{}
	header1.Set("Authorization", "ApiKey 123asdqwe123")
	header2 := http.Header{}
	header2.Set("InvalidAuth", "ApiKey 123asdasdasdqwe123")
	header3 := http.Header{}
	header3.Set("Authorization", "Invalid 123123lkf")

	test_cases := []struct {
		name    string
		header  http.Header
		wantErr bool
	}{
		{
			name:    "Correct Header",
			header:  header1,
			wantErr: false,
		},
		{
			name:    "Invalid Key, Correct Value",
			header:  header2,
			wantErr: true,
		},
		{
			name:    "Valid key, invalid value",
			header:  header3,
			wantErr: false,
		},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetAPIKey(tc.header)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
