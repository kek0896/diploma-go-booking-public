package model

// Responce struct for MVP
type Responce struct {
	HotelID         string      `json:"hotel_id"`
	HotelInternalID string      `json:"hotel_internal_id,omitempty"`
	HotelName       string      `json:"hotel_name"`
	Provider        string      `json:"provider"`
	StructureType   string      `json:"structure_type"`
	Stars           int         `json:"stars"`
	DateTo          string      `json:"date_to,omitempty"`
	DateFrom        string      `json:"date_from,omitempty"`
	Images          []string    `json:"images"`
	Latitude        string      `json:"latitude"`
	Longitude       string      `json:"longitude"`
	Address         *Address    `json:"address"`
	Description     string      `json:"description"`
	Active          bool        `json:"active"`
	Wifi            bool        `json:"wifi"`
	Breakfast       bool        `json:"breakfast"`
	Parking         bool        `json:"parking"`
	Pool            bool        `json:"pool"`
	Playground      bool        `json:"playground"`
	Garden          bool        `json:"garden"`
	CheckIn         string      `json:"check_in"`
	CheckOut        string      `json:"check_out"`
	MinPrice        string      `json:"min_price"`
	Currency        string      `json:"currency"`
	Nights          int         `json:"nights"`
	Properties      []*Property `json:"properties"`
}
