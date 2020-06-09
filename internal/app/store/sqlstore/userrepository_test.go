package sqlstore_test

import (
	"testing"

	"github.com/kek0896/golang-edu/http-rest-api/internal/app/model"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store"
	"github.com/kek0896/golang-edu/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().CreateUser(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example1.org"
	_, err := s.User().FindUserByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().CreateUser(u)
	u, err = s.User().FindUserByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
