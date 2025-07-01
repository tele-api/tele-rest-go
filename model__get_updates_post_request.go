/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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
)

// checks if the GetUpdatesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUpdatesPostRequest{}

// GetUpdatesPostRequest struct for GetUpdatesPostRequest
type GetUpdatesPostRequest struct {
	// Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as [getUpdates](https://core.telegram.org/bots/api/#getupdates) is called with an *offset* higher than its *update\\_id*. The negative offset can be specified to retrieve updates starting from *-offset* update from the end of the updates queue. All previous updates will be forgotten.
	Offset *int32 `json:"offset,omitempty"`
	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int32 `json:"limit,omitempty"`
	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout *int32 `json:"timeout,omitempty"`
	// A JSON-serialized list of the update types you want your bot to receive. For example, specify `[\"message\", \"edited_channel_post\", \"callback_query\"]` to only receive updates of these types. See [Update](https://core.telegram.org/bots/api/#update) for a complete list of available update types. Specify an empty list to receive all update types except *chat\\_member*, *message\\_reaction*, and *message\\_reaction\\_count* (default). If not specified, the previous setting will be used.    Please note that this parameter doesn't affect updates created before the call to getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// NewGetUpdatesPostRequest instantiates a new GetUpdatesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUpdatesPostRequest() *GetUpdatesPostRequest {
	this := GetUpdatesPostRequest{}
	var limit int32 = 100
	this.Limit = &limit
	var timeout int32 = 0
	this.Timeout = &timeout
	return &this
}

// NewGetUpdatesPostRequestWithDefaults instantiates a new GetUpdatesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUpdatesPostRequestWithDefaults() *GetUpdatesPostRequest {
	this := GetUpdatesPostRequest{}
	var limit int32 = 100
	this.Limit = &limit
	var timeout int32 = 0
	this.Timeout = &timeout
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetUpdatesPostRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdatesPostRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetUpdatesPostRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetUpdatesPostRequest) SetOffset(v int32) {
	o.Offset = &v
}


// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetUpdatesPostRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdatesPostRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetUpdatesPostRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetUpdatesPostRequest) SetLimit(v int32) {
	o.Limit = &v
}


// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetUpdatesPostRequest) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdatesPostRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetUpdatesPostRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *GetUpdatesPostRequest) SetTimeout(v int32) {
	o.Timeout = &v
}


// GetAllowedUpdates returns the AllowedUpdates field value if set, zero value otherwise.
func (o *GetUpdatesPostRequest) GetAllowedUpdates() []string {
	if o == nil || IsNil(o.AllowedUpdates) {
		var ret []string
		return ret
	}
	return o.AllowedUpdates
}

// GetAllowedUpdatesOk returns a tuple with the AllowedUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdatesPostRequest) GetAllowedUpdatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedUpdates) {
		return nil, false
	}
	return o.AllowedUpdates, true
}

// HasAllowedUpdates returns a boolean if a field has been set.
func (o *GetUpdatesPostRequest) HasAllowedUpdates() bool {
	if o != nil && !IsNil(o.AllowedUpdates) {
		return true
	}

	return false
}

// SetAllowedUpdates gets a reference to the given []string and assigns it to the AllowedUpdates field.
func (o *GetUpdatesPostRequest) SetAllowedUpdates(v []string) {
	o.AllowedUpdates = v
}


func (o GetUpdatesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUpdatesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.AllowedUpdates) {
		toSerialize["allowed_updates"] = o.AllowedUpdates
	}
	return toSerialize, nil
}

type NullableGetUpdatesPostRequest struct {
	value *GetUpdatesPostRequest
	isSet bool
}

func (v NullableGetUpdatesPostRequest) Get() *GetUpdatesPostRequest {
	return v.value
}

func (v *NullableGetUpdatesPostRequest) Set(val *GetUpdatesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUpdatesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUpdatesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUpdatesPostRequest(val *GetUpdatesPostRequest) *NullableGetUpdatesPostRequest {
	return &NullableGetUpdatesPostRequest{value: val, isSet: true}
}

func (v NullableGetUpdatesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUpdatesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


