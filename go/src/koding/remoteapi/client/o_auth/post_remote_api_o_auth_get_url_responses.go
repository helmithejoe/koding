package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIOAuthGetURLReader is a Reader for the PostRemoteAPIOAuthGetURL structure.
type PostRemoteAPIOAuthGetURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIOAuthGetURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIOAuthGetURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIOAuthGetURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIOAuthGetURLOK creates a PostRemoteAPIOAuthGetURLOK with default headers values
func NewPostRemoteAPIOAuthGetURLOK() *PostRemoteAPIOAuthGetURLOK {
	return &PostRemoteAPIOAuthGetURLOK{}
}

/*PostRemoteAPIOAuthGetURLOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIOAuthGetURLOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIOAuthGetURLOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/OAuth.getUrl][%d] postRemoteApiOAuthGetUrlOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIOAuthGetURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIOAuthGetURLUnauthorized creates a PostRemoteAPIOAuthGetURLUnauthorized with default headers values
func NewPostRemoteAPIOAuthGetURLUnauthorized() *PostRemoteAPIOAuthGetURLUnauthorized {
	return &PostRemoteAPIOAuthGetURLUnauthorized{}
}

/*PostRemoteAPIOAuthGetURLUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIOAuthGetURLUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIOAuthGetURLUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/OAuth.getUrl][%d] postRemoteApiOAuthGetUrlUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIOAuthGetURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
