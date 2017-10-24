package schema

import (
	"github.com/graphql-go/graphql"
)

var (
	groupDatabase = map[string]string{"lunchBunch":"LUNCH", "frisbeeCrew":"SPORTS", "chuckNorrisFanClub":"HOBBIES", "helloWorldGoodbyeSocialLife":"CODING"}
	groups = []string{"lunchBunch", "frisbeeCrew"}
	groupTypes = []string{"FOOD", "SPORTS"}
	groupSuggestions = []string{}
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
		"BOSTON": &graphql.EnumValueConfig{Value: "BOSTON"},
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
	HOME_CAMPUS = "BOSTON"
	CURRENT_CAMPUS = "NEW_YORK_CITY"
)

	var employeeType = graphql.NewObject(graphql.ObjectConfig{
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
			"groupSuggestions": &graphql.Field{
				Type:     graphql.NewList(graphql.String),
				Description: "Suggestions for groups the user might enjoy.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return groupSuggestions, nil
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

	var queryType = graphql.NewObject(graphql.ObjectConfig{
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

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addGroup": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"groupName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"groupType": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := make(map[string]string)
					for k, v := range p.Args {
						args[k] = v.(string)
					}
					addGroup(args)
					return p.Source, nil
				},
			},
		},
	})

func addGroup(args map[string]string) {
	groupName := args["groupName"]
	groupType := args["groupType"]
	groups = append(groups, groupName)
	for _, group := range groupTypes{
		if(group == groupType){
			return
		}
	}
	groupTypes = append(groupTypes, groupType)
}



	var EmployeeSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})


