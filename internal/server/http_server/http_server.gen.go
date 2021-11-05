// Package http_server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package http_server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	LoginScopes = "Login.Scopes"
)

// ErrorInfo defines model for ErrorInfo.
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ListMeta defines model for ListMeta.
type ListMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// TokenInfo defines model for TokenInfo.
type TokenInfo struct {
	ExpiresAt time.Time `json:"expiresAt"`
	IssuedAt  time.Time `json:"issuedAt"`
	Token     string    `json:"token"`
}

// TokenInfos defines model for TokenInfos.
type TokenInfos struct {
	AccessToken  TokenInfo `json:"accessToken"`
	RefreshToken TokenInfo `json:"refreshToken"`
}

// TokenRefresh defines model for TokenRefresh.
type TokenRefresh struct {
	RefreshToken string `json:"refreshToken"`
}

// UserCreate defines model for UserCreate.
type UserCreate struct {
	Email    string   `json:"email"`
	LoginId  string   `json:"loginId"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Role     UserRole `json:"role"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	Email   string   `json:"email"`
	Id      string   `json:"id"`
	LoginId string   `json:"loginId"`
	Phone   string   `json:"phone"`
	Role    UserRole `json:"role"`
}

// UserInfoList defines model for UserInfoList.
type UserInfoList struct {
	Metadata ListMeta   `json:"metadata"`
	Users    []UserInfo `json:"users"`
}

// UserRole defines model for UserRole.
type UserRole string

// List of UserRole
const (
	UserRole_admin UserRole = "admin"
	UserRole_user  UserRole = "user"
)

// UserUpdate defines model for UserUpdate.
type UserUpdate struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Role     UserRole `json:"role"`
}

// Limit defines model for Limit.
type Limit int

// Offset defines model for Offset.
type Offset int

// UserID defines model for UserID.
type UserID string

// PostTokensRefreshJSONBody defines parameters for PostTokensRefresh.
type PostTokensRefreshJSONBody TokenRefresh

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	Offset *Offset `json:"Offset,omitempty"`
	Limit  *Limit  `json:"Limit,omitempty"`
}

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody UserCreate

// PutUsersMeJSONBody defines parameters for PutUsersMe.
type PutUsersMeJSONBody UserUpdate

// PutUsersUserIDJSONBody defines parameters for PutUsersUserID.
type PutUsersUserIDJSONBody UserUpdate

// PostTokensRefreshJSONRequestBody defines body for PostTokensRefresh for application/json ContentType.
type PostTokensRefreshJSONRequestBody PostTokensRefreshJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PutUsersMeJSONRequestBody defines body for PutUsersMe for application/json ContentType.
type PutUsersMeJSONRequestBody PutUsersMeJSONBody

// PutUsersUserIDJSONRequestBody defines body for PutUsersUserID for application/json ContentType.
type PutUsersUserIDJSONRequestBody PutUsersUserIDJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /tokens/login)
	PostTokensLogin(w http.ResponseWriter, r *http.Request)

	// (POST /tokens/refresh)
	PostTokensRefresh(w http.ResponseWriter, r *http.Request)

	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams)

	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)

	// (DELETE /users/me)
	DeleteUsersMe(w http.ResponseWriter, r *http.Request)

	// (GET /users/me)
	GetUsersMe(w http.ResponseWriter, r *http.Request)

	// (PUT /users/me)
	PutUsersMe(w http.ResponseWriter, r *http.Request)

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

