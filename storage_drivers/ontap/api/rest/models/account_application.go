// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountApplication account application
//
// swagger:model account_application
type AccountApplication struct {

	// Applications
	// Enum: [console http ontapi service_processor ssh]
	Application string `json:"application,omitempty"`

	// authentication methods
	AuthenticationMethods []string `json:"authentication_methods,omitempty"`

	// An optional additional authentication method for MFA. This only works with SSH as the application. It is ignored for all other applications.
	// Enum: [none password publickey nsswitch]
	SecondAuthenticationMethod *string `json:"second_authentication_method,omitempty"`
}

// Validate validates this account application
func (m *AccountApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountApplicationTypeApplicationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["console","http","ontapi","service_processor","ssh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationTypeApplicationPropEnum = append(accountApplicationTypeApplicationPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// application
	// Application
	// console
	// END RIPPY DEBUGGING
	// AccountApplicationApplicationConsole captures enum value "console"
	AccountApplicationApplicationConsole string = "console"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// application
	// Application
	// http
	// END RIPPY DEBUGGING
	// AccountApplicationApplicationHTTP captures enum value "http"
	AccountApplicationApplicationHTTP string = "http"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// application
	// Application
	// ontapi
	// END RIPPY DEBUGGING
	// AccountApplicationApplicationOntapi captures enum value "ontapi"
	AccountApplicationApplicationOntapi string = "ontapi"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// application
	// Application
	// service_processor
	// END RIPPY DEBUGGING
	// AccountApplicationApplicationServiceProcessor captures enum value "service_processor"
	AccountApplicationApplicationServiceProcessor string = "service_processor"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// application
	// Application
	// ssh
	// END RIPPY DEBUGGING
	// AccountApplicationApplicationSSH captures enum value "ssh"
	AccountApplicationApplicationSSH string = "ssh"
)

// prop value enum
func (m *AccountApplication) validateApplicationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationTypeApplicationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationEnum("application", "body", m.Application); err != nil {
		return err
	}

	return nil
}

var accountApplicationAuthenticationMethodsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["password","publickey","domain","nsswitch","certificate","saml"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationAuthenticationMethodsItemsEnum = append(accountApplicationAuthenticationMethodsItemsEnum, v)
	}
}

func (m *AccountApplication) validateAuthenticationMethodsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationAuthenticationMethodsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateAuthenticationMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthenticationMethods); i++ {

		// value enum
		if err := m.validateAuthenticationMethodsItemsEnum("authentication_methods"+"."+strconv.Itoa(i), "body", m.AuthenticationMethods[i]); err != nil {
			return err
		}

	}

	return nil
}

var accountApplicationTypeSecondAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","password","publickey","nsswitch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountApplicationTypeSecondAuthenticationMethodPropEnum = append(accountApplicationTypeSecondAuthenticationMethodPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// second_authentication_method
	// SecondAuthenticationMethod
	// none
	// END RIPPY DEBUGGING
	// AccountApplicationSecondAuthenticationMethodNone captures enum value "none"
	AccountApplicationSecondAuthenticationMethodNone string = "none"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// second_authentication_method
	// SecondAuthenticationMethod
	// password
	// END RIPPY DEBUGGING
	// AccountApplicationSecondAuthenticationMethodPassword captures enum value "password"
	AccountApplicationSecondAuthenticationMethodPassword string = "password"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// second_authentication_method
	// SecondAuthenticationMethod
	// publickey
	// END RIPPY DEBUGGING
	// AccountApplicationSecondAuthenticationMethodPublickey captures enum value "publickey"
	AccountApplicationSecondAuthenticationMethodPublickey string = "publickey"

	// BEGIN RIPPY DEBUGGING
	// account_application
	// AccountApplication
	// second_authentication_method
	// SecondAuthenticationMethod
	// nsswitch
	// END RIPPY DEBUGGING
	// AccountApplicationSecondAuthenticationMethodNsswitch captures enum value "nsswitch"
	AccountApplicationSecondAuthenticationMethodNsswitch string = "nsswitch"
)

// prop value enum
func (m *AccountApplication) validateSecondAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountApplicationTypeSecondAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountApplication) validateSecondAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondAuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecondAuthenticationMethodEnum("second_authentication_method", "body", *m.SecondAuthenticationMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account application based on context it is used
func (m *AccountApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountApplication) UnmarshalBinary(b []byte) error {
	var res AccountApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY