package models

import "testing"

func TestCountry_Validate_ShouldReturnError_WhenNameIsEmpty(t *testing.T) {
	country := Country{Name: "", Code: "US"}
	err := country.Validate()

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestCountry_Validate_ShouldReturnError_WhenCodeIsEmpty(t *testing.T) {
	country := Country{Name: "United States", Code: ""}
	err := country.Validate()

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestCountry_Validate_ShouldReturnError_WhenCodeIsNotSupported(t *testing.T) {
	country := Country{Name: "United States", Code: "USA"}
	err := country.Validate()

	if err[0].Error() != "invalid country code: USA" {
		t.Error("Expected error, got nil")
	}
}

func TestCountry_Validate_ShouldReturnError_WhenLanguageIsNotSupported(t *testing.T) {
	country := Country{Name: "United States", Code: "US", Language: "English"}
	err := country.Validate()

	if err[0].Error() != "invalid language code: English" {
		t.Errorf("Expected error, got \"%s\"", err[0].Error())
	}
}
