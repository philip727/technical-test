package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"securigroup/tech-test/database"
	"securigroup/tech-test/types"
	"securigroup/tech-test/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UnauthorizedError struct {
	Msg string
}

func (e *UnauthorizedError) Error() string {
	return e.Msg
}

type JWTPayload struct {
	Id       uuid.UUID
	Username string
	Created  int64
	Expiry   int64
}

// Compares a string to a hash to see if they match
func comparePasswordToHash(pw string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(pw))
	return err == nil
}

// Creates the jwt payload
func createJWTPayload(e database.Employee) (JWTPayload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return JWTPayload{}, err
	}

	payload := &JWTPayload{
		Id:       id,
		Username: e.Username,
		Created:  time.Now().Unix(),
		Expiry:   time.Now().Add(time.Hour * 24 * 14).Unix(),
	}

	return *payload, nil
}

// Creates a jwt token with claims
func createJWTToken(jwtp JWTPayload) (string, error) {
	claims := jwt.MapClaims{
		"id":       jwtp.Id,
		"username": jwtp.Username,
		"exp":      jwtp.Expiry,
		"iat":      jwtp.Created,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := utils.GetEnvVar("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logs in the user and creates a jwt
func LoginEmployeeHanlder(db *sql.DB, p types.LoginPayload) (string, error) {
	// Finds the user with the same username
	query := "SELECT * FROM SecuriGroup.employees WHERE username = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Printf(err.Error())
		return "", errors.New("Failed to prepare SQL statement, contact an admin")
	}

	var employee database.Employee
	err = stmt.QueryRow(p.Username).Scan(
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
		return "", err
	}

	// Checks if the password matches
	if !comparePasswordToHash(p.Password, employee.Password) {
		return "", &UnauthorizedError{"The password provided does not match"}
	}

	payload, err := createJWTPayload(employee)
	if err != nil {
		return "", err
	}

	token, err := createJWTToken(payload)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetAllEmployees(db *sql.DB, filters map[string]interface{}, sort string, amount uint32, page uint32) ([]database.Employee, error) {
	var employees []database.Employee
    query := "SELECT * FROM SecuriGroup.employees"
    whereClauses := []string{}

    if departmentId, ok := filters["departmentIdEquals"].(int); ok {
        whereClauses = append(whereClauses, fmt.Sprint("department_id = ", departmentId))
    }

    if position, ok := filters["positionEquals"].(string); ok {
        whereClauses = append(whereClauses, fmt.Sprint("position = '", position, "'"))
    }


    if len(whereClauses) > 0 {
        query += " WHERE " + strings.Join(whereClauses, " AND ")
    }

    if sort != "" {
        query += " ORDER BY " + sort
    }

    if amount > 0 {
        // Makes sure we are starting from the beginning if nothing provided
        if page <= 0 {
            page = 1
        }

        // Gets the xth page depending on the amount we want
        offset := (page - 1) * amount
        query += fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, amount)
    }

    fmt.Println(query)
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
	row := db.QueryRow("SELECT * FROm SecuriGroup.employees WHERE id = ?", id).Scan(
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
            return employee, errors.New("Employee not found")      
        }
    }

    return employee, row
}
