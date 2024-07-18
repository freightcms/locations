package db

import "github.com/freightcms/locations/models"

// LocationDbContext exposes the method for manaing lookup locations which can be used for Master Data Management.
// This database context can be used in other ways but it has no refernce to other entities.
type LocationDbContext interface {
	// CreateLocation creates a new location in the database as a lookup reference for other data.
	// The location model is expected to have all fields populated except for the identifier property.
	// The function returns the tuple for the ID of the location for lookup afterwards, or
	// an error if one occurs.
	CreateLocation(location *models.AddressModel) (interface{}, error)
	// Revmoes a Location entity from database. Argument type is the identifier of the
	// location.
	DeleteLocation(id interface{}) error
}
