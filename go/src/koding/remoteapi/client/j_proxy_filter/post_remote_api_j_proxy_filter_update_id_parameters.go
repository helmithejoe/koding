package j_proxy_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJProxyFilterUpdateIDParams creates a new PostRemoteAPIJProxyFilterUpdateIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProxyFilterUpdateIDParams() *PostRemoteAPIJProxyFilterUpdateIDParams {
	var ()
	return &PostRemoteAPIJProxyFilterUpdateIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProxyFilterUpdateIDParamsWithTimeout creates a new PostRemoteAPIJProxyFilterUpdateIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProxyFilterUpdateIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProxyFilterUpdateIDParams {
	var ()
	return &PostRemoteAPIJProxyFilterUpdateIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProxyFilterUpdateIDParamsWithContext creates a new PostRemoteAPIJProxyFilterUpdateIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProxyFilterUpdateIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProxyFilterUpdateIDParams {
	var ()
	return &PostRemoteAPIJProxyFilterUpdateIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProxyFilterUpdateIDParams contains all the parameters to send to the API endpoint
for the post remote API j proxy filter update ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProxyFilterUpdateIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProxyFilterUpdateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProxyFilterUpdateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) WithID(id string) *PostRemoteAPIJProxyFilterUpdateIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j proxy filter update ID params
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProxyFilterUpdateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
