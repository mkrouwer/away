package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"log"
)

func main() {

	groups := []string{"lunchBunch", "frisbeeCrew"}
	groupTypes := []string{"FOOD", "SPORTS"}

	var groupTypesList = graphql.EnumValueConfigMap{
		"FOOD": &graphql.EnumValueConfig{Value: "FOOD"},
		"SPORTS": &graphql.EnumValueConfig{Value: "SPORTS"},
		"HOBBIES": &graphql.EnumValueConfig{Value: "HOBBIES"},
		"CODING": &graphql.EnumValueConfig{Value: "CODING"},
	}

	var groupTypesEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:   "OrderBy",
		Values: groupTypesList,
	})


	// Schema
	fields := graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Margot Krouwer", nil
			},
		},
		"role": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Senior Software Engineer", nil
			},
		},
		"homeCampus": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "New York, New York", nil
			},
		},
		"currentCampus": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Cork, Ireland", nil
			},
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return groups, nil
			},
		},
		"groupTypes": &graphql.Field{
			Type: graphql.NewList(groupTypesEnum),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return groupTypes, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}