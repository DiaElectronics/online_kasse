// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Status status
// swagger:model Status
type Status string

const (

	// StatusOffline captures enum value "offline"
	StatusOffline Status = "offline"

	// StatusOnline captures enum value "online"
	StatusOnline Status = "online"
)

// for schema
var statusEnum []interface{}

func init() {
	var res []Status
	if err := json.Unmarshal([]byte(`["offline","online"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusEnum = append(statusEnum, v)
	}
}

func (m Status) validateStatusEnum(path, location string, value Status) error {
	if err := validate.Enum(path, location, value, statusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this status
func (m Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}