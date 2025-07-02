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
	"bytes"
	"fmt"
)

// checks if the PostSavePreparedInlineMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSavePreparedInlineMessageRequest{}

// PostSavePreparedInlineMessageRequest struct for PostSavePreparedInlineMessageRequest
type PostSavePreparedInlineMessageRequest struct {
	// Unique identifier of the target user that can use the prepared message
	UserId int32 `json:"user_id"`
	Result InlineQueryResult `json:"result"`
	// Pass *True* if the message can be sent to private chats with users
	AllowUserChats *bool `json:"allow_user_chats,omitempty"`
	// Pass *True* if the message can be sent to private chats with bots
	AllowBotChats *bool `json:"allow_bot_chats,omitempty"`
	// Pass *True* if the message can be sent to group and supergroup chats
	AllowGroupChats *bool `json:"allow_group_chats,omitempty"`
	// Pass *True* if the message can be sent to channel chats
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty"`
}

type _PostSavePreparedInlineMessageRequest PostSavePreparedInlineMessageRequest

// NewPostSavePreparedInlineMessageRequest instantiates a new PostSavePreparedInlineMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSavePreparedInlineMessageRequest(userId int32, result InlineQueryResult) *PostSavePreparedInlineMessageRequest {
	this := PostSavePreparedInlineMessageRequest{}
	this.UserId = userId
	this.Result = result
	return &this
}

// NewPostSavePreparedInlineMessageRequestWithDefaults instantiates a new PostSavePreparedInlineMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSavePreparedInlineMessageRequestWithDefaults() *PostSavePreparedInlineMessageRequest {
	this := PostSavePreparedInlineMessageRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PostSavePreparedInlineMessageRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostSavePreparedInlineMessageRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetResult returns the Result field value
func (o *PostSavePreparedInlineMessageRequest) GetResult() InlineQueryResult {
	if o == nil {
		var ret InlineQueryResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetResultOk() (*InlineQueryResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PostSavePreparedInlineMessageRequest) SetResult(v InlineQueryResult) {
	o.Result = v
}

// GetAllowUserChats returns the AllowUserChats field value if set, zero value otherwise.
func (o *PostSavePreparedInlineMessageRequest) GetAllowUserChats() bool {
	if o == nil || IsNil(o.AllowUserChats) {
		var ret bool
		return ret
	}
	return *o.AllowUserChats
}

// GetAllowUserChatsOk returns a tuple with the AllowUserChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetAllowUserChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUserChats) {
		return nil, false
	}
	return o.AllowUserChats, true
}

// HasAllowUserChats returns a boolean if a field has been set.
func (o *PostSavePreparedInlineMessageRequest) HasAllowUserChats() bool {
	if o != nil && !IsNil(o.AllowUserChats) {
		return true
	}

	return false
}

// SetAllowUserChats gets a reference to the given bool and assigns it to the AllowUserChats field.
func (o *PostSavePreparedInlineMessageRequest) SetAllowUserChats(v bool) {
	o.AllowUserChats = &v
}


// GetAllowBotChats returns the AllowBotChats field value if set, zero value otherwise.
func (o *PostSavePreparedInlineMessageRequest) GetAllowBotChats() bool {
	if o == nil || IsNil(o.AllowBotChats) {
		var ret bool
		return ret
	}
	return *o.AllowBotChats
}

// GetAllowBotChatsOk returns a tuple with the AllowBotChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetAllowBotChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBotChats) {
		return nil, false
	}
	return o.AllowBotChats, true
}

// HasAllowBotChats returns a boolean if a field has been set.
func (o *PostSavePreparedInlineMessageRequest) HasAllowBotChats() bool {
	if o != nil && !IsNil(o.AllowBotChats) {
		return true
	}

	return false
}

// SetAllowBotChats gets a reference to the given bool and assigns it to the AllowBotChats field.
func (o *PostSavePreparedInlineMessageRequest) SetAllowBotChats(v bool) {
	o.AllowBotChats = &v
}


// GetAllowGroupChats returns the AllowGroupChats field value if set, zero value otherwise.
func (o *PostSavePreparedInlineMessageRequest) GetAllowGroupChats() bool {
	if o == nil || IsNil(o.AllowGroupChats) {
		var ret bool
		return ret
	}
	return *o.AllowGroupChats
}

// GetAllowGroupChatsOk returns a tuple with the AllowGroupChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetAllowGroupChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGroupChats) {
		return nil, false
	}
	return o.AllowGroupChats, true
}

// HasAllowGroupChats returns a boolean if a field has been set.
func (o *PostSavePreparedInlineMessageRequest) HasAllowGroupChats() bool {
	if o != nil && !IsNil(o.AllowGroupChats) {
		return true
	}

	return false
}

// SetAllowGroupChats gets a reference to the given bool and assigns it to the AllowGroupChats field.
func (o *PostSavePreparedInlineMessageRequest) SetAllowGroupChats(v bool) {
	o.AllowGroupChats = &v
}


// GetAllowChannelChats returns the AllowChannelChats field value if set, zero value otherwise.
func (o *PostSavePreparedInlineMessageRequest) GetAllowChannelChats() bool {
	if o == nil || IsNil(o.AllowChannelChats) {
		var ret bool
		return ret
	}
	return *o.AllowChannelChats
}

// GetAllowChannelChatsOk returns a tuple with the AllowChannelChats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSavePreparedInlineMessageRequest) GetAllowChannelChatsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChannelChats) {
		return nil, false
	}
	return o.AllowChannelChats, true
}

// HasAllowChannelChats returns a boolean if a field has been set.
func (o *PostSavePreparedInlineMessageRequest) HasAllowChannelChats() bool {
	if o != nil && !IsNil(o.AllowChannelChats) {
		return true
	}

	return false
}

// SetAllowChannelChats gets a reference to the given bool and assigns it to the AllowChannelChats field.
func (o *PostSavePreparedInlineMessageRequest) SetAllowChannelChats(v bool) {
	o.AllowChannelChats = &v
}


func (o PostSavePreparedInlineMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSavePreparedInlineMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["result"] = o.Result
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

func (o *PostSavePreparedInlineMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"result",
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

	varPostSavePreparedInlineMessageRequest := _PostSavePreparedInlineMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSavePreparedInlineMessageRequest)

	if err != nil {
		return err
	}

	*o = PostSavePreparedInlineMessageRequest(varPostSavePreparedInlineMessageRequest)

	return err
}

type NullablePostSavePreparedInlineMessageRequest struct {
	value *PostSavePreparedInlineMessageRequest
	isSet bool
}

func (v NullablePostSavePreparedInlineMessageRequest) Get() *PostSavePreparedInlineMessageRequest {
	return v.value
}

func (v *NullablePostSavePreparedInlineMessageRequest) Set(val *PostSavePreparedInlineMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSavePreparedInlineMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSavePreparedInlineMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSavePreparedInlineMessageRequest(val *PostSavePreparedInlineMessageRequest) *NullablePostSavePreparedInlineMessageRequest {
	return &NullablePostSavePreparedInlineMessageRequest{value: val, isSet: true}
}

func (v NullablePostSavePreparedInlineMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSavePreparedInlineMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


