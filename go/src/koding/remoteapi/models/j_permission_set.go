package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// JPermissionSet j permission set
// swagger:model JPermissionSet
type JPermissionSet struct {

	// is custom
	IsCustom bool `json:"isCustom,omitempty"`

	// permissions
	Permissions []interface{} `json:"permissions"`
}

// Validate validates this j permission set
func (m *JPermissionSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JPermissionSet) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	return nil
}
