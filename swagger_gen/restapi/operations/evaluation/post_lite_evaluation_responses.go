// Code generated by go-swagger; DO NOT EDIT.

package evaluation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/checkr/flagr/swagger_gen/models"
)

// PostLiteEvaluationOKCode is the HTTP code returned for type PostLiteEvaluationOK
const PostLiteEvaluationOKCode int = 200

/*PostLiteEvaluationOK lite evaluation result

swagger:response postLiteEvaluationOK
*/
type PostLiteEvaluationOK struct {

	/*
	  In: Body
	*/
	Payload *models.LiteEvalResult `json:"body,omitempty"`
}

// NewPostLiteEvaluationOK creates PostLiteEvaluationOK with default headers values
func NewPostLiteEvaluationOK() *PostLiteEvaluationOK {

	return &PostLiteEvaluationOK{}
}

// WithPayload adds the payload to the post lite evaluation o k response
func (o *PostLiteEvaluationOK) WithPayload(payload *models.LiteEvalResult) *PostLiteEvaluationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post lite evaluation o k response
func (o *PostLiteEvaluationOK) SetPayload(payload *models.LiteEvalResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLiteEvaluationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostLiteEvaluationDefault generic error response

swagger:response postLiteEvaluationDefault
*/
type PostLiteEvaluationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostLiteEvaluationDefault creates PostLiteEvaluationDefault with default headers values
func NewPostLiteEvaluationDefault(code int) *PostLiteEvaluationDefault {
	if code <= 0 {
		code = 500
	}

	return &PostLiteEvaluationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post lite evaluation default response
func (o *PostLiteEvaluationDefault) WithStatusCode(code int) *PostLiteEvaluationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post lite evaluation default response
func (o *PostLiteEvaluationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post lite evaluation default response
func (o *PostLiteEvaluationDefault) WithPayload(payload *models.Error) *PostLiteEvaluationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post lite evaluation default response
func (o *PostLiteEvaluationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLiteEvaluationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}