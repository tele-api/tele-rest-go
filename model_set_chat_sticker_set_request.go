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

// checks if the SetChatStickerSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetChatStickerSetRequest{}

// SetChatStickerSetRequest Request parameters for setChatStickerSet
type SetChatStickerSetRequest struct {
	ChatId BotCommandScopeChatChatId `json:"chat_id"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

type _SetChatStickerSetRequest SetChatStickerSetRequest

// NewSetChatStickerSetRequest instantiates a new SetChatStickerSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetChatStickerSetRequest(chatId BotCommandScopeChatChatId, stickerSetName string) *SetChatStickerSetRequest {
	this := SetChatStickerSetRequest{}
	this.ChatId = chatId
	this.StickerSetName = stickerSetName
	return &this
}

// NewSetChatStickerSetRequestWithDefaults instantiates a new SetChatStickerSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetChatStickerSetRequestWithDefaults() *SetChatStickerSetRequest {
	this := SetChatStickerSetRequest{}
	return &this
}

// GetChatId returns the ChatId field value
func (o *SetChatStickerSetRequest) GetChatId() BotCommandScopeChatChatId {
	if o == nil {
		var ret BotCommandScopeChatChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *SetChatStickerSetRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *SetChatStickerSetRequest) SetChatId(v BotCommandScopeChatChatId) {
	o.ChatId = v
}

// GetStickerSetName returns the StickerSetName field value
func (o *SetChatStickerSetRequest) GetStickerSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StickerSetName
}

// GetStickerSetNameOk returns a tuple with the StickerSetName field value
// and a boolean to check if the value has been set.
func (o *SetChatStickerSetRequest) GetStickerSetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StickerSetName, true
}

// SetStickerSetName sets field value
func (o *SetChatStickerSetRequest) SetStickerSetName(v string) {
	o.StickerSetName = v
}

func (o SetChatStickerSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetChatStickerSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	toSerialize["sticker_set_name"] = o.StickerSetName
	return toSerialize, nil
}

func (o *SetChatStickerSetRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSetChatStickerSetRequest := _SetChatStickerSetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetChatStickerSetRequest)

	if err != nil {
		return err
	}

	*o = SetChatStickerSetRequest(varSetChatStickerSetRequest)

	return err
}

type NullableSetChatStickerSetRequest struct {
	value *SetChatStickerSetRequest
	isSet bool
}

func (v NullableSetChatStickerSetRequest) Get() *SetChatStickerSetRequest {
	return v.value
}

func (v *NullableSetChatStickerSetRequest) Set(val *SetChatStickerSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetChatStickerSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetChatStickerSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetChatStickerSetRequest(val *SetChatStickerSetRequest) *NullableSetChatStickerSetRequest {
	return &NullableSetChatStickerSetRequest{value: val, isSet: true}
}

func (v NullableSetChatStickerSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetChatStickerSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


