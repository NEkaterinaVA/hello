package main

import (
	"fmt"
    "github.com/NEkaterinaVA/greetings"
	"log"
)

func main() {
    log.Default().SetPrefix("greetings: ")
    log.SetFlags(0)
    names := []string{"Gladys", "Samantha", "Darrin"}
    
    messages, err := greetings.Hellos(names)
    if err != nil{
        log.Fatal(err)
    }

    fmt.Println(messages)
}