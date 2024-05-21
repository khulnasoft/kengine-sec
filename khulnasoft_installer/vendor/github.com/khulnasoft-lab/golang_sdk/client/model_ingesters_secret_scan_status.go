/*
Khulnasoft Kengine

Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@khulnasoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IngestersSecretScanStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersSecretScanStatus{}

// IngestersSecretScanStatus struct for IngestersSecretScanStatus
type IngestersSecretScanStatus struct {
	ScanId *string `json:"scan_id,omitempty"`
	ScanMessage *string `json:"scan_message,omitempty"`
	ScanStatus *string `json:"scan_status,omitempty"`
}

// NewIngestersSecretScanStatus instantiates a new IngestersSecretScanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersSecretScanStatus() *IngestersSecretScanStatus {
	this := IngestersSecretScanStatus{}
	return &this
}

// NewIngestersSecretScanStatusWithDefaults instantiates a new IngestersSecretScanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersSecretScanStatusWithDefaults() *IngestersSecretScanStatus {
	this := IngestersSecretScanStatus{}
	return &this
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersSecretScanStatus) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretScanStatus) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersSecretScanStatus) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersSecretScanStatus) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanMessage returns the ScanMessage field value if set, zero value otherwise.
func (o *IngestersSecretScanStatus) GetScanMessage() string {
	if o == nil || IsNil(o.ScanMessage) {
		var ret string
		return ret
	}
	return *o.ScanMessage
}

// GetScanMessageOk returns a tuple with the ScanMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretScanStatus) GetScanMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ScanMessage) {
		return nil, false
	}
	return o.ScanMessage, true
}

// HasScanMessage returns a boolean if a field has been set.
func (o *IngestersSecretScanStatus) HasScanMessage() bool {
	if o != nil && !IsNil(o.ScanMessage) {
		return true
	}

	return false
}

// SetScanMessage gets a reference to the given string and assigns it to the ScanMessage field.
func (o *IngestersSecretScanStatus) SetScanMessage(v string) {
	o.ScanMessage = &v
}

// GetScanStatus returns the ScanStatus field value if set, zero value otherwise.
func (o *IngestersSecretScanStatus) GetScanStatus() string {
	if o == nil || IsNil(o.ScanStatus) {
		var ret string
		return ret
	}
	return *o.ScanStatus
}

// GetScanStatusOk returns a tuple with the ScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretScanStatus) GetScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ScanStatus) {
		return nil, false
	}
	return o.ScanStatus, true
}

// HasScanStatus returns a boolean if a field has been set.
func (o *IngestersSecretScanStatus) HasScanStatus() bool {
	if o != nil && !IsNil(o.ScanStatus) {
		return true
	}

	return false
}

// SetScanStatus gets a reference to the given string and assigns it to the ScanStatus field.
func (o *IngestersSecretScanStatus) SetScanStatus(v string) {
	o.ScanStatus = &v
}

func (o IngestersSecretScanStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersSecretScanStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !IsNil(o.ScanMessage) {
		toSerialize["scan_message"] = o.ScanMessage
	}
	if !IsNil(o.ScanStatus) {
		toSerialize["scan_status"] = o.ScanStatus
	}
	return toSerialize, nil
}

type NullableIngestersSecretScanStatus struct {
	value *IngestersSecretScanStatus
	isSet bool
}

func (v NullableIngestersSecretScanStatus) Get() *IngestersSecretScanStatus {
	return v.value
}

func (v *NullableIngestersSecretScanStatus) Set(val *IngestersSecretScanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersSecretScanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersSecretScanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersSecretScanStatus(val *IngestersSecretScanStatus) *NullableIngestersSecretScanStatus {
	return &NullableIngestersSecretScanStatus{value: val, isSet: true}
}

func (v NullableIngestersSecretScanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersSecretScanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


