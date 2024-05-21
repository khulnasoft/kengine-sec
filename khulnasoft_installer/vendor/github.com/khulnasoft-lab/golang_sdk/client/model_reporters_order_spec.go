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
	"bytes"
	"fmt"
)

// checks if the ReportersOrderSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersOrderSpec{}

// ReportersOrderSpec struct for ReportersOrderSpec
type ReportersOrderSpec struct {
	Descending bool `json:"descending"`
	FieldName string `json:"field_name"`
	Size *int32 `json:"size,omitempty"`
}

type _ReportersOrderSpec ReportersOrderSpec

// NewReportersOrderSpec instantiates a new ReportersOrderSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersOrderSpec(descending bool, fieldName string) *ReportersOrderSpec {
	this := ReportersOrderSpec{}
	this.Descending = descending
	this.FieldName = fieldName
	return &this
}

// NewReportersOrderSpecWithDefaults instantiates a new ReportersOrderSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersOrderSpecWithDefaults() *ReportersOrderSpec {
	this := ReportersOrderSpec{}
	return &this
}

// GetDescending returns the Descending field value
func (o *ReportersOrderSpec) GetDescending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value
// and a boolean to check if the value has been set.
func (o *ReportersOrderSpec) GetDescendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descending, true
}

// SetDescending sets field value
func (o *ReportersOrderSpec) SetDescending(v bool) {
	o.Descending = v
}

// GetFieldName returns the FieldName field value
func (o *ReportersOrderSpec) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *ReportersOrderSpec) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *ReportersOrderSpec) SetFieldName(v string) {
	o.FieldName = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReportersOrderSpec) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersOrderSpec) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReportersOrderSpec) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ReportersOrderSpec) SetSize(v int32) {
	o.Size = &v
}

func (o ReportersOrderSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersOrderSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descending"] = o.Descending
	toSerialize["field_name"] = o.FieldName
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

func (o *ReportersOrderSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"descending",
		"field_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReportersOrderSpec := _ReportersOrderSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersOrderSpec)

	if err != nil {
		return err
	}

	*o = ReportersOrderSpec(varReportersOrderSpec)

	return err
}

type NullableReportersOrderSpec struct {
	value *ReportersOrderSpec
	isSet bool
}

func (v NullableReportersOrderSpec) Get() *ReportersOrderSpec {
	return v.value
}

func (v *NullableReportersOrderSpec) Set(val *ReportersOrderSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersOrderSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersOrderSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersOrderSpec(val *ReportersOrderSpec) *NullableReportersOrderSpec {
	return &NullableReportersOrderSpec{value: val, isSet: true}
}

func (v NullableReportersOrderSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersOrderSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


