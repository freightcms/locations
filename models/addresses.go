package models

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

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
