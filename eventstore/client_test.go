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
        t.Errorf("Unexpected error (%s) raised for inputs %v", err.Error(), table)
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

  tables := []struct {
    method     string
    body       string
    statusCode int
    status     string
    err        error
  } {
    {"GET",  `{"data": {}}`, 200, "200 OK",        nil},
    {"POST", `{"data": {}}`, 200, "200 OK",        nil},
    {"GET",  `{"data": {}}`, 200, "200 OK",        errors.New("Something went wrong")},
    {"GET",  `{"data": {}}`, 401, "401 Forbidden", nil},
    {"GET",  `{"data": {}}`, 404, "404 Not Found", nil},
  }

  for _, table := range tables {
    
    httpClient.setHttpClientResponse(table.body, table.statusCode, table.status, table.err)

    _, err := client.makeRequest(table.method, "/mypath", nil)

    if (table.err != nil || table.statusCode != 200) && err == nil {
      t.Errorf("Expected error not raised for inputs %v", table)
      continue
    } else if table.err == nil && table.statusCode == 200 && err != nil {
      t.Errorf("Unexpected error (%s) raised for inputs %v (%s, %d)", err.Error(), table, err.Error(), table.statusCode)
    }
  }
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

func (client *MockHttpClient) setHttpClientResponse(body string, statusCode int, status string, err error) {
  client.response = &http.Response{
    StatusCode: statusCode,
    Status: status,
    Body: io.NopCloser(bytes.NewReader([]byte(body))),
  }

  client.err = err
}

