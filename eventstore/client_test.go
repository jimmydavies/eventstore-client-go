package eventstore

import (
  "testing"
  "net/http"
  "bytes"
  "io"
  "errors"
)

func TestNewClient(t *testing.T) {

  tables := []struct {
    baseUrl string
    userName string
    password string
    expectedPass bool
  } {
    {"http://eventstore.hostname:2113", "myuser", "mypass", true},
    {"://eventstore.hostname:2113", "myuser", "mypass", false},
    {"http://eventstore.hostname:2113", "", "", true},
  }

  for _, table := range tables {
  
    client, err := NewClient(table.baseUrl, table.userName, table.password)

    if table.expectedPass == true {
      if err != nil {
        t.Errorf("Error (%s) raised for inputs %v", err.Error(), table)
        continue
      }

      if client == nil {
        t.Errorf("Client object not created (nil returned) for inputs %v", table)
      }
    } else {
      if err == nil {
        t.Errorf("Expected Client creation to raise error for inputs %v", table)
        continue
      }

      if client != nil {
        t.Errorf("Expected Client creation to fail for inputs %v", table)
      }
    }
  }

}

func TestMakeRequest(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := &MockHttpClient{}

  client.httpClient = httpClient

  httpClient.setHttpClientResponse(`{"data": {}}`, 200, nil)

  client.makeRequest("GET", "/mypath", nil)

  httpClient.setHttpClientResponse(`{"data": {}}`, 200, nil)
  client.makeRequest("POST", "/mypath", nil)
  httpClient.setHttpClientResponse(`{"data": {}}`, 200, errors.New("Something went wrong"))
  client.makeRequest("GET", "/mypath", nil)
}

type MockHttpClient struct {
  response *http.Response
  err error
}

func (client *MockHttpClient) Do(*http.Request) (*http.Response, error) {
  if client.err != nil {
    return nil, client.err
  } else {
    return client.response, nil
  }
}

func (client *MockHttpClient) setHttpClientResponse(body string, statusCode int, err error) {
  client.response = &http.Response{
    StatusCode: statusCode,
    Body: io.NopCloser(bytes.NewReader([]byte(body))),
  }

  client.err = err
}

