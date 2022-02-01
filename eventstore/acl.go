package eventstore

import (
  "log"
  "encoding/json"
)

type StreamACL struct {
  Read          []string `json:"$r"`
  Write         []string `json:"$w"`
  Delete        []string `json:"$d"`
  MetadataRead  []string `json:"$mr"`
  MetadataWrite []string `json:"$mw"`
}

type DefaultACLs struct {
  UserStreamACL   StreamACL `json:"$userStreamAcl"`
  SystemStreamACL StreamACL `json:"$systemStreamAcl"`
}

func (client *Client) ReadDefaultACLs() (*DefaultACLs, error) {
  var data DefaultACLs
  err := client.makeRequest("GET", "/streams/$settings/head", nil, &data)

  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  return &data, nil

}

func (client *Client) WriteDefaultACLs(newACLs DefaultACLs) (*DefaultACLs, error) {
  jsonBody, err := json.Marshal(newACLs)

  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  var response map[string]interface{}
  err = client.makeRequest("POST", "/streams/$settings", jsonBody, &response)

  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  return client.ReadDefaultACLs()
}
