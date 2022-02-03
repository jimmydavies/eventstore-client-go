package eventstore

import (
  "log"
  "encoding/json"
  "errors"
)

type User struct{
  UserName string
  Password string
  FullName string
  Groups []string
  Disabled bool
}

func (client *Client) GetUser(username string) (*User, error) {
  var data map[string]interface{}
  err := client.makeRequest("GET", "/users/" + username, nil, &data)

  if err != nil {
    log.Print(err.Error())

    re, ok := err.(*RequestError)
    if ok && re.StatusCode == 404 {
      //Not Found
      return nil, nil
    }
      
    return nil, err
  }

  var groups []string
  for _, group := range data["data"].(map[string]interface{})["groups"].([]interface{}) {
    groups = append(groups, group.(string))
  }

  return getUserFromMap(data["data"].(map[string]interface{})), nil
}

func (client *Client) GetAllUsers() ([]User, error) {
  var data map[string]interface{}
  err := client.makeRequest("GET", "/users", nil, &data)

  if err != nil {
    log.Print(err.Error())
    return nil, err
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

  var data map[string]interface{}
  err := client.makeRequest("POST", "/users", userData, &data)

  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  if data["success"].(bool) == false {
    err = errors.New(data["error"].(string))
    log.Print(err.Error())
    return nil, err
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) DeleteUser(userName string) (bool) {
  var data map[string]interface{}
  err := client.makeRequest("DELETE", "/users/" + userName, nil, &data)

  if err != nil {
    log.Print(err.Error())
    return false
  }

  return data["success"].(bool)
}

func (client *Client) UpdateUser(userName string, fullName string, groups []string) (*User, error) {
  userData, _ := json.Marshal(map[string]interface{}{
    "FullName": fullName,
    "Groups": groups,
  })

  var data map[string]interface{}
  err := client.makeRequest("PUT", "/users/" + userName, userData, &data)

  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  if data["success"].(bool) == false {
    err = errors.New(data["error"].(string))
    log.Print(err.Error())
    return nil, err
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) EnableUser(userName string) (*User, error) {
  var data map[string]interface{}
  err := client.makeRequest("POST", "/users/" + userName + "/command/enable", nil, &data)
  
  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  if data["success"].(bool) == false {
    err = errors.New(data["error"].(string))
    log.Print(err.Error())
    return nil, err
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) DisableUser(userName string) (*User, error) {
  var data map[string]interface{}
  err := client.makeRequest("POST", "/users/" + userName + "/command/disable", nil, &data)
  
  if err != nil {
    log.Print(err.Error())
    return nil, err
  }

  if data["success"].(bool) == false {
    err = errors.New(data["error"].(string))
    log.Print(err.Error())
    return nil, err
  }

  return client.GetUser(data["loginName"].(string))
}

func (client *Client) SetUserPassword(userName string, password string) (bool) {
  userData, _ := json.Marshal(map[string]interface{}{
    "NewPassword": password,
  })

  var data map[string]interface{}
  err := client.makeRequest("POST", "/users/" + userName + "/command/reset-password", userData, &data)

  if err != nil {
    log.Print(err.Error())
    return false
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
