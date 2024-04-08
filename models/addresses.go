package models

import "errors"

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

type AddressType string

const (
	Physical AddressType = "Physical"
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
	Id          string      `json:"id"`
	Line1       string      `json:"line1"`                 // Street address, P.O. box, company name, c/o
	Line2       *string     `json:"line2,omitempty"`       // Apartment, suite, unit, building, floor, etc.
	Line3       *string     `json:"line3,omitempty"`       // Floor, Bin, Section of Warehouse, Port # etc.
	Local       string      `json:"local"`                 // City, town, village, etc.
	Region      string      `json:"region"`                // State, province within Country
	PostalCode  string      `json:"postalCode"`            // Postal code
	Country     CountryCode `json:"country"`               // Country
	Description *string     `json:"description,omitempty"` // Description of the address
	Attention   *string     `json:"attention,omitempty"`   // Attention of the address
	Type        AddressType `json:"type"`                  // Type of address, e.g. "home", "work", "billing", "shipping", "other"
}
