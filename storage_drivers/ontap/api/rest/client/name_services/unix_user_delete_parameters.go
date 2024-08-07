// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUnixUserDeleteParams creates a new UnixUserDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixUserDeleteParams() *UnixUserDeleteParams {
	return &UnixUserDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixUserDeleteParamsWithTimeout creates a new UnixUserDeleteParams object
// with the ability to set a timeout on a request.
func NewUnixUserDeleteParamsWithTimeout(timeout time.Duration) *UnixUserDeleteParams {
	return &UnixUserDeleteParams{
		timeout: timeout,
	}
}

// NewUnixUserDeleteParamsWithContext creates a new UnixUserDeleteParams object
// with the ability to set a context for a request.
func NewUnixUserDeleteParamsWithContext(ctx context.Context) *UnixUserDeleteParams {
	return &UnixUserDeleteParams{
		Context: ctx,
	}
}

// NewUnixUserDeleteParamsWithHTTPClient creates a new UnixUserDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixUserDeleteParamsWithHTTPClient(client *http.Client) *UnixUserDeleteParams {
	return &UnixUserDeleteParams{
		HTTPClient: client,
	}
}

/*
UnixUserDeleteParams contains all the parameters to send to the API endpoint

	for the unix user delete operation.

	Typically these are written to a http.Request.
*/
type UnixUserDeleteParams struct {

	/* Name.

	   UNIX user name
	*/
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserDeleteParams) WithDefaults() *UnixUserDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unix user delete params
func (o *UnixUserDeleteParams) WithTimeout(timeout time.Duration) *UnixUserDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix user delete params
func (o *UnixUserDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix user delete params
func (o *UnixUserDeleteParams) WithContext(ctx context.Context) *UnixUserDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix user delete params
func (o *UnixUserDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix user delete params
func (o *UnixUserDeleteParams) WithHTTPClient(client *http.Client) *UnixUserDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix user delete params
func (o *UnixUserDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the unix user delete params
func (o *UnixUserDeleteParams) WithName(name string) *UnixUserDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the unix user delete params
func (o *UnixUserDeleteParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the unix user delete params
func (o *UnixUserDeleteParams) WithSvmUUID(svmUUID string) *UnixUserDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the unix user delete params
func (o *UnixUserDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UnixUserDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
