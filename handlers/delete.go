package handlers

import (
	"database/sql"
)

func DeleteEmployee(db *sql.DB, id uint32) (bool, error) {
    query := "DELETE FROM SecuriGroup.employees WHERE id = ?"

    result, err := db.Exec(query, id);
    if err != nil {
        return false, err
    }

    if rows, _ := result.RowsAffected(); rows == 0 {
        return false, nil
    }

    return true, nil
}
