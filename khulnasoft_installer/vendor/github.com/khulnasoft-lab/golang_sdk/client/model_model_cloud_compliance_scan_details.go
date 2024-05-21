/*
Khulnasoft Kengine

Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.2.0
Contact: community@khulnasoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelCloudComplianceScanDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudComplianceScanDetails{}

// ModelCloudComplianceScanDetails struct for ModelCloudComplianceScanDetails
type ModelCloudComplianceScanDetails struct {
	AccountId *string `json:"account_id,omitempty"`
	Benchmarks []ModelCloudComplianceBenchmark `json:"benchmarks,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	ScanTypes []string `json:"scan_types,omitempty"`
	StopRequested *bool `json:"stop_requested,omitempty"`
}

// NewModelCloudComplianceScanDetails instantiates a new ModelCloudComplianceScanDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudComplianceScanDetails() *ModelCloudComplianceScanDetails {
	this := ModelCloudComplianceScanDetails{}
	return &this
}

// NewModelCloudComplianceScanDetailsWithDefaults instantiates a new ModelCloudComplianceScanDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudComplianceScanDetailsWithDefaults() *ModelCloudComplianceScanDetails {
	this := ModelCloudComplianceScanDetails{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelCloudComplianceScanDetails) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanDetails) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelCloudComplianceScanDetails) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelCloudComplianceScanDetails) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBenchmarks returns the Benchmarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudComplianceScanDetails) GetBenchmarks() []ModelCloudComplianceBenchmark {
	if o == nil {
		var ret []ModelCloudComplianceBenchmark
		return ret
	}
	return o.Benchmarks
}

// GetBenchmarksOk returns a tuple with the Benchmarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceScanDetails) GetBenchmarksOk() ([]ModelCloudComplianceBenchmark, bool) {
	if o == nil || IsNil(o.Benchmarks) {
		return nil, false
	}
	return o.Benchmarks, true
}

// HasBenchmarks returns a boolean if a field has been set.
func (o *ModelCloudComplianceScanDetails) HasBenchmarks() bool {
	if o != nil && !IsNil(o.Benchmarks) {
		return true
	}

	return false
}

// SetBenchmarks gets a reference to the given []ModelCloudComplianceBenchmark and assigns it to the Benchmarks field.
func (o *ModelCloudComplianceScanDetails) SetBenchmarks(v []ModelCloudComplianceBenchmark) {
	o.Benchmarks = v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *ModelCloudComplianceScanDetails) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanDetails) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *ModelCloudComplianceScanDetails) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *ModelCloudComplianceScanDetails) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanTypes returns the ScanTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudComplianceScanDetails) GetScanTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ScanTypes
}

// GetScanTypesOk returns a tuple with the ScanTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudComplianceScanDetails) GetScanTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanTypes) {
		return nil, false
	}
	return o.ScanTypes, true
}

// HasScanTypes returns a boolean if a field has been set.
func (o *ModelCloudComplianceScanDetails) HasScanTypes() bool {
	if o != nil && !IsNil(o.ScanTypes) {
		return true
	}

	return false
}

// SetScanTypes gets a reference to the given []string and assigns it to the ScanTypes field.
func (o *ModelCloudComplianceScanDetails) SetScanTypes(v []string) {
	o.ScanTypes = v
}

// GetStopRequested returns the StopRequested field value if set, zero value otherwise.
func (o *ModelCloudComplianceScanDetails) GetStopRequested() bool {
	if o == nil || IsNil(o.StopRequested) {
		var ret bool
		return ret
	}
	return *o.StopRequested
}

// GetStopRequestedOk returns a tuple with the StopRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudComplianceScanDetails) GetStopRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.StopRequested) {
		return nil, false
	}
	return o.StopRequested, true
}

// HasStopRequested returns a boolean if a field has been set.
func (o *ModelCloudComplianceScanDetails) HasStopRequested() bool {
	if o != nil && !IsNil(o.StopRequested) {
		return true
	}

	return false
}

// SetStopRequested gets a reference to the given bool and assigns it to the StopRequested field.
func (o *ModelCloudComplianceScanDetails) SetStopRequested(v bool) {
	o.StopRequested = &v
}

func (o ModelCloudComplianceScanDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudComplianceScanDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Benchmarks != nil {
		toSerialize["benchmarks"] = o.Benchmarks
	}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if o.ScanTypes != nil {
		toSerialize["scan_types"] = o.ScanTypes
	}
	if !IsNil(o.StopRequested) {
		toSerialize["stop_requested"] = o.StopRequested
	}
	return toSerialize, nil
}

type NullableModelCloudComplianceScanDetails struct {
	value *ModelCloudComplianceScanDetails
	isSet bool
}

func (v NullableModelCloudComplianceScanDetails) Get() *ModelCloudComplianceScanDetails {
	return v.value
}

func (v *NullableModelCloudComplianceScanDetails) Set(val *ModelCloudComplianceScanDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudComplianceScanDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudComplianceScanDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudComplianceScanDetails(val *ModelCloudComplianceScanDetails) *NullableModelCloudComplianceScanDetails {
	return &NullableModelCloudComplianceScanDetails{value: val, isSet: true}
}

func (v NullableModelCloudComplianceScanDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudComplianceScanDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


