package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func Test_hashPassword(t *testing.T) {
	var err error
	password := "test"
	hashedPassword, err := hashPassword(password)
	assert.Nil(t, err)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	assert.Nil(t, err)
}
