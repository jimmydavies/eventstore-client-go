package eventstore

import (
  "log"
  "encoding/json"
)

type User struct{
  UserName string
  Password string
  FullName string
  Groups []string
  Disabled bool
}

func (client *Client) GetUser(username string) (*User, error) {
  data, err := client.makeRequest("GET", "/users/" + username, nil)

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
  data, err := client.makeRequest("GET", "/users", nil)

  if err != nil {
    log.Fatal(err)
  }

  var users []User
  for _, user := range data["data"].([]interface{}) {
    users = append(users, *getUserFromMap(user.(map[string]interface{})))
  }

  return users, nil
   
}

func (client *Client) CreateUser(userName string, password string, fullName string, groups []string) (*User, error) {
  userData, _ := json.Marshal(map[string]interface{}{
    "LoginName": userName,
    "Password": password,
    "FullName": fullName,
    "Groups": groups,
  })

  data, err := client.makeRequest("POST", "/users", userData)

  if err != nil {
    log.Fatal(err)
  }

  if data["success"] == false {
    log.Fatal(data["error"])
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) DeleteUser(userName string) (bool) {
  data, err := client.makeRequest("DELETE", "/users/" + userName, nil)

  if err != nil {
    log.Fatal(err)
  }

  return data["success"].(bool)
}

func (client *Client) UpdateUser(userName string, fullName string, groups []string) (*User, error) {
  userData, _ := json.Marshal(map[string]interface{}{
    "FullName": fullName,
    "Groups": groups,
  })

  data, err := client.makeRequest("PUT", "/users/" + userName, userData)

  if err != nil {
    log.Fatal(err)
  }

  if data["success"] == false {
    log.Fatal(data["error"])
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) EnableUser(userName string) (*User, error) {
  data, err := client.makeRequest("POST", "/users/" + userName + "/command/enable", nil)
  
  if err != nil {
    log.Fatal(err)
  }

  if data["success"] == false {
    log.Fatal(data["error"])
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) DisableUser(userName string) (*User, error) {
  data, err := client.makeRequest("POST", "/users/" + userName + "/command/disable", nil)
  
  if err != nil {
    log.Fatal(err)
  }

  if data["success"] == false {
    log.Fatal(data["error"])
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) SetUserPassword(userName string, password string) (bool) {
  userData, _ := json.Marshal(map[string]interface{}{
    "NewPassword": password,
  })

  data, err := client.makeRequest("POST", "/users/" + userName + "/command/reset-password", userData)

  if err != nil {
    log.Fatal(err)
  }

  return data["success"].(bool)
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
    Disabled: userData["disabled"].(bool),
  }
}
