// main_test.go
package main

import (
	"errors"
	"log"
	"net/http"
	"testing"
)

type mockAPIClient struct{}

func (ma *mockAPIClient) Do(req *http.Request) (*http.Response, error) {
	log.Print("mock func called")
	return nil, errors.New("test error")
}

func Test_api_MakeAPI(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name           string
		a              *api
		args           args
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success",
			a: &api{
				httpClient: &mockAPIClient{},
			},
			args: args{
				url: "http://dummy.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatusCode, err := tt.a.MakeAPI(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("api.MakeAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("api.MakeAPI() = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
