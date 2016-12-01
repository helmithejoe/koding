package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPISocialChannelFetchFollowedChannelCountParams creates a new PostRemoteAPISocialChannelFetchFollowedChannelCountParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelFetchFollowedChannelCountParams() *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchFollowedChannelCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelFetchFollowedChannelCountParamsWithTimeout creates a new PostRemoteAPISocialChannelFetchFollowedChannelCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelFetchFollowedChannelCountParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchFollowedChannelCountParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelFetchFollowedChannelCountParamsWithContext creates a new PostRemoteAPISocialChannelFetchFollowedChannelCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelFetchFollowedChannelCountParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchFollowedChannelCountParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelFetchFollowedChannelCountParams contains all the parameters to send to the API endpoint
for the post remote API social channel fetch followed channel count operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelFetchFollowedChannelCountParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelFetchFollowedChannelCountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel fetch followed channel count params
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelFetchFollowedChannelCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
