package sqlstore

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
)

// HistoryRepository ...
type HistoryRepository struct {
	store *Store
}

// StartBooking ...
func (hr *HistoryRepository) StartBooking(b *model.Booking) error {

	query := `SELECT hotel_internal_id, 
				property_id, 
				date_from, 
				date_to, 
				lock
			FROM properties 
			WHERE property_internal_id = ` + b.PropertyInternalID + " and active"

	rows, err := hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	p := &model.Property{}
	err = rows.Scan(p.HotelInternalID, p.PropertyID, p.DateFrom, p.DateTo, p.Lock)

	query = `UPDATE properties
			SET lock = ` + b.Sha1 + `
			WHERE property_id = ` + p.PropertyID + ` and active and hotel_internal_id = ` + p.HotelInternalID + `
			 and date_from = ` + p.DateFrom + ` and date_to = ` + p.DateTo

	rows, err = hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	return nil
}

// EndBooking ...
func (hr *HistoryRepository) EndBooking(b *model.Booking) error {
	return nil
}

// CancelBooking ...
func (hr *HistoryRepository) CancelBooking(b *model.Booking) error {
	return nil
}

// GetBookings ...
func (hr *HistoryRepository) GetBookings(sha1 string) ([]*model.Booking, error) {
	return []*model.Booking{}, nil
}

// Like ...
func (hr *HistoryRepository) Like(l *model.Like) error {
	return nil
}

// GetLikes ...
func (hr *HistoryRepository) GetLikes(sha1 string) ([]*model.Like, error) {
	return []*model.Like{}, nil
}
