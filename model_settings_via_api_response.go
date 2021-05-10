/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.17
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SettingsViaApiResponse The Response for Settings Flows via API
type SettingsViaApiResponse struct {
	Flow SettingsFlow `json:"flow"`
	Identity Identity `json:"identity"`
}

// NewSettingsViaApiResponse instantiates a new SettingsViaApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsViaApiResponse(flow SettingsFlow, identity Identity) *SettingsViaApiResponse {
	this := SettingsViaApiResponse{}
	this.Flow = flow
	this.Identity = identity
	return &this
}

// NewSettingsViaApiResponseWithDefaults instantiates a new SettingsViaApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsViaApiResponseWithDefaults() *SettingsViaApiResponse {
	this := SettingsViaApiResponse{}
	return &this
}

// GetFlow returns the Flow field value
func (o *SettingsViaApiResponse) GetFlow() SettingsFlow {
	if o == nil {
		var ret SettingsFlow
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *SettingsViaApiResponse) GetFlowOk() (*SettingsFlow, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *SettingsViaApiResponse) SetFlow(v SettingsFlow) {
	o.Flow = v
}

// GetIdentity returns the Identity field value
func (o *SettingsViaApiResponse) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *SettingsViaApiResponse) GetIdentityOk() (*Identity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *SettingsViaApiResponse) SetIdentity(v Identity) {
	o.Identity = v
}

func (o SettingsViaApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flow"] = o.Flow
	}
	if true {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsViaApiResponse struct {
	value *SettingsViaApiResponse
	isSet bool
}

func (v NullableSettingsViaApiResponse) Get() *SettingsViaApiResponse {
	return v.value
}

func (v *NullableSettingsViaApiResponse) Set(val *SettingsViaApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsViaApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsViaApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsViaApiResponse(val *SettingsViaApiResponse) *NullableSettingsViaApiResponse {
	return &NullableSettingsViaApiResponse{value: val, isSet: true}
}

func (v NullableSettingsViaApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsViaApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


