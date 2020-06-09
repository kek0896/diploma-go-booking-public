package model

// validation "github.com/go-ozzo/ozzo-validation"
// "github.com/go-ozzo/ozzo-validation/is"

// TODO add order by

// Filter data
type Filter struct {
	Country       string `json:"country,omitempty"`
	City          string `json:"city,omitempty"`
	StructureType string `json:"structureType"` // hotel / apartment etc
	PriceFrom     string `json:"priceFrom,omitempty"`
	PriceTo       string `json:"priceTo,omitempty"`
	Currency      string `json:"currency,omitempty"`
	Stars         string `json:"stars,omitempty"`
	Datefrom      string `json:"datefrom,omitempty"` // date format: 2016-06-23
	Dateto        string `json:"dateto,omitempty"`
	Wifi          string `json:"wifi,omitempty"`
	Breakfast     string `json:"breakfast,omitempty"`
	Parking       string `json:"parking,omitempty"`
	Pool          string `json:"pool,omitempty"`
	Playground    string `json:"playgraund,omitempty"`
	Garden        string `json:"garden,omitempty"`
	RoomsNumber   string `json:"rooms_number"`
	BedsNumber    string `json:"beds_number"`
	Capacity      string `json:"capacity"`
	OrderBy       string `json:"order_by"`
	Desc          string `json:"desc"`
}

// func (f *Filter) Validate() error {
// 	err := validation.ValidateStruct(
// 		f,
// 		validation.Field(&f.Currency, validation.By(requiredIf(f.PriceFrom != 0)), is.CurrencyCode),
// 		validation.Field(&f.Currency, validation.By(requiredIf(f.PriceTo != 0)), is.CurrencyCode),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
