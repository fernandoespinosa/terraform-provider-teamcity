// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewGetFeatureParametersParams creates a new GetFeatureParametersParams object
// with the default values initialized.
func NewGetFeatureParametersParams() *GetFeatureParametersParams {
	var ()
	return &GetFeatureParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeatureParametersParamsWithTimeout creates a new GetFeatureParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFeatureParametersParamsWithTimeout(timeout time.Duration) *GetFeatureParametersParams {
	var ()
	return &GetFeatureParametersParams{

		timeout: timeout,
	}
}

// NewGetFeatureParametersParamsWithContext creates a new GetFeatureParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFeatureParametersParamsWithContext(ctx context.Context) *GetFeatureParametersParams {
	var ()
	return &GetFeatureParametersParams{

		Context: ctx,
	}
}

// NewGetFeatureParametersParamsWithHTTPClient creates a new GetFeatureParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFeatureParametersParamsWithHTTPClient(client *http.Client) *GetFeatureParametersParams {
	var ()
	return &GetFeatureParametersParams{
		HTTPClient: client,
	}
}

/*GetFeatureParametersParams contains all the parameters to send to the API endpoint
for the get feature parameters operation typically these are written to a http.Request
*/
type GetFeatureParametersParams struct {

	/*BtLocator*/
	BtLocator string
	/*FeatureID*/
	FeatureID string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feature parameters params
func (o *GetFeatureParametersParams) WithTimeout(timeout time.Duration) *GetFeatureParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feature parameters params
func (o *GetFeatureParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feature parameters params
func (o *GetFeatureParametersParams) WithContext(ctx context.Context) *GetFeatureParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feature parameters params
func (o *GetFeatureParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feature parameters params
func (o *GetFeatureParametersParams) WithHTTPClient(client *http.Client) *GetFeatureParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feature parameters params
func (o *GetFeatureParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get feature parameters params
func (o *GetFeatureParametersParams) WithBtLocator(btLocator string) *GetFeatureParametersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get feature parameters params
func (o *GetFeatureParametersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFeatureID adds the featureID to the get feature parameters params
func (o *GetFeatureParametersParams) WithFeatureID(featureID string) *GetFeatureParametersParams {
	o.SetFeatureID(featureID)
	return o
}

// SetFeatureID adds the featureId to the get feature parameters params
func (o *GetFeatureParametersParams) SetFeatureID(featureID string) {
	o.FeatureID = featureID
}

// WithFields adds the fields to the get feature parameters params
func (o *GetFeatureParametersParams) WithFields(fields *string) *GetFeatureParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get feature parameters params
func (o *GetFeatureParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeatureParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param featureId
	if err := r.SetPathParam("featureId", o.FeatureID); err != nil {
		return err
	}

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