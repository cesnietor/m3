// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Kubernetes Cloud
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StartMirroringRequest start mirroring request
//
// swagger:model startMirroringRequest
type StartMirroringRequest struct {

	// host source
	// Required: true
	HostSource *string `json:"host_source"`

	// host target
	// Required: true
	HostTarget *string `json:"host_target"`

	// image
	Image string `json:"image,omitempty"`

	// mirror flags
	MirrorFlags []string `json:"mirror_flags"`

	// source
	// Required: true
	Source *string `json:"source"`

	// target
	// Required: true
	Target *string `json:"target"`
}

// Validate validates this start mirroring request
func (m *StartMirroringRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StartMirroringRequest) validateHostSource(formats strfmt.Registry) error {

	if err := validate.Required("host_source", "body", m.HostSource); err != nil {
		return err
	}

	return nil
}

func (m *StartMirroringRequest) validateHostTarget(formats strfmt.Registry) error {

	if err := validate.Required("host_target", "body", m.HostTarget); err != nil {
		return err
	}

	return nil
}

func (m *StartMirroringRequest) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *StartMirroringRequest) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StartMirroringRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartMirroringRequest) UnmarshalBinary(b []byte) error {
	var res StartMirroringRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}