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

// checks if the ConfigInheritedACLItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigInheritedACLItems{}

// ConfigInheritedACLItems Inheritance configuration for a field of type list of _ACLItem_.
type ConfigInheritedACLItems struct {
	// Optional. Inheritance setting for a field. Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// Human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source *string `json:"source,omitempty"`
	// Inherited value.
	Value []ConfigACLItem `json:"value,omitempty"`
}

// NewConfigInheritedACLItems instantiates a new ConfigInheritedACLItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInheritedACLItems() *ConfigInheritedACLItems {
	this := ConfigInheritedACLItems{}
	return &this
}

// NewConfigInheritedACLItemsWithDefaults instantiates a new ConfigInheritedACLItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInheritedACLItemsWithDefaults() *ConfigInheritedACLItems {
	this := ConfigInheritedACLItems{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ConfigInheritedACLItems) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedACLItems) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ConfigInheritedACLItems) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ConfigInheritedACLItems) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConfigInheritedACLItems) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedACLItems) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConfigInheritedACLItems) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConfigInheritedACLItems) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConfigInheritedACLItems) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedACLItems) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConfigInheritedACLItems) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConfigInheritedACLItems) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigInheritedACLItems) GetValue() []ConfigACLItem {
	if o == nil || IsNil(o.Value) {
		var ret []ConfigACLItem
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedACLItems) GetValueOk() ([]ConfigACLItem, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigInheritedACLItems) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []ConfigACLItem and assigns it to the Value field.
func (o *ConfigInheritedACLItems) SetValue(v []ConfigACLItem) {
	o.Value = v
}

func (o ConfigInheritedACLItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigInheritedACLItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableConfigInheritedACLItems struct {
	value *ConfigInheritedACLItems
	isSet bool
}

func (v NullableConfigInheritedACLItems) Get() *ConfigInheritedACLItems {
	return v.value
}

func (v *NullableConfigInheritedACLItems) Set(val *ConfigInheritedACLItems) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInheritedACLItems) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInheritedACLItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInheritedACLItems(val *ConfigInheritedACLItems) *NullableConfigInheritedACLItems {
	return &NullableConfigInheritedACLItems{value: val, isSet: true}
}

func (v NullableConfigInheritedACLItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInheritedACLItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
