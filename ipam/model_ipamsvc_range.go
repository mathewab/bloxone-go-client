/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"time"
)

// checks if the IpamsvcRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcRange{}

// IpamsvcRange A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries.
type IpamsvcRange struct {
	// The description for the range. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// The resource identifier.
	DhcpHost *string `json:"dhcp_host,omitempty"`
	// The list of DHCP options. May be either a specific option or a group of options.
	DhcpOptions []IpamsvcOptionItem `json:"dhcp_options,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.
	DisableDhcp *bool `json:"disable_dhcp,omitempty"`
	// The end IP address of the range.
	End string `json:"end"`
	// The list of all exclusion ranges in the scope of the range.
	ExclusionRanges []IpamsvcExclusionRange `json:"exclusion_ranges,omitempty"`
	// The list of all allow/deny filters of the range.
	Filters []IpamsvcAccessFilter `json:"filters,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The list of the inheritance assigned hosts of the object.
	InheritanceAssignedHosts []InheritanceAssignedHost `json:"inheritance_assigned_hosts,omitempty"`
	// The resource identifier.
	InheritanceParent  *string                        `json:"inheritance_parent,omitempty"`
	InheritanceSources *IpamsvcDHCPOptionsInheritance `json:"inheritance_sources,omitempty"`
	// The name of the range. May contain 1 to 256 characters. Can include UTF-8.
	Name *string `json:"name,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// The type of protocol (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The resource identifier.
	Space string `json:"space"`
	// The start IP address of the range.
	Start string `json:"start"`
	// The tags for the range in JSON format.
	Tags      map[string]interface{}       `json:"tags,omitempty"`
	Threshold *IpamsvcUtilizationThreshold `json:"threshold,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt     NullableTime          `json:"updated_at,omitempty"`
	Utilization   *IpamsvcUtilization   `json:"utilization,omitempty"`
	UtilizationV6 *IpamsvcUtilizationV6 `json:"utilization_v6,omitempty"`
}

// NewIpamsvcRange instantiates a new IpamsvcRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcRange(end string, space string, start string) *IpamsvcRange {
	this := IpamsvcRange{}
	this.End = end
	this.Space = space
	this.Start = start
	return &this
}

// NewIpamsvcRangeWithDefaults instantiates a new IpamsvcRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcRangeWithDefaults() *IpamsvcRange {
	this := IpamsvcRange{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IpamsvcRange) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IpamsvcRange) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IpamsvcRange) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpamsvcRange) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpamsvcRange) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IpamsvcRange) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *IpamsvcRange) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *IpamsvcRange) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *IpamsvcRange) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDhcpHost returns the DhcpHost field value if set, zero value otherwise.
func (o *IpamsvcRange) GetDhcpHost() string {
	if o == nil || IsNil(o.DhcpHost) {
		var ret string
		return ret
	}
	return *o.DhcpHost
}

// GetDhcpHostOk returns a tuple with the DhcpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetDhcpHostOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpHost) {
		return nil, false
	}
	return o.DhcpHost, true
}

// HasDhcpHost returns a boolean if a field has been set.
func (o *IpamsvcRange) HasDhcpHost() bool {
	if o != nil && !IsNil(o.DhcpHost) {
		return true
	}

	return false
}

// SetDhcpHost gets a reference to the given string and assigns it to the DhcpHost field.
func (o *IpamsvcRange) SetDhcpHost(v string) {
	o.DhcpHost = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *IpamsvcRange) GetDhcpOptions() []IpamsvcOptionItem {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []IpamsvcOptionItem
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetDhcpOptionsOk() ([]IpamsvcOptionItem, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *IpamsvcRange) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []IpamsvcOptionItem and assigns it to the DhcpOptions field.
func (o *IpamsvcRange) SetDhcpOptions(v []IpamsvcOptionItem) {
	o.DhcpOptions = v
}

// GetDisableDhcp returns the DisableDhcp field value if set, zero value otherwise.
func (o *IpamsvcRange) GetDisableDhcp() bool {
	if o == nil || IsNil(o.DisableDhcp) {
		var ret bool
		return ret
	}
	return *o.DisableDhcp
}

// GetDisableDhcpOk returns a tuple with the DisableDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetDisableDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDhcp) {
		return nil, false
	}
	return o.DisableDhcp, true
}

// HasDisableDhcp returns a boolean if a field has been set.
func (o *IpamsvcRange) HasDisableDhcp() bool {
	if o != nil && !IsNil(o.DisableDhcp) {
		return true
	}

	return false
}

// SetDisableDhcp gets a reference to the given bool and assigns it to the DisableDhcp field.
func (o *IpamsvcRange) SetDisableDhcp(v bool) {
	o.DisableDhcp = &v
}

