/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-19T09:30:13.278207440Z[Etc/UTC]
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

// checks if the InputTextMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputTextMessageContent{}

// InputTextMessageContent Represents the [content](https://core.telegram.org/bots/api/#inputmessagecontent) of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	// *Optional*. Mode for parsing entities in the message text. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in message text, which can be specified instead of *parse\\_mode*
	Entities []MessageEntity `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type _InputTextMessageContent InputTextMessageContent

// NewInputTextMessageContent instantiates a new InputTextMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputTextMessageContent(messageText string) *InputTextMessageContent {
	this := InputTextMessageContent{}
	this.MessageText = messageText
	return &this
}

// NewInputTextMessageContentWithDefaults instantiates a new InputTextMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputTextMessageContentWithDefaults() *InputTextMessageContent {
	this := InputTextMessageContent{}
	return &this
}

// GetMessageText returns the MessageText field value
func (o *InputTextMessageContent) GetMessageText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageText
}

// GetMessageTextOk returns a tuple with the MessageText field value
// and a boolean to check if the value has been set.
func (o *InputTextMessageContent) GetMessageTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageText, true
}

// SetMessageText sets field value
func (o *InputTextMessageContent) SetMessageText(v string) {
	o.MessageText = v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InputTextMessageContent) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTextMessageContent) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InputTextMessageContent) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InputTextMessageContent) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *InputTextMessageContent) GetEntities() []MessageEntity {
	if o == nil || IsNil(o.Entities) {
		var ret []MessageEntity
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTextMessageContent) GetEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *InputTextMessageContent) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []MessageEntity and assigns it to the Entities field.
func (o *InputTextMessageContent) SetEntities(v []MessageEntity) {
	o.Entities = v
}


// GetLinkPreviewOptions returns the LinkPreviewOptions field value if set, zero value otherwise.
func (o *InputTextMessageContent) GetLinkPreviewOptions() LinkPreviewOptions {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		var ret LinkPreviewOptions
		return ret
	}
	return *o.LinkPreviewOptions
}

// GetLinkPreviewOptionsOk returns a tuple with the LinkPreviewOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputTextMessageContent) GetLinkPreviewOptionsOk() (*LinkPreviewOptions, bool) {
	if o == nil || IsNil(o.LinkPreviewOptions) {
		return nil, false
	}
	return o.LinkPreviewOptions, true
}

// HasLinkPreviewOptions returns a boolean if a field has been set.
func (o *InputTextMessageContent) HasLinkPreviewOptions() bool {
	if o != nil && !IsNil(o.LinkPreviewOptions) {
		return true
	}

	return false
}

// SetLinkPreviewOptions gets a reference to the given LinkPreviewOptions and assigns it to the LinkPreviewOptions field.
func (o *InputTextMessageContent) SetLinkPreviewOptions(v LinkPreviewOptions) {
	o.LinkPreviewOptions = &v
}


func (o InputTextMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputTextMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_text"] = o.MessageText
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.LinkPreviewOptions) {
		toSerialize["link_preview_options"] = o.LinkPreviewOptions
	}
	return toSerialize, nil
}

func (o *InputTextMessageContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_text",
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

	varInputTextMessageContent := _InputTextMessageContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputTextMessageContent)

	if err != nil {
		return err
	}

	*o = InputTextMessageContent(varInputTextMessageContent)

	return err
}

type NullableInputTextMessageContent struct {
	value *InputTextMessageContent
	isSet bool
}

func (v NullableInputTextMessageContent) Get() *InputTextMessageContent {
	return v.value
}

func (v *NullableInputTextMessageContent) Set(val *InputTextMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputTextMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputTextMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputTextMessageContent(val *InputTextMessageContent) *NullableInputTextMessageContent {
	return &NullableInputTextMessageContent{value: val, isSet: true}
}

func (v NullableInputTextMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputTextMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


