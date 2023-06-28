package main

import (
	"log"
	"securigroup/tech-test/database"
	"securigroup/tech-test/graph"
	"securigroup/tech-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    app := fiber.New()

    database, err := database.CreateConnection();
    if err != nil {
        log.Fatal(err)    
    }

    rootQuery := graph.CreateRootQuery(database)
    rootMutation := graph.CreateRootMutation(database)
    schema, err := graphql.NewSchema(graphql.SchemaConfig{
        Query: rootQuery,
        Mutation: rootMutation,
    })

    if err != nil {
        log.Fatal(err.Error())
    }

    routes.CreateEmployeeRoute(app, database, schema);
    log.Fatal(app.Listen(":8080"))
}
