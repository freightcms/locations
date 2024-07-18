package services

import (
	"github.com/freightcms/locations/db"
	"github.com/freightcms/locations/schemas"
)

// LocationService provides an interface for basic operations to perform master data management, e.g. lookup
// information
type LocationService interface {
	// CreateLocation creates a new lookup address. Function returns the identifier or an error
	// that occured during the operation of creating the resource.
	CreateLocation(location *schemas.CreateLocationSchema) (interface{}, error)

	// DeleteLocation removes a location lookup resource
	DeleteLocation(id interface{}) error
}

type locationService struct {
	dbContext *db.LocationDbContext
}

func CreateService(db *db.LocationDbContext) LocationService {
	return &locationService{dbContext: db}
}
