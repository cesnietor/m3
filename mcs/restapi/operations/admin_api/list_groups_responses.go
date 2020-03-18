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

// ListGroupsOKCode is the HTTP code returned for type ListGroupsOK
const ListGroupsOKCode int = 200

/*ListGroupsOK A successful response.

swagger:response listGroupsOK
*/
type ListGroupsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListGroupsResponse `json:"body,omitempty"`
}

// NewListGroupsOK creates ListGroupsOK with default headers values
func NewListGroupsOK() *ListGroupsOK {

	return &ListGroupsOK{}
}

// WithPayload adds the payload to the list groups o k response
func (o *ListGroupsOK) WithPayload(payload *models.ListGroupsResponse) *ListGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list groups o k response
func (o *ListGroupsOK) SetPayload(payload *models.ListGroupsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListGroupsDefault Generic error response.

swagger:response listGroupsDefault
*/
type ListGroupsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListGroupsDefault creates ListGroupsDefault with default headers values
func NewListGroupsDefault(code int) *ListGroupsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListGroupsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list groups default response
func (o *ListGroupsDefault) WithStatusCode(code int) *ListGroupsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list groups default response
func (o *ListGroupsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list groups default response
func (o *ListGroupsDefault) WithPayload(payload *models.Error) *ListGroupsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list groups default response
func (o *ListGroupsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListGroupsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
