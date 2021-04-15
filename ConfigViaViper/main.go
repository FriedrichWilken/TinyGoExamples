package main

import (
	"fmt"

	"github.com/friedrichwilken/tinygoexamples/loadconfigviaviper/config"
)

func main() {
	config, err := config.Load("LoadConfigViaViper")
	if err != nil {
		panic(err)
	}

	for i := 0; i != config.Number; i++ {
		fmt.Printf("%v, %v \n", config.Greeting, config.Name)
	}
}
