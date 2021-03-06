// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// GetLogDrainsLogDrainIDOperationsReader is a Reader for the GetLogDrainsLogDrainIDOperations structure.
type GetLogDrainsLogDrainIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogDrainsLogDrainIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogDrainsLogDrainIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLogDrainsLogDrainIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogDrainsLogDrainIDOperationsOK creates a GetLogDrainsLogDrainIDOperationsOK with default headers values
func NewGetLogDrainsLogDrainIDOperationsOK() *GetLogDrainsLogDrainIDOperationsOK {
	return &GetLogDrainsLogDrainIDOperationsOK{}
}

/*GetLogDrainsLogDrainIDOperationsOK handles this case with default header values.

successful
*/
type GetLogDrainsLogDrainIDOperationsOK struct {
	Payload *models.InlineResponse20029
}

func (o *GetLogDrainsLogDrainIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /log_drains/{log_drain_id}/operations][%d] getLogDrainsLogDrainIdOperationsOK  %+v", 200, o.Payload)
}

func (o *GetLogDrainsLogDrainIDOperationsOK) GetPayload() *models.InlineResponse20029 {
	return o.Payload
}

func (o *GetLogDrainsLogDrainIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20029)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogDrainsLogDrainIDOperationsDefault creates a GetLogDrainsLogDrainIDOperationsDefault with default headers values
func NewGetLogDrainsLogDrainIDOperationsDefault(code int) *GetLogDrainsLogDrainIDOperationsDefault {
	return &GetLogDrainsLogDrainIDOperationsDefault{
		_statusCode: code,
	}
}

/*GetLogDrainsLogDrainIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetLogDrainsLogDrainIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get log drains log drain ID operations default response
func (o *GetLogDrainsLogDrainIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetLogDrainsLogDrainIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /log_drains/{log_drain_id}/operations][%d] GetLogDrainsLogDrainIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogDrainsLogDrainIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetLogDrainsLogDrainIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
