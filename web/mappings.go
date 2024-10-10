package web

import "github.com/freightcms/locations/models"

// AddressFromArgs converts graphql arguments to a location address model. No validation takes place for this
// model to ensure any required or non-required fields are provided. Args can have the listed fields
//
//   - line1
//
//   - line2
//
//   - line3
//
//   - locale
//
//   - countryCode
//
//   - region
//
//   - postalCode
//
//   - attention
//
//   - description
//
//   - notes
func AddressFromArgs(locationType models.AddressType, args map[string]interface{}) *models.AddressModel {
	model := &models.AddressModel{
		Type:  locationType,
		Line1: args["line1"].(string),
	}
	if val, ok := args["line2"]; val != nil && ok {
		model.Line2 = val.(*string)
	}
	if val, ok := args["line3"]; val != nil && ok {
		model.Line3 = val.(*string)
	}
	if val, ok := args["locale"]; ok {
		model.Locale = val.(string)
	}
	if val, ok := args["countryCode"]; ok {
		model.Country = models.CountryCode(val.(string))
	}
	if val, ok := args["region"]; ok {
		model.Region = val.(string)
	}
	if val, ok := args["postalCode"]; ok {
		model.PostalCode = val.(string)
	}
	if val, ok := args["attention"]; val != nil && ok {
		model.Attention = val.(*string)
	}
	if val, ok := args["description"]; val != nil && ok {
		model.Description = val.(*string)
	}
	if val, ok := args["notes"]; val != nil && ok {
		model.Notes = val.(*string)
	}
	return model
}
