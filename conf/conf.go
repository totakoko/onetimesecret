package conf

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	ListenPort int    `default:"5000"`
	LogLevel   string `default:"info"`
	PublicURL  string `default:"http://localhost:5000/"`
	Store      StoreConfig
}

type StoreConfig struct {
	Addr                string        `default:"127.0.0.1:6379"`
	Password            string        `default:""`
	Database            int           `default:"0"`
	Flush               bool          `default:"false"`
	KeyLength           int           `default:"8"`
	MaxSecretSize       int           `default:"10485760"` // 10Mo
	MaxSecretExpiration time.Duration `default:"2678400s"` // 1 month
}

func New() (Configuration, error) {
	var conf Configuration
	err := envconfig.Process("OTS", &conf)
	if err != nil {
		return Configuration{}, err
	}
	return conf, err
}
