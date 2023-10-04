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

// checks if the ConfigViewInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigViewInheritance{}

// ConfigViewInheritance Inheritance configuration specifies how and which fields _View_ object inherits from [ _Global_, _Server_ ] parent.
type ConfigViewInheritance struct {
	AddEdnsOptionInOutgoingQuery      *Inheritance2InheritedBool            `json:"add_edns_option_in_outgoing_query,omitempty"`
	CustomRootNsBlock                 *ConfigInheritedCustomRootNSBlock     `json:"custom_root_ns_block,omitempty"`
	DnssecValidationBlock             *ConfigInheritedDNSSECValidationBlock `json:"dnssec_validation_block,omitempty"`
	EcsBlock                          *ConfigInheritedECSBlock              `json:"ecs_block,omitempty"`
	EdnsUdpSize                       *Inheritance2InheritedUInt32          `json:"edns_udp_size,omitempty"`
	FilterAaaaAcl                     *ConfigInheritedACLItems              `json:"filter_aaaa_acl,omitempty"`
	FilterAaaaOnV4                    *Inheritance2InheritedString          `json:"filter_aaaa_on_v4,omitempty"`
	ForwardersBlock                   *ConfigInheritedForwardersBlock       `json:"forwarders_block,omitempty"`
	GssTsigEnabled                    *Inheritance2InheritedBool            `json:"gss_tsig_enabled,omitempty"`
	LameTtl                           *Inheritance2InheritedUInt32          `json:"lame_ttl,omitempty"`
	MatchRecursiveOnly                *Inheritance2InheritedBool            `json:"match_recursive_only,omitempty"`
	MaxCacheTtl                       *Inheritance2InheritedUInt32          `json:"max_cache_ttl,omitempty"`
	MaxNegativeTtl                    *Inheritance2InheritedUInt32          `json:"max_negative_ttl,omitempty"`
	MaxUdpSize                        *Inheritance2InheritedUInt32          `json:"max_udp_size,omitempty"`
	MinimalResponses                  *Inheritance2InheritedBool            `json:"minimal_responses,omitempty"`
	Notify                            *Inheritance2InheritedBool            `json:"notify,omitempty"`
	QueryAcl                          *ConfigInheritedACLItems              `json:"query_acl,omitempty"`
	RecursionAcl                      *ConfigInheritedACLItems              `json:"recursion_acl,omitempty"`
	RecursionEnabled                  *Inheritance2InheritedBool            `json:"recursion_enabled,omitempty"`
	SortList                          *ConfigInheritedSortListItems         `json:"sort_list,omitempty"`
	SynthesizeAddressRecordsFromHttps *Inheritance2InheritedBool            `json:"synthesize_address_records_from_https,omitempty"`
	TransferAcl                       *ConfigInheritedACLItems              `json:"transfer_acl,omitempty"`
	UpdateAcl                         *ConfigInheritedACLItems              `json:"update_acl,omitempty"`
	UseForwardersForSubzones          *Inheritance2InheritedBool            `json:"use_forwarders_for_subzones,omitempty"`
	ZoneAuthority                     *ConfigInheritedZoneAuthority         `json:"zone_authority,omitempty"`
}

// NewConfigViewInheritance instantiates a new ConfigViewInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigViewInheritance() *ConfigViewInheritance {
	this := ConfigViewInheritance{}
	return &this
}

// NewConfigViewInheritanceWithDefaults instantiates a new ConfigViewInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigViewInheritanceWithDefaults() *ConfigViewInheritance {
	this := ConfigViewInheritance{}
	return &this
}

// GetAddEdnsOptionInOutgoingQuery returns the AddEdnsOptionInOutgoingQuery field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetAddEdnsOptionInOutgoingQuery() Inheritance2InheritedBool {
	if o == nil || IsNil(o.AddEdnsOptionInOutgoingQuery) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.AddEdnsOptionInOutgoingQuery
}

// GetAddEdnsOptionInOutgoingQueryOk returns a tuple with the AddEdnsOptionInOutgoingQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetAddEdnsOptionInOutgoingQueryOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.AddEdnsOptionInOutgoingQuery) {
		return nil, false
	}
	return o.AddEdnsOptionInOutgoingQuery, true
}

// HasAddEdnsOptionInOutgoingQuery returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasAddEdnsOptionInOutgoingQuery() bool {
	if o != nil && !IsNil(o.AddEdnsOptionInOutgoingQuery) {
		return true
	}

	return false
}

