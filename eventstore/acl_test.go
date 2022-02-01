package eventstore

import (
  "testing"
  "log"
)

func TestReadDefaultACLs(t *testing.T) {
  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "$userStreamAcl": {
        "$r": ["$admins"],
        "$w": ["$admins"],
        "$d": ["$admins"],
        "$mr": ["$admins"],
        "$mw": ["$admins"]
      },
      "$systemStreamAcl": {
        "$r": ["$admins"],
        "$w": ["$admins"],
        "$d": ["$admins"],
        "$mr": ["$admins"],
        "$mw": ["$admins"]
      },
      "something": "else"
    }`, 200, "200 OK", nil)

  defaultACLs, err := client.ReadDefaultACLs()

  if err != nil {
    t.Errorf(err.Error())
    return
  }

  log.Print(defaultACLs.UserStreamACL.Read)
}

