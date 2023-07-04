// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPcloudVpnconnectionsNetworksGetParams creates a new PcloudVpnconnectionsNetworksGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVpnconnectionsNetworksGetParams() *PcloudVpnconnectionsNetworksGetParams {
	return &PcloudVpnconnectionsNetworksGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVpnconnectionsNetworksGetParamsWithTimeout creates a new PcloudVpnconnectionsNetworksGetParams object
// with the ability to set a timeout on a request.
func NewPcloudVpnconnectionsNetworksGetParamsWithTimeout(timeout time.Duration) *PcloudVpnconnectionsNetworksGetParams {
	return &PcloudVpnconnectionsNetworksGetParams{
		timeout: timeout,
	}
}

// NewPcloudVpnconnectionsNetworksGetParamsWithContext creates a new PcloudVpnconnectionsNetworksGetParams object
// with the ability to set a context for a request.
func NewPcloudVpnconnectionsNetworksGetParamsWithContext(ctx context.Context) *PcloudVpnconnectionsNetworksGetParams {
	return &PcloudVpnconnectionsNetworksGetParams{
		Context: ctx,
	}
}

// NewPcloudVpnconnectionsNetworksGetParamsWithHTTPClient creates a new PcloudVpnconnectionsNetworksGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVpnconnectionsNetworksGetParamsWithHTTPClient(client *http.Client) *PcloudVpnconnectionsNetworksGetParams {
	return &PcloudVpnconnectionsNetworksGetParams{
		HTTPClient: client,
	}
}

/*
PcloudVpnconnectionsNetworksGetParams contains all the parameters to send to the API endpoint

	for the pcloud vpnconnections networks get operation.

	Typically these are written to a http.Request.
*/
type PcloudVpnconnectionsNetworksGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* VpnConnectionID.

	   ID of a VPN connection
	*/
	VpnConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud vpnconnections networks get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVpnconnectionsNetworksGetParams) WithDefaults() *PcloudVpnconnectionsNetworksGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud vpnconnections networks get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVpnconnectionsNetworksGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) WithTimeout(timeout time.Duration) *PcloudVpnconnectionsNetworksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) WithContext(ctx context.Context) *PcloudVpnconnectionsNetworksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) WithHTTPClient(client *http.Client) *PcloudVpnconnectionsNetworksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVpnconnectionsNetworksGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVpnConnectionID adds the vpnConnectionID to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) WithVpnConnectionID(vpnConnectionID string) *PcloudVpnconnectionsNetworksGetParams {
	o.SetVpnConnectionID(vpnConnectionID)
	return o
}

// SetVpnConnectionID adds the vpnConnectionId to the pcloud vpnconnections networks get params
func (o *PcloudVpnconnectionsNetworksGetParams) SetVpnConnectionID(vpnConnectionID string) {
	o.VpnConnectionID = vpnConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVpnconnectionsNetworksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param vpn_connection_id
	if err := r.SetPathParam("vpn_connection_id", o.VpnConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
