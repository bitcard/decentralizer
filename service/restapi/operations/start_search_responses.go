package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iain17/dht-hello/service/models"
)

/*StartSearchOK Expected response to a valid request

swagger:response startSearchOK
*/
type StartSearchOK struct {
}

// NewStartSearchOK creates StartSearchOK with default headers values
func NewStartSearchOK() *StartSearchOK {
	return &StartSearchOK{}
}

// WriteResponse to the client
func (o *StartSearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*StartSearchDefault unexpected error

swagger:response startSearchDefault
*/
type StartSearchDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewStartSearchDefault creates StartSearchDefault with default headers values
func NewStartSearchDefault(code int) *StartSearchDefault {
	if code <= 0 {
		code = 500
	}

	return &StartSearchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the start search default response
func (o *StartSearchDefault) WithStatusCode(code int) *StartSearchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the start search default response
func (o *StartSearchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the start search default response
func (o *StartSearchDefault) WithPayload(payload *models.Error) *StartSearchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start search default response
func (o *StartSearchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartSearchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
