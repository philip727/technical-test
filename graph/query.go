package graph

import (
	"database/sql"
	"errors"
	"securigroup/tech-test/handlers"

	"github.com/graphql-go/graphql"
)

// Our employee filters
var EmployeeFilterInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeFilter",
	Fields: graphql.InputObjectConfigFieldMap{
		"departmentIdEquals": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"positionEquals": &graphql.InputObjectFieldConfig{
			Type: PositionEnum,
		},
	},
})

// The different types of sorting we can do
var EmployeeSortEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "EmployeeSort",
	Values: graphql.EnumValueConfigMap{
		"ID_ASC": &graphql.EnumValueConfig{
			Value: "id ASC",
		},
		"ID_DESC": &graphql.EnumValueConfig{
			Value: "id DESC",
		},
		"FIRST_NAME_ASC": &graphql.EnumValueConfig{
			Value: "first_name ASC",
		},
		"FIRST_NAME_DESC": &graphql.EnumValueConfig{
			Value: "first_name ASC",
		},
	},
})

// Creates the possible queries
func CreateRootQuery(db *sql.DB) *graphql.Object {
	var RootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
            // Get all employees
			"employees": &graphql.Field{
				Type: graphql.NewList(EmployeeType),
				Args: graphql.FieldConfigArgument{
					"filter": &graphql.ArgumentConfig{
						Type: EmployeeFilterInput,
					},
					"sort": &graphql.ArgumentConfig{
						Type: EmployeeSortEnum,
					},
                    // Pagination
                    "amount": &graphql.ArgumentConfig{
                        Type: graphql.Int,
                    },
                    "page": &graphql.ArgumentConfig{
                        Type: graphql.Int,
                    },
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					filters, ok := p.Args["filter"].(map[string]interface{})
                    if !ok {
                        filters = make(map[string]interface{})
                    }
                    
					sorting, ok := p.Args["sort"].(string)
                    if !ok {
                        sorting = ""
                    }

                    // The amount that will be shown
					amount, ok := p.Args["amount"].(int)
                    if !ok {
                        amount = 0 // Just shows all
                    }

                    // Just goes from the beginning if nothing shown
					page, ok := p.Args["page"].(int)
                    if !ok {
                        page = 1
                    }

					return handlers.GetAllEmployees(db, filters, sorting, uint32(amount), uint32(page))
				},
			},
            // Employee by id
			"employee": &graphql.Field{
				Type: EmployeeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, errors.New("Invalid employee ID")
					}

					return handlers.GetEmployeeById(db, uint32(id))
				},
			},
		},
	})

	return RootQuery
}
