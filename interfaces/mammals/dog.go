package mammals

type Dog struct {
	Nickname string
	Species  string
	Age      int
	FurColor string
}

func NewDog(nickname string, species string, age int, furcolor string) *Dog {
	return &Dog{
		Nickname: nickname,
		Species:  species,
		Age:      age,
		FurColor: furcolor,
	}
}

func (d *Dog) name() string {
	return d.Nickname
}

func (d *Dog) description() string {
	return d.Species
}

func (d *Dog) age() int {
	return d.Age
}
