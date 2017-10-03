package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

var (
	EmployeeSchema graphql.Schema
	groupTypesList = graphql.EnumValueConfigMap{
		"FOOD": &graphql.EnumValueConfig{Value: "FOOD"},
		"SPORTS": &graphql.EnumValueConfig{Value: "SPORTS"},
		"HOBBIES": &graphql.EnumValueConfig{Value: "HOBBIES"},
		"CODING": &graphql.EnumValueConfig{Value: "CODING"},
	}
	groupTypesEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:   "Groups",
		Values: groupTypesList,
	})
	campusList = graphql.EnumValueConfigMap{
		"NEW_YORK_CITY": &graphql.EnumValueConfig{Value: "NEW_YORK_CITY"},
		"CORK": &graphql.EnumValueConfig{Value: "CORK"},
	}
	campusEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:   "Campuses",
		Values: campusList,
	})

)

const (
	ID  = 1
	NAME = "Margot Krouwer"
	ROLE = "Senior Software Engineer"
	HOME_CAMPUS = "CORK"
	CURRENT_CAMPUS = "NEW_YORK_CITY"
)


func init() {
	groups := []string{"lunchBunch", "frisbeeCrew"}
	groupTypes := []string{"FOOD", "SPORTS"}

	employeeType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Employee",
		Description: "A company employee.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "The id of the employee.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ID, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the employee.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return NAME, nil
				},
			},
			"role": &graphql.Field{
				Type:        graphql.String,
				Description: "The job title of the employee.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ROLE, nil
				},
			},
			"homeCampus": &graphql.Field{
				Type:        campusEnum,
				Description: "The empoyee's home campus location.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return HOME_CAMPUS, nil
				},
			},
			"currentCampus": &graphql.Field{
				Type:        campusEnum,
				Description: "The empoyee's current campus location.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return CURRENT_CAMPUS, nil
				},
			},
			"groups": &graphql.Field{
				Type:     graphql.NewList(graphql.String),
				Description: "The names of the group a given employee is a part of.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return groups, nil
				},
			},
			"groupTypes": &graphql.Field{
				Type:        graphql.NewList(groupTypesEnum),
				Description: "The categories of the group a given employee is a part of.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return groupTypes, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: employeeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employeeType.Fields(), nil
				},
			},
		},
	})
	var err error
	EmployeeSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}
