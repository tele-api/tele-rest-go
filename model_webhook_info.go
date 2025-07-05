/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Telegram Bot API service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package tele_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookInfo{}

// WebhookInfo Describes the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`
	// *True*, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int32 `json:"pending_update_count"`
	// *Optional*. Currently used webhook IP address
	IpAddress *string `json:"ip_address,omitempty"`
	// *Optional*. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate *int32 `json:"last_error_date,omitempty"`
	// *Optional*. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage *string `json:"last_error_message,omitempty"`
	// *Optional*. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate *int32 `json:"last_synchronization_error_date,omitempty"`
	// *Optional*. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// *Optional*. A list of update types the bot is subscribed to. Defaults to all update types except *chat\\_member*
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type _WebhookInfo WebhookInfo

// NewWebhookInfo instantiates a new WebhookInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookInfo(url string, hasCustomCertificate bool, pendingUpdateCount int32) *WebhookInfo {
	this := WebhookInfo{}
	this.Url = url
	this.HasCustomCertificate = hasCustomCertificate
	this.PendingUpdateCount = pendingUpdateCount
	return &this
}

// NewWebhookInfoWithDefaults instantiates a new WebhookInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookInfoWithDefaults() *WebhookInfo {
	this := WebhookInfo{}
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookInfo) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookInfo) SetUrl(v string) {
	o.Url = v
}

// GetHasCustomCertificate returns the HasCustomCertificate field value
func (o *WebhookInfo) GetHasCustomCertificate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasCustomCertificate
}

// GetHasCustomCertificateOk returns a tuple with the HasCustomCertificate field value
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetHasCustomCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasCustomCertificate, true
}

// SetHasCustomCertificate sets field value
func (o *WebhookInfo) SetHasCustomCertificate(v bool) {
	o.HasCustomCertificate = v
}

// GetPendingUpdateCount returns the PendingUpdateCount field value
func (o *WebhookInfo) GetPendingUpdateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PendingUpdateCount
}

// GetPendingUpdateCountOk returns a tuple with the PendingUpdateCount field value
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetPendingUpdateCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUpdateCount, true
}

// SetPendingUpdateCount sets field value
func (o *WebhookInfo) SetPendingUpdateCount(v int32) {
	o.PendingUpdateCount = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *WebhookInfo) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *WebhookInfo) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *WebhookInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}


// GetLastErrorDate returns the LastErrorDate field value if set, zero value otherwise.
func (o *WebhookInfo) GetLastErrorDate() int32 {
	if o == nil || IsNil(o.LastErrorDate) {
		var ret int32
		return ret
	}
	return *o.LastErrorDate
}

// GetLastErrorDateOk returns a tuple with the LastErrorDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetLastErrorDateOk() (*int32, bool) {
	if o == nil || IsNil(o.LastErrorDate) {
		return nil, false
	}
	return o.LastErrorDate, true
}

// HasLastErrorDate returns a boolean if a field has been set.
func (o *WebhookInfo) HasLastErrorDate() bool {
	if o != nil && !IsNil(o.LastErrorDate) {
		return true
	}

	return false
}

// SetLastErrorDate gets a reference to the given int32 and assigns it to the LastErrorDate field.
func (o *WebhookInfo) SetLastErrorDate(v int32) {
	o.LastErrorDate = &v
}


// GetLastErrorMessage returns the LastErrorMessage field value if set, zero value otherwise.
func (o *WebhookInfo) GetLastErrorMessage() string {
	if o == nil || IsNil(o.LastErrorMessage) {
		var ret string
		return ret
	}
	return *o.LastErrorMessage
}

// GetLastErrorMessageOk returns a tuple with the LastErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetLastErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LastErrorMessage) {
		return nil, false
	}
	return o.LastErrorMessage, true
}