// SetAddEdnsOptionInOutgoingQuery gets a reference to the given Inheritance2InheritedBool and assigns it to the AddEdnsOptionInOutgoingQuery field.
func (o *ConfigViewInheritance) SetAddEdnsOptionInOutgoingQuery(v Inheritance2InheritedBool) {
	o.AddEdnsOptionInOutgoingQuery = &v
}

// GetCustomRootNsBlock returns the CustomRootNsBlock field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetCustomRootNsBlock() ConfigInheritedCustomRootNSBlock {
	if o == nil || IsNil(o.CustomRootNsBlock) {
		var ret ConfigInheritedCustomRootNSBlock
		return ret
	}
	return *o.CustomRootNsBlock
}

// GetCustomRootNsBlockOk returns a tuple with the CustomRootNsBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetCustomRootNsBlockOk() (*ConfigInheritedCustomRootNSBlock, bool) {
	if o == nil || IsNil(o.CustomRootNsBlock) {
		return nil, false
	}
	return o.CustomRootNsBlock, true
}

// HasCustomRootNsBlock returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasCustomRootNsBlock() bool {
	if o != nil && !IsNil(o.CustomRootNsBlock) {
		return true
	}

	return false
}

// SetCustomRootNsBlock gets a reference to the given ConfigInheritedCustomRootNSBlock and assigns it to the CustomRootNsBlock field.
func (o *ConfigViewInheritance) SetCustomRootNsBlock(v ConfigInheritedCustomRootNSBlock) {
	o.CustomRootNsBlock = &v
}

// GetDnssecValidationBlock returns the DnssecValidationBlock field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetDnssecValidationBlock() ConfigInheritedDNSSECValidationBlock {
	if o == nil || IsNil(o.DnssecValidationBlock) {
		var ret ConfigInheritedDNSSECValidationBlock
		return ret
	}
	return *o.DnssecValidationBlock
}

// GetDnssecValidationBlockOk returns a tuple with the DnssecValidationBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetDnssecValidationBlockOk() (*ConfigInheritedDNSSECValidationBlock, bool) {
	if o == nil || IsNil(o.DnssecValidationBlock) {
		return nil, false
	}
	return o.DnssecValidationBlock, true
}

// HasDnssecValidationBlock returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasDnssecValidationBlock() bool {
	if o != nil && !IsNil(o.DnssecValidationBlock) {
		return true
	}

	return false
}

// SetDnssecValidationBlock gets a reference to the given ConfigInheritedDNSSECValidationBlock and assigns it to the DnssecValidationBlock field.
func (o *ConfigViewInheritance) SetDnssecValidationBlock(v ConfigInheritedDNSSECValidationBlock) {
	o.DnssecValidationBlock = &v
}

// GetEcsBlock returns the EcsBlock field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetEcsBlock() ConfigInheritedECSBlock {
	if o == nil || IsNil(o.EcsBlock) {
		var ret ConfigInheritedECSBlock
		return ret
	}
	return *o.EcsBlock
}

// GetEcsBlockOk returns a tuple with the EcsBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetEcsBlockOk() (*ConfigInheritedECSBlock, bool) {
	if o == nil || IsNil(o.EcsBlock) {
		return nil, false
	}
	return o.EcsBlock, true
}

// HasEcsBlock returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasEcsBlock() bool {
	if o != nil && !IsNil(o.EcsBlock) {
		return true
	}

	return false
}

// SetEcsBlock gets a reference to the given ConfigInheritedECSBlock and assigns it to the EcsBlock field.
func (o *ConfigViewInheritance) SetEcsBlock(v ConfigInheritedECSBlock) {
	o.EcsBlock = &v
}

// GetEdnsUdpSize returns the EdnsUdpSize field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetEdnsUdpSize() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.EdnsUdpSize) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.EdnsUdpSize
}

// GetEdnsUdpSizeOk returns a tuple with the EdnsUdpSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetEdnsUdpSizeOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.EdnsUdpSize) {
		return nil, false
	}
	return o.EdnsUdpSize, true
}

// HasEdnsUdpSize returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasEdnsUdpSize() bool {
	if o != nil && !IsNil(o.EdnsUdpSize) {
		return true
	}

	return false
}

