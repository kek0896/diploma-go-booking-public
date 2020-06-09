package model

import (
	"testing"
)

// TestUser to test user creation
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

// TestHotel to test hotel creation
func TestHotel(t *testing.T) *Hotel {
	return &Hotel{
		StructureType: "HOTEL",
		Provider:      "Dominick",
		// Nights:        8,
		Stars: 5,
		// Image:         "https://i.ibb.co/W5hkWSd/newyork3.jpg",
		// Datefrom:      "2020-09-01",
		// Dateto:        "2020-09-09",
		Latitude:    "40.725507",
		Longitude:   "-74.005576",
		Address:     &Address{Line1: "246 Spring Street, SoHo", Line2: "", City: "New York", Country: "USA", PostalCode: "10013"},
		Description: "The Dominick Hotel provides relaxing, contemporary rooms with custom-made furniture, a 42-inch flat-screen TV and iPod docking station. An Italian marble bath and a bathroom vanity with TV are included. A seasonal 7th floor outdoor pool overlooks the city and features a connected bar. A fitness center and business services are also available. Complimentary WiFi is offered to all guests. A premium or platinum WiFi with more bandwidth is also available at a surcharge. The Washington Square Park is less than 1 km from the The Dominick Hotel. The Film Forum is 500 m away.",
		Wifi:        true,
		Breakfast:   true,
		Parking:     true,
		Pool:        true,
		Playground:  true,
		Garden:      true,
		CheckIn:     "From 15:00",
		CheckOut:    "12:00 AM",
	}
}
