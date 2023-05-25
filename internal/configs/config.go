package configs

import (
	"os"
	"siteavliable/internal/models"
	redisclient "siteavliable/pkg/client/redis"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - ...
type Config struct {
	Port               string
	LogPath            string
	FilePath           string
	UpdateTimeout      int
	SaveMetricsTimeout int
	Redis              redisclient.Config
	AdminAuth          models.AuthCred
}

// LoadConfig load config if path is not empty.
// If path isn't correct, it returns error
func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if len(path) != 0 {
		godotenv.Load(path)
	}
	cfg.Port = os.Getenv("PORT")
	cfg.LogPath = os.Getenv("LOG_PATH")
	cfg.FilePath = os.Getenv("URL_FILE_PATH")
	t, err := strconv.Atoi(os.Getenv("UPDATE_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	cfg.UpdateTimeout = t
	t, err = strconv.Atoi(os.Getenv("SAVE_METRICS_TIME"))
	if err != nil {
		return nil, err
	}
	cfg.SaveMetricsTimeout = t
	cfg.Redis.Addr = os.Getenv("REDIS_ADDR")
	cfg.Redis.Set = os.Getenv("REDIS_SET")
	t, err = strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}
	cfg.Redis.DB = t

	cfg.AdminAuth.Pass = os.Getenv("ADMIN_PASS")
	cfg.AdminAuth.User = os.Getenv("ADMIN_USER")

	return &cfg, nil
}
