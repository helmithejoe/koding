package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialChannelUnpinMessageReader is a Reader for the PostRemoteAPISocialChannelUnpinMessage structure.
type PostRemoteAPISocialChannelUnpinMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelUnpinMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelUnpinMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelUnpinMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelUnpinMessageOK creates a PostRemoteAPISocialChannelUnpinMessageOK with default headers values
func NewPostRemoteAPISocialChannelUnpinMessageOK() *PostRemoteAPISocialChannelUnpinMessageOK {
	return &PostRemoteAPISocialChannelUnpinMessageOK{}
}

/*PostRemoteAPISocialChannelUnpinMessageOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPISocialChannelUnpinMessageOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelUnpinMessageOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.unpinMessage][%d] postRemoteApiSocialChannelUnpinMessageOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelUnpinMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelUnpinMessageUnauthorized creates a PostRemoteAPISocialChannelUnpinMessageUnauthorized with default headers values
func NewPostRemoteAPISocialChannelUnpinMessageUnauthorized() *PostRemoteAPISocialChannelUnpinMessageUnauthorized {
	return &PostRemoteAPISocialChannelUnpinMessageUnauthorized{}
}

/*PostRemoteAPISocialChannelUnpinMessageUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelUnpinMessageUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelUnpinMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.unpinMessage][%d] postRemoteApiSocialChannelUnpinMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelUnpinMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
