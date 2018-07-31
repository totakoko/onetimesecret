package store

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/conf"
	"gitlab.dreau.fr/home/onetimesecret/helpers"
)

// la configuration se trouve dans la package conf, car sinon il y a des cycles d'imports
type Store struct {
	config      conf.StoreConfig
	redisClient *redis.Client
}

func New(config conf.StoreConfig) *Store {
	return &Store{
		config: config,
		redisClient: redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.Database,
		}),
	}
}

/*
Init : establishes a connection to the redis server.
*/
func (s *Store) Init() error {
	err := s.redisClient.Ping().Err()
	if err != nil {
		return err
	}
	if s.config.Flush {
		err = s.redisClient.FlushDB().Err()
		return err
	}
	return nil
}

func (s *Store) StoreSecret(secret string, expiration time.Duration) (string, error) {
	key := helpers.Rand(s.config.KeyLength)
	_, err := s.redisClient.Set(secretPath(key), secret, expiration).Result()
	log.Info().Msgf("Stored new secret at %s", key)
	return key, err
}

// pas de conversion en objet, car il est de toute façon reserialisé pour être renvoyé au client
func (s *Store) GetSecret(key string) (string, error) {
	secretStr, err := s.redisClient.Get(secretPath(key)).Result()
	log.Info().Msgf("Reading secret %s", key)
	return secretStr, err
}

func secretPath(key string) string {
	return "ots:secrets:" + key
}
