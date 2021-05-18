/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginInterfaceType PluginInterfaceType plugin interface type
type PluginInterfaceType struct {
	// capability
	Capability string `json:"Capability"`
	// prefix
	Prefix string `json:"Prefix"`
	// version
	Version string `json:"Version"`
}

// NewPluginInterfaceType instantiates a new PluginInterfaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginInterfaceType(capability string, prefix string, version string) *PluginInterfaceType {
	this := PluginInterfaceType{}
	this.Capability = capability
	this.Prefix = prefix
	this.Version = version
	return &this
}

// NewPluginInterfaceTypeWithDefaults instantiates a new PluginInterfaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginInterfaceTypeWithDefaults() *PluginInterfaceType {
	this := PluginInterfaceType{}
	return &this
}

// GetCapability returns the Capability field value
func (o *PluginInterfaceType) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *PluginInterfaceType) GetCapabilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *PluginInterfaceType) SetCapability(v string) {
	o.Capability = v
}

// GetPrefix returns the Prefix field value
func (o *PluginInterfaceType) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *PluginInterfaceType) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *PluginInterfaceType) SetPrefix(v string) {
	o.Prefix = v
}

// GetVersion returns the Version field value
func (o *PluginInterfaceType) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PluginInterfaceType) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PluginInterfaceType) SetVersion(v string) {
	o.Version = v
}

func (o PluginInterfaceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Capability"] = o.Capability
	}
	if true {
		toSerialize["Prefix"] = o.Prefix
	}
	if true {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePluginInterfaceType struct {
	value *PluginInterfaceType
	isSet bool
}

func (v NullablePluginInterfaceType) Get() *PluginInterfaceType {
	return v.value
}

func (v *NullablePluginInterfaceType) Set(val *PluginInterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginInterfaceType(val *PluginInterfaceType) *NullablePluginInterfaceType {
	return &NullablePluginInterfaceType{value: val, isSet: true}
}

func (v NullablePluginInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


