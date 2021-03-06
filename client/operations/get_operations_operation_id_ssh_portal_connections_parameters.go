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
)

// NewGetOperationsOperationIDSSHPortalConnectionsParams creates a new GetOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized.
func NewGetOperationsOperationIDSSHPortalConnectionsParams() *GetOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &GetOperationsOperationIDSSHPortalConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperationsOperationIDSSHPortalConnectionsParamsWithTimeout creates a new GetOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperationsOperationIDSSHPortalConnectionsParamsWithTimeout(timeout time.Duration) *GetOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &GetOperationsOperationIDSSHPortalConnectionsParams{

		timeout: timeout,
	}
}

// NewGetOperationsOperationIDSSHPortalConnectionsParamsWithContext creates a new GetOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperationsOperationIDSSHPortalConnectionsParamsWithContext(ctx context.Context) *GetOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &GetOperationsOperationIDSSHPortalConnectionsParams{

		Context: ctx,
	}
}

// NewGetOperationsOperationIDSSHPortalConnectionsParamsWithHTTPClient creates a new GetOperationsOperationIDSSHPortalConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperationsOperationIDSSHPortalConnectionsParamsWithHTTPClient(client *http.Client) *GetOperationsOperationIDSSHPortalConnectionsParams {
	var ()
	return &GetOperationsOperationIDSSHPortalConnectionsParams{
		HTTPClient: client,
	}
}

/*GetOperationsOperationIDSSHPortalConnectionsParams contains all the parameters to send to the API endpoint
for the get operations operation ID SSH portal connections operation typically these are written to a http.Request
*/
type GetOperationsOperationIDSSHPortalConnectionsParams struct {

	/*OperationID
	  operation_id

	*/
	OperationID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WithTimeout(timeout time.Duration) *GetOperationsOperationIDSSHPortalConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WithContext(ctx context.Context) *GetOperationsOperationIDSSHPortalConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WithHTTPClient(client *http.Client) *GetOperationsOperationIDSSHPortalConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperationID adds the operationID to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WithOperationID(operationID int64) *GetOperationsOperationIDSSHPortalConnectionsParams {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) SetOperationID(operationID int64) {
	o.OperationID = operationID
}

// WithPage adds the page to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WithPage(page *int64) *GetOperationsOperationIDSSHPortalConnectionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get operations operation ID SSH portal connections params
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperationsOperationIDSSHPortalConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operation_id
	if err := r.SetPathParam("operation_id", swag.FormatInt64(o.OperationID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
