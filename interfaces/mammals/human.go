package mammals

type Human struct {
	FristName string
	LastName  string
	Age       int
	HairColor string
}

func NewHuman(firstname, lastName string, age int) *Human {
	return &Human{
		FristName: firstname,
		LastName:  lastName,
		Age:       age,
	}
}

func (h *Human) name() string {
	return h.FristName + " " + h.LastName
}

func (h *Human) description() string {
	return "human"
}

func (h *Human) age() int {
	return h.Age
}