// GetEnd returns the End field value
func (o *IpamsvcRange) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *IpamsvcRange) SetEnd(v string) {
	o.End = v
}

// GetExclusionRanges returns the ExclusionRanges field value if set, zero value otherwise.
func (o *IpamsvcRange) GetExclusionRanges() []IpamsvcExclusionRange {
	if o == nil || IsNil(o.ExclusionRanges) {
		var ret []IpamsvcExclusionRange
		return ret
	}
	return o.ExclusionRanges
}

// GetExclusionRangesOk returns a tuple with the ExclusionRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetExclusionRangesOk() ([]IpamsvcExclusionRange, bool) {
	if o == nil || IsNil(o.ExclusionRanges) {
		return nil, false
	}
	return o.ExclusionRanges, true
}

// HasExclusionRanges returns a boolean if a field has been set.
func (o *IpamsvcRange) HasExclusionRanges() bool {
	if o != nil && !IsNil(o.ExclusionRanges) {
		return true
	}

	return false
}

// SetExclusionRanges gets a reference to the given []IpamsvcExclusionRange and assigns it to the ExclusionRanges field.
func (o *IpamsvcRange) SetExclusionRanges(v []IpamsvcExclusionRange) {
	o.ExclusionRanges = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *IpamsvcRange) GetFilters() []IpamsvcAccessFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []IpamsvcAccessFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetFiltersOk() ([]IpamsvcAccessFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *IpamsvcRange) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []IpamsvcAccessFilter and assigns it to the Filters field.
func (o *IpamsvcRange) SetFilters(v []IpamsvcAccessFilter) {
	o.Filters = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IpamsvcRange) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IpamsvcRange) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IpamsvcRange) SetId(v string) {
	o.Id = &v
}

// GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field value if set, zero value otherwise.
func (o *IpamsvcRange) GetInheritanceAssignedHosts() []InheritanceAssignedHost {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		var ret []InheritanceAssignedHost
		return ret
	}
	return o.InheritanceAssignedHosts
}

// GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetInheritanceAssignedHostsOk() ([]InheritanceAssignedHost, bool) {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		return nil, false
	}
	return o.InheritanceAssignedHosts, true
}

// HasInheritanceAssignedHosts returns a boolean if a field has been set.
func (o *IpamsvcRange) HasInheritanceAssignedHosts() bool {
	if o != nil && !IsNil(o.InheritanceAssignedHosts) {
		return true
	}

	return false
}

// SetInheritanceAssignedHosts gets a reference to the given []InheritanceAssignedHost and assigns it to the InheritanceAssignedHosts field.
func (o *IpamsvcRange) SetInheritanceAssignedHosts(v []InheritanceAssignedHost) {
	o.InheritanceAssignedHosts = v
}

// GetInheritanceParent returns the InheritanceParent field value if set, zero value otherwise.
func (o *IpamsvcRange) GetInheritanceParent() string {
	if o == nil || IsNil(o.InheritanceParent) {
		var ret string
		return ret
	}
	return *o.InheritanceParent
}

// GetInheritanceParentOk returns a tuple with the InheritanceParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetInheritanceParentOk() (*string, bool) {
	if o == nil || IsNil(o.InheritanceParent) {
		return nil, false
	}
	return o.InheritanceParent, true
}

// HasInheritanceParent returns a boolean if a field has been set.
func (o *IpamsvcRange) HasInheritanceParent() bool {
	if o != nil && !IsNil(o.InheritanceParent) {
		return true
	}

	return false
}

// SetInheritanceParent gets a reference to the given string and assigns it to the InheritanceParent field.
func (o *IpamsvcRange) SetInheritanceParent(v string) {
	o.InheritanceParent = &v
}

// GetInheritanceSources returns the InheritanceSources field value if set, zero value otherwise.
func (o *IpamsvcRange) GetInheritanceSources() IpamsvcDHCPOptionsInheritance {
	if o == nil || IsNil(o.InheritanceSources) {
		var ret IpamsvcDHCPOptionsInheritance
		return ret
	}
	return *o.InheritanceSources
}

// GetInheritanceSourcesOk returns a tuple with the InheritanceSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetInheritanceSourcesOk() (*IpamsvcDHCPOptionsInheritance, bool) {
	if o == nil || IsNil(o.InheritanceSources) {
		return nil, false
	}
	return o.InheritanceSources, true
}

// HasInheritanceSources returns a boolean if a field has been set.
func (o *IpamsvcRange) HasInheritanceSources() bool {
	if o != nil && !IsNil(o.InheritanceSources) {
		return true
	}

	return false
}

