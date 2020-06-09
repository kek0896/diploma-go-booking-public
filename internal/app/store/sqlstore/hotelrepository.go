package sqlstore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
	"github.com/lib/pq"
)

// HotelRepository ...
type HotelRepository struct {
	store *Store
}

// CreateHotels ...
func (hr *HotelRepository) CreateHotels(h *model.Hotel) error {

	if err := h.Validate(); err != nil {
		return err
	}

	jAddress, err := json.Marshal(h.Address)
	if err != nil {
		return err
	}

	// jImages, err := json.Marshal(h.Images)
	// if err != nil {
	// 	return err
	// }

	hotelInternalID := h.Provider + "_" + h.HotelID

	err = hr.store.db.QueryRow(
		`INSERT INTO hotels_v2 (
			hotel_id, hotel_name, hotel_internal_id, provider,
			structure_type, min_nights, max_nights, images,
			start_date, active_day_period, stars, latitude, 
			longitude, address, description, active, 
			wifi, breakfast, parking, pool, 
			playground, garden, check_in, check_out) 
		VALUES ($1, $2, $3, $4, 
			$5, $6, $7, $8, 
			$9, $10, $11, $12, 
			$13, $14, $15, $16, 
			$17, $18, $19, $20, 
			$21, $22, $23, $24)`,
		h.HotelID, h.HotelName, hotelInternalID, h.Provider,
		h.StructureType, h.MinNights, h.MaxNights, pq.Array(h.Images),
		h.StartDate, h.ActiveDayPeriod, h.Stars, h.Latitude,
		h.Longitude, jAddress, h.Description, h.Active,
		h.Wifi, h.Breakfast, h.Parking, h.Pool,
		h.Playground, h.Garden, h.CheckIn, h.CheckOut).Scan()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}

	return nil

}

// CreateProperty ...
func (hr *HotelRepository) CreateProperty(p *model.Property) error {

	if err := p.Validate(); err != nil {
		return err
	}

	hotelInternalID := p.Provider + "_" + p.HotelID
	p.PropertyInternalID = hotelInternalID + "_" + p.PropertyID

	startDate := ""
	activeDayPeriod := ""
	min := ""
	max := ""

	// get hotel data for generation
	err := hr.store.db.QueryRow(
		`SELECT start_date, active_day_period, min_nights, max_nights FROM hotels_v2 WHERE hotel_internal_id = $1`,
		hotelInternalID).Scan(&startDate, &activeDayPeriod, &min, &max)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}

	layout := "2006-01-02"
	dateFrom, err := time.Parse(layout, startDate)

	if err != nil {
		return err
	}

	adp, err := strconv.ParseInt(activeDayPeriod, 10, 32)
	minn, err := strconv.ParseInt(min, 10, 32)
	maxn, err := strconv.ParseInt(max, 10, 32)

	if err != nil {
		return err
	}

	// start dates generation
	endDate := dateFrom.AddDate(0, 0, int(adp)) // save the end date
	for dateFrom.Before(endDate) {              // while the start date is less than saved end date
		for x := minn; x <= maxn; x++ { // days amount to add
			dateTo := dateFrom.AddDate(0, 0, int(x))
			if dateTo.After(dateFrom.AddDate(0, 0, int(adp))) {
				return nil
			}
			err = hr.store.db.QueryRow(
				`INSERT INTO properties (
					hotel_internal_id, property_id, property_internal_id, property_name, 
					price, currency, nights, image, 
					date_from, date_to, description, active, 
					rooms_number, beds_number, size_m, capacity, 
					lock) 
				VALUES ($1, $2, $3, $4, 
					$5, $6, $7, $8, 
					$9, $10, $11, $12, 
					$13, $14, $15, $16, $17)`,
				hotelInternalID, p.PropertyID, p.PropertyInternalID+"_"+dateFrom.Format(layout)+"_"+dateTo.Format(layout), p.PropertyName,
				p.Price*float32(x), p.Currency, x, p.Image,
				dateFrom.Format(layout), dateTo.Format(layout), p.Description, p.Active,
				p.RoomsNumber, p.BedsNumber, p.SizeM, p.Capacity, "").Scan()
		}
		dateFrom = dateFrom.AddDate(0, 0, 1)
	}

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}

	return nil

}

