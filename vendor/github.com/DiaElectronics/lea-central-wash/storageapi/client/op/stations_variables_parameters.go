// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStationsVariablesParams creates a new StationsVariablesParams object
// with the default values initialized.
func NewStationsVariablesParams() *StationsVariablesParams {

	return &StationsVariablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStationsVariablesParamsWithTimeout creates a new StationsVariablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStationsVariablesParamsWithTimeout(timeout time.Duration) *StationsVariablesParams {

	return &StationsVariablesParams{

		timeout: timeout,
	}
}

// NewStationsVariablesParamsWithContext creates a new StationsVariablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStationsVariablesParamsWithContext(ctx context.Context) *StationsVariablesParams {

	return &StationsVariablesParams{

		Context: ctx,
	}
}

// NewStationsVariablesParamsWithHTTPClient creates a new StationsVariablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStationsVariablesParamsWithHTTPClient(client *http.Client) *StationsVariablesParams {

	return &StationsVariablesParams{
		HTTPClient: client,
	}
}

/*StationsVariablesParams contains all the parameters to send to the API endpoint
for the stations variables operation typically these are written to a http.Request
*/
type StationsVariablesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stations variables params
func (o *StationsVariablesParams) WithTimeout(timeout time.Duration) *StationsVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stations variables params
func (o *StationsVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stations variables params
func (o *StationsVariablesParams) WithContext(ctx context.Context) *StationsVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stations variables params
func (o *StationsVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stations variables params
func (o *StationsVariablesParams) WithHTTPClient(client *http.Client) *StationsVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stations variables params
func (o *StationsVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StationsVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