// SetEdnsUdpSize gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the EdnsUdpSize field.
func (o *ConfigViewInheritance) SetEdnsUdpSize(v Inheritance2InheritedUInt32) {
	o.EdnsUdpSize = &v
}

// GetFilterAaaaAcl returns the FilterAaaaAcl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetFilterAaaaAcl() ConfigInheritedACLItems {
	if o == nil || IsNil(o.FilterAaaaAcl) {
		var ret ConfigInheritedACLItems
		return ret
	}
	return *o.FilterAaaaAcl
}

// GetFilterAaaaAclOk returns a tuple with the FilterAaaaAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetFilterAaaaAclOk() (*ConfigInheritedACLItems, bool) {
	if o == nil || IsNil(o.FilterAaaaAcl) {
		return nil, false
	}
	return o.FilterAaaaAcl, true
}

// HasFilterAaaaAcl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasFilterAaaaAcl() bool {
	if o != nil && !IsNil(o.FilterAaaaAcl) {
		return true
	}

	return false
}

// SetFilterAaaaAcl gets a reference to the given ConfigInheritedACLItems and assigns it to the FilterAaaaAcl field.
func (o *ConfigViewInheritance) SetFilterAaaaAcl(v ConfigInheritedACLItems) {
	o.FilterAaaaAcl = &v
}

// GetFilterAaaaOnV4 returns the FilterAaaaOnV4 field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetFilterAaaaOnV4() Inheritance2InheritedString {
	if o == nil || IsNil(o.FilterAaaaOnV4) {
		var ret Inheritance2InheritedString
		return ret
	}
	return *o.FilterAaaaOnV4
}

// GetFilterAaaaOnV4Ok returns a tuple with the FilterAaaaOnV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetFilterAaaaOnV4Ok() (*Inheritance2InheritedString, bool) {
	if o == nil || IsNil(o.FilterAaaaOnV4) {
		return nil, false
	}
	return o.FilterAaaaOnV4, true
}

// HasFilterAaaaOnV4 returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasFilterAaaaOnV4() bool {
	if o != nil && !IsNil(o.FilterAaaaOnV4) {
		return true
	}

	return false
}

// SetFilterAaaaOnV4 gets a reference to the given Inheritance2InheritedString and assigns it to the FilterAaaaOnV4 field.
func (o *ConfigViewInheritance) SetFilterAaaaOnV4(v Inheritance2InheritedString) {
	o.FilterAaaaOnV4 = &v
}

// GetForwardersBlock returns the ForwardersBlock field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetForwardersBlock() ConfigInheritedForwardersBlock {
	if o == nil || IsNil(o.ForwardersBlock) {
		var ret ConfigInheritedForwardersBlock
		return ret
	}
	return *o.ForwardersBlock
}

// GetForwardersBlockOk returns a tuple with the ForwardersBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetForwardersBlockOk() (*ConfigInheritedForwardersBlock, bool) {
	if o == nil || IsNil(o.ForwardersBlock) {
		return nil, false
	}
	return o.ForwardersBlock, true
}

// HasForwardersBlock returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasForwardersBlock() bool {
	if o != nil && !IsNil(o.ForwardersBlock) {
		return true
	}

	return false
}

// SetForwardersBlock gets a reference to the given ConfigInheritedForwardersBlock and assigns it to the ForwardersBlock field.
func (o *ConfigViewInheritance) SetForwardersBlock(v ConfigInheritedForwardersBlock) {
	o.ForwardersBlock = &v
}

// GetGssTsigEnabled returns the GssTsigEnabled field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetGssTsigEnabled() Inheritance2InheritedBool {
	if o == nil || IsNil(o.GssTsigEnabled) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.GssTsigEnabled
}

// GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetGssTsigEnabledOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.GssTsigEnabled) {
		return nil, false
	}
	return o.GssTsigEnabled, true
}

// HasGssTsigEnabled returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasGssTsigEnabled() bool {
	if o != nil && !IsNil(o.GssTsigEnabled) {
		return true
	}

	return false
}

// SetGssTsigEnabled gets a reference to the given Inheritance2InheritedBool and assigns it to the GssTsigEnabled field.
func (o *ConfigViewInheritance) SetGssTsigEnabled(v Inheritance2InheritedBool) {
	o.GssTsigEnabled = &v
}

// GetLameTtl returns the LameTtl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetLameTtl() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.LameTtl) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.LameTtl
}

