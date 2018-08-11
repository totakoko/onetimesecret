package store

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/common/errors"
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
		log.Warn().Msg("Flushing database")
		err = s.redisClient.FlushDB().Err()
		return err
	}
	return nil
}

func (s *Store) StoreSecret(secret string, expiration time.Duration) (string, error) {
	key := helpers.GenerateRandomString(s.config.KeyLength)
	_, err := s.redisClient.Set(secretPath(key), secret, expiration).Result()
	log.Info().Msgf("Stored new secret at %s (exp %s)", key, expiration.String())
	return key, err
}

// pas de conversion en objet, car il est de toute façon reserialisé pour être renvoyé au client
func (s *Store) GetSecret(key string) (string, error) {
	secretFullKey := secretPath(key)

	pipeline := s.redisClient.TxPipeline()
	get := pipeline.Get(secretFullKey)
	pipeline.Del(secretFullKey)

	_, err := pipeline.Exec()
	switch err {
	case redis.Nil:
		return "", errors.MissingResource("missing secret")
	case nil:
		log.Info().Msgf("Reading secret at %s", key)
		return get.Val(), nil
	default:
		return "", err
	}
}

func secretPath(key string) string {
	return "ots:secrets:" + key
}
