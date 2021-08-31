/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.4
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AuthenticateOKBody AuthenticateOKBody authenticate o k body
type AuthenticateOKBody struct {
	// An opaque token used to authenticate a user after a successful login
	IdentityToken string `json:"IdentityToken"`
	// The status of the authentication
	Status string `json:"Status"`
}

// NewAuthenticateOKBody instantiates a new AuthenticateOKBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateOKBody(identityToken string, status string) *AuthenticateOKBody {
	this := AuthenticateOKBody{}
	this.IdentityToken = identityToken
	this.Status = status
	return &this
}

// NewAuthenticateOKBodyWithDefaults instantiates a new AuthenticateOKBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateOKBodyWithDefaults() *AuthenticateOKBody {
	this := AuthenticateOKBody{}
	return &this
}

// GetIdentityToken returns the IdentityToken field value
func (o *AuthenticateOKBody) GetIdentityToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityToken
}

// GetIdentityTokenOk returns a tuple with the IdentityToken field value
// and a boolean to check if the value has been set.
func (o *AuthenticateOKBody) GetIdentityTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityToken, true
}

// SetIdentityToken sets field value
func (o *AuthenticateOKBody) SetIdentityToken(v string) {
	o.IdentityToken = v
}

// GetStatus returns the Status field value
func (o *AuthenticateOKBody) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AuthenticateOKBody) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AuthenticateOKBody) SetStatus(v string) {
	o.Status = v
}

func (o AuthenticateOKBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["IdentityToken"] = o.IdentityToken
	}
	if true {
		toSerialize["Status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateOKBody struct {
	value *AuthenticateOKBody
	isSet bool
}

func (v NullableAuthenticateOKBody) Get() *AuthenticateOKBody {
	return v.value
}

func (v *NullableAuthenticateOKBody) Set(val *AuthenticateOKBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateOKBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateOKBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateOKBody(val *AuthenticateOKBody) *NullableAuthenticateOKBody {
	return &NullableAuthenticateOKBody{value: val, isSet: true}
}

func (v NullableAuthenticateOKBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateOKBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


