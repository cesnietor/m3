// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/m3/mcs/models"
)

// ProfilingStartCreatedCode is the HTTP code returned for type ProfilingStartCreated
const ProfilingStartCreatedCode int = 201

/*ProfilingStartCreated A successful response.

swagger:response profilingStartCreated
*/
type ProfilingStartCreated struct {

	/*
	  In: Body
	*/
	Payload *models.StartProfilingList `json:"body,omitempty"`
}

// NewProfilingStartCreated creates ProfilingStartCreated with default headers values
func NewProfilingStartCreated() *ProfilingStartCreated {

	return &ProfilingStartCreated{}
}

// WithPayload adds the payload to the profiling start created response
func (o *ProfilingStartCreated) WithPayload(payload *models.StartProfilingList) *ProfilingStartCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profiling start created response
func (o *ProfilingStartCreated) SetPayload(payload *models.StartProfilingList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfilingStartCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ProfilingStartDefault Generic error response.

swagger:response profilingStartDefault
*/
type ProfilingStartDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewProfilingStartDefault creates ProfilingStartDefault with default headers values
func NewProfilingStartDefault(code int) *ProfilingStartDefault {
	if code <= 0 {
		code = 500
	}

	return &ProfilingStartDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the profiling start default response
func (o *ProfilingStartDefault) WithStatusCode(code int) *ProfilingStartDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the profiling start default response
func (o *ProfilingStartDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the profiling start default response
func (o *ProfilingStartDefault) WithPayload(payload *models.Error) *ProfilingStartDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the profiling start default response
func (o *ProfilingStartDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProfilingStartDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
