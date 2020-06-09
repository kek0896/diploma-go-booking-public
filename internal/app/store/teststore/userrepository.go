package teststore

import (
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// InsertGeoIP fake
// func (r *UserRepository) InsertGeoIP(geonameID string, countryISOCode string, countryName string, cityName string) error {

// 	return nil

// }

// CreateUser ...
func (r *UserRepository) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// WriteSha1 ...
func (r *UserRepository) WriteSha1(u *model.User) error {

	return nil
}

// FindUserByEmail ...
func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
