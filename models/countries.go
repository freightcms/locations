package models

// Country Model meant for providing lookup data for countries. Optionally you extend the address model
// and add the country Id and override the Address.Country property with a reference to this model.
type Country struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Code             string  `json:"code"`
	Description      *string `json:"description,omitempty"`
	Currency         *string `json:"currency,omitempty"`
	Language         *string `json:"language,omitempty"`
	Region           *string `json:"region,omitempty"`
	SubRegion        *string `json:"subRegion,omitempty"`
	Capital          *string `json:"capital,omitempty"`
	CallingCode      *string `json:"callingCode,omitempty"`
	Flag             *string `json:"flag,omitempty"`
	FlagEmoji        *string `json:"flagEmoji,omitempty"`
	FlagEmojiUnicode *string `json:"flagEmojiUnicode,omitempty"`
}
