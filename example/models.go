package example

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Company struct {
	Name      string
	Employees int
	Revenue   float64
	Public    bool
}
