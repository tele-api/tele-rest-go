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

// checks if the GetMyCommandsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMyCommandsRequest{}

// GetMyCommandsRequest Request parameters for getMyCommands
type GetMyCommandsRequest struct {
	Scope *BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
}

// NewGetMyCommandsRequest instantiates a new GetMyCommandsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMyCommandsRequest() *GetMyCommandsRequest {
	this := GetMyCommandsRequest{}
	return &this
}

// NewGetMyCommandsRequestWithDefaults instantiates a new GetMyCommandsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMyCommandsRequestWithDefaults() *GetMyCommandsRequest {
	this := GetMyCommandsRequest{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetMyCommandsRequest) GetScope() BotCommandScope {
	if o == nil || IsNil(o.Scope) {
		var ret BotCommandScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyCommandsRequest) GetScopeOk() (*BotCommandScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetMyCommandsRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given BotCommandScope and assigns it to the Scope field.
func (o *GetMyCommandsRequest) SetScope(v BotCommandScope) {
	o.Scope = &v
}


// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *GetMyCommandsRequest) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyCommandsRequest) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *GetMyCommandsRequest) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *GetMyCommandsRequest) SetLanguageCode(v string) {
	o.LanguageCode = &v
}


func (o GetMyCommandsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMyCommandsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["language_code"] = o.LanguageCode
	}
	return toSerialize, nil
}

type NullableGetMyCommandsRequest struct {
	value *GetMyCommandsRequest
	isSet bool
}

func (v NullableGetMyCommandsRequest) Get() *GetMyCommandsRequest {
	return v.value
}

func (v *NullableGetMyCommandsRequest) Set(val *GetMyCommandsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMyCommandsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMyCommandsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMyCommandsRequest(val *GetMyCommandsRequest) *NullableGetMyCommandsRequest {
	return &NullableGetMyCommandsRequest{value: val, isSet: true}
}

func (v NullableGetMyCommandsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMyCommandsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


