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

// ImageSummary ImageSummary image summary
type ImageSummary struct {
	// containers
	Containers int64 `json:"Containers"`
	// created
	Created int64 `json:"Created"`
	// Id
	Id string `json:"Id"`
	// labels
	Labels map[string]string `json:"Labels"`
	// parent Id
	ParentId string `json:"ParentId"`
	// repo digests
	RepoDigests []string `json:"RepoDigests"`
	// repo tags
	RepoTags []string `json:"RepoTags"`
	// shared size
	SharedSize int64 `json:"SharedSize"`
	// size
	Size int64 `json:"Size"`
	// virtual size
	VirtualSize int64 `json:"VirtualSize"`
}

// NewImageSummary instantiates a new ImageSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageSummary(containers int64, created int64, id string, labels map[string]string, parentId string, repoDigests []string, repoTags []string, sharedSize int64, size int64, virtualSize int64) *ImageSummary {
	this := ImageSummary{}
	this.Containers = containers
	this.Created = created
	this.Id = id
	this.Labels = labels
	this.ParentId = parentId
	this.RepoDigests = repoDigests
	this.RepoTags = repoTags
	this.SharedSize = sharedSize
	this.Size = size
	this.VirtualSize = virtualSize
	return &this
}

// NewImageSummaryWithDefaults instantiates a new ImageSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageSummaryWithDefaults() *ImageSummary {
	this := ImageSummary{}
	return &this
}

// GetContainers returns the Containers field value
func (o *ImageSummary) GetContainers() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetContainersOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Containers, true
}

// SetContainers sets field value
func (o *ImageSummary) SetContainers(v int64) {
	o.Containers = v
}

// GetCreated returns the Created field value
func (o *ImageSummary) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetCreatedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ImageSummary) SetCreated(v int64) {
	o.Created = v
}

// GetId returns the Id field value
func (o *ImageSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageSummary) SetId(v string) {
	o.Id = v
}

// GetLabels returns the Labels field value
func (o *ImageSummary) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetLabelsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *ImageSummary) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetParentId returns the ParentId field value
func (o *ImageSummary) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetParentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ImageSummary) SetParentId(v string) {
	o.ParentId = v
}

// GetRepoDigests returns the RepoDigests field value
func (o *ImageSummary) GetRepoDigests() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RepoDigests
}

// GetRepoDigestsOk returns a tuple with the RepoDigests field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetRepoDigestsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RepoDigests, true
}

// SetRepoDigests sets field value
func (o *ImageSummary) SetRepoDigests(v []string) {
	o.RepoDigests = v
}

// GetRepoTags returns the RepoTags field value
func (o *ImageSummary) GetRepoTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RepoTags
}

// GetRepoTagsOk returns a tuple with the RepoTags field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetRepoTagsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RepoTags, true
}

// SetRepoTags sets field value
func (o *ImageSummary) SetRepoTags(v []string) {
	o.RepoTags = v
}

// GetSharedSize returns the SharedSize field value
func (o *ImageSummary) GetSharedSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SharedSize
}

// GetSharedSizeOk returns a tuple with the SharedSize field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetSharedSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SharedSize, true
}

// SetSharedSize sets field value
func (o *ImageSummary) SetSharedSize(v int64) {
	o.SharedSize = v
}

// GetSize returns the Size field value
func (o *ImageSummary) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ImageSummary) SetSize(v int64) {
	o.Size = v
}

// GetVirtualSize returns the VirtualSize field value
func (o *ImageSummary) GetVirtualSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VirtualSize
}

// GetVirtualSizeOk returns a tuple with the VirtualSize field value
// and a boolean to check if the value has been set.
func (o *ImageSummary) GetVirtualSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualSize, true
}

// SetVirtualSize sets field value
func (o *ImageSummary) SetVirtualSize(v int64) {
	o.VirtualSize = v
}

func (o ImageSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Containers"] = o.Containers
	}
	if true {
		toSerialize["Created"] = o.Created
	}
	if true {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Labels"] = o.Labels
	}
	if true {
		toSerialize["ParentId"] = o.ParentId
	}
	if true {
		toSerialize["RepoDigests"] = o.RepoDigests
	}
	if true {
		toSerialize["RepoTags"] = o.RepoTags
	}
	if true {
		toSerialize["SharedSize"] = o.SharedSize
	}
	if true {
		toSerialize["Size"] = o.Size
	}
	if true {
		toSerialize["VirtualSize"] = o.VirtualSize
	}
	return json.Marshal(toSerialize)
}

type NullableImageSummary struct {
	value *ImageSummary
	isSet bool
}

func (v NullableImageSummary) Get() *ImageSummary {
	return v.value
}

func (v *NullableImageSummary) Set(val *ImageSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSummary(val *ImageSummary) *NullableImageSummary {
	return &NullableImageSummary{value: val, isSet: true}
}

func (v NullableImageSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


