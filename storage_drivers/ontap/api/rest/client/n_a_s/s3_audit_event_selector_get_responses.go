// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3AuditEventSelectorGetReader is a Reader for the S3AuditEventSelectorGet structure.
type S3AuditEventSelectorGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3AuditEventSelectorGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3AuditEventSelectorGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3AuditEventSelectorGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3AuditEventSelectorGetOK creates a S3AuditEventSelectorGetOK with default headers values
func NewS3AuditEventSelectorGetOK() *S3AuditEventSelectorGetOK {
	return &S3AuditEventSelectorGetOK{}
}

/* S3AuditEventSelectorGetOK describes a response with status code 200, with default header values.

OK
*/
type S3AuditEventSelectorGetOK struct {
	Payload *models.S3AuditEventSelector
}

func (o *S3AuditEventSelectorGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/event-selectors/{svm.uuid}/{bucket}][%d] s3AuditEventSelectorGetOK  %+v", 200, o.Payload)
}
func (o *S3AuditEventSelectorGetOK) GetPayload() *models.S3AuditEventSelector {
	return o.Payload
}

func (o *S3AuditEventSelectorGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3AuditEventSelector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3AuditEventSelectorGetDefault creates a S3AuditEventSelectorGetDefault with default headers values
func NewS3AuditEventSelectorGetDefault(code int) *S3AuditEventSelectorGetDefault {
	return &S3AuditEventSelectorGetDefault{
		_statusCode: code,
	}
}

/* S3AuditEventSelectorGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3AuditEventSelectorGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 audit event selector get default response
func (o *S3AuditEventSelectorGetDefault) Code() int {
	return o._statusCode
}

func (o *S3AuditEventSelectorGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/event-selectors/{svm.uuid}/{bucket}][%d] s3_audit_event_selector_get default  %+v", o._statusCode, o.Payload)
}
func (o *S3AuditEventSelectorGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3AuditEventSelectorGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}