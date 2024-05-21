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

// checks if the DiagnosisDiagnosticLogsStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisDiagnosticLogsStatus{}

// DiagnosisDiagnosticLogsStatus struct for DiagnosisDiagnosticLogsStatus
type DiagnosisDiagnosticLogsStatus struct {
	Message *string `json:"message,omitempty"`
	Status string `json:"status"`
}

type _DiagnosisDiagnosticLogsStatus DiagnosisDiagnosticLogsStatus

// NewDiagnosisDiagnosticLogsStatus instantiates a new DiagnosisDiagnosticLogsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisDiagnosticLogsStatus(status string) *DiagnosisDiagnosticLogsStatus {
	this := DiagnosisDiagnosticLogsStatus{}
	this.Status = status
	return &this
}

// NewDiagnosisDiagnosticLogsStatusWithDefaults instantiates a new DiagnosisDiagnosticLogsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisDiagnosticLogsStatusWithDefaults() *DiagnosisDiagnosticLogsStatus {
	this := DiagnosisDiagnosticLogsStatus{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DiagnosisDiagnosticLogsStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DiagnosisDiagnosticLogsStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DiagnosisDiagnosticLogsStatus) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value
func (o *DiagnosisDiagnosticLogsStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosisDiagnosticLogsStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DiagnosisDiagnosticLogsStatus) SetStatus(v string) {
	o.Status = v
}

func (o DiagnosisDiagnosticLogsStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisDiagnosticLogsStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *DiagnosisDiagnosticLogsStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varDiagnosisDiagnosticLogsStatus := _DiagnosisDiagnosticLogsStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiagnosisDiagnosticLogsStatus)

	if err != nil {
		return err
	}

	*o = DiagnosisDiagnosticLogsStatus(varDiagnosisDiagnosticLogsStatus)

	return err
}

type NullableDiagnosisDiagnosticLogsStatus struct {
	value *DiagnosisDiagnosticLogsStatus
	isSet bool
}

func (v NullableDiagnosisDiagnosticLogsStatus) Get() *DiagnosisDiagnosticLogsStatus {
	return v.value
}

func (v *NullableDiagnosisDiagnosticLogsStatus) Set(val *DiagnosisDiagnosticLogsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisDiagnosticLogsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisDiagnosticLogsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisDiagnosticLogsStatus(val *DiagnosisDiagnosticLogsStatus) *NullableDiagnosisDiagnosticLogsStatus {
	return &NullableDiagnosisDiagnosticLogsStatus{value: val, isSet: true}
}

func (v NullableDiagnosisDiagnosticLogsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisDiagnosticLogsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


