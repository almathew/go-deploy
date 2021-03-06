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

// GetAppsAppIDVhostsReader is a Reader for the GetAppsAppIDVhosts structure.
type GetAppsAppIDVhostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppIDVhostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppsAppIDVhostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppsAppIDVhostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppsAppIDVhostsOK creates a GetAppsAppIDVhostsOK with default headers values
func NewGetAppsAppIDVhostsOK() *GetAppsAppIDVhostsOK {
	return &GetAppsAppIDVhostsOK{}
}

/*GetAppsAppIDVhostsOK handles this case with default header values.

successful
*/
type GetAppsAppIDVhostsOK struct {
	Payload *models.InlineResponse20038
}

func (o *GetAppsAppIDVhostsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/vhosts][%d] getAppsAppIdVhostsOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppIDVhostsOK) GetPayload() *models.InlineResponse20038 {
	return o.Payload
}

func (o *GetAppsAppIDVhostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20038)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppIDVhostsDefault creates a GetAppsAppIDVhostsDefault with default headers values
func NewGetAppsAppIDVhostsDefault(code int) *GetAppsAppIDVhostsDefault {
	return &GetAppsAppIDVhostsDefault{
		_statusCode: code,
	}
}

/*GetAppsAppIDVhostsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAppsAppIDVhostsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get apps app ID vhosts default response
func (o *GetAppsAppIDVhostsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppsAppIDVhostsDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/vhosts][%d] GetAppsAppIDVhosts default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppsAppIDVhostsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAppsAppIDVhostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
