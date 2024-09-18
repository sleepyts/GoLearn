package lib1

type A interface {
	GetName() string
	GetAge() int
	SetName(name string)
	SetAge(age int)
}
type Person struct {
	Name     string
	Age      int
	Password string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}