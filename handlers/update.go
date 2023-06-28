package handlers

import (
	"database/sql"
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

        newQuery := fmt.Sprint("SET ", utils.ConvertCamelToSnake(column), " = '", newValue, "'");
        setQueries = append(setQueries, newQuery);
    }

    query += strings.Join(setQueries, ", ") + fmt.Sprint(" WHERE id = ", args["id"].(int))

    if _, err := db.Exec(query); err != nil {
        return "", err
    }

    return "Successfully updated employee", nil
}
