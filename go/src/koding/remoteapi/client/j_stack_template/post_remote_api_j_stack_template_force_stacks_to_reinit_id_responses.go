package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJStackTemplateForceStacksToReinitIDReader is a Reader for the PostRemoteAPIJStackTemplateForceStacksToReinitID structure.
type PostRemoteAPIJStackTemplateForceStacksToReinitIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK creates a PostRemoteAPIJStackTemplateForceStacksToReinitIDOK with default headers values
func NewPostRemoteAPIJStackTemplateForceStacksToReinitIDOK() *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK {
	return &PostRemoteAPIJStackTemplateForceStacksToReinitIDOK{}
}

/*PostRemoteAPIJStackTemplateForceStacksToReinitIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJStackTemplateForceStacksToReinitIDOK struct {
	Payload *models.JStackTemplate
}

func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.forceStacksToReinit/{id}][%d] postRemoteApiJStackTemplateForceStacksToReinitIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJStackTemplateForceStacksToReinitIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JStackTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
