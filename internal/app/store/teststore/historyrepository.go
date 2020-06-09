package teststore

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
)

// HistoryRepository ...
type HistoryRepository struct {
	store    *Store
	bookings map[string]*model.Booking
	likes    map[string]*model.Like
}

// StartBooking ...
func (r *HistoryRepository) StartBooking(b *model.Booking) error {
	return nil
}

// EndBooking ...
func (r *HistoryRepository) EndBooking(b *model.Booking) error {
	return nil
}

// CancelBooking ...
func (r *HistoryRepository) CancelBooking(b *model.Booking) error {
	return nil
}

// GetBookings ...
func (r *HistoryRepository) GetBookings(sha1 string) ([]*model.Booking, error) {
	return []*model.Booking{}, nil
}

// Like ...
func (r *HistoryRepository) Like(l *model.Like) error {
	return nil
}

// GetLikes ...
func (r *HistoryRepository) GetLikes(sha1 string) ([]*model.Like, error) {
	return []*model.Like{}, nil
}
