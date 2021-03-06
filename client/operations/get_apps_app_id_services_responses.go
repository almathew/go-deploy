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

// GetAppsAppIDServicesReader is a Reader for the GetAppsAppIDServices structure.
type GetAppsAppIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppsAppIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppsAppIDServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppsAppIDServicesOK creates a GetAppsAppIDServicesOK with default headers values
func NewGetAppsAppIDServicesOK() *GetAppsAppIDServicesOK {
	return &GetAppsAppIDServicesOK{}
}

/*GetAppsAppIDServicesOK handles this case with default header values.

successful
*/
type GetAppsAppIDServicesOK struct {
	Payload *models.InlineResponse20033
}

func (o *GetAppsAppIDServicesOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/services][%d] getAppsAppIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppIDServicesOK) GetPayload() *models.InlineResponse20033 {
	return o.Payload
}

func (o *GetAppsAppIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20033)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppIDServicesDefault creates a GetAppsAppIDServicesDefault with default headers values
func NewGetAppsAppIDServicesDefault(code int) *GetAppsAppIDServicesDefault {
	return &GetAppsAppIDServicesDefault{
		_statusCode: code,
	}
}

/*GetAppsAppIDServicesDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAppsAppIDServicesDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get apps app ID services default response
func (o *GetAppsAppIDServicesDefault) Code() int {
	return o._statusCode
}

func (o *GetAppsAppIDServicesDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/services][%d] GetAppsAppIDServices default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppsAppIDServicesDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAppsAppIDServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
