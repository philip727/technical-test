package graph

import (
	"database/sql"
	"errors"
	"securigroup/tech-test/handlers"

	"github.com/graphql-go/graphql"
)

var CreateEmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateEmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"dateOfBirth": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"departmentId": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"position": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(PositionEnum),
		},
	},
})

var UpdateEmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateEmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"username": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"dateOfBirth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"departmentId": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"position": &graphql.InputObjectFieldConfig{
			Type: PositionEnum,
		},
	},
})

func CreateRootMutation(db *sql.DB) *graphql.Object {
	var RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createEmployee": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(CreateEmployeeInputType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args, ok := p.Args["input"].(map[string]interface{})
					if !ok {
						return nil, errors.New("Invalid input arguments")
					}

					return handlers.CreateEmployee(db, args)
                },
			},
			"updateEmployee": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(UpdateEmployeeInputType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args, ok := p.Args["input"].(map[string]interface{})
					if !ok {
						return nil, errors.New("Invalid input arguments")
					}

					return handlers.UpdateEmployee(db, args)
				},
			},
			"deleteEmployee": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return false, errors.New("Invalid input arguments")
					}

					ok, err := handlers.DeleteEmployee(db, uint32(id))
					if err != nil {
						return false, err
					}
                    
                    // This library is pretty limited and you can not send custom status codes from it
					if !ok {
						return false, errors.New("No employee exists with the id provided")
					}

					return true, nil
				},
			},
		},
	})

	return RootMutation
}
