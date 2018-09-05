package helpers

import (
	"testing"

	"gitlab.com/totakoko/onetimesecret/helpers/tests"
)

func Test_GenerateRandomString(t *testing.T) {
	assert := tests.SetupTest(t)

	assert.Equal("ONRhfKsU", GenerateRandomString(8))
	assert.Equal("OHoF", GenerateRandomString(4))
	assert.Equal("8iVdiJs5", GenerateRandomString(8))
	assert.Equal("huhtgVW8Q5MTfldn", GenerateRandomString(16))
}
