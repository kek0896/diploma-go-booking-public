package sqlstore

import (
	"database/sql"

	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	db                *sql.DB
	userRepository    *UserRepository
	hotelRepository   *HotelRepository
	historyRepository *HistoryRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

// Hotel ...
func (s *Store) Hotel() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store: s,
	}
	return s.hotelRepository
}

// Property ...
func (s *Store) Property() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store: s,
	}
	return s.hotelRepository
}

// Responce ...
func (s *Store) Responce() store.HotelRepository {
	if s.hotelRepository != nil {
		return s.hotelRepository
	}

	s.hotelRepository = &HotelRepository{
		store: s,
	}
	return s.hotelRepository
}

// Booking ...
func (s *Store) Booking() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}

	s.historyRepository = &HistoryRepository{
		store: s,
	}
	return s.historyRepository
}

// Like ...
func (s *Store) Like() store.HistoryRepository {
	if s.historyRepository != nil {
		return s.historyRepository
	}

	s.historyRepository = &HistoryRepository{
		store: s,
	}
	return s.historyRepository
}
