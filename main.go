package main

import (
    "fmt"
    "os"

    "github.com/madedotcom/eventstore-client-go/eventstore"
)

func main() {
    fmt.Println("Hello, Modules!")

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

}
