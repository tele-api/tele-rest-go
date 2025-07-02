/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T09:17:05.586815301Z[Etc/UTC]
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

// checks if the ReplyParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplyParameters{}

// ReplyParameters Describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	// Identifier of the message that will be replied to in the current chat, or in the chat *chat\\_id* if it is specified
	MessageId int32 `json:"message_id"`
	ChatId *ReplyParametersChatId `json:"chat_id,omitempty"`
	// *Optional*. Pass *True* if the message should be sent even if the specified message to be replied to is not found. Always *False* for replies in another chat or forum topic. Always *True* for messages sent on behalf of a business account.
	AllowSendingWithoutReply *bool `json:"allow_sending_without_reply,omitempty"`
	// *Optional*. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities. The message will fail to send if the quote isn't found in the original message.
	Quote *string `json:"quote,omitempty"`
	// *Optional*. Mode for parsing entities in the quote. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	QuoteParseMode *string `json:"quote_parse_mode,omitempty"`
	// *Optional*. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of *quote\\_parse\\_mode*.
	QuoteEntities []MessageEntity `json:"quote_entities,omitempty"`
	// *Optional*. Position of the quote in the original message in UTF-16 code units
	QuotePosition *int32 `json:"quote_position,omitempty"`
}

type _ReplyParameters ReplyParameters

// NewReplyParameters instantiates a new ReplyParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplyParameters(messageId int32) *ReplyParameters {
	this := ReplyParameters{}
	this.MessageId = messageId
	return &this
}

// NewReplyParametersWithDefaults instantiates a new ReplyParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplyParametersWithDefaults() *ReplyParameters {
	this := ReplyParameters{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *ReplyParameters) GetMessageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *ReplyParameters) SetMessageId(v int32) {
	o.MessageId = v
}

// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *ReplyParameters) GetChatId() ReplyParametersChatId {
	if o == nil || IsNil(o.ChatId) {
		var ret ReplyParametersChatId
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetChatIdOk() (*ReplyParametersChatId, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *ReplyParameters) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given ReplyParametersChatId and assigns it to the ChatId field.
func (o *ReplyParameters) SetChatId(v ReplyParametersChatId) {
	o.ChatId = &v
}


// GetAllowSendingWithoutReply returns the AllowSendingWithoutReply field value if set, zero value otherwise.
func (o *ReplyParameters) GetAllowSendingWithoutReply() bool {
	if o == nil || IsNil(o.AllowSendingWithoutReply) {
		var ret bool
		return ret
	}
	return *o.AllowSendingWithoutReply
}

// GetAllowSendingWithoutReplyOk returns a tuple with the AllowSendingWithoutReply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetAllowSendingWithoutReplyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSendingWithoutReply) {
		return nil, false
	}
	return o.AllowSendingWithoutReply, true
}

// HasAllowSendingWithoutReply returns a boolean if a field has been set.
func (o *ReplyParameters) HasAllowSendingWithoutReply() bool {
	if o != nil && !IsNil(o.AllowSendingWithoutReply) {
		return true
	}

	return false
}

// SetAllowSendingWithoutReply gets a reference to the given bool and assigns it to the AllowSendingWithoutReply field.
func (o *ReplyParameters) SetAllowSendingWithoutReply(v bool) {
	o.AllowSendingWithoutReply = &v
}


// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *ReplyParameters) GetQuote() string {
	if o == nil || IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *ReplyParameters) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *ReplyParameters) SetQuote(v string) {
	o.Quote = &v
}


// GetQuoteParseMode returns the QuoteParseMode field value if set, zero value otherwise.
func (o *ReplyParameters) GetQuoteParseMode() string {
	if o == nil || IsNil(o.QuoteParseMode) {
		var ret string
		return ret
	}
	return *o.QuoteParseMode
}

// GetQuoteParseModeOk returns a tuple with the QuoteParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetQuoteParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteParseMode) {
		return nil, false
	}
	return o.QuoteParseMode, true
}

// HasQuoteParseMode returns a boolean if a field has been set.
func (o *ReplyParameters) HasQuoteParseMode() bool {
	if o != nil && !IsNil(o.QuoteParseMode) {
		return true
	}

	return false
}

// SetQuoteParseMode gets a reference to the given string and assigns it to the QuoteParseMode field.
func (o *ReplyParameters) SetQuoteParseMode(v string) {
	o.QuoteParseMode = &v
}


// GetQuoteEntities returns the QuoteEntities field value if set, zero value otherwise.
func (o *ReplyParameters) GetQuoteEntities() []MessageEntity {
	if o == nil || IsNil(o.QuoteEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.QuoteEntities
}

// GetQuoteEntitiesOk returns a tuple with the QuoteEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetQuoteEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.QuoteEntities) {
		return nil, false
	}
	return o.QuoteEntities, true
}

// HasQuoteEntities returns a boolean if a field has been set.
func (o *ReplyParameters) HasQuoteEntities() bool {
	if o != nil && !IsNil(o.QuoteEntities) {
		return true
	}

	return false
}

// SetQuoteEntities gets a reference to the given []MessageEntity and assigns it to the QuoteEntities field.
func (o *ReplyParameters) SetQuoteEntities(v []MessageEntity) {
	o.QuoteEntities = v
}


// GetQuotePosition returns the QuotePosition field value if set, zero value otherwise.
func (o *ReplyParameters) GetQuotePosition() int32 {
	if o == nil || IsNil(o.QuotePosition) {
		var ret int32
		return ret
	}
	return *o.QuotePosition
}

// GetQuotePositionOk returns a tuple with the QuotePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplyParameters) GetQuotePositionOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotePosition) {
		return nil, false
	}
	return o.QuotePosition, true
}

// HasQuotePosition returns a boolean if a field has been set.
func (o *ReplyParameters) HasQuotePosition() bool {
	if o != nil && !IsNil(o.QuotePosition) {
		return true
	}

	return false
}

// SetQuotePosition gets a reference to the given int32 and assigns it to the QuotePosition field.
func (o *ReplyParameters) SetQuotePosition(v int32) {
	o.QuotePosition = &v
}


func (o ReplyParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplyParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_id"] = o.MessageId
	if !IsNil(o.ChatId) {
		toSerialize["chat_id"] = o.ChatId
	}
	if !IsNil(o.AllowSendingWithoutReply) {
		toSerialize["allow_sending_without_reply"] = o.AllowSendingWithoutReply
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.QuoteParseMode) {
		toSerialize["quote_parse_mode"] = o.QuoteParseMode
	}
	if !IsNil(o.QuoteEntities) {
		toSerialize["quote_entities"] = o.QuoteEntities
	}
	if !IsNil(o.QuotePosition) {
		toSerialize["quote_position"] = o.QuotePosition
	}
	return toSerialize, nil
}

func (o *ReplyParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_id",
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

	varReplyParameters := _ReplyParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplyParameters)

	if err != nil {
		return err
	}

	*o = ReplyParameters(varReplyParameters)

	return err
}

type NullableReplyParameters struct {
	value *ReplyParameters
	isSet bool
}

func (v NullableReplyParameters) Get() *ReplyParameters {
	return v.value
}

func (v *NullableReplyParameters) Set(val *ReplyParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableReplyParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableReplyParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplyParameters(val *ReplyParameters) *NullableReplyParameters {
	return &NullableReplyParameters{value: val, isSet: true}
}

func (v NullableReplyParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplyParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


