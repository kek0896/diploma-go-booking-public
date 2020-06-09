package store

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	CreateUser(*model.User) error
	// WriteSha1(*model.User) error
	FindUserByEmail(string) (*model.User, error)
	// InsertGeoIP(geonameID string, countryISOCode string, countryName string, cityName string) error // fake helper method
}

// HotelRepository ...
type HotelRepository interface {
	CreateHotels(*model.Hotel) error
	CreateProperty(*model.Property) error
	Search(*model.Filter) ([]*model.Responce, error)
	// FindHotelByFilter(*model.Filter) ([]*model.Hotel, error)
	CleanHotels() error
	CleanProperties() error
}

// HistoryRepository ...
type HistoryRepository interface {
	StartBooking(b *model.Booking) error
	EndBooking(b *model.Booking) error
	CancelBooking(*model.Booking) error
	GetBookings(sha1 string) ([]*model.Booking, error)
	Like(*model.Like) error
	GetLikes(sha1 string) ([]*model.Like, error)
}
