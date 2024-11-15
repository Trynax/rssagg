package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load()

	portString:=os.Getenv("PORT")

	if portString==""{
	   log.Fatal("PORT is not found in the enviroment")
	}
	fmt.Println("Port:", portString)
}