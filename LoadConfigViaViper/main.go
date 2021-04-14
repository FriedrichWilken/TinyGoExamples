package main

import (
	"fmt"

	"github.com/FriedrichWilken/TinyGoExamples/LoadConfigViaViper/config"
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
