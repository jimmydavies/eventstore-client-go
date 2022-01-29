package main

import (
    "fmt"
    "os"

    "github.com/madedotcom/eventstore-client-go/eventstore"
)

func main() {
    client, err := eventstore.NewClient("http://eventstore.service.test.consul:2113", "admin", os.Getenv("EVENTSTORE_PASSWORD"))
    
    if err != nil {
      fmt.Println("Error creating client")
      os.Exit(1)
    }

    user, err := client.GetUser("admin")

    fmt.Println(user.FullName)
    fmt.Println(user.UserName)
    fmt.Println(user.Groups)

    users, err := client.GetAllUsers()

    for _,user := range users {
      fmt.Println(user.FullName)
    }

    newUser, err := client.CreateUser("test-123", "test-123", "Test 123", []string{"developers"})

    fmt.Println(newUser.FullName)

    updatedUser, err := client.UpdateUser("test-123", "Test 456", []string{"support"})

    fmt.Println(updatedUser.Groups)

    updatedUser, err = client.DisableUser("test-123")
    fmt.Println(updatedUser.Disabled)
    updatedUser, err = client.EnableUser("test-123")
    fmt.Println(updatedUser.Disabled)
    fmt.Println(client.SetUserPassword("test-123", "new-password"))

    fmt.Println(client.DeleteUser("test-123"))
}
