package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
	"time"
)

// AppConfig - конфигурация приложения.
type AppConfig struct {
	Name                string `yaml:"name"`
	ReadTimeout         int    `yaml:"read_timeout"`
	WriteTimeout        int    `yaml:"write_timeout"`
	EnableRequestID     bool   `yaml:"enable_request_id"`
	EnableLogging       bool   `yaml:"enable_logging"`
	EnableErrorHandling bool   `yaml:"enable_error_handling"`
}

// MongoConfig конфигурация MongoDB.
type MongoConfig struct {
	Host           string        `yaml:"host" validate:"required"`
	Port           int           `yaml:"port" validate:"required"`
	User           string        `yaml:"user" validate:"required"`
	Password       string        `yaml:"password" validate:"required"`
	AuthSource     string        `yaml:"auth_source" validate:"required"`
	Database       string        `yaml:"database" validate:"required"`
	ConnectTimeout time.Duration `yaml:"connect_timeout" validate:"required"`
}

// URI формирование строки подключения к БД.
func (c *MongoConfig) URI() *url.URL {
	hostPost := fmt.Sprintf("%s:%d", c.Host, c.Port)

	return &url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(c.User, c.Password),
		Host:   hostPost,
	}
}

// HTTPConfig - конфигурация HTTP сервера.
type HTTPConfig struct {
	Host string `yaml:"host" validate:"required"`
	Port int    `yaml:"port" validate:"required"`
}

// LoggingConfig - конфигурация логирования.
type LoggingConfig struct {
	Level          string `yaml:"level"`
	Format         string `yaml:"format"`
	EnableHTTPLogs bool   `yaml:"enable_http_logs"`
}

// Config основная конфигурация.
type Config struct {
	App     AppConfig     `yaml:"app"`
	HTTP    HTTPConfig    `yaml:"http"`
	MongoDB MongoConfig   `yaml:"mongodb"`
	Logging LoggingConfig `yaml:"logging"`
}

// LoadConfig загрузка конфига из файла.
func LoadConfig(configPath string) (Config, error) {
	_ = godotenv.Load()

	raw, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	// Подставляем переменные окружения
	expanded := os.ExpandEnv(string(raw))

	// Парсим YAML
	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config yaml: %w", err)
	}
	if err = cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// Validate валидация конфига.
func (c *Config) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("Config.Validate: %w", err)
	}

	return nil
}
