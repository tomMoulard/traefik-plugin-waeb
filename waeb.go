// Package waeb builds a middleware that works like a web server
package traefik_plugin_waeb

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Root string `json:"root,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Root: ".",
	}
}

// New created a new Waeb plugin.
func New(_ context.Context, _ http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.FileServer(http.Dir(config.Root)), nil
}
