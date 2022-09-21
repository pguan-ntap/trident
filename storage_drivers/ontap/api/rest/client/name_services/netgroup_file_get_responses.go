// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetgroupFileGetReader is a Reader for the NetgroupFileGet structure.
type NetgroupFileGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupFileGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupFileGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupFileGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupFileGetOK creates a NetgroupFileGetOK with default headers values
func NewNetgroupFileGetOK() *NetgroupFileGetOK {
	return &NetgroupFileGetOK{}
}

/* NetgroupFileGetOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupFileGetOK struct {
	Payload *models.NetgroupFile
}

func (o *NetgroupFileGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroupFileGetOK  %+v", 200, o.Payload)
}
func (o *NetgroupFileGetOK) GetPayload() *models.NetgroupFile {
	return o.Payload
}

func (o *NetgroupFileGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetgroupFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetgroupFileGetDefault creates a NetgroupFileGetDefault with default headers values
func NewNetgroupFileGetDefault(code int) *NetgroupFileGetDefault {
	return &NetgroupFileGetDefault{
		_statusCode: code,
	}
}

/* NetgroupFileGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetgroupFileGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the netgroup file get default response
func (o *NetgroupFileGetDefault) Code() int {
	return o._statusCode
}

func (o *NetgroupFileGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/netgroup-files/{svm.uuid}][%d] netgroup_file_get default  %+v", o._statusCode, o.Payload)
}
func (o *NetgroupFileGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupFileGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}