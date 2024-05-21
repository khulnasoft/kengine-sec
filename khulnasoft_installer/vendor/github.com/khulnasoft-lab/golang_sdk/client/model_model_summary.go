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

// checks if the ModelSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSummary{}

// ModelSummary struct for ModelSummary
type ModelSummary struct {
	Images *int32 `json:"images,omitempty"`
	Registries *int32 `json:"registries,omitempty"`
	Repositories *int32 `json:"repositories,omitempty"`
	ScansComplete *int32 `json:"scans_complete,omitempty"`
	ScansInProgress *int32 `json:"scans_in_progress,omitempty"`
	ScansTotal *int32 `json:"scans_total,omitempty"`
}

// NewModelSummary instantiates a new ModelSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSummary() *ModelSummary {
	this := ModelSummary{}
	return &this
}

// NewModelSummaryWithDefaults instantiates a new ModelSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSummaryWithDefaults() *ModelSummary {
	this := ModelSummary{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ModelSummary) GetImages() int32 {
	if o == nil || IsNil(o.Images) {
		var ret int32
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetImagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ModelSummary) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given int32 and assigns it to the Images field.
func (o *ModelSummary) SetImages(v int32) {
	o.Images = &v
}

// GetRegistries returns the Registries field value if set, zero value otherwise.
func (o *ModelSummary) GetRegistries() int32 {
	if o == nil || IsNil(o.Registries) {
		var ret int32
		return ret
	}
	return *o.Registries
}

// GetRegistriesOk returns a tuple with the Registries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetRegistriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Registries) {
		return nil, false
	}
	return o.Registries, true
}

// HasRegistries returns a boolean if a field has been set.
func (o *ModelSummary) HasRegistries() bool {
	if o != nil && !IsNil(o.Registries) {
		return true
	}

	return false
}

// SetRegistries gets a reference to the given int32 and assigns it to the Registries field.
func (o *ModelSummary) SetRegistries(v int32) {
	o.Registries = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ModelSummary) GetRepositories() int32 {
	if o == nil || IsNil(o.Repositories) {
		var ret int32
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetRepositoriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ModelSummary) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given int32 and assigns it to the Repositories field.
func (o *ModelSummary) SetRepositories(v int32) {
	o.Repositories = &v
}

// GetScansComplete returns the ScansComplete field value if set, zero value otherwise.
func (o *ModelSummary) GetScansComplete() int32 {
	if o == nil || IsNil(o.ScansComplete) {
		var ret int32
		return ret
	}
	return *o.ScansComplete
}

// GetScansCompleteOk returns a tuple with the ScansComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetScansCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.ScansComplete) {
		return nil, false
	}
	return o.ScansComplete, true
}

// HasScansComplete returns a boolean if a field has been set.
func (o *ModelSummary) HasScansComplete() bool {
	if o != nil && !IsNil(o.ScansComplete) {
		return true
	}

	return false
}

// SetScansComplete gets a reference to the given int32 and assigns it to the ScansComplete field.
func (o *ModelSummary) SetScansComplete(v int32) {
	o.ScansComplete = &v
}

// GetScansInProgress returns the ScansInProgress field value if set, zero value otherwise.
func (o *ModelSummary) GetScansInProgress() int32 {
	if o == nil || IsNil(o.ScansInProgress) {
		var ret int32
		return ret
	}
	return *o.ScansInProgress
}

// GetScansInProgressOk returns a tuple with the ScansInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetScansInProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.ScansInProgress) {
		return nil, false
	}
	return o.ScansInProgress, true
}

// HasScansInProgress returns a boolean if a field has been set.
func (o *ModelSummary) HasScansInProgress() bool {
	if o != nil && !IsNil(o.ScansInProgress) {
		return true
	}

	return false
}

// SetScansInProgress gets a reference to the given int32 and assigns it to the ScansInProgress field.
func (o *ModelSummary) SetScansInProgress(v int32) {
	o.ScansInProgress = &v
}

// GetScansTotal returns the ScansTotal field value if set, zero value otherwise.
func (o *ModelSummary) GetScansTotal() int32 {
	if o == nil || IsNil(o.ScansTotal) {
		var ret int32
		return ret
	}
	return *o.ScansTotal
}

// GetScansTotalOk returns a tuple with the ScansTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSummary) GetScansTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.ScansTotal) {
		return nil, false
	}
	return o.ScansTotal, true
}

// HasScansTotal returns a boolean if a field has been set.
func (o *ModelSummary) HasScansTotal() bool {
	if o != nil && !IsNil(o.ScansTotal) {
		return true
	}

	return false
}

// SetScansTotal gets a reference to the given int32 and assigns it to the ScansTotal field.
func (o *ModelSummary) SetScansTotal(v int32) {
	o.ScansTotal = &v
}

func (o ModelSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Registries) {
		toSerialize["registries"] = o.Registries
	}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	if !IsNil(o.ScansComplete) {
		toSerialize["scans_complete"] = o.ScansComplete
	}
	if !IsNil(o.ScansInProgress) {
		toSerialize["scans_in_progress"] = o.ScansInProgress
	}
	if !IsNil(o.ScansTotal) {
		toSerialize["scans_total"] = o.ScansTotal
	}
	return toSerialize, nil
}

type NullableModelSummary struct {
	value *ModelSummary
	isSet bool
}

func (v NullableModelSummary) Get() *ModelSummary {
	return v.value
}

func (v *NullableModelSummary) Set(val *ModelSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSummary(val *ModelSummary) *NullableModelSummary {
	return &NullableModelSummary{value: val, isSet: true}
}

func (v NullableModelSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


