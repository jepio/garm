// Code generated by go-swagger; DO NOT EDIT.

package enterprises

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
)

// DeleteEnterprisePoolReader is a Reader for the DeleteEnterprisePool structure.
type DeleteEnterprisePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnterprisePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteEnterprisePoolDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteEnterprisePoolDefault creates a DeleteEnterprisePoolDefault with default headers values
func NewDeleteEnterprisePoolDefault(code int) *DeleteEnterprisePoolDefault {
	return &DeleteEnterprisePoolDefault{
		_statusCode: code,
	}
}

/*
DeleteEnterprisePoolDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type DeleteEnterprisePoolDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this delete enterprise pool default response has a 2xx status code
func (o *DeleteEnterprisePoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete enterprise pool default response has a 3xx status code
func (o *DeleteEnterprisePoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete enterprise pool default response has a 4xx status code
func (o *DeleteEnterprisePoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete enterprise pool default response has a 5xx status code
func (o *DeleteEnterprisePoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete enterprise pool default response a status code equal to that given
func (o *DeleteEnterprisePoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete enterprise pool default response
func (o *DeleteEnterprisePoolDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEnterprisePoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /enterprises/{enterpriseID}/pools/{poolID}][%d] DeleteEnterprisePool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEnterprisePoolDefault) String() string {
	return fmt.Sprintf("[DELETE /enterprises/{enterpriseID}/pools/{poolID}][%d] DeleteEnterprisePool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEnterprisePoolDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *DeleteEnterprisePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
