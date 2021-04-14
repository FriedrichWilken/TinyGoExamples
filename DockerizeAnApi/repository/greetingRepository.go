package repository

import "github.com/FriedrichWilken/TinyGoExamples/DockerizeAnApi/models"

type GreetingRepository models.Greeting

func New(id int, greeting string) *GreetingRepository {
	return &GreetingRepository{
		Id:       id,
		Greeting: greeting,
	}
}