// GetLameTtlOk returns a tuple with the LameTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetLameTtlOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.LameTtl) {
		return nil, false
	}
	return o.LameTtl, true
}

// HasLameTtl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasLameTtl() bool {
	if o != nil && !IsNil(o.LameTtl) {
		return true
	}

	return false
}

// SetLameTtl gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the LameTtl field.
func (o *ConfigViewInheritance) SetLameTtl(v Inheritance2InheritedUInt32) {
	o.LameTtl = &v
}

// GetMatchRecursiveOnly returns the MatchRecursiveOnly field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetMatchRecursiveOnly() Inheritance2InheritedBool {
	if o == nil || IsNil(o.MatchRecursiveOnly) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.MatchRecursiveOnly
}

// GetMatchRecursiveOnlyOk returns a tuple with the MatchRecursiveOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetMatchRecursiveOnlyOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.MatchRecursiveOnly) {
		return nil, false
	}
	return o.MatchRecursiveOnly, true
}

// HasMatchRecursiveOnly returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasMatchRecursiveOnly() bool {
	if o != nil && !IsNil(o.MatchRecursiveOnly) {
		return true
	}

	return false
}

// SetMatchRecursiveOnly gets a reference to the given Inheritance2InheritedBool and assigns it to the MatchRecursiveOnly field.
func (o *ConfigViewInheritance) SetMatchRecursiveOnly(v Inheritance2InheritedBool) {
	o.MatchRecursiveOnly = &v
}

// GetMaxCacheTtl returns the MaxCacheTtl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetMaxCacheTtl() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.MaxCacheTtl) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.MaxCacheTtl
}

// GetMaxCacheTtlOk returns a tuple with the MaxCacheTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetMaxCacheTtlOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.MaxCacheTtl) {
		return nil, false
	}
	return o.MaxCacheTtl, true
}

// HasMaxCacheTtl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasMaxCacheTtl() bool {
	if o != nil && !IsNil(o.MaxCacheTtl) {
		return true
	}

	return false
}

// SetMaxCacheTtl gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the MaxCacheTtl field.
func (o *ConfigViewInheritance) SetMaxCacheTtl(v Inheritance2InheritedUInt32) {
	o.MaxCacheTtl = &v
}

// GetMaxNegativeTtl returns the MaxNegativeTtl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetMaxNegativeTtl() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.MaxNegativeTtl) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.MaxNegativeTtl
}

// GetMaxNegativeTtlOk returns a tuple with the MaxNegativeTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetMaxNegativeTtlOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.MaxNegativeTtl) {
		return nil, false
	}
	return o.MaxNegativeTtl, true
}

// HasMaxNegativeTtl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasMaxNegativeTtl() bool {
	if o != nil && !IsNil(o.MaxNegativeTtl) {
		return true
	}

	return false
}

// SetMaxNegativeTtl gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the MaxNegativeTtl field.
func (o *ConfigViewInheritance) SetMaxNegativeTtl(v Inheritance2InheritedUInt32) {
	o.MaxNegativeTtl = &v
}

// GetMaxUdpSize returns the MaxUdpSize field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetMaxUdpSize() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.MaxUdpSize) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.MaxUdpSize
}

// GetMaxUdpSizeOk returns a tuple with the MaxUdpSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetMaxUdpSizeOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.MaxUdpSize) {
		return nil, false
	}
	return o.MaxUdpSize, true
}

// HasMaxUdpSize returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasMaxUdpSize() bool {
	if o != nil && !IsNil(o.MaxUdpSize) {
		return true
	}

	return false
}

// SetMaxUdpSize gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the MaxUdpSize field.
func (o *ConfigViewInheritance) SetMaxUdpSize(v Inheritance2InheritedUInt32) {
	o.MaxUdpSize = &v
}

// GetMinimalResponses returns the MinimalResponses field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetMinimalResponses() Inheritance2InheritedBool {
	if o == nil || IsNil(o.MinimalResponses) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.MinimalResponses
}

// GetMinimalResponsesOk returns a tuple with the MinimalResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetMinimalResponsesOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.MinimalResponses) {
		return nil, false
	}
	return o.MinimalResponses, true
}

// HasMinimalResponses returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasMinimalResponses() bool {
	if o != nil && !IsNil(o.MinimalResponses) {
		return true
	}

	return false
}

