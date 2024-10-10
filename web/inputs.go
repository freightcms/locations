package web

import "github.com/graphql-go/graphql"

var (
	CreateLocationInput = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateLocationInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"line1": &graphql.InputObjectFieldConfig{
				Description: Line1Field.Description,
				Type:        Line1Field.Type,
			},
			"line2": &graphql.InputObjectFieldConfig{
				Description:  Line2Field.Description,
				Type:         Line2Field.Type,
				DefaultValue: nil,
			},
			"line3": &graphql.InputObjectFieldConfig{
				Description:  "Typically a bin within a warehouse or a floor number and building section",
				Type:         graphql.String,
				DefaultValue: nil,
			},
			"locale": &graphql.InputObjectFieldConfig{
				Description: LocaleField.Description,
				Type:        LocaleField.Type,
			},
			"region": &graphql.InputObjectFieldConfig{
				Description: RegionField.Description,
				Type:        RegionField.Type,
			},
			"postalCode": &graphql.InputObjectFieldConfig{
				Description: PostalCodeField.Description,
				Type:        PostalCodeField.Type,
			},
			"countryCode": &graphql.InputObjectFieldConfig{
				Description: CountryField.Description,
				Type:        CountryField.Type,
			},
			"attention": &graphql.InputObjectFieldConfig{
				Description: AttentionField.Description,
				Type:        AttentionField.Type,
			},
		},
	})
)
