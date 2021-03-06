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

// GetLogDrainsLogDrainIDContainersReader is a Reader for the GetLogDrainsLogDrainIDContainers structure.
type GetLogDrainsLogDrainIDContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogDrainsLogDrainIDContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogDrainsLogDrainIDContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLogDrainsLogDrainIDContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogDrainsLogDrainIDContainersOK creates a GetLogDrainsLogDrainIDContainersOK with default headers values
func NewGetLogDrainsLogDrainIDContainersOK() *GetLogDrainsLogDrainIDContainersOK {
	return &GetLogDrainsLogDrainIDContainersOK{}
}

/*GetLogDrainsLogDrainIDContainersOK handles this case with default header values.

successful
*/
type GetLogDrainsLogDrainIDContainersOK struct {
	Payload *models.InlineResponse2008
}

func (o *GetLogDrainsLogDrainIDContainersOK) Error() string {
	return fmt.Sprintf("[GET /log_drains/{log_drain_id}/containers][%d] getLogDrainsLogDrainIdContainersOK  %+v", 200, o.Payload)
}

func (o *GetLogDrainsLogDrainIDContainersOK) GetPayload() *models.InlineResponse2008 {
	return o.Payload
}

func (o *GetLogDrainsLogDrainIDContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2008)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogDrainsLogDrainIDContainersDefault creates a GetLogDrainsLogDrainIDContainersDefault with default headers values
func NewGetLogDrainsLogDrainIDContainersDefault(code int) *GetLogDrainsLogDrainIDContainersDefault {
	return &GetLogDrainsLogDrainIDContainersDefault{
		_statusCode: code,
	}
}

/*GetLogDrainsLogDrainIDContainersDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetLogDrainsLogDrainIDContainersDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get log drains log drain ID containers default response
func (o *GetLogDrainsLogDrainIDContainersDefault) Code() int {
	return o._statusCode
}

func (o *GetLogDrainsLogDrainIDContainersDefault) Error() string {
	return fmt.Sprintf("[GET /log_drains/{log_drain_id}/containers][%d] GetLogDrainsLogDrainIDContainers default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogDrainsLogDrainIDContainersDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetLogDrainsLogDrainIDContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
