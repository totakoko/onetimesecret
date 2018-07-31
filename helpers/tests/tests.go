package tests

import (
	"math/rand"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"gitlab.dreau.fr/home/onetimesecret/conf"
)

// basic setup before each test
func SetupTest(t require.TestingT) (*require.Assertions, conf.Configuration) {
	// freeze the random generator
	rand.Seed(1)
	// filter all logs from the output
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	assert := require.New(t)
	config := conf.New()
	return assert, config
}
