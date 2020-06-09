package teststore

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
)

// HotelRepository ...
type HotelRepository struct {
	store      *Store
	hotels     map[string]*model.Hotel
	properties map[string]*model.Property
	responces  map[string]*model.Responce
}

// CreateHotels ...
func (hr *HotelRepository) CreateHotels(h *model.Hotel) error {

	if err := h.Validate(); err != nil {
		return err
	}

	// hr.hotels[h.PropertyID] = h
	return nil

}

// CreateProperty ...
func (hr *HotelRepository) CreateProperty(h *model.Property) error {

	if err := h.Validate(); err != nil {
		return err
	}

	// hr.hotels[h.PropertyID] = h
	return nil

}

// Search ...
func (hr *HotelRepository) Search(filter *model.Filter) ([]*model.Responce, error) {
	h := []*model.Responce{}
	return h, nil
}

// FindHotelByFilter ...
func (hr *HotelRepository) FindHotelByFilter(filter *model.Filter) ([]*model.Hotel, error) {
	h := []*model.Hotel{}
	return h, nil
}

// CleanHotels ...
func (hr *HotelRepository) CleanHotels() error {

	return nil

}

// CleanProperties ...
func (hr *HotelRepository) CleanProperties() error {

	return nil

}
