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

// Waeb a plugin to serve web content.
type Waeb struct {
	next   http.Handler
	name   string
	config *Config
}

// New created a new Waeb plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Waeb{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (a *Waeb) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	http.FileServer(http.Dir(a.config.Root)).ServeHTTP(rw, req)
}
