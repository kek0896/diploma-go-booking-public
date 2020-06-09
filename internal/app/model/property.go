package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Property struct for MVP
type Property struct {
	HotelID            string  `json:"hotel_id,omitempty"`
	HotelInternalID    string  `json:"hotel_internal_id,omitempty"`
	PropertyInternalID string  `json:"property_internal_id,omitempty"`
	Provider           string  `json:"provider"`
	PropertyID         string  `json:"property_id"`
	PropertyName       string  `json:"property_name"`
	Price              float32 `json:"price"`
	Currency           string  `json:"currency"`
	Image              string  `json:"image"`
	Description        string  `json:"description"`
	Active             bool    `json:"active"`
	RoomsNumber        int     `json:"rooms_number"`
	BedsNumber         int     `json:"beds_number"`
	Capacity           int     `json:"capacity"`
	SizeM              float32 `json:"size_m"`
	DateTo             string  `json:"date_to,omitempty"`
	DateFrom           string  `json:"date_from,omitempty"`
	Nights             int     `json:"nights,omitempty"`
	Lock               string  `json:"lock,omitempty"`
}

// Validate validate struct
func (h *Property) Validate() error {

	// t := time.Now()

	err := validation.ValidateStruct(
		h,
		validation.Field(&h.Image, validation.Required, is.URL),
		// validation.Field(&h.CheckInYear, validation.Required, &validation.Min(t.Year),
		validation.Field(&h.Currency, validation.By(requiredIf(h.Price > 0)), is.CurrencyCode),
	)
	if err != nil {
		return err
	}

	return nil

}
