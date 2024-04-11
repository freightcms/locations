package graph

import (
	"github.com/freightcms/locations/models"
	"github.com/graphql-go/graphql"
)

var (
	AddressType = graphql.NewEnum(graphql.EnumConfig{
		Name: "AddressType",
		Values: graphql.EnumValueConfigMap{
			"Physical": &graphql.EnumValueConfig{
				Value: models.Physical,
			},
			"Mailing": &graphql.EnumValueConfig{
				Value: models.Mailing,
			},
			"Billing": &graphql.EnumValueConfig{
				Value: models.Billing,
			},
			"Shipping": &graphql.EnumValueConfig{
				Value: models.Shipping,
			},
			"Work": &graphql.EnumValueConfig{
				Value: models.Work,
			},
			"Home": &graphql.EnumValueConfig{
				Value: models.Home,
			},
			"Virtual": &graphql.EnumValueConfig{
				Value: models.Virtual,
			},
			"Other": &graphql.EnumValueConfig{
				Value: models.Other,
			},
		},
	})
	CoordinatesType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Coordinates",
		Fields: graphql.Fields{
			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
			"longitude": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})
	AddressModelType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AddressModel",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"line1": &graphql.Field{
				Type: graphql.String,
			},
			"line2": &graphql.Field{
				Type: graphql.String,
			},
			"line3": &graphql.Field{
				Type: graphql.String,
			},
			"local": &graphql.Field{
				Type: graphql.String,
			},
			"region": &graphql.Field{

				Type: graphql.String,
			},
			"postalCode": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: CountryCodeType,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"attention": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: AddressType,
			},
		},
	})
)
