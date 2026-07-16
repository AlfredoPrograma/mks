package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Environment string

const (
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

var (
	ErrCannotLoadConfig        = errors.New("cannot load config")
	ErrRequiredFieldMissing    = errors.New("required field missing")
	ErrInvalidEnvironmentField = errors.New("invalid environment field")
	ErrInvalidIntField         = errors.New("invalid integer field")
)

func parseEnvironment(key string) (Environment, error) {
	env := Environment(os.Getenv(key))

	switch env {
	case Development, Staging, Production:
		return env, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrInvalidEnvironmentField, key)
	}
}

func parseIntConfig(key string) (int, error) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidIntField, key)
	}

	return val, nil
}

func getRequiredConfig(key string) (string, error) {
	val, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("%w: %s", ErrRequiredFieldMissing, key)
	}

	return val, nil
}

func formatErrors(errors []error) string {
	strs := make([]string, len(errors))

	for i, err := range errors {
		strs[i] = "    - " + err.Error()
	}

	return strings.Join(strs, "\n")
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Config struct {
	DB          DBConfig
	Environment Environment
}

func Load() (Config, error) {
	errors := []error{}

	env, err := parseEnvironment("ENVIRONMENT")
	if err != nil {
		errors = append(errors, err)
	}

	dbHost, err := getRequiredConfig("DB_HOST")
	if err != nil {
		errors = append(errors, err)
	}

	dbUser, err := getRequiredConfig("DB_USER")
	if err != nil {
		errors = append(errors, err)
	}

	dbPassword, err := getRequiredConfig("DB_PASSWORD")
	if err != nil {
		errors = append(errors, err)
	}

	dbPort, err := parseIntConfig("DB_PORT")
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Config{}, fmt.Errorf("%w:\n%s", ErrCannotLoadConfig, formatErrors(errors))
	}

	return Config{
		Environment: env,
		DB: DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
		},
	}, nil
}