// SetInheritanceSources gets a reference to the given IpamsvcDHCPOptionsInheritance and assigns it to the InheritanceSources field.
func (o *IpamsvcRange) SetInheritanceSources(v IpamsvcDHCPOptionsInheritance) {
	o.InheritanceSources = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IpamsvcRange) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IpamsvcRange) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IpamsvcRange) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *IpamsvcRange) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *IpamsvcRange) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *IpamsvcRange) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IpamsvcRange) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IpamsvcRange) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *IpamsvcRange) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSpace returns the Space field value
func (o *IpamsvcRange) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *IpamsvcRange) SetSpace(v string) {
	o.Space = v
}

// GetStart returns the Start field value
func (o *IpamsvcRange) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *IpamsvcRange) SetStart(v string) {
	o.Start = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpamsvcRange) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpamsvcRange) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *IpamsvcRange) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *IpamsvcRange) GetThreshold() IpamsvcUtilizationThreshold {
	if o == nil || IsNil(o.Threshold) {
		var ret IpamsvcUtilizationThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetThresholdOk() (*IpamsvcUtilizationThreshold, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *IpamsvcRange) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given IpamsvcUtilizationThreshold and assigns it to the Threshold field.
func (o *IpamsvcRange) SetThreshold(v IpamsvcUtilizationThreshold) {
	o.Threshold = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpamsvcRange) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpamsvcRange) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IpamsvcRange) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *IpamsvcRange) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *IpamsvcRange) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *IpamsvcRange) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *IpamsvcRange) GetUtilization() IpamsvcUtilization {
	if o == nil || IsNil(o.Utilization) {
		var ret IpamsvcUtilization
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetUtilizationOk() (*IpamsvcUtilization, bool) {
	if o == nil || IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *IpamsvcRange) HasUtilization() bool {
	if o != nil && !IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given IpamsvcUtilization and assigns it to the Utilization field.
func (o *IpamsvcRange) SetUtilization(v IpamsvcUtilization) {
	o.Utilization = &v
}

// GetUtilizationV6 returns the UtilizationV6 field value if set, zero value otherwise.
func (o *IpamsvcRange) GetUtilizationV6() IpamsvcUtilizationV6 {
	if o == nil || IsNil(o.UtilizationV6) {
		var ret IpamsvcUtilizationV6
		return ret
	}
	return *o.UtilizationV6
}

// GetUtilizationV6Ok returns a tuple with the UtilizationV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcRange) GetUtilizationV6Ok() (*IpamsvcUtilizationV6, bool) {
	if o == nil || IsNil(o.UtilizationV6) {
		return nil, false
	}
	return o.UtilizationV6, true
}

// HasUtilizationV6 returns a boolean if a field has been set.
func (o *IpamsvcRange) HasUtilizationV6() bool {
	if o != nil && !IsNil(o.UtilizationV6) {
		return true
	}

	return false
}

// SetUtilizationV6 gets a reference to the given IpamsvcUtilizationV6 and assigns it to the UtilizationV6 field.
func (o *IpamsvcRange) SetUtilizationV6(v IpamsvcUtilizationV6) {
	o.UtilizationV6 = &v
}

func (o IpamsvcRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if !IsNil(o.DhcpHost) {
		toSerialize["dhcp_host"] = o.DhcpHost
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.DisableDhcp) {
		toSerialize["disable_dhcp"] = o.DisableDhcp
	}
	toSerialize["end"] = o.End
	if !IsNil(o.ExclusionRanges) {
		toSerialize["exclusion_ranges"] = o.ExclusionRanges
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InheritanceAssignedHosts) {
		toSerialize["inheritance_assigned_hosts"] = o.InheritanceAssignedHosts
	}
	if !IsNil(o.InheritanceParent) {
		toSerialize["inheritance_parent"] = o.InheritanceParent
	}
	if !IsNil(o.InheritanceSources) {
		toSerialize["inheritance_sources"] = o.InheritanceSources
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["space"] = o.Space
	toSerialize["start"] = o.Start
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	if !IsNil(o.UtilizationV6) {
		toSerialize["utilization_v6"] = o.UtilizationV6
	}
	return toSerialize, nil
}

type NullableIpamsvcRange struct {
	value *IpamsvcRange
	isSet bool
}

func (v NullableIpamsvcRange) Get() *IpamsvcRange {
	return v.value
}

func (v *NullableIpamsvcRange) Set(val *IpamsvcRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcRange(val *IpamsvcRange) *NullableIpamsvcRange {
	return &NullableIpamsvcRange{value: val, isSet: true}
}

func (v NullableIpamsvcRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
