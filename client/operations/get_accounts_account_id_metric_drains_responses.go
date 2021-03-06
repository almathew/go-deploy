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

// GetAccountsAccountIDMetricDrainsReader is a Reader for the GetAccountsAccountIDMetricDrains structure.
type GetAccountsAccountIDMetricDrainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDMetricDrainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDMetricDrainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountsAccountIDMetricDrainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountsAccountIDMetricDrainsOK creates a GetAccountsAccountIDMetricDrainsOK with default headers values
func NewGetAccountsAccountIDMetricDrainsOK() *GetAccountsAccountIDMetricDrainsOK {
	return &GetAccountsAccountIDMetricDrainsOK{}
}

/*GetAccountsAccountIDMetricDrainsOK handles this case with default header values.

successful
*/
type GetAccountsAccountIDMetricDrainsOK struct {
	Payload *models.InlineResponse20027
}

func (o *GetAccountsAccountIDMetricDrainsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/metric_drains][%d] getAccountsAccountIdMetricDrainsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDMetricDrainsOK) GetPayload() *models.InlineResponse20027 {
	return o.Payload
}

func (o *GetAccountsAccountIDMetricDrainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20027)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDMetricDrainsDefault creates a GetAccountsAccountIDMetricDrainsDefault with default headers values
func NewGetAccountsAccountIDMetricDrainsDefault(code int) *GetAccountsAccountIDMetricDrainsDefault {
	return &GetAccountsAccountIDMetricDrainsDefault{
		_statusCode: code,
	}
}

/*GetAccountsAccountIDMetricDrainsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAccountsAccountIDMetricDrainsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get accounts account ID metric drains default response
func (o *GetAccountsAccountIDMetricDrainsDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountsAccountIDMetricDrainsDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/metric_drains][%d] GetAccountsAccountIDMetricDrains default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountsAccountIDMetricDrainsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAccountsAccountIDMetricDrainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
