/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_provision

import (
	"encoding/json"
	"time"
)

// checks if the HostactivationJoinToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostactivationJoinToken{}

// HostactivationJoinToken struct for HostactivationJoinToken
type HostactivationJoinToken struct {
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	// The resource identifier.
	Id         *string                   `json:"id,omitempty"`
	LastUsedAt *time.Time                `json:"last_used_at,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Status     *JoinTokenJoinTokenStatus `json:"status,omitempty"`
	Tags       *TypesJSONValue           `json:"tags,omitempty"`
	// first half of the token.
	TokenId    *string `json:"token_id,omitempty"`
	UseCounter *int64  `json:"use_counter,omitempty"`
}

// NewHostactivationJoinToken instantiates a new HostactivationJoinToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostactivationJoinToken() *HostactivationJoinToken {
	this := HostactivationJoinToken{}
	var status JoinTokenJoinTokenStatus = JOINTOKENJOINTOKENSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewHostactivationJoinTokenWithDefaults instantiates a new HostactivationJoinToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostactivationJoinTokenWithDefaults() *HostactivationJoinToken {
	this := HostactivationJoinToken{}
	var status JoinTokenJoinTokenStatus = JOINTOKENJOINTOKENSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *HostactivationJoinToken) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HostactivationJoinToken) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *HostactivationJoinToken) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostactivationJoinToken) SetId(v string) {
	o.Id = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *HostactivationJoinToken) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostactivationJoinToken) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetStatus() JoinTokenJoinTokenStatus {
	if o == nil || IsNil(o.Status) {
		var ret JoinTokenJoinTokenStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetStatusOk() (*JoinTokenJoinTokenStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given JoinTokenJoinTokenStatus and assigns it to the Status field.
func (o *HostactivationJoinToken) SetStatus(v JoinTokenJoinTokenStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetTags() TypesJSONValue {
	if o == nil || IsNil(o.Tags) {
		var ret TypesJSONValue
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetTagsOk() (*TypesJSONValue, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given TypesJSONValue and assigns it to the Tags field.
func (o *HostactivationJoinToken) SetTags(v TypesJSONValue) {
	o.Tags = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *HostactivationJoinToken) SetTokenId(v string) {
	o.TokenId = &v
}

// GetUseCounter returns the UseCounter field value if set, zero value otherwise.
func (o *HostactivationJoinToken) GetUseCounter() int64 {
	if o == nil || IsNil(o.UseCounter) {
		var ret int64
		return ret
	}
	return *o.UseCounter
}

// GetUseCounterOk returns a tuple with the UseCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostactivationJoinToken) GetUseCounterOk() (*int64, bool) {
	if o == nil || IsNil(o.UseCounter) {
		return nil, false
	}
	return o.UseCounter, true
}

// HasUseCounter returns a boolean if a field has been set.
func (o *HostactivationJoinToken) HasUseCounter() bool {
	if o != nil && !IsNil(o.UseCounter) {
		return true
	}

	return false
}

// SetUseCounter gets a reference to the given int64 and assigns it to the UseCounter field.
func (o *HostactivationJoinToken) SetUseCounter(v int64) {
	o.UseCounter = &v
}

func (o HostactivationJoinToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostactivationJoinToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUsedAt) {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.UseCounter) {
		toSerialize["use_counter"] = o.UseCounter
	}
	return toSerialize, nil
}

type NullableHostactivationJoinToken struct {
	value *HostactivationJoinToken
	isSet bool
}

func (v NullableHostactivationJoinToken) Get() *HostactivationJoinToken {
	return v.value
}

func (v *NullableHostactivationJoinToken) Set(val *HostactivationJoinToken) {
	v.value = val
	v.isSet = true
}

func (v NullableHostactivationJoinToken) IsSet() bool {
	return v.isSet
}

func (v *NullableHostactivationJoinToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostactivationJoinToken(val *HostactivationJoinToken) *NullableHostactivationJoinToken {
	return &NullableHostactivationJoinToken{value: val, isSet: true}
}

func (v NullableHostactivationJoinToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostactivationJoinToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
