package database

type Position string

const (
	Junior Position = "Junior"
	Senior Position = "Senior"
	Leader Position = "Leader"
)

type Employee struct {
	Id           uint32
	FirstName    string
	LastName     string
	Password     string
	Email        string
	DateOfBirth  string
	DepartmentId uint32
	Position     string
    Username     string
}

func (e Employee) WithoutPassword() Employee {
    e.Password = ""
    return e
}
