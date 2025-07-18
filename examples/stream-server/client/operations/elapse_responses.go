// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	stderrors "errors"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ElapseReader is a Reader for the Elapse structure.
type ElapseReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ElapseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
	switch response.Code() {
	case 200:
		result := NewElapseOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewElapseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /elapse/{length}] elapse", response, response.Code())
	}
}

// NewElapseOK creates a ElapseOK with default headers values
func NewElapseOK(writer io.Writer) *ElapseOK {
	return &ElapseOK{

		Payload: writer,
	}
}

/*
ElapseOK describes a response with status code 200, with default header values.

Secondly update on remaining time
*/
type ElapseOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this elapse o k response has a 2xx status code
func (o *ElapseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this elapse o k response has a 3xx status code
func (o *ElapseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this elapse o k response has a 4xx status code
func (o *ElapseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this elapse o k response has a 5xx status code
func (o *ElapseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this elapse o k response a status code equal to that given
func (o *ElapseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the elapse o k response
func (o *ElapseOK) Code() int {
	return 200
}

func (o *ElapseOK) Error() string {
	return fmt.Sprintf("[GET /elapse/{length}][%d] elapseOK", 200)
}

func (o *ElapseOK) String() string {
	return fmt.Sprintf("[GET /elapse/{length}][%d] elapseOK", 200)
}

func (o *ElapseOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ElapseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && !stderrors.Is(err, io.EOF) {
		return err
	}

	return nil
}

// NewElapseForbidden creates a ElapseForbidden with default headers values
func NewElapseForbidden() *ElapseForbidden {
	return &ElapseForbidden{}
}

/*
ElapseForbidden describes a response with status code 403, with default header values.

Contrived - thrown when length of 11 is chosen
*/
type ElapseForbidden struct {
}

// IsSuccess returns true when this elapse forbidden response has a 2xx status code
func (o *ElapseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this elapse forbidden response has a 3xx status code
func (o *ElapseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this elapse forbidden response has a 4xx status code
func (o *ElapseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this elapse forbidden response has a 5xx status code
func (o *ElapseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this elapse forbidden response a status code equal to that given
func (o *ElapseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the elapse forbidden response
func (o *ElapseForbidden) Code() int {
	return 403
}

func (o *ElapseForbidden) Error() string {
	return fmt.Sprintf("[GET /elapse/{length}][%d] elapseForbidden", 403)
}

func (o *ElapseForbidden) String() string {
	return fmt.Sprintf("[GET /elapse/{length}][%d] elapseForbidden", 403)
}

func (o *ElapseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
