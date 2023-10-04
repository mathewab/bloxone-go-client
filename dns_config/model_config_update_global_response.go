/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigUpdateGlobalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigUpdateGlobalResponse{}

// ConfigUpdateGlobalResponse The Global object update response format.
type ConfigUpdateGlobalResponse struct {
	Result *ConfigGlobal `json:"result,omitempty"`
}

// NewConfigUpdateGlobalResponse instantiates a new ConfigUpdateGlobalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUpdateGlobalResponse() *ConfigUpdateGlobalResponse {
	this := ConfigUpdateGlobalResponse{}
	return &this
}

// NewConfigUpdateGlobalResponseWithDefaults instantiates a new ConfigUpdateGlobalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUpdateGlobalResponseWithDefaults() *ConfigUpdateGlobalResponse {
	this := ConfigUpdateGlobalResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigUpdateGlobalResponse) GetResult() ConfigGlobal {
	if o == nil || IsNil(o.Result) {
		var ret ConfigGlobal
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUpdateGlobalResponse) GetResultOk() (*ConfigGlobal, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigUpdateGlobalResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigGlobal and assigns it to the Result field.
func (o *ConfigUpdateGlobalResponse) SetResult(v ConfigGlobal) {
	o.Result = &v
}

func (o ConfigUpdateGlobalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigUpdateGlobalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigUpdateGlobalResponse struct {
	value *ConfigUpdateGlobalResponse
	isSet bool
}

func (v NullableConfigUpdateGlobalResponse) Get() *ConfigUpdateGlobalResponse {
	return v.value
}

func (v *NullableConfigUpdateGlobalResponse) Set(val *ConfigUpdateGlobalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUpdateGlobalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUpdateGlobalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUpdateGlobalResponse(val *ConfigUpdateGlobalResponse) *NullableConfigUpdateGlobalResponse {
	return &NullableConfigUpdateGlobalResponse{value: val, isSet: true}
}

func (v NullableConfigUpdateGlobalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigUpdateGlobalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
