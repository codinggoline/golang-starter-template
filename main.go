package main

import (
	"fmt"
	"golang_starter_template/pkg/session"
)

func main() {
	fmt.Println(string(session.SecretKey))
}
