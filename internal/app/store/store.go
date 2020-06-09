package store

// Store ...
type Store interface {
	User() UserRepository
	Hotel() HotelRepository
	Property() HotelRepository
	Responce() HotelRepository
	Booking() HistoryRepository
	Like() HistoryRepository
}
