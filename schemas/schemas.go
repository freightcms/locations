package schemas

import "github.com/freightcms/locations/models"

// Schema for binding against an HTTP POST or HTTP PUT method.
// All properties are in pascal casing. If you are using gin
// you can bind this also to a form as long-as the form adheres to the
// pascal casing naming convention.
type CreateLocationSchema struct {
	Line1       string             `json:"line1" xml:"line1" binding:"required" form:"line1"`          // Street address, P.O. box, company name, c/o
	Line2       *string            `json:"line2,omitempty" xml:"line2" form:"line2"`                   // Apartment, suite, unit, building, floor, etc.
	Line3       *string            `json:"line3,omitempty" xml:"line3" form:"line3"`                   // Floor, Bin, Section of Warehouse, Port # etc.
	Locale      string             `json:"locale" xml:"locale" form:"locale"`                          // City, town, village, etc.
	Region      string             `json:"region" xml:"region" form:"region"`                          // State, province within Country
	PostalCode  string             `json:"postalCode" xml:"postalCode" form:"postalCode"`              // Postal code
	Country     models.CountryCode `json:"country" xml:"country" form:"country" binding:"required"`    // Country
	Description *string            `json:"description,omitempty" xml:"description" form:"description"` // Description of the address
	Attention   *string            `json:"attention,omitempty" xml:"attention" form:"attention"`       // Attention of the address
	Type        models.AddressType `json:"type" xml:"type" form:"type"`                                // Type of address, e.g. "home", "work", "billing", "shipping", "other"
	Notes       *string            `json:"notes" xml:"notes" form:"notes"`                             // Notes are details about a delivery such as leaving by a back door, leave at a paticular location
}

type LocationSchema struct {
	Id string `json:"id" xml:"id" form:"id" binding:"required"`
	CreateLocationSchema
}
