// Package http_server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package http_server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ErrorInfo defines model for ErrorInfo.
type ErrorInfo struct {
	Code      string `json:"code"`
	RequestId string `json:"requestId"`
}

// ListMeta defines model for ListMeta.
type ListMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email    string `json:"email"`
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	Email   string `json:"email"`
	Id      string `json:"id"`
	LoginId string `json:"loginId"`
	Phone   string `json:"phone"`
}

// UserInfoList defines model for UserInfoList.
type UserInfoList struct {
	Metadata ListMeta   `json:"metadata"`
	Users    []UserInfo `json:"users"`
}

// UserUpdate defines model for UserUpdate.
type UserUpdate struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// Limit defines model for Limit.
type Limit int

// Offset defines model for Offset.
type Offset int

// UserID defines model for UserID.
type UserID string

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	Offset *Offset `json:"Offset,omitempty"`
	Limit  *Limit  `json:"Limit,omitempty"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody UserCreate

// PutUsersUserIDJSONBody defines parameters for PutUsersUserID.
type PutUsersUserIDJSONBody UserUpdate

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PutUsersUserIDJSONRequestBody defines body for PutUsersUserID for application/json ContentType.
type PutUsersUserIDJSONRequestBody PutUsersUserIDJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams)

	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)

	// (DELETE /users/{UserID})
	DeleteUsersUserID(w http.ResponseWriter, r *http.Request, userID UserID)

	// (GET /users/{UserID})
	GetUsersUserID(w http.ResponseWriter, r *http.Request, userID UserID)

	// (PUT /users/{UserID})
	PutUsersUserID(w http.ResponseWriter, r *http.Request, userID UserID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersParams

	// ------------- Optional query parameter "Offset" -------------
	if paramValue := r.URL.Query().Get("Offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "Offset", r.URL.Query(), &params.Offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter Offset: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "Limit" -------------
	if paramValue := r.URL.Query().Get("Limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "Limit", r.URL.Query(), &params.Limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter Limit: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUsers(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteUsersUserID operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "UserID" -------------
	var userID UserID

	err = runtime.BindStyledParameter("simple", false, "UserID", chi.URLParam(r, "UserID"), &userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter UserID: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersUserID(w, r, userID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUsersUserID operation middleware
func (siw *ServerInterfaceWrapper) GetUsersUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "UserID" -------------
	var userID UserID

	err = runtime.BindStyledParameter("simple", false, "UserID", chi.URLParam(r, "UserID"), &userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter UserID: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersUserID(w, r, userID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutUsersUserID operation middleware
func (siw *ServerInterfaceWrapper) PutUsersUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "UserID" -------------
	var userID UserID

	err = runtime.BindStyledParameter("simple", false, "UserID", chi.URLParam(r, "UserID"), &userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter UserID: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUsersUserID(w, r, userID)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.PostUsers)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/{UserID}", wrapper.DeleteUsersUserID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{UserID}", wrapper.GetUsersUserID)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/users/{UserID}", wrapper.PutUsersUserID)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXXU/jOhD9K2jufcxNegtc6eZpdwGtKrEf0oon1AdvMm0N8Qf2uFVV5b+v7LhpoCm0",
	"ErCrZd8gYx+fmTk+466gUEIriZIs5CvQzDCBhCb8d8kFJ/8Hl5DDnUOzhAQkEwh5DCZgixkK5leVOGGu",
	"IshPBwnQUvtVXBJO0UBdJ/BlMrG4Ey9Gu4CCSy6cgLwf78qiGZ23eJrRbAMXgwkYvHPcYAk5GYdd+Ahp",
	"yXA5hdpDNsGQ/IUxyozkRIW6GKXREMcQKlSJPQDNWWhpVPbBd6lcNxjdHeM2R/X9BgvyeJfc0icktk2h",
	"WrfmYV0SUG2Zt2OkiFV9oQfsqthctW5Ks7GPoy/0mUFGuM0SBeNVb6UqNeWyt04JaGbtQpkdwZmS+HR5",
	"1/gdtPXeJNLalUx/z3enwsvDM9wvCe5JdzLZm77XzXYKAomVrFHT3wYnkMNf2eb+Z1H8Wau6OgFnoxdw",
	"QmGf2tmWr27JMWPYciuzBjbZUNqVzZUuD1QWfzlR8QP15PfzqKZCSWIFdbiDddo6ffrfyfDd1H9KCyU8",
	"mxJtYbgmrrytWev08BbpiDmaHVk0c174QyteoLSBdbS895oVMzwapgNIwBl/xIxI51m2WCxSFqKpMtMs",
	"brXZ5ejs4vO3i3+G6SCdkahC3zhV+Mi5czS2YfZvOkgHwXE0SqY55HAcPiXBi0OfslZA08aTfBOZT81f",
	"DfiIdBWl0B091/0y2yzJ4rSokydXNnOqHvtWWq183p7HcDBYtwVlYMa0rngRuGU31me46gyLfWQfrl3o",
	"+v0W+hqdPON5m9G087CT1zvs9PUy8/JkU7t2EBj7u6tsj7C+KtsqK87YD6pcPmvL49Sr68YnXlhcv4Sw",
	"fmqv6yTaSbZqXnd18+qssJkQ9xVwHr4HDXTegttNetu351FXfrRuv5+435RrHjRtoxD8ENWuz23dlmBe",
	"xnLjc3C35f6R0GuZsf+1jGa+VtB+j9Z7r1KmeVreKjs/Tv83TqRFAfW4/hEAAP//isYwFZMQAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