// HasLastErrorMessage returns a boolean if a field has been set.
func (o *WebhookInfo) HasLastErrorMessage() bool {
	if o != nil && !IsNil(o.LastErrorMessage) {
		return true
	}

	return false
}

// SetLastErrorMessage gets a reference to the given string and assigns it to the LastErrorMessage field.
func (o *WebhookInfo) SetLastErrorMessage(v string) {
	o.LastErrorMessage = &v
}


// GetLastSynchronizationErrorDate returns the LastSynchronizationErrorDate field value if set, zero value otherwise.
func (o *WebhookInfo) GetLastSynchronizationErrorDate() int32 {
	if o == nil || IsNil(o.LastSynchronizationErrorDate) {
		var ret int32
		return ret
	}
	return *o.LastSynchronizationErrorDate
}

// GetLastSynchronizationErrorDateOk returns a tuple with the LastSynchronizationErrorDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetLastSynchronizationErrorDateOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSynchronizationErrorDate) {
		return nil, false
	}
	return o.LastSynchronizationErrorDate, true
}

// HasLastSynchronizationErrorDate returns a boolean if a field has been set.
func (o *WebhookInfo) HasLastSynchronizationErrorDate() bool {
	if o != nil && !IsNil(o.LastSynchronizationErrorDate) {
		return true
	}

	return false
}

// SetLastSynchronizationErrorDate gets a reference to the given int32 and assigns it to the LastSynchronizationErrorDate field.
func (o *WebhookInfo) SetLastSynchronizationErrorDate(v int32) {
	o.LastSynchronizationErrorDate = &v
}


// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *WebhookInfo) GetMaxConnections() int32 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *WebhookInfo) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *WebhookInfo) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}


// GetAllowedUpdates returns the AllowedUpdates field value if set, zero value otherwise.
func (o *WebhookInfo) GetAllowedUpdates() []string {
	if o == nil || IsNil(o.AllowedUpdates) {
		var ret []string
		return ret
	}
	return o.AllowedUpdates
}

// GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookInfo) GetAllowedUpdatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedUpdates) {
		return nil, false
	}
	return o.AllowedUpdates, true
}

// HasAllowedUpdates returns a boolean if a field has been set.
func (o *WebhookInfo) HasAllowedUpdates() bool {
	if o != nil && !IsNil(o.AllowedUpdates) {
		return true
	}

	return false
}

// SetAllowedUpdates gets a reference to the given []string and assigns it to the AllowedUpdates field.
func (o *WebhookInfo) SetAllowedUpdates(v []string) {
	o.AllowedUpdates = v
}


func (o WebhookInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["has_custom_certificate"] = o.HasCustomCertificate
	toSerialize["pending_update_count"] = o.PendingUpdateCount
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.LastErrorDate) {
		toSerialize["last_error_date"] = o.LastErrorDate
	}
	if !IsNil(o.LastErrorMessage) {
		toSerialize["last_error_message"] = o.LastErrorMessage
	}
	if !IsNil(o.LastSynchronizationErrorDate) {
		toSerialize["last_synchronization_error_date"] = o.LastSynchronizationErrorDate
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["max_connections"] = o.MaxConnections
	}
	if !IsNil(o.AllowedUpdates) {
		toSerialize["allowed_updates"] = o.AllowedUpdates
	}
	return toSerialize, nil
}

func (o *WebhookInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"has_custom_certificate",
		"pending_update_count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookInfo := _WebhookInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookInfo)

	if err != nil {
		return err
	}

	*o = WebhookInfo(varWebhookInfo)

	return err
}

type NullableWebhookInfo struct {
	value *WebhookInfo
	isSet bool
}

func (v NullableWebhookInfo) Get() *WebhookInfo {
	return v.value
}

func (v *NullableWebhookInfo) Set(val *WebhookInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookInfo(val *WebhookInfo) *NullableWebhookInfo {
	return &NullableWebhookInfo{value: val, isSet: true}
}

func (v NullableWebhookInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


