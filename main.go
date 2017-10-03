package main

import (
	"github.com/graphql-go/handler"
	"net/http"
	"./schema"
)

func main() {

	h := handler.New(&handler.Config{
		Schema: &schema.EmployeeSchema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}