package main

import (
	"dockerk8s/dockers"
	"fmt"
)

func main() {
	fmt.Println(">>>>START")

	// err := dockers.PullImage("ubuntu:latest")
	l, err := dockers.RunContainer("ubuntu:latest", "test6", "example_byfn", "/bin/bash", []string{"username=testme"}, "/", []string{"8002"}, []string{"/opt"}, []string{"/opt"})
	fmt.Println(err, l)
}
