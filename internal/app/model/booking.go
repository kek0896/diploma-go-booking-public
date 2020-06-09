package model

// Booking ...
type Booking struct {
	BookingKey         string      `json:"booking_key"`
	Sha1               string      `json:"sha1"`
	Email              string      `json:"email"`
	Phone              string      `json:"phone,omitempty"`
	Name               string      `json:"name,omitempty"`
	Surname            string      `json:"surname,omitempty"`
	CreditCard         *CreditCard `json:"credit_card,omitempty"`
	PaymentID          string      `json:"payment_id"`
	Status             string      `json:"status"`
	Timestamp          string      `json:"timestamp"`
	PropertyInternalID string      `json:"property_internal_id"`
	// PropertyID         string      `json:"property_id"`
	// HotelInternalID    string      `json:"hotel_internal_id"`
}
