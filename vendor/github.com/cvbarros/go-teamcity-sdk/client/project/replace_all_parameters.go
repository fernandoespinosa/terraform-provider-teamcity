// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewReplaceAllParams creates a new ReplaceAllParams object
// with the default values initialized.
func NewReplaceAllParams() *ReplaceAllParams {
	var ()
	return &ReplaceAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAllParamsWithTimeout creates a new ReplaceAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAllParamsWithTimeout(timeout time.Duration) *ReplaceAllParams {
	var ()
	return &ReplaceAllParams{

		timeout: timeout,
	}
}

// NewReplaceAllParamsWithContext creates a new ReplaceAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAllParamsWithContext(ctx context.Context) *ReplaceAllParams {
	var ()
	return &ReplaceAllParams{

		Context: ctx,
	}
}

// NewReplaceAllParamsWithHTTPClient creates a new ReplaceAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAllParamsWithHTTPClient(client *http.Client) *ReplaceAllParams {
	var ()
	return &ReplaceAllParams{
		HTTPClient: client,
	}
}

/*ReplaceAllParams contains all the parameters to send to the API endpoint
for the replace all operation typically these are written to a http.Request
*/
type ReplaceAllParams struct {

	/*Body*/
	Body *models.ProjectFeatures
	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace all params
func (o *ReplaceAllParams) WithTimeout(timeout time.Duration) *ReplaceAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace all params
func (o *ReplaceAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace all params
func (o *ReplaceAllParams) WithContext(ctx context.Context) *ReplaceAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace all params
func (o *ReplaceAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace all params
func (o *ReplaceAllParams) WithHTTPClient(client *http.Client) *ReplaceAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace all params
func (o *ReplaceAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace all params
func (o *ReplaceAllParams) WithBody(body *models.ProjectFeatures) *ReplaceAllParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace all params
func (o *ReplaceAllParams) SetBody(body *models.ProjectFeatures) {
	o.Body = body
}

// WithFields adds the fields to the replace all params
func (o *ReplaceAllParams) WithFields(fields *string) *ReplaceAllParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace all params
func (o *ReplaceAllParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the replace all params
func (o *ReplaceAllParams) WithProjectLocator(projectLocator string) *ReplaceAllParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the replace all params
func (o *ReplaceAllParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}