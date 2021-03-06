/*
 * memo-go
 *
 * メモのCRUD用API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteMemo",
			strings.ToUpper("Delete"),
			"/memo/{id}",
			c.DeleteMemo,
		},
		{
			"GetMemo",
			strings.ToUpper("Get"),
			"/memo/{id}",
			c.GetMemo,
		},
		{
			"GetMemos",
			strings.ToUpper("Get"),
			"/memos",
			c.GetMemos,
		},
		{
			"PostMemo",
			strings.ToUpper("Post"),
			"/memo/{id}",
			c.PostMemo,
		},
		{
			"PostMemos",
			strings.ToUpper("Post"),
			"/memos",
			c.PostMemos,
		},
	}
}

// DeleteMemo - メモの削除
func (c *DefaultApiController) DeleteMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseInt32Parameter(params["id"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.DeleteMemo(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetMemo - メモの詳細
func (c *DefaultApiController) GetMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseInt32Parameter(params["id"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetMemo(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetMemos - メモの一覧取得
func (c *DefaultApiController) GetMemos(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetMemos(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PostMemo - メモの更新
func (c *DefaultApiController) PostMemo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseInt32Parameter(params["id"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	inlineObject1Param := InlineObject1{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObject1Param); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObject1Required(inlineObject1Param); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostMemo(r.Context(), idParam, inlineObject1Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PostMemos - メモの作成
func (c *DefaultApiController) PostMemos(w http.ResponseWriter, r *http.Request) {
	inlineObjectParam := InlineObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObjectParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObjectRequired(inlineObjectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostMemos(r.Context(), inlineObjectParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
