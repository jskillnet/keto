// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpsertOryAccessControlPolicyRoleOK UpsertOryAccessControlPolicyRoleOK UpsertOryAccessControlPolicyRoleOK UpsertOryAccessControlPolicyRoleOK handles this case with default header values.
//
// oryAccessControlPolicyRole
// swagger:model UpsertOryAccessControlPolicyRoleOK
type UpsertOryAccessControlPolicyRoleOK struct {

	// payload
	Payload *Role `json:"Payload,omitempty"`
}

// Validate validates this upsert ory access control policy role o k
func (m *UpsertOryAccessControlPolicyRoleOK) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpsertOryAccessControlPolicyRoleOK) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpsertOryAccessControlPolicyRoleOK) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpsertOryAccessControlPolicyRoleOK) UnmarshalBinary(b []byte) error {
	var res UpsertOryAccessControlPolicyRoleOK
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
