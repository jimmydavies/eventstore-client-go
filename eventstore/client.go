package eventstore

import (
  "net/http"
  "net/url"
  "log"
  "io/ioutil"
  "io"
  "encoding/json"
  "bytes"
  "os"
  "errors"
)

type Client struct {
  baseUrl    string
  username   string
  password   string
  httpClient interface{
     Do (req *http.Request) (*http.Response, error)
  }
}

func NewClient(baseUrl string, userName string, password string) (*Client, error) {
  _, err := url.ParseRequestURI(baseUrl)

  if err != nil {
    return nil, err
  }

  if userName == "" {
    userName = os.Getenv("EVENTSTORE_USER")
  }

  if password == "" {
    password = os.Getenv("EVENTSTORE_PASSWORD")
  }

  return &Client{
    baseUrl: baseUrl,
    username: userName,
    password: password,
  }, nil
}

func (client *Client) makeRequest(requestType string, path string, body []byte) (map[string]interface{}, error) {

  if client.httpClient == nil {
    client.httpClient = &http.Client{}
  }

  var buffer io.Reader = nil

  if body != nil {
    buffer = bytes.NewBuffer([]byte(body))
  }

  req, err := http.NewRequest(requestType, client.baseUrl + path, buffer)

  if err != nil {
    log.Print(err)
    return nil, err
  }

  if client.username != "" {
    req.SetBasicAuth(client.username, client.password)
  }

  if requestType != "GET" {
    req.Header.Set("Content-Type", "application/json")
  }

  resp, err := client.httpClient.Do(req)

  if err != nil {
    log.Print(err)
    return nil, err
  }

  defer resp.Body.Close()

  // Handle Status Codes
  if resp.StatusCode != 200 {
    return nil, errors.New("Request failed with status code " + resp.Status)
  }

  bodyText, err := ioutil.ReadAll(resp.Body)

  var data map[string]interface{}
  json.Unmarshal([]byte(string(bodyText)), &data)

  return data, nil
}

  
