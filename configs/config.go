package configs

import (
	"os"

	"github.com/golobby/dotenv"
)

type AppConfig struct {
	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}
	Server struct {
		Host string `env:"HTTP_HOST"`
		Port string `env:"HTTP_PORT"`
	}

	Mongo struct {
		Url string `env:"MONGO_URL"`
		Db  string `env:"MONGO_DB"`
	}
	Postgres struct {
		Dsn string `env:"POSTGRES_DSN"`
	}
	Cache struct {
		DefaultExpireTimeSec int    `env:"CACHE_DEFAULT_EXPIRE_TIME_SEC"`
		CleanupIntervalHour  int    `env:"CACHE_CLEANUP_INTERVAL_HOUR"`
		RedisHost            string `env:"REDIS_HOST"`
		RedisPort            string `env:"REDIS_PORT"`
		RedisPassword        string `env:"REDIS_PASSWORD"`
		RedisDB              int    `env:"REDIS_DB"`
	}
	Auth struct {
		ignoreMethods []string `env:"AUTH_IGNORE_METHODS"`
		JWT           struct {
			SecretKey        string `env:"JWT_SECRET_KEY"`
			RefreshSecretKey string `env:"JWT_REFRESH_SECRET_KEY"`
			Issuer           string `env:"JWT_ISSUER"`
			ExpiredTime      int64  `env:"JWT_EXPIRED_TIME"`
			RefreshExpired   int64  `env:"JWT_REFRESH_EXPIRED_TIME"`
		}
		IgnoreMethods map[string]bool
		Google        struct {
			ClientID     string `env:"GOOGLE_CLIENT_ID"`
			ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
		}
	}
}

func NewAppConfig(envDir string) (*AppConfig, error) {
	config := &AppConfig{}
	file, err := os.Open(envDir)
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(config)
	if err != nil {
		return config, err
	}

	config.Auth.IgnoreMethods = make(map[string]bool)
	for _, method := range config.Auth.ignoreMethods {
		config.Auth.IgnoreMethods[method] = true
	}

	return config, err
}
