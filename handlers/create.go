package handlers

import (
	"database/sql"
	"fmt"
	"securigroup/tech-test/smtp"
	"securigroup/tech-test/utils"
)

func CreateEmployee(db *sql.DB, args map[string]interface{}) (string, error) {
	query := "INSERT INTO SecuriGroup.employees"
	columnsString := ""
	valuesString := ""

	for index, pair := range utils.Enumerate(args) {
		column := pair.Key
		value := pair.Value

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
		return "", err
	}

	subject := fmt.Sprintf("Welcome to SecuriGroup, %s%s", args["firstName"].(string), "!")
	// The name is completely random :)
	body := fmt.Sprintf(
		`Dear %s 

On behalf of the entire team at SecuriGroup, I would like to extend a warm welcome to you!
We are thrilled to have you apart of the team.

Please take the time to familiarize yourself with our company culture, values and policies.
If you need anything please contact our HR team: "HR@SecuriGroup.co.uk" and they'll be happy to help.

Once again, welcome to SecuriGroup, we are excited to have you on the team and look forward to the
contributions you will make.

Best regards,

SecuriGroup
`, args["firstName"].(string))

	if err := smtp.SendNoreplyMail(args["email"].(string), subject, body); err != nil {
		fmt.Printf("Failed to send an email to %s", args["email"].(string))
	}

	return "Successfully created a new employee", nil
}
