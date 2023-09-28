package brycpt_test

import (
	encrypt "service-user/internal/infra/bcrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bcrypt(t *testing.T) {
	test, err := encrypt.EncriptarPassword("asdfasgsag")
	assert.NotNil(t, test)
	assert.Nil(t, err)
}

func Test_Verify(t *testing.T) {
	err := encrypt.VerificarPassword("asdfasgsag", "$2a$10$ieAHeghMaE6.mfrtdbI0E.9aFbakjAy237vIriN09e7y9BAQcUiIq")
	assert.Nil(t, err)
}
