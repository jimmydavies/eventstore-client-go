package eventstore

import (
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
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

func (client *Client) makeRequest(requestType string, path string, body string) (map[string]interface{}, error) {

  httpClient := &http.Client{}

  req, err := http.NewRequest(requestType, client.baseUrl + "/" + path, nil)

  if err != nil {
    log.Fatal(err)
  }

  if client.username != "" {
    req.SetBasicAuth(client.username, client.password)
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

  
