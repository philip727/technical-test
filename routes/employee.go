package routes

import (
	"database/sql"
	"errors"
	"regexp"
	"securigroup/tech-test/handlers"
	"securigroup/tech-test/middleware/jwtauth"
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
		return c.SendStatus(400)
	}

	if !validateRequiredFields(payload) {
		return c.Status(400).SendString("Please fill in all required fields")
	}

	if !validateUsername(payload.Username) {
		return c.Status(400).SendString("Invalid username, a username can only consist of alphanumerical letters")
	}

    // Gets a token from logging in
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

// Creates a route, ideally i would put them in a group called "/employee"
func CreateEmployeeRoute(f *fiber.App, db *sql.DB, s graphql.Schema) {
	f.Post("/login", func(c *fiber.Ctx) error {
		return loginEmployee(c, db)
	})

    // Uses jwt auth for this route
	f.Post("/employee", jwtauth.New(), func(c *fiber.Ctx) error {
        var requestBody struct {
            Query string `json:"query"`
        }

        if err := c.BodyParser(&requestBody); err != nil {
            return err
        }
        
		result := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: requestBody.Query,
		})

		if len(result.Errors) > 0 {
            mainError := result.Errors[0]

			return c.Status(404).SendString(mainError.Message)
		}

		return c.JSON(result)
	})
}
