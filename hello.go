package main

import (
	"fmt"
    "github.com/NEkaterinaVA/greetings"
	"log"
)

func main() {
    log.Default().SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("Mary")
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(message)
}