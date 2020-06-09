package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter ...
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")                 // registration&auth
	s.router.HandleFunc("/like", s.handleLike()).Methods("POST")                         // like hotel
	s.router.HandleFunc("/getliked", s.handleGetLiked()).Methods("GET")                  // get likes
	s.router.HandleFunc("/getbooked", s.handleGetBooked()).Methods("GET")                // get bookings
	s.router.HandleFunc("/createhotels", s.handleHotelsCreate()).Methods("POST")         // create hotels (for partners)
	s.router.HandleFunc("/createproperties", s.handlePropertiesCreate()).Methods("POST") // create properties (for partners)
	s.router.HandleFunc("/search", s.handleSearch()).Methods("GET")                      // search for supply by given filters
	s.router.HandleFunc("/startbook", s.handleStartBookingProperty()).Methods("POST")    // start booking - user sends some data and locks properties
	s.router.HandleFunc("/approvebook", s.handleBooking()).Methods("POST")               // book property - add booking to db and lock properties forever
	s.router.HandleFunc("/cancelbooking", s.handleCancelBooking()).Methods("POST")       // unbook property
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")           // for web version (future plans)
	s.router.HandleFunc("/cleanhotels", s.handleCleanHotels()).Methods("POST")           // recreate hotels DB (for tests)
	s.router.HandleFunc("/cleanproperties", s.handleCleanProperties()).Methods("POST")   // recreate properties DB (for tests)
}

// s.router.HandleFunc("/polyatest", s.handlePolyaTest()).Methods("GET")
// s.router.HandleFunc("/createhotel", s.handleHotelCreate()).Methods("POST")
// s.router.HandleFunc("/fillgeoip/{id}", s.handleTableCreate()).Methods("GET") // fake method for adding

func (s *server) handleUsersCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().CreateUser(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)

	}
}

func (s *server) handleLike() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusCreated, "TODO")

	}
}

func (s *server) handleGetLiked() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusCreated, "TODO")

	}
}

func (s *server) handleGetBooked() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusCreated, "TODO")

	}
}

func (s *server) handleSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		u, _ := url.Parse(r.URL.String())
		values, _ := url.ParseQuery(u.RawQuery)

		var f model.Filter
		f.Datefrom = values.Get("datefrom")
		f.Dateto = values.Get("dateto")
		f.City = values.Get("city")
		f.Country = values.Get("country")
		f.Currency = values.Get("curr")
		f.StructureType = values.Get("type")
		f.Stars = values.Get("stars")
		f.PriceFrom = values.Get("pricefrom")
		f.PriceTo = values.Get("priceto")
		f.Wifi = values.Get("wifi")
		f.Breakfast = values.Get("breakfast")
		f.Garden = values.Get("garden")
		f.Parking = values.Get("parking")
		f.Playground = values.Get("playground")
		f.Pool = values.Get("pool")
		f.BedsNumber = values.Get("bedsnumber")
		f.RoomsNumber = values.Get("roomsnumber")
		f.Capacity = values.Get("capacity")
		f.OrderBy = values.Get("orderby")

		h, err := s.store.Hotel().Search(&f)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		s.respond(w, r, http.StatusOK, h)

	}
}

