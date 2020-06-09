package sqlstore

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"time"

	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// InsertGeoIP ...
// func (r *UserRepository) InsertGeoIP(geonameID string, countryISOCode string, countryName string, cityName string) error {

// 	return r.store.db.QueryRow(
// 		"INSERT INTO geoip2 (geoname_id, country_iso_code, country_name, city_name) VALUES ($1, $2, $3, $4)",
// 		geonameID, countryISOCode, countryName, cityName,
// 	).Scan()

// }

// CreateUser ...
func (r *UserRepository) CreateUser(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	s := u.Email + u.EncryptedPassword + string(time.Now().UnixNano())

	h := sha1.New()
	h.Write([]byte(s))
	u.Sha1 = fmt.Sprintf("%x", h.Sum(nil))

	return r.store.db.QueryRow(
		"INSERT INTO users_v2 (sha1, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Sha1,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

}

// FindUserByEmail ...
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
