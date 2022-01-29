package eventstore

import (
  "net/http"
  "log"
  "io/ioutil"
  "io"
  "encoding/json"
  "bytes"
)

type Client struct {
  baseUrl    string
  username   string
  password   string
}

func NewClient(baseUrl string, username string, password string) (*Client, error) {
  return &Client{
    baseUrl: baseUrl,
    username: username,
    password: password,
  }, nil
}

func (client *Client) makeRequest(requestType string, path string, body []byte) (map[string]interface{}, error) {

  httpClient := &http.Client{}

  var buffer io.Reader = nil

  if body != nil {
    buffer = bytes.NewBuffer([]byte(body))
  }

  req, err := http.NewRequest(requestType, client.baseUrl + path, buffer)

  if err != nil {
    log.Fatal(err)
  }

  if client.username != "" {
    req.SetBasicAuth(client.username, client.password)
  }

  if requestType != "GET" {
    req.Header.Set("Content-Type", "application/json")
  }

  resp, err := httpClient.Do(req)
  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }


  bodyText, err := ioutil.ReadAll(resp.Body)

  var data map[string]interface{}
  json.Unmarshal([]byte(string(bodyText)), &data)

  return data, nil
}

  
