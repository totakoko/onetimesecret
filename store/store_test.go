package store

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
	"gitlab.dreau.fr/home/onetimesecret/conf"
	"gitlab.dreau.fr/home/onetimesecret/helpers/tests"
)

var (
	existingSecretKey string
)

func Test_StoreInit_OK(t *testing.T) {
	assert := tests.SetupTest(t)
	config, _ := conf.New()
	store := New(config.Store)

	err := store.Init()
	assert.NoError(err)
}
func Test_StoreInit_InvalidAddr(t *testing.T) {
	assert := tests.SetupTest(t)
	config, _ := conf.New()
	config.Store.Addr = "127.0.0.1:9999"
	store := New(config.Store)

	err := store.Init()
	assert.EqualError(err, "dial tcp 127.0.0.1:9999: getsockopt: connection refused")
}

func Test_StoreStoreSecret_OK(t *testing.T) {
	assert, store := SetupValidStore(t)

	key, err := store.StoreSecret("top-secret", 10*time.Second)
	assert.NoError(err)
	assert.Equal("OHoF8iVd", key) // first call to rand is constant with the same seed
}

func Test_Store_Expiration(t *testing.T) {
	assert, store := SetupValidStore(t)

	expectedKey := "OHoF8iVd"
	key, err := store.StoreSecret("top-secret", 50*time.Millisecond)
	assert.NoError(err)
	assert.Equal(expectedKey, key) // first call to rand is constant with the same seed
	time.Sleep(30 * time.Millisecond)

	secret, err := store.GetSecret(expectedKey)
	assert.NoError(err)
	assert.Equal("top-secret", secret)
	time.Sleep(30 * time.Millisecond)

	secret, err = store.GetSecret(expectedKey)
	assert.Equal(redis.Nil, err)
	assert.Empty(secret)
}

func Test_StoreGetSecret_OK(t *testing.T) {
	assert, store := SetupValidStore(t)

	secret, err := store.GetSecret(existingSecretKey)
	assert.NoError(err)
	assert.Equal("existing top-secret", secret)
}

func Test_StoreGetSecret_Missing(t *testing.T) {
	assert, store := SetupValidStore(t)

	secret, err := store.GetSecret("non-existing key")
	assert.Equal(redis.Nil, err)
	assert.Empty(secret)
}

func SetupValidStore(t require.TestingT) (*require.Assertions, *Store) {
	assert := tests.SetupTest(t)
	config, err := conf.New()
	assert.NoError(err)
	config.Store.Flush = true

	store := New(config.Store)
	err = store.Init()
	assert.NoError(err)

	existingSecretKey, err = store.StoreSecret("existing top-secret", 10*time.Second)
	assert.NoError(err)
	return assert, store
}
