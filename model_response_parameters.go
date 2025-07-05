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
)

// checks if the ResponseParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseParameters{}

// ResponseParameters Describes why a request was unsuccessful.
type ResponseParameters struct {
	// *Optional*. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId *int32 `json:"migrate_to_chat_id,omitempty"`
	// *Optional*. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter *int32 `json:"retry_after,omitempty"`
}

// NewResponseParameters instantiates a new ResponseParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseParameters() *ResponseParameters {
	this := ResponseParameters{}
	return &this
}

// NewResponseParametersWithDefaults instantiates a new ResponseParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseParametersWithDefaults() *ResponseParameters {
	this := ResponseParameters{}
	return &this
}

// GetMigrateToChatId returns the MigrateToChatId field value if set, zero value otherwise.
func (o *ResponseParameters) GetMigrateToChatId() int32 {
	if o == nil || IsNil(o.MigrateToChatId) {
		var ret int32
		return ret
	}
	return *o.MigrateToChatId
}

// GetMigrateToChatIdOk returns a tuple with the MigrateToChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParameters) GetMigrateToChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MigrateToChatId) {
		return nil, false
	}
	return o.MigrateToChatId, true
}

// HasMigrateToChatId returns a boolean if a field has been set.
func (o *ResponseParameters) HasMigrateToChatId() bool {
	if o != nil && !IsNil(o.MigrateToChatId) {
		return true
	}

	return false
}

// SetMigrateToChatId gets a reference to the given int32 and assigns it to the MigrateToChatId field.
func (o *ResponseParameters) SetMigrateToChatId(v int32) {
	o.MigrateToChatId = &v
}


// GetRetryAfter returns the RetryAfter field value if set, zero value otherwise.
func (o *ResponseParameters) GetRetryAfter() int32 {
	if o == nil || IsNil(o.RetryAfter) {
		var ret int32
		return ret
	}
	return *o.RetryAfter
}

// GetRetryAfterOk returns a tuple with the RetryAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseParameters) GetRetryAfterOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryAfter) {
		return nil, false
	}
	return o.RetryAfter, true
}

// HasRetryAfter returns a boolean if a field has been set.
func (o *ResponseParameters) HasRetryAfter() bool {
	if o != nil && !IsNil(o.RetryAfter) {
		return true
	}

	return false
}

// SetRetryAfter gets a reference to the given int32 and assigns it to the RetryAfter field.
func (o *ResponseParameters) SetRetryAfter(v int32) {
	o.RetryAfter = &v
}


func (o ResponseParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MigrateToChatId) {
		toSerialize["migrate_to_chat_id"] = o.MigrateToChatId
	}
	if !IsNil(o.RetryAfter) {
		toSerialize["retry_after"] = o.RetryAfter
	}
	return toSerialize, nil
}

type NullableResponseParameters struct {
	value *ResponseParameters
	isSet bool
}

func (v NullableResponseParameters) Get() *ResponseParameters {
	return v.value
}

func (v *NullableResponseParameters) Set(val *ResponseParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseParameters(val *ResponseParameters) *NullableResponseParameters {
	return &NullableResponseParameters{value: val, isSet: true}
}

func (v NullableResponseParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


