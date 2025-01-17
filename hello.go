package main

import (
	"fmt"
    "github.com/NEkaterinaVA/greetings"
	"log"
    "math/rand"
)

func main() {
    log.Default().SetPrefix("greetings: ")
    log.SetFlags(0)
    
    message, err := greetings.Hello("Mary", rand.Intn(3))
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(message)
}