package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"securigroup/tech-test/smtp"
	"securigroup/tech-test/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func isValidEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Match the email against the pattern
	return regex.MatchString(email)
}

// Creates an employee in the mssql db and emails them with a welcome message
func CreateEmployee(db *sql.DB, args map[string]interface{}) (string, error) {
	query := "INSERT INTO SecuriGroup.employees"
	columnsString := ""
	valuesString := ""
    

    // These are values that are used further one thats why they are here
    // Either way the db will throw an error if its not a correct type
    firstName, ok := args["firstName"].(string) 
    if !ok {
        return "", errors.New("First name is invalid, please provide a string")
    }

    email, ok := args["email"].(string) 
    if !ok || !isValidEmail(email) {
        return "", errors.New("Email is invalid, please provide a valid email")
    }



    // Creates the query for the columns and values that need to be added
	for index, pair := range utils.Enumerate(args) {
		column := pair.Key
		value := pair.Value

        // Hashes the password
        if strings.ToLower(column) == "password" {
            pw, ok := value.(string) 
            if !ok {
                return "", errors.New("Invalid password provided, could not cast to string") 
            }


            hash, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
            if err != nil {
                return "", errors.New("Failed to hash password")
            }

            value = string(hash)
        }

		// Formats for db
		columnString := utils.ConvertCamelToSnake(column)
		valueString := fmt.Sprint("'", value, "'")

		// Makes sure that the query is valid
		if index < (len(args) - 1) {
			columnString += ", "
			valueString += ", "
		}

		columnsString += columnString
		valuesString += valueString
	}

	query += "(" + columnsString + ") VALUES (" + valuesString + ")"

	if _, err := db.Exec(query); err != nil {
		return "", errors.New("Failed to create user")
	}

	subject := fmt.Sprintf("Welcome to SecuriGroup, %s%s", firstName, "!")

	// The name is completely random :)
    // A warm welcome subject
	body := fmt.Sprintf(
		`Dear %s,

On behalf of the entire team at SecuriGroup, I would like to extend a warm welcome to you!
We are thrilled to have you apart of the team.

Please take the time to familiarize yourself with our company culture, values and policies.
If you need anything please contact our HR team: "HR@SecuriGroup.co.uk" and they'll be happy to help.

Once again, welcome to SecuriGroup, we are excited to have you on the team and look forward to the
contributions you will make.

Best regards,

SecuriGroup
`, firstName)

    // This doesn't stop execution or stop the handler, it strictly just prints an error
	if err := smtp.SendNoreplyMail(email, subject, body); err != nil {
		fmt.Printf("Failed to send an email to %s", email)
	}

	return "Successfully created a new employee", nil
}
