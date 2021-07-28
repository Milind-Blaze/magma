// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutWifiNetworkIDGatewaysGatewayIDReader is a Reader for the PutWifiNetworkIDGatewaysGatewayID structure.
type PutWifiNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDNoContent creates a PutWifiNetworkIDGatewaysGatewayIDNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDNoContent() *PutWifiNetworkIDGatewaysGatewayIDNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDNoContent{}
}

/*PutWifiNetworkIDGatewaysGatewayIDNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}][%d] putWifiNetworkIdGatewaysGatewayIdNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDDefault creates a PutWifiNetworkIDGatewaysGatewayIDDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID default response
func (o *PutWifiNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}][%d] PutWifiNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}