// TODO takes email / (in future uuid) and property id
// Writes new booking to booked db (uuid-property data)
// Sets property active flag to false
func (s *server) handleStartBookingProperty() http.HandlerFunc {

	type request struct {
		Sha1               string            `json:"sha1"`
		Email              string            `json:"email"`
		Phone              string            `json:"phone,omitempty"`
		Name               string            `json:"name,omitempty"`
		Surname            string            `json:"surname,omitempty"`
		CreditCard         *model.CreditCard `json:"credit_card,omitempty"`
		Status             string            `json:"status"`
		Timestamp          string            `json:"timestamp"`
		PropertyInternalID string            `json:"property_internal_id"`
		HotelInternalID    string            `json:"hotel_internal_id"`
		PropertyID         string            `json:"property_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		b := &model.Booking{
			Sha1:               req.Sha1,
			Email:              req.Email,
			Phone:              req.Phone,
			Name:               req.Name,
			Surname:            req.Surname,
			CreditCard:         req.CreditCard,
			Status:             req.Status,
			Timestamp:          string(time.Now().UnixNano()),
			PropertyInternalID: req.PropertyInternalID,
		}

		if err := s.store.Booking().StartBooking(b); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "waiting for final approval")

	}
}

func (s *server) handleBooking() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusCreated, "TODO")

	}
}

func (s *server) handleCancelBooking() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusCreated, "TODO")

	}
}

func (s *server) handleHotelsCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var hotels []model.Hotel
		if err := json.NewDecoder(r.Body).Decode(&hotels); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		for _, h := range hotels {
			if err := s.store.Hotel().CreateHotels(&h); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}
		s.respond(w, r, http.StatusCreated, hotels)

	}
}

func (s *server) handleCleanHotels() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := s.store.Hotel().CleanHotels(); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, "hotels_v2 now cleaned")

	}
}

func (s *server) handleCleanProperties() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := s.store.Hotel().CleanProperties(); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, "properties now cleaned")

	}
}

func (s *server) handlePropertiesCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var properties []model.Property
		if err := json.NewDecoder(r.Body).Decode(&properties); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		for _, p := range properties {
			if err := s.store.Property().CreateProperty(&p); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}
		s.respond(w, r, http.StatusCreated, properties)

	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindUserByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusOK, nil)

	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// handleTableCreate - helper func for writing table geoip2
// func (s *server) handleTableCreate() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)

// 		csvfile, err := os.Open("geoip2-files/out" + params["id"] + ".csv")
// 		if err != nil {
// 			s.respond(w, r, http.StatusBadRequest, "Couldn't open the csv file")
// 		}

// 		// Parse the file
// 		c := csv.NewReader(csvfile)
// 		//r := csv.NewReader(bufio.NewReader(csvfile))

// 		// Iterate through the records
// 		for {
// 			// Read each record from csv
// 			record, err := c.Read()
// 			if err == io.EOF {
// 				break
// 			}
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			///////////

// 			s.store.User().InsertGeoIP(record[1], record[2], record[3], record[4])

// 			//////////
// 		}
// 		s.respond(w, r, http.StatusCreated, "sucessfully filled table geoip2 from csv")

// 	}
// }

// func (s *server) handlePolyaTest() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		h := []*model.Hotel{
// 			{
// 				PropertyID:    "1234567",
// 				PropertyName:  "Sunbath Milan Resort",
// 				StructureType: "COTTAGE",
// 				Provider:      "Milan Resort",
// 				Price:         65350,
// 				Currency:      "RUB",
// 				Nights:        41,
// 				Stars:         5,
// 				Image:         "https://i.ibb.co/SXqC237/rclhmcc02jun19-800x600.jpg",
// 				Datefrom:      "2020-07-20",
// 				Dateto:        "2020-07-29",
// 				Latitude:      "45.469854",
// 				Longitude:     "9.182881",
// 				Address:       &model.Adress{Line1: "12 Ires Road", Line2: "", City: "Glasgow", Country: "Scotland", PostalCode: ""},
// 				Description:   "Located on the hills of Nice, a short distance from the famous Promenade des Anglais, Hotel Anis is one of the hotels in the Costa Azzurra (or Blue Coast) able to satisfy the different needs of its guests with comfort and first rate services. It is only 2 km from the airport and from highway exits. The hotel has a large parking area , a real luxury in a city like Nice.",
// 				Wifi:          true,
// 				Breakfast:     true,
// 				Parking:       false,
// 				Pool:          false,
// 				Playground:    false,
// 				Garden:        true,
// 				CheckIn:       "11:00 AM - 11:00 PM",
// 				CheckOut:      "10:00 AM",
// 			},
// 			{
// 				PropertyID:    "6574532",
// 				PropertyName:  "The Dominick Hotel",
// 				StructureType: "HOTEL",
// 				Provider:      "Dominick",
// 				Price:         216884,
// 				Currency:      "RUB",
// 				Nights:        8,
// 				Stars:         5,
// 				Image:         "https://i.ibb.co/W5hkWSd/newyork3.jpg",
// 				Datefrom:      "2020-09-01",
// 				Dateto:        "2020-09-09",
// 				Latitude:      "40.725507",
// 				Longitude:     "-74.005576",
// 				Address:       &model.Adress{Line1: "246 Spring Street, SoHo", Line2: "", City: "New York", Country: "USA", PostalCode: "10013"},
// 				Description:   "The Dominick Hotel provides relaxing, contemporary rooms with custom-made furniture, a 42-inch flat-screen TV and iPod docking station. An Italian marble bath and a bathroom vanity with TV are included. A seasonal 7th floor outdoor pool overlooks the city and features a connected bar. A fitness center and business services are also available. Complimentary WiFi is offered to all guests. A premium or platinum WiFi with more bandwidth is also available at a surcharge. The Washington Square Park is less than 1 km from the The Dominick Hotel. The Film Forum is 500 m away.",
// 				Wifi:          true,
// 				Breakfast:     true,
// 				Parking:       true,
// 				Pool:          true,
// 				Playground:    true,
// 				Garden:        true,
// 				CheckIn:       "From 15:00",
// 				CheckOut:      "12:00 AM",
// 			},
// 		}

// 		s.respond(w, r, http.StatusOK, h)

// 	}
// }
