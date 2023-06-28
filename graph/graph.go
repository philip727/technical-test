package graph

import (
	"github.com/graphql-go/graphql"
)

var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
			// I feel like it makes sense to remove the password from the queries
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"dateOfBirth": &graphql.Field{
			Type: graphql.String,
		},
		"departmentId": &graphql.Field{
			Type: graphql.Int,
		},
		"position": &graphql.Field{
			Type: PositionEnum,
		},
	},
})

var PositionEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "Position",
	Values: graphql.EnumValueConfigMap{
		"JUNIOR": &graphql.EnumValueConfig{
			Value: "Junior",
		},
		"SENIOR": &graphql.EnumValueConfig{
			Value: "Senior",
		},
		"LEADER": &graphql.EnumValueConfig{
			Value: "Leader",
		},
	},
})
