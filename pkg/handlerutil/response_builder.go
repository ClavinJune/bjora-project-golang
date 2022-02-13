// Copyright 2022 ClavinJune/bjora
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlerutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Response initiates response builder
func Response() *ResponseBuilder {
	return &ResponseBuilder{
		header: make(http.Header),
	}
}

// ResponseBuilder builds response
type ResponseBuilder struct {
	statusCode int
	header     http.Header
	body       []byte
	err        error
}

// Code set statusCode
func (r *ResponseBuilder) Code(statusCode int) *ResponseBuilder {
	if r.err != nil {
		return r
	}

	if http.StatusText(statusCode) == "" {
		r.err = fmt.Errorf("ResponseBuilder.Code: unknown status code %d", statusCode)
	} else {
		r.statusCode = statusCode
	}

	return r
}

// Error set error
func (r *ResponseBuilder) Error(err error) *ResponseBuilder {
	r.err = err
	return r
}

// JSON set body response as JSON
func (r *ResponseBuilder) JSON(i interface{}) *ResponseBuilder {
	if r.err != nil {
		return r
	}

	r.header.Set("Content-Type", "application/json")

	b := bytes.NewBuffer(r.body)
	b.Reset()
	if err := json.NewEncoder(b).Encode(i); err != nil {
		r.err = fmt.Errorf("ResponseBuilder.json: %w", err)
	} else {
		r.body = b.Bytes()
	}
	return r
}

// Write writes response builder to the http.ResponseWriter
func (r *ResponseBuilder) Write(w http.ResponseWriter) {
	if r.err != nil {
		e := r.err
		r.err = nil

		r.JSON(struct {
			Error string `json:"error"`
		}{
			Error: e.Error(),
		})

		for k, v := range r.header {
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(http.StatusInternalServerError)

		_, _ = fmt.Fprint(w, string(r.body))
	}
}
