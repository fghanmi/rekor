// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
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
// */
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchIndex search index
//
// swagger:model SearchIndex
type SearchIndex struct {

	// hash
	// Pattern: ^[0-9a-fA-F]{64}$
	Hash string `json:"hash,omitempty"`

	// public key
	PublicKey *SearchIndexPublicKey `json:"publicKey,omitempty"`
}

// Validate validates this search index
func (m *SearchIndex) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIndex) validateHash(formats strfmt.Registry) error {

	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if err := validate.Pattern("hash", "body", string(m.Hash), `^[0-9a-fA-F]{64}$`); err != nil {
		return err
	}

	return nil
}

func (m *SearchIndex) validatePublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchIndex) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchIndex) UnmarshalBinary(b []byte) error {
	var res SearchIndex
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SearchIndexPublicKey search index public key
//
// swagger:model SearchIndexPublicKey
type SearchIndexPublicKey struct {

	// content
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// format
	// Required: true
	// Enum: [pgp x509 minisign]
	Format *string `json:"format"`

	// url
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this search index public key
func (m *SearchIndexPublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var searchIndexPublicKeyTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pgp","x509","minisign"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchIndexPublicKeyTypeFormatPropEnum = append(searchIndexPublicKeyTypeFormatPropEnum, v)
	}
}

const (

	// SearchIndexPublicKeyFormatPgp captures enum value "pgp"
	SearchIndexPublicKeyFormatPgp string = "pgp"

	// SearchIndexPublicKeyFormatX509 captures enum value "x509"
	SearchIndexPublicKeyFormatX509 string = "x509"

	// SearchIndexPublicKeyFormatMinisign captures enum value "minisign"
	SearchIndexPublicKeyFormatMinisign string = "minisign"
)

// prop value enum
func (m *SearchIndexPublicKey) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchIndexPublicKeyTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchIndexPublicKey) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("publicKey"+"."+"format", "body", m.Format); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormatEnum("publicKey"+"."+"format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *SearchIndexPublicKey) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("publicKey"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchIndexPublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchIndexPublicKey) UnmarshalBinary(b []byte) error {
	var res SearchIndexPublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
