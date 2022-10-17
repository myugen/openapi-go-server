// Package petstore provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package petstore

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx echo.Context) error
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id int64) error
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"urvf/S0AAP//PxZu8FUSAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