// PostTokensLogin operation middleware
func (siw *ServerInterfaceWrapper) PostTokensLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, LoginScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTokensLogin(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostTokensRefresh operation middleware
func (siw *ServerInterfaceWrapper) PostTokensRefresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTokensRefresh(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

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

// DeleteUsersMe operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersMe(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUsersMe operation middleware
func (siw *ServerInterfaceWrapper) GetUsersMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersMe(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutUsersMe operation middleware
func (siw *ServerInterfaceWrapper) PutUsersMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUsersMe(w, r)
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
		r.Post(options.BaseURL+"/tokens/login", wrapper.PostTokensLogin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/tokens/refresh", wrapper.PostTokensRefresh)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.PostUsers)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/me", wrapper.DeleteUsersMe)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/me", wrapper.GetUsersMe)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/users/me", wrapper.PutUsersMe)
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

	"H4sIAAAAAAAC/+xZX2/bNhD/KgG3R9VyvWTA9LSuDYYA6TakzZORB1Y622xFkeVRyYxA3304krJkW/If",
	"IM6KRG+27nj/+Lv7kdIjS5XUqoDCIksemeaGS7Bg3L9rIYWlH6JgCfteglmyiBVcAkuCMGKYLkBy0spg",
	"xsvcsuRiHDG71KQlCgtzMKyqIvb3bIbQay9I2walKIQsJUu67d0imKsPK3ua20VjLggjZuB7KQxkLLGm",
	"hLb5YBKtEcWcVWTSC13yl8Yoc1XMlKuLURqMFeBEqcqgw0DEJCDyeZesagcy9RYa/btVfurLV0gt2boW",
	"aD+C5dvu83pbNmsSMbUq8bbMKsvzLtFGbHnYWFVviF/YFeNn9Q2K7hrBv1oYwHcumJkykluWsIxbeGOF",
	"pOS3qicQS8iOWWHJ//5qe7WWg6gV3s68cDsxnqaA+Ln2/LOBGUvYT3HTSHFAUdyUxwU0M4CLYxdupNL2",
	"vmGzN5Ebr7WdymZIu6u41xm13HsD3EIHHCQXeWfP5GouiqusU6Y54oMyPcKFKrrb0Kgc9pWYgr0hvS34",
	"h3ha3oPF2mUUsumrQU9D9FZAZMcX5jS5C8q1KcDRWdPM2s5cguUZ95NsV1iriVdFrMTAQcKCxEMSqvss",
	"BMeN4cutBL3ZqAmpL5ubUEYoiH+mjGdSUMeRgdaipu606FZnR4L//wD48bAmYoS0NMIuP5F1n9U14YR+",
	"OI+k/4WjSJs5vbBWe1oVoSNSVVie2lZNGJYaS33x6/nk9zk9GqVKUooZYGqEtkIRtyOWevIN7Bkv7eIM",
	"wdyLlCLORQoFuloE3n+nebqAs8loTLtl8hBHEscPDw8j7qQjZeZxWIrx9dX7y78+Xb6ZjMajhZW5A5Gw",
	"Oezwew8GfWRvR+PR2FGvhoJrwRL2i3sUuQOJq1TsCAjjvK6YVr5PCCScUqQ2Z/8otG66oi8tbRtqRUGS",
	"8mQ8rmsIhVvOtc5F6gzEX1E1m8EPJhj0G7RebUrn/Am9NUepXmdvn8/ZxfNl1moclkxXLTO9q6jJ+Byb",
	"08kdKddAMS3G3gOVmtx9jwPaP1S2fFqc1C6qyo+SU2NygORpIdkFvBXfzqEDbn+CvQ3M2b4hTrsjaVTi",
	"cKmror2a/jpJfXEygK2dUl4dxs7H5y8d0P5wRien3qFZw/gUw7J1/znxqGzOuz8Ein97JcBaDcpYgn/h",
	"lYM/768D7YN77qD2EXqOccM0eLZpsJPSejfoJTbqALDT0E3ZxTblGsBOQzfhjUM/3QyQe9lE9Oi/c1QH",
	"0lHrq8hAST8qJe3cpIGWBpAdSktHXdUD6ugGvpPP1j6sDpw24PWpOM29LDX3NVwP+wCw9oa/VvLfEO6q",
	"/wIAAP//t6e7794gAAA=",
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
