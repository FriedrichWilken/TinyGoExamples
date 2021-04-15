package mammals

import "fmt"

type Informer interface {
	description() string
	name() string
	age() int
}

func Introduce(i Informer) {
	fmt.Printf("Hello, my name is %v and I am a %v. I am %v years old \n", i.name(), i.description(), i.age())
}
