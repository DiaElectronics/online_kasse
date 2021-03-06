// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RelayReport relay report
// swagger:model RelayReport
type RelayReport struct {

	// hash
	// Required: true
	Hash Hash `json:"hash"`

	// relay stats
	RelayStats []*RelayStat `json:"relayStats"`
}

// Validate validates this relay report
func (m *RelayReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelayStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelayReport) validateHash(formats strfmt.Registry) error {

	if err := m.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hash")
		}
		return err
	}

	return nil
}

func (m *RelayReport) validateRelayStats(formats strfmt.Registry) error {

	if swag.IsZero(m.RelayStats) { // not required
		return nil
	}

	for i := 0; i < len(m.RelayStats); i++ {
		if swag.IsZero(m.RelayStats[i]) { // not required
			continue
		}

		if m.RelayStats[i] != nil {
			if err := m.RelayStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relayStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelayReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelayReport) UnmarshalBinary(b []byte) error {
	var res RelayReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
