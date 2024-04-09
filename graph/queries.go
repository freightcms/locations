package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	CurrencyObjectConfig = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Currency",
			Fields: graphql.Fields{
				"code": &graphql.Field{
					Name: "code",
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Name: "name",
					Type: graphql.String,
				},
				"symbol": &graphql.Field{
					Name: "symbol",
					Type: graphql.String,
				},
			},
		})
	CountryFields = &graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Name: "code",
			Type: graphql.EnumValueType,
		},
		"language": &graphql.Field{
			Name: "language",
			Type: graphql.String,
		},
		"region": &graphql.Field{
			Name: "region",
			Type: graphql.String,
		},
		"subregion": &graphql.Field{
			Name: "subregion",
			Type: graphql.String,
		},
		"capital": &graphql.Field{
			Name: "capital",
			Type: graphql.String,
		},
		"callingCode": &graphql.Field{
			Name: "callingCode",
			Type: graphql.String,
		},
		"flag": &graphql.Field{
			Name: "flag",
			Type: graphql.String,
		},
		"flagEmoji": &graphql.Field{
			Name: "flagEmoji",
			Type: graphql.String,
		},
		"flagEmojiUnicode": &graphql.Field{
			Name: "flagEmojiUnicode",
			Type: graphql.String,
		},
		"currency": &graphql.Field{
			Name: "currency",
			Type: CurrencyObjectConfig,
		},
	}
)
