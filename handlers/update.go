package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"securigroup/tech-test/utils"
	"strings"
)

// Finds the user by id and updates them with new values
func UpdateEmployee(db *sql.DB, args map[string]interface{}) (string, error) {
    query := "UPDATE SecuriGroup.employees "
    setQueries := make([]string, 0)
    for column, newValue := range args {
        if column == "id" {
            continue
        }

        newQuery := fmt.Sprintf("SET %s = '%s'", utils.ConvertCamelToSnake(column), newValue);
        setQueries = append(setQueries, newQuery);
    }

    id, ok := args["id"].(int) 
    if !ok {
        return "", errors.New("The provided id is invalid, please provide a valid id")
    }

    query += strings.Join(setQueries, ", ") + fmt.Sprintf(" WHERE id = %d", id)
    if _, err := db.Exec(query); err != nil {
        return "", err
    }

    return "Successfully updated employee", nil
}
