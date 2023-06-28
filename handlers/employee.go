package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"securigroup/tech-test/database"
	"strings"
)

type UnauthorizedError struct {
	Msg string
}

func (e *UnauthorizedError) Error() string {
	return e.Msg
}

type CustomError struct {
	Msg    string
	Status int
}

func (e *CustomError) Error() string {
	return e.Msg
}

// Gets all employees with specific filters
func GetAllEmployees(db *sql.DB, filters map[string]interface{}, sort string, amount uint32, page uint32) ([]database.Employee, error) {
	var employees []database.Employee
	query := "SELECT * FROM SecuriGroup.employees"
	whereClauses := []string{}

	// Gets the users with a specific department id
	if departmentId, ok := filters["departmentIdEquals"].(int); ok {
		whereClauses = append(whereClauses, fmt.Sprint("department_id = ", departmentId))
	}

	// Gets the users with a specific position
	if position, ok := filters["positionEquals"].(string); ok {
		whereClauses = append(whereClauses, fmt.Sprint("position = '", position, "'"))
	}

	// Makes sure theres all the where clauses
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Sorts them by specifics
	if sort != "" {
		query += " ORDER BY " + sort
	}

	// Pages
	if amount > 0 {
		// Makes sure we are starting from the beginning if nothing provided
		if page <= 0 {
			page = 1
		}

		// Gets the xth page depending on the amount we want
		offset := (page - 1) * amount
		query += fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, amount)
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee database.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.FirstName,
			&employee.LastName,
			&employee.Password,
			&employee.Email,
			&employee.DateOfBirth,
			&employee.DepartmentId,
			&employee.Position,
			&employee.Username,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func GetEmployeeById(db *sql.DB, id uint32) (database.Employee, error) {
	var employee database.Employee
	row := db.QueryRow("SELECT * FROM SecuriGroup.employees WHERE id = ?", id).Scan(
		&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.Password,
		&employee.Email,
		&employee.DateOfBirth,
		&employee.DepartmentId,
		&employee.Position,
		&employee.Username,
	)

	if row != nil {
		if errors.Is(row, sql.ErrNoRows) {
			return employee, &CustomError{
                Msg: fmt.Sprintf("Could not find employee with the id: %d", id),
                Status: 404,
            }
		}
	}

	return employee, nil
}
