package tests

import (
	"math/rand"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// basic setup before each test
func SetupTest(t require.TestingT) *require.Assertions {
	// freeze the random generator
	rand.Seed(1)
	// filter all logs from the output
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	assert := require.New(t)
	return assert
}
