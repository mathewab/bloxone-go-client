/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamsvcAsmGrowthBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcAsmGrowthBlock{}

// IpamsvcAsmGrowthBlock ASM growth group of fields.
type IpamsvcAsmGrowthBlock struct {
	// Either the number or percentage of addresses to grow by.
	GrowthFactor *int64 `json:"growth_factor,omitempty"`
	// The type of factor to use: _percent_ or _count_.
	GrowthType *string `json:"growth_type,omitempty"`
}

// NewIpamsvcAsmGrowthBlock instantiates a new IpamsvcAsmGrowthBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcAsmGrowthBlock() *IpamsvcAsmGrowthBlock {
	this := IpamsvcAsmGrowthBlock{}
	return &this
}

// NewIpamsvcAsmGrowthBlockWithDefaults instantiates a new IpamsvcAsmGrowthBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcAsmGrowthBlockWithDefaults() *IpamsvcAsmGrowthBlock {
	this := IpamsvcAsmGrowthBlock{}
	return &this
}

// GetGrowthFactor returns the GrowthFactor field value if set, zero value otherwise.
func (o *IpamsvcAsmGrowthBlock) GetGrowthFactor() int64 {
	if o == nil || IsNil(o.GrowthFactor) {
		var ret int64
		return ret
	}
	return *o.GrowthFactor
}

// GetGrowthFactorOk returns a tuple with the GrowthFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAsmGrowthBlock) GetGrowthFactorOk() (*int64, bool) {
	if o == nil || IsNil(o.GrowthFactor) {
		return nil, false
	}
	return o.GrowthFactor, true
}

// HasGrowthFactor returns a boolean if a field has been set.
func (o *IpamsvcAsmGrowthBlock) HasGrowthFactor() bool {
	if o != nil && !IsNil(o.GrowthFactor) {
		return true
	}

	return false
}

// SetGrowthFactor gets a reference to the given int64 and assigns it to the GrowthFactor field.
func (o *IpamsvcAsmGrowthBlock) SetGrowthFactor(v int64) {
	o.GrowthFactor = &v
}

// GetGrowthType returns the GrowthType field value if set, zero value otherwise.
func (o *IpamsvcAsmGrowthBlock) GetGrowthType() string {
	if o == nil || IsNil(o.GrowthType) {
		var ret string
		return ret
	}
	return *o.GrowthType
}

// GetGrowthTypeOk returns a tuple with the GrowthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcAsmGrowthBlock) GetGrowthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrowthType) {
		return nil, false
	}
	return o.GrowthType, true
}

// HasGrowthType returns a boolean if a field has been set.
func (o *IpamsvcAsmGrowthBlock) HasGrowthType() bool {
	if o != nil && !IsNil(o.GrowthType) {
		return true
	}

	return false
}

// SetGrowthType gets a reference to the given string and assigns it to the GrowthType field.
func (o *IpamsvcAsmGrowthBlock) SetGrowthType(v string) {
	o.GrowthType = &v
}

func (o IpamsvcAsmGrowthBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcAsmGrowthBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GrowthFactor) {
		toSerialize["growth_factor"] = o.GrowthFactor
	}
	if !IsNil(o.GrowthType) {
		toSerialize["growth_type"] = o.GrowthType
	}
	return toSerialize, nil
}

type NullableIpamsvcAsmGrowthBlock struct {
	value *IpamsvcAsmGrowthBlock
	isSet bool
}

func (v NullableIpamsvcAsmGrowthBlock) Get() *IpamsvcAsmGrowthBlock {
	return v.value
}

func (v *NullableIpamsvcAsmGrowthBlock) Set(val *IpamsvcAsmGrowthBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcAsmGrowthBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcAsmGrowthBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcAsmGrowthBlock(val *IpamsvcAsmGrowthBlock) *NullableIpamsvcAsmGrowthBlock {
	return &NullableIpamsvcAsmGrowthBlock{value: val, isSet: true}
}

func (v NullableIpamsvcAsmGrowthBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcAsmGrowthBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
