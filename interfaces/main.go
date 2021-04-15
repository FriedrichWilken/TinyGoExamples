package main

import "github.com/friedrichwilken/tinygoexamples/interfaces/mammals"

func main() {
	peter := mammals.NewDog("Pete", "Husky", 5, "grey")
	apple := mammals.NewHuman("Lauri", "Apple", 25)

	mammals.Introduce(peter)
	mammals.Introduce(apple)
}
