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
		return false, "No authorisation token was provided"
    }

	if !parsedToken.Valid {
		return false, "Token is invalid"
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
    if !ok {
        return false, "Failed to parse claims"
    }

    exp, ok := claims["exp"].(float64)
    if !ok {
        return false, "Invalid expiry time"
    }

	if int64(exp) < time.Now().Unix() {
		return false, "Token has expired, relogin"
	}

	// Allows us to use the username in current request context
	c.Locals("userId", claims["uid"])

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
