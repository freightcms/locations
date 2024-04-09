package models

import (
	"testing"
)

func TestCoordinates_Validate(t *testing.T) {
	tests := []struct {
		name    string
		coords  Coordinates
		wantErr bool
	}{
		{
			name:    "Valid coordinates",
			coords:  Coordinates{Latitude: 45.0, Longitude: 90.0},
			wantErr: false,
		},
		{
			name:    "Invalid latitude",
			coords:  Coordinates{Latitude: 100.0, Longitude: 90.0},
			wantErr: true,
		},
		{
			name:    "Invalid longitude",
			coords:  Coordinates{Latitude: 45.0, Longitude: 200.0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.coords.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Coordinates.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