// Search ...
func (hr *HotelRepository) Search(f *model.Filter) ([]*model.Responce, error) {

	hotels, err := hr.FindHotelByFilter(f)
	if err != nil {
		return nil, err
	}

	var responce []*model.Responce
	for _, h := range hotels {

		properties, err := hr.FindPropertyByFilter(f, h)
		if err != nil {
			return nil, err
		}

		if len(properties) == 0 {
			continue
		}

		minPrice := float32(1000000000)
		curr := ""
		nights := 1
		for _, p := range properties {
			if p.Price < minPrice {
				minPrice = p.Price
				curr = p.Currency
				nights = p.Nights
			}
		}

		responce = append(responce, &model.Responce{
			HotelID:         h.HotelID,
			HotelInternalID: h.HotelInternalID,
			HotelName:       h.HotelName,
			Provider:        h.Provider,
			StructureType:   h.StructureType,
			Stars:           h.Stars,
			DateFrom:        properties[0].DateFrom,
			DateTo:          properties[0].DateTo,
			Images:          h.Images,
			Latitude:        h.Latitude,
			Longitude:       h.Longitude,
			Address:         h.Address,
			Description:     h.Description,
			Active:          h.Active,
			Wifi:            h.Wifi,
			Breakfast:       h.Breakfast,
			Parking:         h.Parking,
			Pool:            h.Pool,
			Playground:      h.Playground,
			Garden:          h.Garden,
			CheckIn:         h.CheckIn,
			CheckOut:        h.CheckOut,
			MinPrice:        fmt.Sprintf("%f", minPrice),
			Nights:          nights,
			Currency:        curr,
			Properties:      properties,
		})
	}

	return responce, nil
}

