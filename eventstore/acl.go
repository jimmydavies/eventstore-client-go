package eventstore

import (
  "log"
)

type DefaultACLs struct {
  UserStreamACL struct {
    Read          []string `json:"$r"`
    Write         []string `json:"$w"`
    Delete        []string `json:"$d"`
    MetadataRead  []string `json:"$mr"`
    MetadataWrite []string `json:"$mw"`
  } `json:"$userStreamAcl"`
  SystemStreamACL struct {
    Read          []string `json:"$r"`
    Write         []string `json:"$w"`
    Delete        []string `json:"$d"`
    MetadataRead  []string `json:"$mr"`
    MetadataWrite []string `json:"$mw"`
  } `json:"$systemStreamAcl"`
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
 
