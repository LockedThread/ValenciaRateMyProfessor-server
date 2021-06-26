package models

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"strconv"
)

type Professor struct {
	Department      string `bson:"department"`
	InstitutionName string `bson:"institutionName"`
	FirstName       string `bson:"firstName"`
	MiddleName      string `bson:"middleName"`
	LastName        string `bson:"lastName"`
	TeacherID       int    `bson:"teacherId"`
	RatingsCount    int    `bson:"ratingCount"`
	RatingClass     string `bson:"ratingClass"`
	OverallRating   string `bson:"overallRating"`
}

func (p Professor) CalculateRating() float64 {
	float, err := strconv.ParseFloat(p.OverallRating, 64)
	if err != nil {
		return 0.0
	}
	return float * float64(p.RatingsCount)
}

func (p Professor) String() string {
	return fmt.Sprintf("Department: %s,  InstitutionName: %s, FirstName: %s, MiddleName: %s, LastName: %s, TeacherID: %d, RatingsCount: %d, RatingClass: %s, OverallRating: %s",
		p.Department,
		p.InstitutionName,
		p.FirstName,
		p.MiddleName,
		p.LastName,
		p.TeacherID,
		p.RatingsCount,
		p.RatingClass,
		p.OverallRating,
	)
}

var ProfessorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Professor",
	Fields: graphql.Fields{
		"department": &graphql.Field{
			Type: graphql.String,
		},
		"ratingClass": &graphql.Field{
			Type: graphql.String,
		},
		"ratingsCount": &graphql.Field{
			Type: graphql.Int,
		},
		"overallRating": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"middleName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
	},
})
