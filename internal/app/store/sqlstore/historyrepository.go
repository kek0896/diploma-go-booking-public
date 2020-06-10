package sqlstore

import (
	"database/sql"
	"time"

	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
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
			WHERE property_internal_id = "` + b.PropertyID + `" and active`

	rows, err := hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	p := &model.Property{}
	err = rows.Scan(p.HotelInternalID, p.PropertyID, p.DateFrom, p.DateTo, p.Lock)

	query = `UPDATE properties
			SET lock = "` + b.Sha1 + `", timestamp = "` + string(time.Now().UnixNano()) + `"
			WHERE property_id = "` + p.PropertyID + `" and active and hotel_internal_id = "` + p.HotelInternalID + `"
			 and date_from = "` + p.DateFrom + `" and date_to = "` + p.DateTo + `"`

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

	err := hr.store.db.QueryRow(
		`INSERT INTO bookings(
			booking_key, sha1, email, phone, 
			name, surname, credit_card, payment_id, 
			status, timestamp, property_internal_id, property_id, 
			hotel_internal_id, date_from, date_to)
		VALUES ($1, $2, $3, $4, 
			$5, $6, $7, $8, 
			$9, $10, $11, $12, 
			$13, $14, $15)`,
		b.BookingKey, b.Sha1, b.Email, b.Phone,
		b.Name, b.Surname, b.CreditCard, b.PaymentID,
		b.Status, b.Timestamp, b.PropertyInternalID, b.PropertyID,
		b.HotelInternalID, b.DateFrom, b.DateTo).Scan()

	if err != nil {
		return err
	}

	query := `UPDATE properties
			SET timestamp = ''
			WHERE timestamp != '' and property_id = '` + b.PropertyID + `' and lock = '` + b.Sha1 + `'`

	_, err = hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	return nil
}

// CancelBooking ...
func (hr *HistoryRepository) CancelBooking(b *model.Booking) error {

	query := `UPDATE properties
			SET lock = '', timestamp = ''
			WHERE property_id = '` + b.PropertyID + `' and active and lock = '` + b.Sha1 + `' and hotel_internal_id = '` + b.HotelInternalID + `'
			and (date_from BETWEEN '` + b.DateFrom + `' and '` + b.DateTo + `' or date_to BETWEEN '` + b.DateFrom + `' and '` + b.DateTo + `'` // TODO add dates conditions
	_, err := hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	query = `DELETE FROM bookings
	WHERE booking_key = '` + b.BookingKey + `'`

	_, err = hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetBookings ...
func (hr *HistoryRepository) GetBookings(sha1 string) ([]*model.Booking, error) {

	query := `SELECT 
				booking_key, sha1, email, phone,
				name, surname, payment_id, status,
				timestamp, property_internal_id, date_from, date_to,
				property_id, hotel_internal_id
			FROM bookings 
			WHERE sha1 = '` + sha1 + `'`

	rows, err := hr.store.db.Query(
		query,
	)

	if err != nil {
		return []*model.Booking{}, err
	}

	var bookings []*model.Booking
	for rows.Next() {
		b := &model.Booking{}
		err := rows.Scan(
			b.BookingKey, b.Sha1, b.Email, b.Phone,
			b.Name, b.Surname, b.PaymentID, b.Status,
			b.Timestamp, b.PropertyInternalID, b.DateFrom, b.DateTo,
			b.PropertyID, b.HotelInternalID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

// Like ...
func (hr *HistoryRepository) Like(l *model.Like) error {

	err := hr.store.db.QueryRow(
		`SELECT 
			hotel_internal_id, sha1
		FROM likes 
		WHERE hotel_internal_id = '` + l.HotelInternalID + `' and sha1 = '` + l.Sha1 + `'`).Scan()

	if err != nil {
		if err == sql.ErrNoRows {
			err := hr.store.db.QueryRow(
				`INSERT INTO likes(
					hotel_internal_id, sha1)
				VALUES ($1, $2)`,
				l.HotelInternalID, l.Sha1).Scan()
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	query := `DELETE FROM likes
	WHERE hotel_internal_id = '` + l.HotelInternalID + `' and sha1 = ` + l.Sha1 + `'`

	_, err = hr.store.db.Query(
		query,
	)

	if err != nil {
		return err
	}

	return nil

}

// GetLikes ...
func (hr *HistoryRepository) GetLikes(sha1 string) ([]*model.Like, error) {

	query := `SELECT 
				hotel_internal_id
			FROM likes 
			WHERE sha1 = '` + sha1 + `'`

	rows, err := hr.store.db.Query(
		query,
	)

	if err != nil {
		return []*model.Like{}, err
	}

	var likes []*model.Like
	for rows.Next() {
		l := &model.Like{}
		err := rows.Scan(
			l.HotelInternalID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		likes = append(likes, l)
	}
	return likes, nil
}
