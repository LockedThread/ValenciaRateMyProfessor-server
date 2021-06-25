package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"server/schema"
)

func main() {
	// SchemaObj
	schemaConfig := graphql.SchemaConfig{
		Query:    schema.QueryType,
		Types:    []graphql.Type{schema.ProfessorType},
	}
	sc, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalln(err)
	}
	h := handler.New(&handler.Config{
		Schema:     &sc,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Running server on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		_ = fmt.Errorf("unable to handle exception: %s", err)
	}
}
