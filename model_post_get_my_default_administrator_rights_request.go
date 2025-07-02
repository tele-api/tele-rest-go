/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the PostGetMyDefaultAdministratorRightsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGetMyDefaultAdministratorRightsRequest{}

// PostGetMyDefaultAdministratorRightsRequest struct for PostGetMyDefaultAdministratorRightsRequest
type PostGetMyDefaultAdministratorRightsRequest struct {
	// Pass *True* to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels *bool `json:"for_channels,omitempty"`
}

// NewPostGetMyDefaultAdministratorRightsRequest instantiates a new PostGetMyDefaultAdministratorRightsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGetMyDefaultAdministratorRightsRequest() *PostGetMyDefaultAdministratorRightsRequest {
	this := PostGetMyDefaultAdministratorRightsRequest{}
	return &this
}

// NewPostGetMyDefaultAdministratorRightsRequestWithDefaults instantiates a new PostGetMyDefaultAdministratorRightsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGetMyDefaultAdministratorRightsRequestWithDefaults() *PostGetMyDefaultAdministratorRightsRequest {
	this := PostGetMyDefaultAdministratorRightsRequest{}
	return &this
}

// GetForChannels returns the ForChannels field value if set, zero value otherwise.
func (o *PostGetMyDefaultAdministratorRightsRequest) GetForChannels() bool {
	if o == nil || IsNil(o.ForChannels) {
		var ret bool
		return ret
	}
	return *o.ForChannels
}

// GetForChannelsOk returns a tuple with the ForChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostGetMyDefaultAdministratorRightsRequest) GetForChannelsOk() (*bool, bool) {
	if o == nil || IsNil(o.ForChannels) {
		return nil, false
	}
	return o.ForChannels, true
}

// HasForChannels returns a boolean if a field has been set.
func (o *PostGetMyDefaultAdministratorRightsRequest) HasForChannels() bool {
	if o != nil && !IsNil(o.ForChannels) {
		return true
	}

	return false
}

// SetForChannels gets a reference to the given bool and assigns it to the ForChannels field.
func (o *PostGetMyDefaultAdministratorRightsRequest) SetForChannels(v bool) {
	o.ForChannels = &v
}


func (o PostGetMyDefaultAdministratorRightsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGetMyDefaultAdministratorRightsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForChannels) {
		toSerialize["for_channels"] = o.ForChannels
	}
	return toSerialize, nil
}

type NullablePostGetMyDefaultAdministratorRightsRequest struct {
	value *PostGetMyDefaultAdministratorRightsRequest
	isSet bool
}

func (v NullablePostGetMyDefaultAdministratorRightsRequest) Get() *PostGetMyDefaultAdministratorRightsRequest {
	return v.value
}

func (v *NullablePostGetMyDefaultAdministratorRightsRequest) Set(val *PostGetMyDefaultAdministratorRightsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGetMyDefaultAdministratorRightsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGetMyDefaultAdministratorRightsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGetMyDefaultAdministratorRightsRequest(val *PostGetMyDefaultAdministratorRightsRequest) *NullablePostGetMyDefaultAdministratorRightsRequest {
	return &NullablePostGetMyDefaultAdministratorRightsRequest{value: val, isSet: true}
}

func (v NullablePostGetMyDefaultAdministratorRightsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGetMyDefaultAdministratorRightsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