// SetMinimalResponses gets a reference to the given Inheritance2InheritedBool and assigns it to the MinimalResponses field.
func (o *ConfigViewInheritance) SetMinimalResponses(v Inheritance2InheritedBool) {
	o.MinimalResponses = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetNotify() Inheritance2InheritedBool {
	if o == nil || IsNil(o.Notify) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetNotifyOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given Inheritance2InheritedBool and assigns it to the Notify field.
func (o *ConfigViewInheritance) SetNotify(v Inheritance2InheritedBool) {
	o.Notify = &v
}

// GetQueryAcl returns the QueryAcl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetQueryAcl() ConfigInheritedACLItems {
	if o == nil || IsNil(o.QueryAcl) {
		var ret ConfigInheritedACLItems
		return ret
	}
	return *o.QueryAcl
}

// GetQueryAclOk returns a tuple with the QueryAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetQueryAclOk() (*ConfigInheritedACLItems, bool) {
	if o == nil || IsNil(o.QueryAcl) {
		return nil, false
	}
	return o.QueryAcl, true
}

// HasQueryAcl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasQueryAcl() bool {
	if o != nil && !IsNil(o.QueryAcl) {
		return true
	}

	return false
}

// SetQueryAcl gets a reference to the given ConfigInheritedACLItems and assigns it to the QueryAcl field.
func (o *ConfigViewInheritance) SetQueryAcl(v ConfigInheritedACLItems) {
	o.QueryAcl = &v
}

// GetRecursionAcl returns the RecursionAcl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetRecursionAcl() ConfigInheritedACLItems {
	if o == nil || IsNil(o.RecursionAcl) {
		var ret ConfigInheritedACLItems
		return ret
	}
	return *o.RecursionAcl
}

// GetRecursionAclOk returns a tuple with the RecursionAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetRecursionAclOk() (*ConfigInheritedACLItems, bool) {
	if o == nil || IsNil(o.RecursionAcl) {
		return nil, false
	}
	return o.RecursionAcl, true
}

// HasRecursionAcl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasRecursionAcl() bool {
	if o != nil && !IsNil(o.RecursionAcl) {
		return true
	}

	return false
}

// SetRecursionAcl gets a reference to the given ConfigInheritedACLItems and assigns it to the RecursionAcl field.
func (o *ConfigViewInheritance) SetRecursionAcl(v ConfigInheritedACLItems) {
	o.RecursionAcl = &v
}

// GetRecursionEnabled returns the RecursionEnabled field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetRecursionEnabled() Inheritance2InheritedBool {
	if o == nil || IsNil(o.RecursionEnabled) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.RecursionEnabled
}

// GetRecursionEnabledOk returns a tuple with the RecursionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetRecursionEnabledOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.RecursionEnabled) {
		return nil, false
	}
	return o.RecursionEnabled, true
}

// HasRecursionEnabled returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasRecursionEnabled() bool {
	if o != nil && !IsNil(o.RecursionEnabled) {
		return true
	}

	return false
}

// SetRecursionEnabled gets a reference to the given Inheritance2InheritedBool and assigns it to the RecursionEnabled field.
func (o *ConfigViewInheritance) SetRecursionEnabled(v Inheritance2InheritedBool) {
	o.RecursionEnabled = &v
}

// GetSortList returns the SortList field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetSortList() ConfigInheritedSortListItems {
	if o == nil || IsNil(o.SortList) {
		var ret ConfigInheritedSortListItems
		return ret
	}
	return *o.SortList
}

// GetSortListOk returns a tuple with the SortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetSortListOk() (*ConfigInheritedSortListItems, bool) {
	if o == nil || IsNil(o.SortList) {
		return nil, false
	}
	return o.SortList, true
}

// HasSortList returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasSortList() bool {
	if o != nil && !IsNil(o.SortList) {
		return true
	}

	return false
}

// SetSortList gets a reference to the given ConfigInheritedSortListItems and assigns it to the SortList field.
func (o *ConfigViewInheritance) SetSortList(v ConfigInheritedSortListItems) {
	o.SortList = &v
}

// GetSynthesizeAddressRecordsFromHttps returns the SynthesizeAddressRecordsFromHttps field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetSynthesizeAddressRecordsFromHttps() Inheritance2InheritedBool {
	if o == nil || IsNil(o.SynthesizeAddressRecordsFromHttps) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.SynthesizeAddressRecordsFromHttps
}

