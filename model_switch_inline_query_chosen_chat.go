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

// checks if the SwitchInlineQueryChosenChat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchInlineQueryChosenChat{}

// SwitchInlineQueryChosenChat This object represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	// *Optional*. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	Query *string `json:"query,omitempty"`
	// *Optional*. True, if private chats with users can be chosen
	AllowUserChats *bool `json:"allow_user_chats,omitempty"`
	// *Optional*. True, if private chats with bots can be chosen
	AllowBotChats *bool `json:"allow_bot_chats,omitempty"`
	// *Optional*. True, if group and supergroup chats can be chosen
	AllowGroupChats *bool `json:"allow_group_chats,omitempty"`
	// *Optional*. True, if channel chats can be chosen
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty"`
}

// NewSwitchInlineQueryChosenChat instantiates a new SwitchInlineQueryChosenChat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchInlineQueryChosenChat() *SwitchInlineQueryChosenChat {
	this := SwitchInlineQueryChosenChat{}
	return &this
}

// NewSwitchInlineQueryChosenChatWithDefaults instantiates a new SwitchInlineQueryChosenChat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchInlineQueryChosenChatWithDefaults() *SwitchInlineQueryChosenChat {
	this := SwitchInlineQueryChosenChat{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SwitchInlineQueryChosenChat) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchInlineQueryChosenChat) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SwitchInlineQueryChosenChat) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SwitchInlineQueryChosenChat) SetQuery(v string) {
	o.Query = &v
}


// GetAllowUserChats returns the AllowUserChats field value if set, zero value otherwise.
func (o *SwitchInlineQueryChosenChat) GetAllowUserChats() bool {
	if o == nil || IsNil(o.AllowUserChats) {
		var ret bool
		return ret
	}
	return *o.AllowUserChats
}

// GetAllowUserChatsOk returns a tuple with the AllowUserChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchInlineQueryChosenChat) GetAllowUserChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUserChats) {
		return nil, false
	}
	return o.AllowUserChats, true
}

// HasAllowUserChats returns a boolean if a field has been set.
func (o *SwitchInlineQueryChosenChat) HasAllowUserChats() bool {
	if o != nil && !IsNil(o.AllowUserChats) {
		return true
	}

	return false
}

// SetAllowUserChats gets a reference to the given bool and assigns it to the AllowUserChats field.
func (o *SwitchInlineQueryChosenChat) SetAllowUserChats(v bool) {
	o.AllowUserChats = &v
}


// GetAllowBotChats returns the AllowBotChats field value if set, zero value otherwise.
func (o *SwitchInlineQueryChosenChat) GetAllowBotChats() bool {
	if o == nil || IsNil(o.AllowBotChats) {
		var ret bool
		return ret
	}
	return *o.AllowBotChats
}

// GetAllowBotChatsOk returns a tuple with the AllowBotChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchInlineQueryChosenChat) GetAllowBotChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBotChats) {
		return nil, false
	}
	return o.AllowBotChats, true
}

// HasAllowBotChats returns a boolean if a field has been set.
func (o *SwitchInlineQueryChosenChat) HasAllowBotChats() bool {
	if o != nil && !IsNil(o.AllowBotChats) {
		return true
	}

	return false
}

// SetAllowBotChats gets a reference to the given bool and assigns it to the AllowBotChats field.
func (o *SwitchInlineQueryChosenChat) SetAllowBotChats(v bool) {
	o.AllowBotChats = &v
}


// GetAllowGroupChats returns the AllowGroupChats field value if set, zero value otherwise.
func (o *SwitchInlineQueryChosenChat) GetAllowGroupChats() bool {
	if o == nil || IsNil(o.AllowGroupChats) {
		var ret bool
		return ret
	}
	return *o.AllowGroupChats
}

// GetAllowGroupChatsOk returns a tuple with the AllowGroupChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchInlineQueryChosenChat) GetAllowGroupChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGroupChats) {
		return nil, false
	}
	return o.AllowGroupChats, true
}

// HasAllowGroupChats returns a boolean if a field has been set.
func (o *SwitchInlineQueryChosenChat) HasAllowGroupChats() bool {
	if o != nil && !IsNil(o.AllowGroupChats) {
		return true
	}

	return false
}

// SetAllowGroupChats gets a reference to the given bool and assigns it to the AllowGroupChats field.
func (o *SwitchInlineQueryChosenChat) SetAllowGroupChats(v bool) {
	o.AllowGroupChats = &v
}


// GetAllowChannelChats returns the AllowChannelChats field value if set, zero value otherwise.
func (o *SwitchInlineQueryChosenChat) GetAllowChannelChats() bool {
	if o == nil || IsNil(o.AllowChannelChats) {
		var ret bool
		return ret
	}
	return *o.AllowChannelChats
}

// GetAllowChannelChatsOk returns a tuple with the AllowChannelChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchInlineQueryChosenChat) GetAllowChannelChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChannelChats) {
		return nil, false
	}
	return o.AllowChannelChats, true
}

// HasAllowChannelChats returns a boolean if a field has been set.
func (o *SwitchInlineQueryChosenChat) HasAllowChannelChats() bool {
	if o != nil && !IsNil(o.AllowChannelChats) {
		return true
	}

	return false
}

// SetAllowChannelChats gets a reference to the given bool and assigns it to the AllowChannelChats field.
func (o *SwitchInlineQueryChosenChat) SetAllowChannelChats(v bool) {
	o.AllowChannelChats = &v
}


func (o SwitchInlineQueryChosenChat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchInlineQueryChosenChat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.AllowUserChats) {
		toSerialize["allow_user_chats"] = o.AllowUserChats
	}
	if !IsNil(o.AllowBotChats) {
		toSerialize["allow_bot_chats"] = o.AllowBotChats
	}
	if !IsNil(o.AllowGroupChats) {
		toSerialize["allow_group_chats"] = o.AllowGroupChats
	}
	if !IsNil(o.AllowChannelChats) {
		toSerialize["allow_channel_chats"] = o.AllowChannelChats
	}
	return toSerialize, nil
}

type NullableSwitchInlineQueryChosenChat struct {
	value *SwitchInlineQueryChosenChat
	isSet bool
}

func (v NullableSwitchInlineQueryChosenChat) Get() *SwitchInlineQueryChosenChat {
	return v.value
}

func (v *NullableSwitchInlineQueryChosenChat) Set(val *SwitchInlineQueryChosenChat) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchInlineQueryChosenChat) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchInlineQueryChosenChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchInlineQueryChosenChat(val *SwitchInlineQueryChosenChat) *NullableSwitchInlineQueryChosenChat {
	return &NullableSwitchInlineQueryChosenChat{value: val, isSet: true}
}

func (v NullableSwitchInlineQueryChosenChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchInlineQueryChosenChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


