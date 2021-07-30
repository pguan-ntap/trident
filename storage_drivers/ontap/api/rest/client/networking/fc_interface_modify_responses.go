// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcInterfaceModifyReader is a Reader for the FcInterfaceModify structure.
type FcInterfaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcInterfaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcInterfaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcInterfaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcInterfaceModifyOK creates a FcInterfaceModifyOK with default headers values
func NewFcInterfaceModifyOK() *FcInterfaceModifyOK {
	return &FcInterfaceModifyOK{}
}

/* FcInterfaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type FcInterfaceModifyOK struct {
}

func (o *FcInterfaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/fc/interfaces/{uuid}][%d] fcInterfaceModifyOK ", 200)
}

func (o *FcInterfaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFcInterfaceModifyDefault creates a FcInterfaceModifyDefault with default headers values
func NewFcInterfaceModifyDefault(code int) *FcInterfaceModifyDefault {
	return &FcInterfaceModifyDefault{
		_statusCode: code,
	}
}

/* FcInterfaceModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1966140 | An interface with the same name already exists. |
| 1966217 | The specified port is not valid on the node provided. |
| 1966238 | The node or port of an active SAN data interface cannot be changed. |
| 1966702 | The destination node is not healthy. |
| 5374579 | The SAN Kernel Agent on the node is unavailable. |
| 5374870 | A partial failure occurred; renaming the interface failed. Correct the error and resubmit the request. |
| 5374871 | The Fibre Channel port identified by the specified UUID does not refer to the same port as that identified by the specified node name and/or port name. |
| 5374872 | If either `location.port.node.name` or `location.port.name` is supplied, both properties must be supplied. |
| 72089674 | You cannot move a Fibre Channel interface configured for the NVMe over FC data protocol. |

*/
type FcInterfaceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc interface modify default response
func (o *FcInterfaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *FcInterfaceModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/fc/interfaces/{uuid}][%d] fc_interface_modify default  %+v", o._statusCode, o.Payload)
}
func (o *FcInterfaceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcInterfaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}