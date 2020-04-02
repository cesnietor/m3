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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ProfilerType profiler type
//
// swagger:model profilerType
type ProfilerType string

const (

	// ProfilerTypeCPU captures enum value "cpu"
	ProfilerTypeCPU ProfilerType = "cpu"

	// ProfilerTypeMem captures enum value "mem"
	ProfilerTypeMem ProfilerType = "mem"

	// ProfilerTypeBlock captures enum value "block"
	ProfilerTypeBlock ProfilerType = "block"

	// ProfilerTypeMutex captures enum value "mutex"
	ProfilerTypeMutex ProfilerType = "mutex"

	// ProfilerTypeTrace captures enum value "trace"
	ProfilerTypeTrace ProfilerType = "trace"

	// ProfilerTypeThreads captures enum value "threads"
	ProfilerTypeThreads ProfilerType = "threads"

	// ProfilerTypeGoroutines captures enum value "goroutines"
	ProfilerTypeGoroutines ProfilerType = "goroutines"
)

// for schema
var profilerTypeEnum []interface{}

func init() {
	var res []ProfilerType
	if err := json.Unmarshal([]byte(`["cpu","mem","block","mutex","trace","threads","goroutines"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		profilerTypeEnum = append(profilerTypeEnum, v)
	}
}

func (m ProfilerType) validateProfilerTypeEnum(path, location string, value ProfilerType) error {
	if err := validate.Enum(path, location, value, profilerTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this profiler type
func (m ProfilerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProfilerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