// FindHotelByFilter ...
func (hr *HotelRepository) FindHotelByFilter(f *model.Filter) ([]*model.Hotel, error) {

	query := `SELECT  
	hotel_id, hotel_name, hotel_internal_id, provider,
    structure_type, min_nights, max_nights, stars,
    images, start_date, active_day_period, latitude,
    longitude, address, description, active,
    wifi, breakfast, parking, pool,
	playground, garden, check_in, check_out 
	FROM hotels_v2`

	// "address"->>'country' = 'Scotland';

	var conditions []string

	if f.Country != "" {
		conditions = append(conditions, "\"address\"->>'country' ilike '"+f.Country+"'")
	}

	if f.City != "" {
		conditions = append(conditions, "\"address\"->>'city' ilike '"+f.City+"'")
	}

	if f.Stars != "" {
		conditions = append(conditions, "stars = '"+f.Stars+"'")
	}

	if f.StructureType != "" {
		conditions = append(conditions, "structure_type ilike '"+f.StructureType+"'")
	}

	if f.Wifi != "" {
		conditions = append(conditions, "wifi")
	}

	if f.Parking != "" {
		conditions = append(conditions, "parking")
	}

	if f.Pool != "" {
		conditions = append(conditions, "pool")
	}

	if f.Playground != "" {
		conditions = append(conditions, "playground")
	}

	if f.Breakfast != "" {
		conditions = append(conditions, "breakfast")
	}

	if f.Garden != "" {
		conditions = append(conditions, "garden")
	}

	conditions = append(conditions, "active")

	if len(conditions) != 0 {
		query += " WHERE "
		query += conditions[0]
		for _, i := range conditions[1:] {
			query += " and " + i
		}
	}

	rows, err := hr.store.db.Query(
		query,
	)
	if err != nil {
		return []*model.Hotel{}, err
	}

	var address string
	var hotels []*model.Hotel
	for rows.Next() {
		h := &model.Hotel{}
		err := rows.Scan(
			&h.HotelID, &h.HotelName, &h.HotelInternalID, &h.Provider,
			&h.StructureType, &h.MinNights, &h.MaxNights, &h.Stars,
			pq.Array(&h.Images), &h.StartDate, &h.ActiveDayPeriod, &h.Latitude,
			&h.Longitude, &address, &h.Description, &h.Active,
			&h.Wifi, &h.Breakfast, &h.Parking, &h.Pool,
			&h.Playground, &h.Garden, &h.CheckIn, &h.CheckOut,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		json.Unmarshal([]byte(address), &h.Address)
		hotels = append(hotels, h)
	}
	return hotels, nil
}

// FindPropertyByFilter ...
func (hr *HotelRepository) FindPropertyByFilter(f *model.Filter, hotel *model.Hotel) ([]*model.Property, error) {

	query := `SELECT 
	hotel_internal_id, property_id, property_internal_id, property_name,
	price, currency, nights, date_from,
	date_to, description, active, beds_number,
	rooms_number, size_m, lock, 
	capacity FROM properties`

	query += " WHERE hotel_internal_id = '" + hotel.HotelInternalID + "'"

	var conditions []string

	if f.Datefrom != "" && len(f.Datefrom) == 10 {
		conditions = append(conditions, "date_from = '"+f.Datefrom+"'")
	}

	if f.Dateto != "" && len(f.Dateto) == 10 {
		conditions = append(conditions, "date_to = '"+f.Dateto+"'")
	}

	if f.PriceFrom != "" {
		conditions = append(conditions, "price >= "+f.PriceFrom)
	}

	if f.PriceTo != "" {
		conditions = append(conditions, "price <= "+f.PriceTo)
	}

	if f.Currency != "" {
		conditions = append(conditions, "currency ilike '"+f.Currency+"'")
	}

	if f.Capacity != "" {
		conditions = append(conditions, "capacity = '"+f.Capacity+"'")
	}

	if f.BedsNumber != "" {
		conditions = append(conditions, "beds_number = '"+f.BedsNumber+"'")
	}

	if f.RoomsNumber != "" {
		conditions = append(conditions, "rooms_number = '"+f.RoomsNumber+"'")
	}

	conditions = append(conditions, "active")

	conditions = append(conditions, "lock = ''")

	if len(conditions) != 0 {
		for _, i := range conditions {
			query += " and " + i
		}
	}

	// var canOrderBy map[string]bool

	canOrderBy := map[string]bool{
		"price":        true,
		"nights":       true,
		"date_from":    true,
		"date_to":      true,
		"beds_number":  true,
		"rooms_number": true,
		"size_m":       true,
		"capacity":     true,
	}

	if f.OrderBy != "" && canOrderBy[f.OrderBy] {
		query += " ORDER BY " + f.OrderBy
	}

	rows, err := hr.store.db.Query(
		query,
	)
	if err != nil {
		return []*model.Property{}, err // error is here
	}

	var properties []*model.Property
	for rows.Next() {
		p := &model.Property{}
		err := rows.Scan(
			&p.HotelInternalID,
			&p.PropertyID,
			&p.PropertyInternalID,
			&p.PropertyName,
			&p.Price,
			&p.Currency,
			&p.Nights,
			&p.DateFrom,
			&p.DateTo,
			&p.Description,
			&p.Active,
			&p.BedsNumber,
			&p.RoomsNumber,
			&p.SizeM,
			&p.Lock,
			&p.Capacity,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		properties = append(properties, p)
	}
	return properties, nil
}

// CleanHotels ...
func (hr *HotelRepository) CleanHotels() error {

	err := hr.store.db.QueryRow(
		`DROP TABLE hotels_v2; COMMIT; ` + `CREATE TABLE hotels_v2 (
			hotel_id VARCHAR NOT NULL,
			hotel_name VARCHAR NOT NULL,
			hotel_internal_id VARCHAR NOT NULL unique,
			provider VARCHAR NOT NULL,
			structure_type VARCHAR NOT NULL,
			min_nights INT NOT NULL,
			max_nights INT NOT NULL,
			stars INT NOT NULL,
			images text ARRAY NOT NULL,
			start_date VARCHAR NOT NULL,
			active_day_period VARCHAR NOT NULL,
			latitude VARCHAR NOT NULL,
			longitude VARCHAR NOT NULL,
			address JSON NOT NULL,
			description VARCHAR NOT NULL,
			active BOOLEAN NOT NULL,
			wifi BOOLEAN NOT NULL,
			breakfast BOOLEAN NOT NULL,
			parking BOOLEAN NOT NULL,
			pool BOOLEAN NOT NULL,
			playground BOOLEAN NOT NULL,
			garden BOOLEAN NOT NULL,
			check_in VARCHAR NOT NULL,
			check_out VARCHAR NOT NULL
		); COMMIT;`).Scan()

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}

	return nil

}

// CleanProperties ...
func (hr *HotelRepository) CleanProperties() error {

	err := hr.store.db.QueryRow(
		`DROP TABLE properties; COMMIT; ` + `CREATE TABLE properties (
			hotel_internal_id VARCHAR NOT NULL,
			property_id VARCHAR NOT NULL,
			property_internal_id VARCHAR NOT NULL unique,
			property_name VARCHAR NOT NULL, 
			price FLOAT NOT NULL,
			currency VARCHAR NOT NULL,
			nights INT NOT NULL,
			image VARCHAR NOT NULL,
			date_from VARCHAR NOT NULL,
			date_to  VARCHAR NOT NULL,
			description VARCHAR NOT NULL,
			active boolean NOT NULL,
			rooms_number int NOT NULL,
			beds_number VARCHAR NOT NULL,
			size_m FLOAT NOT NULL, 
			lock VARCHAR NOT NULL, 
			capacity int NOT NULL
		); COMMIT;`).Scan()

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}

	return nil

}
