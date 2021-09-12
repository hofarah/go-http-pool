package go_http_pool

import (
	"net/http"
	"time"
)

type Config struct {
	MaxIdleConnections        int
	MaxIdleConnectionsPerHost int
	MaxConnectionsPerHost     int
	Timeout                   int
}
type App struct {
	config Config
}

var client *http.Client

func New(config ...Config) *App {
	app := &App{}
	if len(config) > 0 {
		app.config = config[0]
	}
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = app.config.MaxIdleConnections
	t.MaxConnsPerHost = app.config.MaxConnectionsPerHost
	t.MaxIdleConnsPerHost = app.config.MaxIdleConnectionsPerHost

	client = &http.Client{
		Timeout:   time.Duration(app.config.Timeout) * time.Second,
		Transport: t,
	}
	return app
}

func (app *App) GetClient() *http.Client {
	return client
}
