package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	petstoreapi "github.com/myugen/openapi-go-server/api/petstore"
)

const apiBasePath = "/api"

func updateSwaggerPathsWithBaseURL(swagger *openapi3.T, pathPrefix string) {
	var updatedPaths = make(openapi3.Paths)

	for key, value := range swagger.Paths {
		updatedPaths[pathPrefix+key] = value
	}

	swagger.Paths = updatedPaths
}

func main() {
	port := flag.Int("port", 8080, "Port to run server [default: 8080].")
	flag.Parse()
	petStoreSwagger, err := petstoreapi.GetSwagger()
	if err != nil {
		log.Panicf("Error loading swagger petstore spec\n: %s", err)
	}
	updateSwaggerPathsWithBaseURL(petStoreSwagger, apiBasePath)

	// Clear out the servers array in the swagger pet store spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	petStoreSwagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStoreHandlers := petstoreapi.NewHandlers()

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(petStoreSwagger))

	apiRoute := e.Group(apiBasePath)
	// We now register our pet store handlers
	petstoreapi.RegisterHandlers(apiRoute, petStoreHandlers)

	// And we serve HTTP until the world ends.
	go func() {
		if err := e.Start(fmt.Sprintf("0.0.0.0:%d", *port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
