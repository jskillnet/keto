// Code generated by go-swagger; DO NOT EDIT.

package engines

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

// NewListOryAccessControlPoliciesParams creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized.
func NewListOryAccessControlPoliciesParams() *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOryAccessControlPoliciesParamsWithTimeout creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOryAccessControlPoliciesParamsWithTimeout(timeout time.Duration) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		timeout: timeout,
	}
}

// NewListOryAccessControlPoliciesParamsWithContext creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOryAccessControlPoliciesParamsWithContext(ctx context.Context) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{

		Context: ctx,
	}
}

// NewListOryAccessControlPoliciesParamsWithHTTPClient creates a new ListOryAccessControlPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOryAccessControlPoliciesParamsWithHTTPClient(client *http.Client) *ListOryAccessControlPoliciesParams {
	var ()
	return &ListOryAccessControlPoliciesParams{
		HTTPClient: client,
	}
}

/*ListOryAccessControlPoliciesParams contains all the parameters to send to the API endpoint
for the list ory access control policies operation typically these are written to a http.Request
*/
type ListOryAccessControlPoliciesParams struct {

	/*Flavor
	  The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact"

	*/
	Flavor string
	/*Limit
	  The maximum amount of policies returned.

	*/
	Limit *int64
	/*Offset
	  The offset from where to start looking.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithTimeout(timeout time.Duration) *ListOryAccessControlPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithContext(ctx context.Context) *ListOryAccessControlPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithHTTPClient(client *http.Client) *ListOryAccessControlPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlavor adds the flavor to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithFlavor(flavor string) *ListOryAccessControlPoliciesParams {
	o.SetFlavor(flavor)
	return o
}

// SetFlavor adds the flavor to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetFlavor(flavor string) {
	o.Flavor = flavor
}

// WithLimit adds the limit to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithLimit(limit *int64) *ListOryAccessControlPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) WithOffset(offset *int64) *ListOryAccessControlPoliciesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list ory access control policies params
func (o *ListOryAccessControlPoliciesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListOryAccessControlPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flavor
	if err := r.SetPathParam("flavor", o.Flavor); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
