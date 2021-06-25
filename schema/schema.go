package schema

import (
	"github.com/graphql-go/graphql"
	"server/database"
	"server/models"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"professorById": &graphql.Field{
				Type:        models.ProfessorType,
				Description: "Get professor by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					professorById := database.GetProfessorById(id)
					return professorById, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(models.ProfessorType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return database.GetProfessors(), nil
				},
			},
		},
	},
)
