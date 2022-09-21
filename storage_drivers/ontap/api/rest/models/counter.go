// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Counter Representation of a counter and contains one of the following:
// * Scalar counter populates the 'name' and 'value' fields.
// * A 1D array populates the 'name', 'labels' and 'values' fields.
// * A 2D array is represented as a list of counter entries.
// ```
// "counters": [
//     // Scalar counter
//     {
//         "name": "memory",
//         "value": 4480
//     },
//     // one dimensional array "sys_read_latency_hist"
//     {
//         "name": "sys_read_latency_hist",
//         "labels": ["0 - <1ms", "1 - <2ms", ...],
//         "values": [0, 0, ...]
//     },
//     // Two dimensional array "foo" with ["Label 1", "Label 2"] as the first
//     // array dimension and labels ["w", "x", "y"] for the 2nd dimension
//     {
//         "name": "foo",
//         "labels": ["Label 1", "Label 2"],
//         "counters": [
//             {
//                 "label": "x",
//                 "values": [0, 0]
//             },
//             {
//                 "label": "y",
//                 "values": [0, 0]
//             },
//             {
//                 "label": "z",
//                 "values": [0, 0]
//             }
//         ]
//     }
// ```
//
//
// swagger:model counter
type Counter struct {

	// List of labels and values for the second dimension.
	Counters []*Counter2d `json:"counters,omitempty"`

	// List of labels for the first dimension.
	Labels []string `json:"labels,omitempty"`

	// Counter name.
	Name string `json:"name,omitempty"`

	// Scalar value.
	Value int64 `json:"value,omitempty"`

	// List of values in a one-dimensional counter.
	Values []int64 `json:"values,omitempty"`
}

// Validate validates this counter
func (m *Counter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Counter) validateCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.Counters) { // not required
		return nil
	}

	for i := 0; i < len(m.Counters); i++ {
		if swag.IsZero(m.Counters[i]) { // not required
			continue
		}

		if m.Counters[i] != nil {
			if err := m.Counters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this counter based on the context it is used
func (m *Counter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Counter) contextValidateCounters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Counters); i++ {

		if m.Counters[i] != nil {
			if err := m.Counters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Counter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Counter) UnmarshalBinary(b []byte) error {
	var res Counter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}