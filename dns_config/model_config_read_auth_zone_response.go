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

// checks if the ConfigReadAuthZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigReadAuthZoneResponse{}

// ConfigReadAuthZoneResponse The Authoritative Zone object read response format.
type ConfigReadAuthZoneResponse struct {
	Result *ConfigAuthZone `json:"result,omitempty"`
}

// NewConfigReadAuthZoneResponse instantiates a new ConfigReadAuthZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigReadAuthZoneResponse() *ConfigReadAuthZoneResponse {
	this := ConfigReadAuthZoneResponse{}
	return &this
}

// NewConfigReadAuthZoneResponseWithDefaults instantiates a new ConfigReadAuthZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigReadAuthZoneResponseWithDefaults() *ConfigReadAuthZoneResponse {
	this := ConfigReadAuthZoneResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConfigReadAuthZoneResponse) GetResult() ConfigAuthZone {
	if o == nil || IsNil(o.Result) {
		var ret ConfigAuthZone
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigReadAuthZoneResponse) GetResultOk() (*ConfigAuthZone, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConfigReadAuthZoneResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ConfigAuthZone and assigns it to the Result field.
func (o *ConfigReadAuthZoneResponse) SetResult(v ConfigAuthZone) {
	o.Result = &v
}

func (o ConfigReadAuthZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigReadAuthZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableConfigReadAuthZoneResponse struct {
	value *ConfigReadAuthZoneResponse
	isSet bool
}

func (v NullableConfigReadAuthZoneResponse) Get() *ConfigReadAuthZoneResponse {
	return v.value
}

func (v *NullableConfigReadAuthZoneResponse) Set(val *ConfigReadAuthZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigReadAuthZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigReadAuthZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigReadAuthZoneResponse(val *ConfigReadAuthZoneResponse) *NullableConfigReadAuthZoneResponse {
	return &NullableConfigReadAuthZoneResponse{value: val, isSet: true}
}

func (v NullableConfigReadAuthZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigReadAuthZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
