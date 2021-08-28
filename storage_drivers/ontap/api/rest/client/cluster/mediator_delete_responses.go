// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// MediatorDeleteReader is a Reader for the MediatorDelete structure.
type MediatorDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediatorDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMediatorDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediatorDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediatorDeleteAccepted creates a MediatorDeleteAccepted with default headers values
func NewMediatorDeleteAccepted() *MediatorDeleteAccepted {
	return &MediatorDeleteAccepted{}
}

/* MediatorDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MediatorDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *MediatorDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediatorDeleteAccepted  %+v", 202, o.Payload)
}
func (o *MediatorDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MediatorDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediatorDeleteDefault creates a MediatorDeleteDefault with default headers values
func NewMediatorDeleteDefault(code int) *MediatorDeleteDefault {
	return &MediatorDeleteDefault{
		_statusCode: code,
	}
}

/* MediatorDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 13369377    | Mediator field "mediator.id" does not exist.|

*/
type MediatorDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the mediator delete default response
func (o *MediatorDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MediatorDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /cluster/mediators/{uuid}][%d] mediator_delete default  %+v", o._statusCode, o.Payload)
}
func (o *MediatorDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MediatorDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}