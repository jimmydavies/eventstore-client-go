package eventstore

import (
  "testing"
)

func TestGetUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := &MockHttpClient{}

  client.httpClient = httpClient

  httpClient.setHttpClientResponse(`
    {
      "data": {
        "loginName": "test-123",
        "fullName": "Test User 123",
        "groups": ["developer"],
        "disabled": false
      }
    }`, 200, "200 OK", nil)

  client.GetUser("test-123")
}

func TestGetAllUsers(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := &MockHttpClient{}

  client.httpClient = httpClient

  httpClient.setHttpClientResponse(`
    {
      "data": [
        {
          "loginName": "test-123",
          "fullName": "Test User 123",
          "groups": ["developer"],
          "disabled": false
        },
        {
          "loginName": "test-456",
          "fullName": "Test User 456",
          "groups": ["support"],
          "disabled": false
        }
      ]
    }`, 200, "200 OK", nil)

  users, err := client.GetAllUsers()

  if err != nil {
    t.Errorf("Unexpected error (%s)", err.Error())
  }

  if users[0].UserName != "test-123" || users[1].UserName != "test-456" {
    t.Errorf("Unexpected users returned (%v)", users)
  }
}

