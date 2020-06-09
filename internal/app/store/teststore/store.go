package teststore

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	userRepository    *UserRepository
	hotelRepository   *HotelRepository
	historyRepository *HistoryRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}

// Hotel ...
func (s *Store) Hotel() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store:  s,
		hotels: make(map[string]*model.Hotel),
	}
	return s.hotelRepository
}

// Property ...
func (s *Store) Property() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store:      s,
		properties: make(map[string]*model.Property),
	}
	return s.hotelRepository
}

// Responce ...
func (s *Store) Responce() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store:     s,
		responces: make(map[string]*model.Responce),
	}
	return s.hotelRepository
}

// Booking ...
func (s *Store) Booking() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}

	s.historyRepository = &HistoryRepository{
		store:    s,
		bookings: make(map[string]*model.Booking),
	}
	return s.historyRepository
}

// Booking ...
func (s *Store) Like() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}

	s.historyRepository = &HistoryRepository{
		store: s,
		likes: make(map[string]*model.Like),
	}
	return s.historyRepository
}
