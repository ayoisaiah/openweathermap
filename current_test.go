package openweathermap

import (
	"os"
	"testing"
)

type query struct {
	Name   string
	ID     int
	Coords *Coord
}

var citiesTable = []query{
	{
		Name: "Ilorin",
		ID:   2337639,
		Coords: &Coord{
			Lat: 8.5,
			Lon: 4.5,
		},
	},
	{
		Name: "Hamilton County",
		ID:   4513583,
		Coords: &Coord{
			Lat: 39.183392,
			Lon: -84.533279,
		},
	},
	{
		Name: "Lagos",
		ID:   2332459,
		Coords: &Coord{
			Lat: 6.45306,
			Lon: 3.39583,
		},
	},
}

func TestGetCurrentByCoords(t *testing.T) {
	owm, err := New(os.Getenv("OPENWEATHERMAP_KEY"))
	if err != nil {
		t.Error(err)
	}

	for _, city := range citiesTable {
		data, err := owm.GetCurrentByCoords(*city.Coords)
		if err != nil {
			t.Error(err)
		}

		if data.ID != city.ID {
			t.Errorf("Expected city ID to match %d. Got %d", city.ID, data.ID)
		}
	}
}

func TestGetCurrentByCityName(t *testing.T) {
	owm, err := New(os.Getenv("OPENWEATHERMAP_KEY"))
	if err != nil {
		t.Error(err)
	}

	for _, city := range citiesTable {
		data, err := owm.GetCurrentByCityName(city.Name)
		if err != nil {
			t.Error(err)
		}

		if data.ID != city.ID {
			t.Errorf("Expected city ID to match %d. Got %d", city.ID, data.ID)
		}
	}
}

func TestGetCurrentByID(t *testing.T) {
	owm, err := New(os.Getenv("OPENWEATHERMAP_KEY"))
	if err != nil {
		t.Error(err)
	}

	for _, city := range citiesTable {
		data, err := owm.GetCurrentByID(city.ID)
		if err != nil {
			t.Error(err)
		}

		if data.ID != city.ID {
			t.Errorf("Expected city ID to match %d. Got %d", city.ID, data.ID)
		}
	}
}

func TestGetCurrentByZipCode(t *testing.T) {
	owm, err := New(os.Getenv("OPENWEATHERMAP_KEY"))
	if err != nil {
		t.Error(err)
	}

	if _, err := owm.GetCurrentByZipCode(94040, "us"); err != nil {
		t.Error(err)
	}
}
