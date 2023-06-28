package jwtauth

import (
	"securigroup/tech-test/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Verifies the jwt in the authorisation header
func verifyJWT(token string, c *fiber.Ctx) (bool, string) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnvVar("JWT_SECRET")), nil
	}

	parsedToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return false, err.Error()
	}

	if !parsedToken.Valid {
		return false, "Token is invalid"
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
    exp, ok := claims["exp"].(int64)
    if !ok {
        return false, "Invalid expiry time"
    }

	if exp > time.Now().Unix() {
		return false, "Token has expired, relogin"
	}

	// Allows us to use the username in current request context
	c.Locals("Username", claims["username"])

	return true, ""
}

// Creates the jwt auth middleware
func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetReqHeaders()["Authorisation"]
        
        authorized, reason := verifyJWT(token, c)
        if !authorized {
            return c.Status(401).SendString(reason)
        }

        // Goes on to the rest of the request
        return c.Next()
	}
}
