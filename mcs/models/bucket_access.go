// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2019 MinIO, Inc.
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

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BucketAccess bucket access
// swagger:model bucketAccess
type BucketAccess string

const (

	// BucketAccessPRIVATE captures enum value "PRIVATE"
	BucketAccessPRIVATE BucketAccess = "PRIVATE"

	// BucketAccessPUBLIC captures enum value "PUBLIC"
	BucketAccessPUBLIC BucketAccess = "PUBLIC"

	// BucketAccessCUSTOM captures enum value "CUSTOM"
	BucketAccessCUSTOM BucketAccess = "CUSTOM"
)

// for schema
var bucketAccessEnum []interface{}

func init() {
	var res []BucketAccess
	if err := json.Unmarshal([]byte(`["PRIVATE","PUBLIC","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bucketAccessEnum = append(bucketAccessEnum, v)
	}
}

func (m BucketAccess) validateBucketAccessEnum(path, location string, value BucketAccess) error {
	if err := validate.Enum(path, location, value, bucketAccessEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this bucket access
func (m BucketAccess) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBucketAccessEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
