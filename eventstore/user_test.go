package eventstore

import (
  "testing"
)

func TestGetUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
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

  httpClient := NewMockHttpClient()
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

  httpClient := NewMockHttpClient()
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
        "groups": ["developers"],
        "disabled": false
      }
    }`, 200, "200 OK", nil)

  client.CreateUser("test-123", "mypass", "Test User 123", []string{"developers"})
}

func TestDeleteUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "loginName": "test-123",
      "success": true,
      "error": "Success"
    }`, 200, "200 OK", nil)

  client.DeleteUser("test-123")
}

func TestUpdateUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
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
        "fullName": "Test User",
        "groups": ["support"],
        "disabled": false
      }
    }`, 200, "200 OK", nil)

  client.UpdateUser("test-123", "Test User", []string{"support"})
}

func TestEnableUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
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
        "fullName": "Test User",
        "groups": ["support"],
        "disabled": false
      }
    }`, 200, "200 OK", nil)

  client.EnableUser("test-123")
}

func TestDisableUser(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
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
        "fullName": "Test User",
        "groups": ["support"],
        "disabled": true
      }
    }`, 200, "200 OK", nil)

  client.DisableUser("test-123")
}

func TestSetUserPassword(t *testing.T) {

  client, _ := NewClient("http://eventstore.hostname:2113", "myuser", "mypass")

  httpClient := NewMockHttpClient()
  client.httpClient = httpClient

  httpClient.addHttpClientResponse(`
    {
      "loginName": "test-123",
      "success": true,
      "error": "Success"
    }`, 200, "200 OK", nil)

  client.SetUserPassword("test-123", "new-password")
}
