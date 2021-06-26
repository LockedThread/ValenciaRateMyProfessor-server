package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"os"
	"server/database"
	"server/models"
	"server/schema"
	"strconv"
)

func main() {
	database.Client = *database.Connect()

	// SchemaObj
	schemaConfig := graphql.SchemaConfig{
		Query: schema.QueryType,
		Types: []graphql.Type{models.ProfessorType},
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
	port, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 0, 16)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Started server on port: %d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		_ = fmt.Errorf("unable to handle exception: %s", err)
	}
}
