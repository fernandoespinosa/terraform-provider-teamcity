// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLicenseKeysParams creates a new GetLicenseKeysParams object
// with the default values initialized.
func NewGetLicenseKeysParams() *GetLicenseKeysParams {
	var ()
	return &GetLicenseKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseKeysParamsWithTimeout creates a new GetLicenseKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLicenseKeysParamsWithTimeout(timeout time.Duration) *GetLicenseKeysParams {
	var ()
	return &GetLicenseKeysParams{

		timeout: timeout,
	}
}

// NewGetLicenseKeysParamsWithContext creates a new GetLicenseKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLicenseKeysParamsWithContext(ctx context.Context) *GetLicenseKeysParams {
	var ()
	return &GetLicenseKeysParams{

		Context: ctx,
	}
}

// NewGetLicenseKeysParamsWithHTTPClient creates a new GetLicenseKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLicenseKeysParamsWithHTTPClient(client *http.Client) *GetLicenseKeysParams {
	var ()
	return &GetLicenseKeysParams{
		HTTPClient: client,
	}
}

/*GetLicenseKeysParams contains all the parameters to send to the API endpoint
for the get license keys operation typically these are written to a http.Request
*/
type GetLicenseKeysParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get license keys params
func (o *GetLicenseKeysParams) WithTimeout(timeout time.Duration) *GetLicenseKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license keys params
func (o *GetLicenseKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license keys params
func (o *GetLicenseKeysParams) WithContext(ctx context.Context) *GetLicenseKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license keys params
func (o *GetLicenseKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license keys params
func (o *GetLicenseKeysParams) WithHTTPClient(client *http.Client) *GetLicenseKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license keys params
func (o *GetLicenseKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get license keys params
func (o *GetLicenseKeysParams) WithFields(fields *string) *GetLicenseKeysParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get license keys params
func (o *GetLicenseKeysParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}