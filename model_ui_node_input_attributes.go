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

// UiNodeInputAttributes InputAttributes represents the attributes of an input node
type UiNodeInputAttributes struct {
	// Sets the input's disabled field to true or false.
	Disabled bool `json:"disabled"`
	Label *UiText `json:"label,omitempty"`
	// The input's element name.
	Name string `json:"name"`
	// The input's pattern.
	Pattern *string `json:"pattern,omitempty"`
	// Mark this input field as required.
	Required *bool `json:"required,omitempty"`
	Type string `json:"type"`
	Value *UiNodeInputAttributesValue `json:"value,omitempty"`
}

// NewUiNodeInputAttributes instantiates a new UiNodeInputAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeInputAttributes(disabled bool, name string, type_ string) *UiNodeInputAttributes {
	this := UiNodeInputAttributes{}
	this.Disabled = disabled
	this.Name = name
	this.Type = type_
	return &this
}

// NewUiNodeInputAttributesWithDefaults instantiates a new UiNodeInputAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeInputAttributesWithDefaults() *UiNodeInputAttributes {
	this := UiNodeInputAttributes{}
	return &this
}

// GetDisabled returns the Disabled field value
func (o *UiNodeInputAttributes) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *UiNodeInputAttributes) SetDisabled(v bool) {
	o.Disabled = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetLabel() UiText {
	if o == nil || o.Label == nil {
		var ret UiText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetLabelOk() (*UiText, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given UiText and assigns it to the Label field.
func (o *UiNodeInputAttributes) SetLabel(v UiText) {
	o.Label = &v
}

// GetName returns the Name field value
func (o *UiNodeInputAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UiNodeInputAttributes) SetName(v string) {
	o.Name = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *UiNodeInputAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UiNodeInputAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value
func (o *UiNodeInputAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UiNodeInputAttributes) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetValue() UiNodeInputAttributesValue {
	if o == nil || o.Value == nil {
		var ret UiNodeInputAttributesValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetValueOk() (*UiNodeInputAttributesValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given UiNodeInputAttributesValue and assigns it to the Value field.
func (o *UiNodeInputAttributes) SetValue(v UiNodeInputAttributesValue) {
	o.Value = &v
}

func (o UiNodeInputAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUiNodeInputAttributes struct {
	value *UiNodeInputAttributes
	isSet bool
}

func (v NullableUiNodeInputAttributes) Get() *UiNodeInputAttributes {
	return v.value
}

func (v *NullableUiNodeInputAttributes) Set(val *UiNodeInputAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeInputAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeInputAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeInputAttributes(val *UiNodeInputAttributes) *NullableUiNodeInputAttributes {
	return &NullableUiNodeInputAttributes{value: val, isSet: true}
}

func (v NullableUiNodeInputAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeInputAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


