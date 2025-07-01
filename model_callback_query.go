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

// checks if the CallbackQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallbackQuery{}

// CallbackQuery This object represents an incoming callback query from a callback button in an [inline keyboard](https://core.telegram.org/bots/features#inline-keyboards). If the button that originated the query was attached to a message sent by the bot, the field *message* will be present. If the button was attached to a message sent via the bot (in [inline mode](https://core.telegram.org/bots/api/#inline-mode)), the field *inline\\_message\\_id* will be present. Exactly one of the fields *data* or *game\\_short\\_name* will be present.
type CallbackQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	From User `json:"from"`
	Message *MaybeInaccessibleMessage `json:"message,omitempty"`
	// *Optional*. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in [games](https://core.telegram.org/bots/api/#games).
	ChatInstance string `json:"chat_instance"`
	// *Optional*. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	Data *string `json:"data,omitempty"`
	// *Optional*. Short name of a [Game](https://core.telegram.org/bots/api/#games) to be returned, serves as the unique identifier for the game
	GameShortName *string `json:"game_short_name,omitempty"`
}

type _CallbackQuery CallbackQuery

// NewCallbackQuery instantiates a new CallbackQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallbackQuery(id string, from User, chatInstance string) *CallbackQuery {
	this := CallbackQuery{}
	this.Id = id
	this.From = from
	this.ChatInstance = chatInstance
	return &this
}

// NewCallbackQueryWithDefaults instantiates a new CallbackQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackQueryWithDefaults() *CallbackQuery {
	this := CallbackQuery{}
	return &this
}

// GetId returns the Id field value
func (o *CallbackQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CallbackQuery) SetId(v string) {
	o.Id = v
}

// GetFrom returns the From field value
func (o *CallbackQuery) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CallbackQuery) SetFrom(v User) {
	o.From = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CallbackQuery) GetMessage() MaybeInaccessibleMessage {
	if o == nil || IsNil(o.Message) {
		var ret MaybeInaccessibleMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetMessageOk() (*MaybeInaccessibleMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CallbackQuery) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given MaybeInaccessibleMessage and assigns it to the Message field.
func (o *CallbackQuery) SetMessage(v MaybeInaccessibleMessage) {
	o.Message = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *CallbackQuery) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *CallbackQuery) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *CallbackQuery) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


// GetChatInstance returns the ChatInstance field value
func (o *CallbackQuery) GetChatInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChatInstance
}

// GetChatInstanceOk returns a tuple with the ChatInstance field value
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetChatInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatInstance, true
}

// SetChatInstance sets field value
func (o *CallbackQuery) SetChatInstance(v string) {
	o.ChatInstance = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CallbackQuery) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CallbackQuery) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *CallbackQuery) SetData(v string) {
	o.Data = &v
}


// GetGameShortName returns the GameShortName field value if set, zero value otherwise.
func (o *CallbackQuery) GetGameShortName() string {
	if o == nil || IsNil(o.GameShortName) {
		var ret string
		return ret
	}
	return *o.GameShortName
}

// GetGameShortNameOk returns a tuple with the GameShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallbackQuery) GetGameShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.GameShortName) {
		return nil, false
	}
	return o.GameShortName, true
}

// HasGameShortName returns a boolean if a field has been set.
func (o *CallbackQuery) HasGameShortName() bool {
	if o != nil && !IsNil(o.GameShortName) {
		return true
	}

	return false
}

// SetGameShortName gets a reference to the given string and assigns it to the GameShortName field.
func (o *CallbackQuery) SetGameShortName(v string) {
	o.GameShortName = &v
}


func (o CallbackQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallbackQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["from"] = o.From
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.InlineMessageId) {
		toSerialize["inline_message_id"] = o.InlineMessageId
	}
	toSerialize["chat_instance"] = o.ChatInstance
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.GameShortName) {
		toSerialize["game_short_name"] = o.GameShortName
	}
	return toSerialize, nil
}

func (o *CallbackQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"from",
		"chat_instance",
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

	varCallbackQuery := _CallbackQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCallbackQuery)

	if err != nil {
		return err
	}

	*o = CallbackQuery(varCallbackQuery)

	return err
}

type NullableCallbackQuery struct {
	value *CallbackQuery
	isSet bool
}

func (v NullableCallbackQuery) Get() *CallbackQuery {
	return v.value
}

func (v *NullableCallbackQuery) Set(val *CallbackQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackQuery(val *CallbackQuery) *NullableCallbackQuery {
	return &NullableCallbackQuery{value: val, isSet: true}
}

func (v NullableCallbackQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


