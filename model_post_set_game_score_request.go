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

// checks if the PostSetGameScoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSetGameScoreRequest{}

// PostSetGameScoreRequest struct for PostSetGameScoreRequest
type PostSetGameScoreRequest struct {
	// User identifier
	UserId int32 `json:"user_id"`
	// New score, must be non-negative
	Score int32 `json:"score"`
	// Pass *True* if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	Force *bool `json:"force,omitempty"`
	// Pass *True* if the game message should not be automatically edited to include the current scoreboard
	DisableEditMessage *bool `json:"disable_edit_message,omitempty"`
	// Required if *inline\\_message\\_id* is not specified. Unique identifier for the target chat
	ChatId *int32 `json:"chat_id,omitempty"`
	// Required if *inline\\_message\\_id* is not specified. Identifier of the sent message
	MessageId *int32 `json:"message_id,omitempty"`
	// Required if *chat\\_id* and *message\\_id* are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}

type _PostSetGameScoreRequest PostSetGameScoreRequest

// NewPostSetGameScoreRequest instantiates a new PostSetGameScoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSetGameScoreRequest(userId int32, score int32) *PostSetGameScoreRequest {
	this := PostSetGameScoreRequest{}
	this.UserId = userId
	this.Score = score
	return &this
}

// NewPostSetGameScoreRequestWithDefaults instantiates a new PostSetGameScoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSetGameScoreRequestWithDefaults() *PostSetGameScoreRequest {
	this := PostSetGameScoreRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *PostSetGameScoreRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostSetGameScoreRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetScore returns the Score field value
func (o *PostSetGameScoreRequest) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *PostSetGameScoreRequest) SetScore(v int32) {
	o.Score = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *PostSetGameScoreRequest) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *PostSetGameScoreRequest) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *PostSetGameScoreRequest) SetForce(v bool) {
	o.Force = &v
}


// GetDisableEditMessage returns the DisableEditMessage field value if set, zero value otherwise.
func (o *PostSetGameScoreRequest) GetDisableEditMessage() bool {
	if o == nil || IsNil(o.DisableEditMessage) {
		var ret bool
		return ret
	}
	return *o.DisableEditMessage
}

// GetDisableEditMessageOk returns a tuple with the DisableEditMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetDisableEditMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableEditMessage) {
		return nil, false
	}
	return o.DisableEditMessage, true
}

// HasDisableEditMessage returns a boolean if a field has been set.
func (o *PostSetGameScoreRequest) HasDisableEditMessage() bool {
	if o != nil && !IsNil(o.DisableEditMessage) {
		return true
	}

	return false
}

// SetDisableEditMessage gets a reference to the given bool and assigns it to the DisableEditMessage field.
func (o *PostSetGameScoreRequest) SetDisableEditMessage(v bool) {
	o.DisableEditMessage = &v
}


// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *PostSetGameScoreRequest) GetChatId() int32 {
	if o == nil || IsNil(o.ChatId) {
		var ret int32
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetChatIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *PostSetGameScoreRequest) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given int32 and assigns it to the ChatId field.
func (o *PostSetGameScoreRequest) SetChatId(v int32) {
	o.ChatId = &v
}


// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *PostSetGameScoreRequest) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId) {
		var ret int32
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetMessageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *PostSetGameScoreRequest) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int32 and assigns it to the MessageId field.
func (o *PostSetGameScoreRequest) SetMessageId(v int32) {
	o.MessageId = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *PostSetGameScoreRequest) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSetGameScoreRequest) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *PostSetGameScoreRequest) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *PostSetGameScoreRequest) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


func (o PostSetGameScoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSetGameScoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["score"] = o.Score
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.DisableEditMessage) {
		toSerialize["disable_edit_message"] = o.DisableEditMessage
	}
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.InlineMessageId) {
		toSerialize["inline_message_id"] = o.InlineMessageId
	}
	return toSerialize, nil
}

func (o *PostSetGameScoreRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"score",
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

	varPostSetGameScoreRequest := _PostSetGameScoreRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostSetGameScoreRequest)

	if err != nil {
		return err
	}

	*o = PostSetGameScoreRequest(varPostSetGameScoreRequest)

	return err
}

type NullablePostSetGameScoreRequest struct {
	value *PostSetGameScoreRequest
	isSet bool
}

func (v NullablePostSetGameScoreRequest) Get() *PostSetGameScoreRequest {
	return v.value
}

func (v *NullablePostSetGameScoreRequest) Set(val *PostSetGameScoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSetGameScoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSetGameScoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSetGameScoreRequest(val *PostSetGameScoreRequest) *NullablePostSetGameScoreRequest {
	return &NullablePostSetGameScoreRequest{value: val, isSet: true}
}

func (v NullablePostSetGameScoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSetGameScoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


