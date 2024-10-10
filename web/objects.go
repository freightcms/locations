package web

import "github.com/graphql-go/graphql"

var (
	LocationObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id":         &IDField,
			"line1":      &Line1Field,
			"line2":      &Line2Field,
			"line3":      &Line3Field,
			"locale":     &LocaleField,
			"region":     &RegionField,
			"country":    &CountryField,
			"postalCode": &PostalCodeField,
			"attention":  &AttentionField,
		},
	})
)
