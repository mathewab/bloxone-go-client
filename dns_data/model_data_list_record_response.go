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

// checks if the DataListRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataListRecordResponse{}

// DataListRecordResponse The response format to retrieve __Record__ objects.
type DataListRecordResponse struct {
	// The list of Record objects.
	Results []DataRecord `json:"results,omitempty"`
}

// NewDataListRecordResponse instantiates a new DataListRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataListRecordResponse() *DataListRecordResponse {
	this := DataListRecordResponse{}
	return &this
}

// NewDataListRecordResponseWithDefaults instantiates a new DataListRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataListRecordResponseWithDefaults() *DataListRecordResponse {
	this := DataListRecordResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DataListRecordResponse) GetResults() []DataRecord {
	if o == nil || IsNil(o.Results) {
		var ret []DataRecord
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListRecordResponse) GetResultsOk() ([]DataRecord, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DataListRecordResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DataRecord and assigns it to the Results field.
func (o *DataListRecordResponse) SetResults(v []DataRecord) {
	o.Results = v
}

func (o DataListRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataListRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDataListRecordResponse struct {
	value *DataListRecordResponse
	isSet bool
}

func (v NullableDataListRecordResponse) Get() *DataListRecordResponse {
	return v.value
}

func (v *NullableDataListRecordResponse) Set(val *DataListRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataListRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataListRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataListRecordResponse(val *DataListRecordResponse) *NullableDataListRecordResponse {
	return &NullableDataListRecordResponse{value: val, isSet: true}
}

func (v NullableDataListRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataListRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
