package database


// Our employee table
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
