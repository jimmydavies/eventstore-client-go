package eventstore

import (
  "log"
)

type User struct{
  UserName string
  FullName string
  Groups []string
}

func (client *Client) GetUser(username string) (*User, error) {
  data, err := client.makeRequest("GET", "users/" + username, "")

  if err != nil {
    log.Fatal(err)
  }

  var groups []string
  for _, group := range data["data"].(map[string]interface{})["groups"].([]interface{}) {
    groups = append(groups, group.(string))
  }

  return getUserFromMap(data["data"].(map[string]interface{})), nil
}

func (client *Client) GetAllUsers() ([]User, error) {
  data, err := client.makeRequest("GET", "users", "")

  if err != nil {
    log.Fatal(err)
  }

  var users []User
  for _, user := range data["data"].([]interface{}) {
    users = append(users, *getUserFromMap(user.(map[string]interface{})))
  }

  return users, nil
   
}

func getUserFromMap(userData map[string]interface{}) (*User) {
  var groups []string
  for _, group := range userData["groups"].([]interface{}) {
    groups = append(groups, group.(string))
  }

  return &User{
    UserName: userData["loginName"].(string),
    FullName: userData["fullName"].(string),
    Groups: groups,
  }
}
