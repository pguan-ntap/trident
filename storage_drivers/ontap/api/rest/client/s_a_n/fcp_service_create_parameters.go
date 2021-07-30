// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFcpServiceCreateParams creates a new FcpServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcpServiceCreateParams() *FcpServiceCreateParams {
	return &FcpServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcpServiceCreateParamsWithTimeout creates a new FcpServiceCreateParams object
// with the ability to set a timeout on a request.
func NewFcpServiceCreateParamsWithTimeout(timeout time.Duration) *FcpServiceCreateParams {
	return &FcpServiceCreateParams{
		timeout: timeout,
	}
}

// NewFcpServiceCreateParamsWithContext creates a new FcpServiceCreateParams object
// with the ability to set a context for a request.
func NewFcpServiceCreateParamsWithContext(ctx context.Context) *FcpServiceCreateParams {
	return &FcpServiceCreateParams{
		Context: ctx,
	}
}

// NewFcpServiceCreateParamsWithHTTPClient creates a new FcpServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcpServiceCreateParamsWithHTTPClient(client *http.Client) *FcpServiceCreateParams {
	return &FcpServiceCreateParams{
		HTTPClient: client,
	}
}

/* FcpServiceCreateParams contains all the parameters to send to the API endpoint
   for the fcp service create operation.

   Typically these are written to a http.Request.
*/
type FcpServiceCreateParams struct {

	/* Info.

	   The property values for the new FC Protocol service.

	*/
	Info *models.FcpService

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fcp service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceCreateParams) WithDefaults() *FcpServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fcp service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcpServiceCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FcpServiceCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fcp service create params
func (o *FcpServiceCreateParams) WithTimeout(timeout time.Duration) *FcpServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fcp service create params
func (o *FcpServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fcp service create params
func (o *FcpServiceCreateParams) WithContext(ctx context.Context) *FcpServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fcp service create params
func (o *FcpServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fcp service create params
func (o *FcpServiceCreateParams) WithHTTPClient(client *http.Client) *FcpServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fcp service create params
func (o *FcpServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fcp service create params
func (o *FcpServiceCreateParams) WithInfo(info *models.FcpService) *FcpServiceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fcp service create params
func (o *FcpServiceCreateParams) SetInfo(info *models.FcpService) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the fcp service create params
func (o *FcpServiceCreateParams) WithReturnRecords(returnRecords *bool) *FcpServiceCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fcp service create params
func (o *FcpServiceCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *FcpServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}