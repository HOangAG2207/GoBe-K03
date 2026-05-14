package config

import (
	"log"
	"net/http"

	"github.com/HOangAG2207/GoBe-K03/internal/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Engine defines the contract for starting the application
type Engine interface {
	// Start runs the HTTP server
	Start() error
	// ServeHTTP allows the engine to handle HTTP requests.
	// It satisfies the http.Handler interface.
	// Parameters:
	//   - w: The http.ResponseWriter to write the response
	//   - r: The incoming http.Request to be handled
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// engine is the concrete implementation of Engine
type engine struct {
	// app is the Echo instance used to handle HTTP requests
	app *echo.Echo
}

// NewEngine initializes and configures the application
func NewEngine() Engine {

	// ===== 1. Initialize Echo instance =====
	app := echo.New()

	// ===== 2. Register middlewares =====

	// Enable CORS for all origins
	// Note: In production, you should restrict allowed origins
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Log incoming HTTP requests
	app.Use(middleware.RequestLogger())

	// Recover from panics and prevent server crash
	app.Use(middleware.Recover())

	// ===== 3. Create engine instance =====
	e := &engine{
		app: app,
	}

	// ===== 4. Initialize routes =====
	e.InitRoutes()

	return e
}

// Start runs the HTTP server
func (e *engine) Start() error {

	// ===== 1. Define server port =====

	// Default port (can be replaced by config/env)
	port := "8080"

	// Ensure port has ":" prefix (required by Echo)
	if port[0] != ':' {
		port = ":" + port
	}

	// ===== 2. Start server =====
	log.Printf("Server running at %s\n", port)

	// Start Echo server
	return e.app.Start(port)
}

// InitRoutes registers all application routes
func (e *engine) InitRoutes() {
	// Group API routes under /api prefix
	route.RegisterRoutes(e.app)
}
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}
