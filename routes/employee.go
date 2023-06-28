package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"securigroup/tech-test/handlers"
	"securigroup/tech-test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

// Makes sure all fields have been provided
func validateRequiredFields(lp types.LoginPayload) bool {
    return len(lp.Username) > 0 && len(lp.Password) > 0
}

// Makes sure the username only consists of alphanumerical characters
func validateUsername(u string) bool {
    regex := "^[a-zA-Z0-9]+$"
    match, _ := regexp.MatchString(regex, u)
    return match
}

func loginEmployee(c *fiber.Ctx, db *sql.DB) error {
    var payload types.LoginPayload

    if err := c.BodyParser(&payload); err != nil {
        return c.SendStatus(400);
    }

    if !validateRequiredFields(payload) {
        return c.Status(400).SendString("Please fill in all required fields")
    }

    if !validateUsername(payload.Username) {
        return c.Status(400).SendString("Invalid username, a username can only consist of alphanumerical letters")
    }

    token, err := handlers.LoginEmployeeHanlder(db, payload)
    if err != nil {
        var unauthorizedError *handlers.UnauthorizedError
        if errors.As(err, &unauthorizedError) {
            return c.Status(401).SendString(err.Error())
        }

        return c.Status(500).SendString(err.Error())
    }

	return c.Status(200).SendString(token)
}

func CreateEmployeeRoute(f *fiber.App, db *sql.DB, s graphql.Schema) {
	f.Post("/login", func(c *fiber.Ctx) error {
		return loginEmployee(c, db)
	})

    f.Post("/employee", func(c *fiber.Ctx) error {
        result := graphql.Do(graphql.Params {
            Schema: s,
            RequestString: string(c.Body()),
        })

        if len(result.Errors) > 0 {
            fmt.Print(result.Errors)
            return c.Status(400).SendString(fmt.Sprintf("GraphQL query errors: %v", result.Errors))
        }

        return c.JSON(result);
    })
}
