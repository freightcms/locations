package mongodb

import (
	"context"

	"github.com/freightcms/locations/db"
	"github.com/freightcms/locations/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type locationDb struct {
	session mongo.Session
}

// LocationDbcontext creates a new database context which exposes basic APIs
// for modifying, creating, or deleting entities.
func LocationDbContext(session mongo.Session) db.LocationDbContext {
	return &locationDb{
		session: session,
	}
}

// CreateLocation creates a new location (address) entity. The first argumet returned
// is the ID of the newly created element to be fetched later. The second argument
// is any error bubled up from the native tools used to communicate with the database.
// e.g. drivers, HTTP Web Clinets, WCF Services, SOAP services, etc.
func (l *locationDb) CreateLocation(model *models.AddressModel) (interface{}, error) {
	if err := l.session.StartTransaction(); err != nil {
		return nil, err
	}
	result, err := l.session.Client().Database("freightcms").Collection("locations").InsertOne(context.TODO(), &model)
	if err != nil {
		return nil, err
	}

	if err := l.session.CommitTransaction(context.TODO()); err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// DeleteLocation
func (l *locationDb) DeleteLocation(id interface{}) error {
	return nil
}
