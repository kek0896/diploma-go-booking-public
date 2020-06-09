package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// TODO add extra fields

// Hotel struct for MVP
type Hotel struct {
	HotelID         string   `json:"hotel_id"`
	HotelInternalID string   `json:"hotel_internal_id,omitempty"`
	HotelName       string   `json:"hotel_name"`
	Provider        string   `json:"provider"`
	StructureType   string   `json:"structure_type"`
	Stars           int      `json:"stars"`
	MinNights       int      `json:"min_nights"`
	MaxNights       int      `json:"max_nights"`
	StartDate       string   `json:"start_date"`
	ActiveDayPeriod string   `json:"active_day_period"`
	Images          []string `json:"images"`
	Latitude        string   `json:"latitude"`
	Longitude       string   `json:"longitude"`
	Address         *Address `json:"address"`
	Description     string   `json:"description"`
	Active          bool     `json:"active"`
	Wifi            bool     `json:"wifi"`
	Breakfast       bool     `json:"breakfast"`
	Parking         bool     `json:"parking"`
	Pool            bool     `json:"pool"`
	Playground      bool     `json:"playground"`
	Garden          bool     `json:"garden"`
	CheckIn         string   `json:"check_in"`
	CheckOut        string   `json:"check_out"`
}

// PropertyID    string  `json:"property_id"`
// PropertyName  string  `json:"property_name"`

// Nights        int     `json:"nights"`
// Datefrom      string  `json:"datefrom"`
// Dateto        string  `json:"dateto"`

// Validate hotel struct
func (h *Hotel) Validate() error {
	err := validation.ValidateStruct(
		h,
		//validation.Field(&h.Images, validation.Required, is.URL),
		validation.Field(&h.Latitude, validation.Required, is.Latitude),
		validation.Field(&h.Longitude, validation.Required, is.Longitude),
	)
	if err != nil {
		return err
	}
	err = validation.ValidateStruct(
		h.Address,
		validation.Field(&h.Address.Line1, validation.Required, validation.Length(5, 2000)),
	)
	if err != nil {
		return err
	}
	return nil
}

// // Hotel struct final version
// type Hotel struct {
// 	ProviderPropertyID             string                           `json:"providerPropertyID"` // Not updatable
// 	Addresses                      *[]Adress                        `json:"addresses"`
// 	Attributes                     *[]Attribute                     `json:"attributes"`
// 	BillingCurrencyCode            string                           `json:"billingCurrencyCode"` // Not updatable
// 	Contacts                       *[]Contact                       `json:"contacts"`
// 	Contents                       *[]Content                       `json:"contents"`
// 	CurrencyCode                   string                           `json:"currencyCode"` // Not updatable
// 	HideAddress                    bool                             `json:"hideAddress"`
// 	HideExactLocation              bool                             `json:"hideExactLocation"`
// 	InventorySettings              *InventorySetting                `json:"inventorySettings"`
// 	IsVacationRental               bool                             `json:"isVacationRental"`
// 	Latitude                       string                           `json:"latitude"`  // Not updatable
// 	Longitude                      string                           `json:"longitude"` // Not updatable
// 	Name                           string                           `json:"name"`
// 	Policies                       *[]Policy                        `json:"policies"`
// 	PropertyCollectedMandatoryFees *[]PropertyCollectedMandatoryFee `json:"propertyCollectedMandatoryFees"`
// 	Provider                       string                           `json:"provider"` // Not updatable
// 	ProviderPropertyURL            string                           `json:"providerPropertyURL"`
// 	Ratings                        *[]Rating                        `json:"ratings"`
// 	StructureType                  string                           `json:"structureType"`
// 	Taxes                          *[]Tax                           `json:"taxes"`
// 	TimeZone                       string                           `json:"timeZone"`
// }

// Address struct
type Address struct {
	Line1      string `json:"line1"` // This is the address used by travelers to reach your property. Minimum 5 Characters Required. If your address is longer than 40 characters, we recommend you continue with Line2.
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`                 // Town/City
	State      string `json:"state,omitempty"`      // State/Province
	PostalCode string `json:"postalCode,omitempty"` // ZIP/Postal Code
	Country    string `json:"country"`              // Use ISO 3166-1 alpha-2 or alpha-3. Stored in Expedia system as alpha-3.
}

// // Rating struct
// type Rating struct {
// 	Score       string `json:"score"`                 // Normalized Score, e.g. 4.0
// 	MaxScore    string `json:"maxScore"`              // The maximum score possible for the rating scale, e.g. 5.0 if highest rating is 5 stars.
// 	Source      string `json:"source,omitempty"`      // Indicate here if national or non-national rating and affiliation.
// 	Description string `json:"description,omitempty"` // Only supported value is "Stars"
// }

// // Attribute struct
// type Attribute struct {
// 	Code  string `json:"code"`
// 	Value string `json:"value,omitempty"`
// }

// // Contact struct
// type Contact struct {
// 	Position     string   `json:"position"`
// 	PhoneNumbers []string `json:"phoneNumbers,omitempty"`
// 	Emails       []string `json:"emails,omitempty"`
// }

// // Content struct
// type Content struct {
// 	Locale    string     `json:"locale"`
// 	Name      string     `json:"name"`
// 	Images    []string   `json:"images"`
// 	Amenities *[]Amenity `json:"amenities"`
// 	Paragraph Paragraph  `json:"paragraph"`
// }

// // Paragraph to describe property
// type Amenity struct {
// 	Code       string `json:"code"`
// 	DetailCode string `json:"detailCode"`
// 	Value      string `json:"value"`
// }

// // Paragraph to describe property
// type Paragraph struct {
// 	Code  string `json:"code"`
// 	Value string `json:"value"`
// }

// // InventorySetting struct
// type InventorySetting struct {
// 	Position     string   `json:"position"`
// 	PhoneNumbers string   `json:"phoneNumbers,omitempty"`
// 	Emails       []string `json:"emails,omitempty"`
// }

// // Policy struct
// type Policy struct {
// 	Code       string `json:"code"`
// 	Value      string `json:"value"`
// 	DetailCode string `json:"detailCode,omitempty"`
// }

// // Tax struct
// type Tax struct {
// 	Code       string `json:"code"`
// 	Value      string `json:"value"`
// 	DetailCode string `json:"detailCode,omitempty"`
// }

// // PropertyCollectedMandatoryFee struct
// type PropertyCollectedMandatoryFee struct {
// 	Code      string `json:"code"`
// 	Value     string `json:"value"`
// 	Scope     string `json:"scope"`
// 	Duration  string `json:"duration,omitempty"`
// 	StartDate string `json:"startDate,omitempty"`
// 	EndDate   string `json:"endDate,omitempty"`
// }
