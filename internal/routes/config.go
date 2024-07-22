package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func Newconfig() *Config {
	return &Config{}
}

// Cors is a method of the Config struct that adds Cross-Origin Resource Sharing (CORS) support to the HTTP handler.
// It creates a new CORS middleware with the specified options and wraps it around the provided HTTP handler.
func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           5,
	}).Handler(next)
}

// SetTimeOut sets the timeout value for the Config struct.
// It takes an integer representing the time in seconds and converts it to a time.Duration in seconds.
// It then assigns this duration to the timeout field of the Config struct.
// It returns a pointer to the modified Config struct.
func (c *Config) SetTimeOut(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}

// GetTimeOut returns the timeout duration set in the Config struct.
func (c *Config) GetTimeOut() time.Duration {
	return c.timeout
}
