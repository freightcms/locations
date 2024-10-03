package models

import "errors"

// Coordinates is the location geographically based on latitude and longitude of
// where a location is. All properties are serialized to pascal casing for `json` and `bson`
type Coordinates struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

// Validate checks if the coordinates are within the valid range for latitude and longitude
func (c *Coordinates) Validate() error {
	if c.Latitude < -90 || c.Latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}
	if c.Longitude < -180 || c.Longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}
	return nil
}

// AddressType tells how to handle the address related to billing, mailing, rollup, etc.
type AddressType string

const (
	Physical AddressType = "Physical" // Physical location for a resource/entity
	Mailing  AddressType = "Mailing"  // Mailing address
	Billing  AddressType = "Billing"  // Billing address
	Shipping AddressType = "Shipping" // Shipping address
	Work     AddressType = "Work"     // Work
	Home     AddressType = "Home"     // Home
	Virtual  AddressType = "Virtual"  // Virtual address
	Other    AddressType = "Other"    // Other
)

// AddressModel is a model meant for providing lookup data for addresses. Optionally you extend the address model
type AddressModel struct {
	Id          interface{} `json:"id" bson:"id"`                             // unique identifier of the address model.
	Line1       string      `json:"line1" bson:"line1"`                       // Street address, P.O. box, company name, c/o
	Line2       *string     `json:"line2,omitempty" bson:"line2"`             // Apartment, suite, unit, building, floor, etc.
	Line3       *string     `json:"line3,omitempty" bson:"line3"`             // Floor, Bin, Section of Warehouse, Port # etc.
	Locale      string      `json:"locale" bson:"locale"`                     // City, town, village, etc.
	Region      string      `json:"region" bson:"region"`                     // State, province within Country
	PostalCode  string      `json:"postalCode" bson:"postalCode"`             // Postal code
	Country     CountryCode `json:"country" bson:"country"`                   // Country
	Description *string     `json:"description,omitempty" bson:"description"` // Description of the address
	Attention   *string     `json:"attention,omitempty" bson:"attention"`     // Attention of the address
	Type        AddressType `json:"type" bson:"type"`                         // Type of address, e.g. "home", "work", "billing", "shipping", "other"
	Notes       *string     `json:"notes" bson:"notes"`                       // Any additional notes or special instructions about the delivery
}
