package handlers

import (
	"database/sql"
	"errors"
	"securigroup/tech-test/database"
	"securigroup/tech-test/types"
	"securigroup/tech-test/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// The json that is held in the JWt
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
