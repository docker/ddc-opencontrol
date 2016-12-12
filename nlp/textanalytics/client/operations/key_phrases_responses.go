package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/docker/ddc-opencontrol/nlp/textanalytics/models"
)

// KeyPhrasesReader is a Reader for the KeyPhrases structure.
type KeyPhrasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyPhrasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewKeyPhrasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewKeyPhrasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewKeyPhrasesOK creates a KeyPhrasesOK with default headers values
func NewKeyPhrasesOK() *KeyPhrasesOK {
	return &KeyPhrasesOK{}
}

/*KeyPhrasesOK handles this case with default header values.

OK
*/
type KeyPhrasesOK struct {
	Payload *models.KeyPhraseBatchResultV2
}

func (o *KeyPhrasesOK) Error() string {
	return fmt.Sprintf("[POST /keyPhrases][%d] keyPhrasesOK  %+v", 200, o.Payload)
}

func (o *KeyPhrasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyPhraseBatchResultV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyPhrasesBadRequest creates a KeyPhrasesBadRequest with default headers values
func NewKeyPhrasesBadRequest() *KeyPhrasesBadRequest {
	return &KeyPhrasesBadRequest{}
}

/*KeyPhrasesBadRequest handles this case with default header values.

BadRequest
*/
type KeyPhrasesBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *KeyPhrasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /keyPhrases][%d] keyPhrasesBadRequest  %+v", 400, o.Payload)
}

func (o *KeyPhrasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}