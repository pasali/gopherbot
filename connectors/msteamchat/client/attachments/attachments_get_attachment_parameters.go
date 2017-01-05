package attachments

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

// NewAttachmentsGetAttachmentParams creates a new AttachmentsGetAttachmentParams object
// with the default values initialized.
func NewAttachmentsGetAttachmentParams() *AttachmentsGetAttachmentParams {
	var ()
	return &AttachmentsGetAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAttachmentsGetAttachmentParamsWithTimeout creates a new AttachmentsGetAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAttachmentsGetAttachmentParamsWithTimeout(timeout time.Duration) *AttachmentsGetAttachmentParams {
	var ()
	return &AttachmentsGetAttachmentParams{

		timeout: timeout,
	}
}

// NewAttachmentsGetAttachmentParamsWithContext creates a new AttachmentsGetAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewAttachmentsGetAttachmentParamsWithContext(ctx context.Context) *AttachmentsGetAttachmentParams {
	var ()
	return &AttachmentsGetAttachmentParams{

		Context: ctx,
	}
}

/*AttachmentsGetAttachmentParams contains all the parameters to send to the API endpoint
for the attachments get attachment operation typically these are written to a http.Request
*/
type AttachmentsGetAttachmentParams struct {

	/*AttachmentID
	  attachment id

	*/
	AttachmentID string
	/*ViewID
	  View id from attachmentInfo

	*/
	ViewID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) WithTimeout(timeout time.Duration) *AttachmentsGetAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) WithContext(ctx context.Context) *AttachmentsGetAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAttachmentID adds the attachmentID to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) WithAttachmentID(attachmentID string) *AttachmentsGetAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) SetAttachmentID(attachmentID string) {
	o.AttachmentID = attachmentID
}

// WithViewID adds the viewID to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) WithViewID(viewID string) *AttachmentsGetAttachmentParams {
	o.SetViewID(viewID)
	return o
}

// SetViewID adds the viewId to the attachments get attachment params
func (o *AttachmentsGetAttachmentParams) SetViewID(viewID string) {
	o.ViewID = viewID
}

// WriteToRequest writes these params to a swagger request
func (o *AttachmentsGetAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param attachmentId
	if err := r.SetPathParam("attachmentId", o.AttachmentID); err != nil {
		return err
	}

	// path param viewId
	if err := r.SetPathParam("viewId", o.ViewID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}