// GetSynthesizeAddressRecordsFromHttpsOk returns a tuple with the SynthesizeAddressRecordsFromHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetSynthesizeAddressRecordsFromHttpsOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.SynthesizeAddressRecordsFromHttps) {
		return nil, false
	}
	return o.SynthesizeAddressRecordsFromHttps, true
}

// HasSynthesizeAddressRecordsFromHttps returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasSynthesizeAddressRecordsFromHttps() bool {
	if o != nil && !IsNil(o.SynthesizeAddressRecordsFromHttps) {
		return true
	}

	return false
}

// SetSynthesizeAddressRecordsFromHttps gets a reference to the given Inheritance2InheritedBool and assigns it to the SynthesizeAddressRecordsFromHttps field.
func (o *ConfigViewInheritance) SetSynthesizeAddressRecordsFromHttps(v Inheritance2InheritedBool) {
	o.SynthesizeAddressRecordsFromHttps = &v
}

// GetTransferAcl returns the TransferAcl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetTransferAcl() ConfigInheritedACLItems {
	if o == nil || IsNil(o.TransferAcl) {
		var ret ConfigInheritedACLItems
		return ret
	}
	return *o.TransferAcl
}

// GetTransferAclOk returns a tuple with the TransferAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetTransferAclOk() (*ConfigInheritedACLItems, bool) {
	if o == nil || IsNil(o.TransferAcl) {
		return nil, false
	}
	return o.TransferAcl, true
}

// HasTransferAcl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasTransferAcl() bool {
	if o != nil && !IsNil(o.TransferAcl) {
		return true
	}

	return false
}

// SetTransferAcl gets a reference to the given ConfigInheritedACLItems and assigns it to the TransferAcl field.
func (o *ConfigViewInheritance) SetTransferAcl(v ConfigInheritedACLItems) {
	o.TransferAcl = &v
}

// GetUpdateAcl returns the UpdateAcl field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetUpdateAcl() ConfigInheritedACLItems {
	if o == nil || IsNil(o.UpdateAcl) {
		var ret ConfigInheritedACLItems
		return ret
	}
	return *o.UpdateAcl
}

// GetUpdateAclOk returns a tuple with the UpdateAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetUpdateAclOk() (*ConfigInheritedACLItems, bool) {
	if o == nil || IsNil(o.UpdateAcl) {
		return nil, false
	}
	return o.UpdateAcl, true
}

// HasUpdateAcl returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasUpdateAcl() bool {
	if o != nil && !IsNil(o.UpdateAcl) {
		return true
	}

	return false
}

// SetUpdateAcl gets a reference to the given ConfigInheritedACLItems and assigns it to the UpdateAcl field.
func (o *ConfigViewInheritance) SetUpdateAcl(v ConfigInheritedACLItems) {
	o.UpdateAcl = &v
}

// GetUseForwardersForSubzones returns the UseForwardersForSubzones field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetUseForwardersForSubzones() Inheritance2InheritedBool {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		var ret Inheritance2InheritedBool
		return ret
	}
	return *o.UseForwardersForSubzones
}

// GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetUseForwardersForSubzonesOk() (*Inheritance2InheritedBool, bool) {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		return nil, false
	}
	return o.UseForwardersForSubzones, true
}

// HasUseForwardersForSubzones returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasUseForwardersForSubzones() bool {
	if o != nil && !IsNil(o.UseForwardersForSubzones) {
		return true
	}

	return false
}

// SetUseForwardersForSubzones gets a reference to the given Inheritance2InheritedBool and assigns it to the UseForwardersForSubzones field.
func (o *ConfigViewInheritance) SetUseForwardersForSubzones(v Inheritance2InheritedBool) {
	o.UseForwardersForSubzones = &v
}

// GetZoneAuthority returns the ZoneAuthority field value if set, zero value otherwise.
func (o *ConfigViewInheritance) GetZoneAuthority() ConfigInheritedZoneAuthority {
	if o == nil || IsNil(o.ZoneAuthority) {
		var ret ConfigInheritedZoneAuthority
		return ret
	}
	return *o.ZoneAuthority
}

// GetZoneAuthorityOk returns a tuple with the ZoneAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigViewInheritance) GetZoneAuthorityOk() (*ConfigInheritedZoneAuthority, bool) {
	if o == nil || IsNil(o.ZoneAuthority) {
		return nil, false
	}
	return o.ZoneAuthority, true
}

