package models

import "errors"

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Validate checks if the coordinates are within the valid range for latitude and longitude
func (c *Coordinates) Validate() error {
	if c.Latitude < -90 || c.Latitude > 90 {
		return errors.New("Latitude must be between -90 and 90")
	}
	if c.Longitude < -180 || c.Longitude > 180 {
		return errors.New("Longitude must be between -180 and 180")
	}
	return nil
}

// AddressModel is a model meant for providing lookup data for addresses. Optionally you extend the address model
type AddressModel struct {
	Id          string  `json:"id"`
	Line1       string  `json:"line1"`
	Line2       *string `json:"line2,omitempty"`
	Line3       *string `json:"line3,omitempty"`
	Local       string  `json:"local"`
	Region      string  `json:"region"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Description *string `json:"description,omitempty"`
	Attention   *string `json:"attention,omitempty"`
	Type        string  `json:"type"`
}
