package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ConfNew_Defaults(t *testing.T) {
	assert := require.New(t)

	config, err := New()

	assert.NoError(err)
	assert.Equal("127.0.0.1:6379", config.Store.Addr)
	assert.Equal(0, config.Store.Database)
	assert.Equal(false, config.Store.Flush)
	assert.Equal(8, config.Store.KeyLength)
	assert.Equal("", config.Store.Password)
}

func Test_ConfNew_WithOverrides(t *testing.T) {
	assert := require.New(t)
	os.Setenv("OTS_STORE_ADDR", "127.0.0.1:8000")
	os.Setenv("OTS_STORE_PASSWORD", "securepassword")
	defer os.Unsetenv("OTS_STORE_ADDR")
	defer os.Unsetenv("OTS_STORE_PASSWORD")

	config, err := New()

	assert.NoError(err)
	assert.Equal("127.0.0.1:8000", config.Store.Addr)
	assert.Equal(0, config.Store.Database)
	assert.Equal(false, config.Store.Flush)
	assert.Equal(8, config.Store.KeyLength)
	assert.Equal("securepassword", config.Store.Password)
}

func Test_ConfNew_WithBadOverrides(t *testing.T) {
	assert := require.New(t)
	os.Setenv("OTS_STORE_DATABASE", "mybaddatabase") // string instead of int
	defer os.Unsetenv("OTS_STORE_DATABASE")

	_, err := New()

	assert.EqualError(err, "envconfig.Process: assigning OTS_STORE_DATABASE to Database: converting 'mybaddatabase' to type int. details: strconv.ParseInt: parsing \"mybaddatabase\": invalid syntax")
}
