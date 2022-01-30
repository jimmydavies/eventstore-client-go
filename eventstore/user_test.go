package eventstore

import (
  "testing"
  "container/list"
)

func TestGetUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := &MockHttpClient{
    responses: list.New(),
    errs: list.New(),
  }

  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
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

  httpClient := &MockHttpClient{
    responses: list.New(),
    errs: list.New(),
  }

  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
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

func TestCreateUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := &MockHttpClient{
    responses: list.New(),
    errs: list.New(),
  }

  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "loginName": "test-123",
      "success": true,
      "error": "Success"
    }`, 200, "200 OK", nil)

  httpClient.addHttpClientResponse(`
    {
      "data": {
        "loginName": "test-123",
        "fullName": "Test User 123",
        "groups": ["developer"],
        "disabled": false
      }
    }`, 200, "200 OK", nil)

  client.CreateUser("test-123", "mypass", "Test 123", []string{"developers"})
}