// HasZoneAuthority returns a boolean if a field has been set.
func (o *ConfigViewInheritance) HasZoneAuthority() bool {
	if o != nil && !IsNil(o.ZoneAuthority) {
		return true
	}

	return false
}

// SetZoneAuthority gets a reference to the given ConfigInheritedZoneAuthority and assigns it to the ZoneAuthority field.
func (o *ConfigViewInheritance) SetZoneAuthority(v ConfigInheritedZoneAuthority) {
	o.ZoneAuthority = &v
}

func (o ConfigViewInheritance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigViewInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddEdnsOptionInOutgoingQuery) {
		toSerialize["add_edns_option_in_outgoing_query"] = o.AddEdnsOptionInOutgoingQuery
	}
	if !IsNil(o.CustomRootNsBlock) {
		toSerialize["custom_root_ns_block"] = o.CustomRootNsBlock
	}
	if !IsNil(o.DnssecValidationBlock) {
		toSerialize["dnssec_validation_block"] = o.DnssecValidationBlock
	}
	if !IsNil(o.EcsBlock) {
		toSerialize["ecs_block"] = o.EcsBlock
	}
	if !IsNil(o.EdnsUdpSize) {
		toSerialize["edns_udp_size"] = o.EdnsUdpSize
	}
	if !IsNil(o.FilterAaaaAcl) {
		toSerialize["filter_aaaa_acl"] = o.FilterAaaaAcl
	}
	if !IsNil(o.FilterAaaaOnV4) {
		toSerialize["filter_aaaa_on_v4"] = o.FilterAaaaOnV4
	}
	if !IsNil(o.ForwardersBlock) {
		toSerialize["forwarders_block"] = o.ForwardersBlock
	}
	if !IsNil(o.GssTsigEnabled) {
		toSerialize["gss_tsig_enabled"] = o.GssTsigEnabled
	}
	if !IsNil(o.LameTtl) {
		toSerialize["lame_ttl"] = o.LameTtl
	}
	if !IsNil(o.MatchRecursiveOnly) {
		toSerialize["match_recursive_only"] = o.MatchRecursiveOnly
	}
	if !IsNil(o.MaxCacheTtl) {
		toSerialize["max_cache_ttl"] = o.MaxCacheTtl
	}
	if !IsNil(o.MaxNegativeTtl) {
		toSerialize["max_negative_ttl"] = o.MaxNegativeTtl
	}
	if !IsNil(o.MaxUdpSize) {
		toSerialize["max_udp_size"] = o.MaxUdpSize
	}
	if !IsNil(o.MinimalResponses) {
		toSerialize["minimal_responses"] = o.MinimalResponses
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.QueryAcl) {
		toSerialize["query_acl"] = o.QueryAcl
	}
	if !IsNil(o.RecursionAcl) {
		toSerialize["recursion_acl"] = o.RecursionAcl
	}
	if !IsNil(o.RecursionEnabled) {
		toSerialize["recursion_enabled"] = o.RecursionEnabled
	}
	if !IsNil(o.SortList) {
		toSerialize["sort_list"] = o.SortList
	}
	if !IsNil(o.SynthesizeAddressRecordsFromHttps) {
		toSerialize["synthesize_address_records_from_https"] = o.SynthesizeAddressRecordsFromHttps
	}
	if !IsNil(o.TransferAcl) {
		toSerialize["transfer_acl"] = o.TransferAcl
	}
	if !IsNil(o.UpdateAcl) {
		toSerialize["update_acl"] = o.UpdateAcl
	}
	if !IsNil(o.UseForwardersForSubzones) {
		toSerialize["use_forwarders_for_subzones"] = o.UseForwardersForSubzones
	}
	if !IsNil(o.ZoneAuthority) {
		toSerialize["zone_authority"] = o.ZoneAuthority
	}
	return toSerialize, nil
}

type NullableConfigViewInheritance struct {
	value *ConfigViewInheritance
	isSet bool
}

func (v NullableConfigViewInheritance) Get() *ConfigViewInheritance {
	return v.value
}

func (v *NullableConfigViewInheritance) Set(val *ConfigViewInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigViewInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigViewInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigViewInheritance(val *ConfigViewInheritance) *NullableConfigViewInheritance {
	return &NullableConfigViewInheritance{value: val, isSet: true}
}

func (v NullableConfigViewInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigViewInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
