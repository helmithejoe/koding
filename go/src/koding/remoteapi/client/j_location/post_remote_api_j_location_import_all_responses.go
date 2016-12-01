package j_location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJLocationImportAllReader is a Reader for the PostRemoteAPIJLocationImportAll structure.
type PostRemoteAPIJLocationImportAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJLocationImportAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJLocationImportAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJLocationImportAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJLocationImportAllOK creates a PostRemoteAPIJLocationImportAllOK with default headers values
func NewPostRemoteAPIJLocationImportAllOK() *PostRemoteAPIJLocationImportAllOK {
	return &PostRemoteAPIJLocationImportAllOK{}
}

/*PostRemoteAPIJLocationImportAllOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJLocationImportAllOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJLocationImportAllOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.importAll][%d] postRemoteApiJLocationImportAllOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJLocationImportAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJLocationImportAllUnauthorized creates a PostRemoteAPIJLocationImportAllUnauthorized with default headers values
func NewPostRemoteAPIJLocationImportAllUnauthorized() *PostRemoteAPIJLocationImportAllUnauthorized {
	return &PostRemoteAPIJLocationImportAllUnauthorized{}
}

/*PostRemoteAPIJLocationImportAllUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJLocationImportAllUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJLocationImportAllUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JLocation.importAll][%d] postRemoteApiJLocationImportAllUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJLocationImportAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
