/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the SetChatStickerSetPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetChatStickerSetPostRequest{}

// SetChatStickerSetPostRequest struct for SetChatStickerSetPostRequest
type SetChatStickerSetPostRequest struct {
	ChatId RestrictChatMemberPostRequestChatId `json:"chat_id"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

type _SetChatStickerSetPostRequest SetChatStickerSetPostRequest

// NewSetChatStickerSetPostRequest instantiates a new SetChatStickerSetPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetChatStickerSetPostRequest(chatId RestrictChatMemberPostRequestChatId, stickerSetName string) *SetChatStickerSetPostRequest {
	this := SetChatStickerSetPostRequest{}
	this.ChatId = chatId
	this.StickerSetName = stickerSetName
	return &this
}

// NewSetChatStickerSetPostRequestWithDefaults instantiates a new SetChatStickerSetPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetChatStickerSetPostRequestWithDefaults() *SetChatStickerSetPostRequest {
	this := SetChatStickerSetPostRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *SetChatStickerSetPostRequest) GetChatId() RestrictChatMemberPostRequestChatId {
	if o == nil {
		var ret RestrictChatMemberPostRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SetChatStickerSetPostRequest) GetChatIdOk() (*RestrictChatMemberPostRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SetChatStickerSetPostRequest) SetChatId(v RestrictChatMemberPostRequestChatId) {
	o.ChatId = v
}

// GetStickerSetName returns the StickerSetName field value
func (o *SetChatStickerSetPostRequest) GetStickerSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StickerSetName
}

// GetStickerSetNameOk returns a tuple with the StickerSetName field value
// and a boolean to check if the value has been set.
func (o *SetChatStickerSetPostRequest) GetStickerSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StickerSetName, true
}

// SetStickerSetName sets field value
func (o *SetChatStickerSetPostRequest) SetStickerSetName(v string) {
	o.StickerSetName = v
}

func (o SetChatStickerSetPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetChatStickerSetPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["sticker_set_name"] = o.StickerSetName
	return toSerialize, nil
}

func (o *SetChatStickerSetPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
		"sticker_set_name",
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

	varSetChatStickerSetPostRequest := _SetChatStickerSetPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetChatStickerSetPostRequest)

	if err != nil {
		return err
	}

	*o = SetChatStickerSetPostRequest(varSetChatStickerSetPostRequest)

	return err
}

type NullableSetChatStickerSetPostRequest struct {
	value *SetChatStickerSetPostRequest
	isSet bool
}

func (v NullableSetChatStickerSetPostRequest) Get() *SetChatStickerSetPostRequest {
	return v.value
}

func (v *NullableSetChatStickerSetPostRequest) Set(val *SetChatStickerSetPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetChatStickerSetPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetChatStickerSetPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetChatStickerSetPostRequest(val *SetChatStickerSetPostRequest) *NullableSetChatStickerSetPostRequest {
	return &NullableSetChatStickerSetPostRequest{value: val, isSet: true}
}

func (v NullableSetChatStickerSetPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetChatStickerSetPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


