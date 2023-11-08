/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_data

import (
	"encoding/json"
)

// checks if the DataReadRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataReadRecordResponse{}

// DataReadRecordResponse The response format to retrieve the __Record__ object.
type DataReadRecordResponse struct {
	Result *DataRecord `json:"result,omitempty"`
}

// NewDataReadRecordResponse instantiates a new DataReadRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReadRecordResponse() *DataReadRecordResponse {
	this := DataReadRecordResponse{}
	return &this
}

// NewDataReadRecordResponseWithDefaults instantiates a new DataReadRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReadRecordResponseWithDefaults() *DataReadRecordResponse {
	this := DataReadRecordResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DataReadRecordResponse) GetResult() DataRecord {
	if o == nil || IsNil(o.Result) {
		var ret DataRecord
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReadRecordResponse) GetResultOk() (*DataRecord, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DataReadRecordResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DataRecord and assigns it to the Result field.
func (o *DataReadRecordResponse) SetResult(v DataRecord) {
	o.Result = &v
}

func (o DataReadRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataReadRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableDataReadRecordResponse struct {
	value *DataReadRecordResponse
	isSet bool
}

func (v NullableDataReadRecordResponse) Get() *DataReadRecordResponse {
	return v.value
}

func (v *NullableDataReadRecordResponse) Set(val *DataReadRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReadRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReadRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReadRecordResponse(val *DataReadRecordResponse) *NullableDataReadRecordResponse {
	return &NullableDataReadRecordResponse{value: val, isSet: true}
}

func (v NullableDataReadRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReadRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
