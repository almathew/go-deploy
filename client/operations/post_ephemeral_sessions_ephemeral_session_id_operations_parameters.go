// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPostEphemeralSessionsEphemeralSessionIDOperationsParams creates a new PostEphemeralSessionsEphemeralSessionIDOperationsParams object
// with the default values initialized.
func NewPostEphemeralSessionsEphemeralSessionIDOperationsParams() *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	var ()
	return &PostEphemeralSessionsEphemeralSessionIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithTimeout creates a new PostEphemeralSessionsEphemeralSessionIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithTimeout(timeout time.Duration) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	var ()
	return &PostEphemeralSessionsEphemeralSessionIDOperationsParams{

		timeout: timeout,
	}
}

// NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithContext creates a new PostEphemeralSessionsEphemeralSessionIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithContext(ctx context.Context) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	var ()
	return &PostEphemeralSessionsEphemeralSessionIDOperationsParams{

		Context: ctx,
	}
}

// NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithHTTPClient creates a new PostEphemeralSessionsEphemeralSessionIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEphemeralSessionsEphemeralSessionIDOperationsParamsWithHTTPClient(client *http.Client) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	var ()
	return &PostEphemeralSessionsEphemeralSessionIDOperationsParams{
		HTTPClient: client,
	}
}

/*PostEphemeralSessionsEphemeralSessionIDOperationsParams contains all the parameters to send to the API endpoint
for the post ephemeral sessions ephemeral session ID operations operation typically these are written to a http.Request
*/
type PostEphemeralSessionsEphemeralSessionIDOperationsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest22
	/*EphemeralSessionID
	  ephemeral_session_id

	*/
	EphemeralSessionID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WithTimeout(timeout time.Duration) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WithContext(ctx context.Context) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WithHTTPClient(client *http.Client) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WithAppRequest(appRequest *models.AppRequest22) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) SetAppRequest(appRequest *models.AppRequest22) {
	o.AppRequest = appRequest
}

// WithEphemeralSessionID adds the ephemeralSessionID to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WithEphemeralSessionID(ephemeralSessionID int64) *PostEphemeralSessionsEphemeralSessionIDOperationsParams {
	o.SetEphemeralSessionID(ephemeralSessionID)
	return o
}

// SetEphemeralSessionID adds the ephemeralSessionId to the post ephemeral sessions ephemeral session ID operations params
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) SetEphemeralSessionID(ephemeralSessionID int64) {
	o.EphemeralSessionID = ephemeralSessionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostEphemeralSessionsEphemeralSessionIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param ephemeral_session_id
	if err := r.SetPathParam("ephemeral_session_id", swag.FormatInt64(o.EphemeralSessionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
