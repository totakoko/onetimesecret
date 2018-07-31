package store

import (
	"math/rand"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"gitlab.dreau.fr/home/onetimesecret/common"
	"gitlab.dreau.fr/home/onetimesecret/conf"
)

var (
	existingSecretKey string
)

func Test_StoreStoreSecret_OK(t *testing.T) {
	assert, store := SetupStoreTest(t)

	key, err := store.StoreSecret("top-secret", 10*time.Second)
	assert.NoError(err)
	assert.Equal("OHoF8iVd", key) // first call to rand is constant with the same seed
	// assert.Equal("ONRhfKsU", key) // first call to rand is constant with the same seed
}

func Test_StoreGetSecret_OK(t *testing.T) {
	assert, store := SetupStoreTest(t)

	secret, err := store.GetSecret(existingSecretKey)
	assert.NoError(err)
	assert.Equal("existing top-secret", secret)
}

func SetupStoreTest(t require.TestingT) (*require.Assertions, common.Store) {
	rand.Seed(1)
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	assert := require.New(t)
	config := conf.New()
	config.Store.Flush = true
	store := New(config.Store)

	var err error
	existingSecretKey, err = store.StoreSecret("existing top-secret", 10*time.Second)
	assert.NoError(err)
	return assert, store
}
