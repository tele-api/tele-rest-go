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

// checks if the InputChecklistTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputChecklistTask{}

// InputChecklistTask Describes a task to add to a checklist.
type InputChecklistTask struct {
	// Unique identifier of the task; must be positive and unique among all task identifiers currently present in the checklist
	Id int32 `json:"id"`
	// Text of the task; 1-100 characters after entities parsing
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the text. See [formatting options](https://core.telegram.org/bots/api#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the text, which can be specified instead of parse\\_mode. Currently, only *bold*, *italic*, *underline*, *strikethrough*, *spoiler*, and *custom\\_emoji* entities are allowed.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type _InputChecklistTask InputChecklistTask

// NewInputChecklistTask instantiates a new InputChecklistTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputChecklistTask(id int32, text string) *InputChecklistTask {
	this := InputChecklistTask{}
	this.Id = id
	this.Text = text
	return &this
}

// NewInputChecklistTaskWithDefaults instantiates a new InputChecklistTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputChecklistTaskWithDefaults() *InputChecklistTask {
	this := InputChecklistTask{}
	return &this
}

// GetId returns the Id field value
func (o *InputChecklistTask) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InputChecklistTask) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InputChecklistTask) SetId(v int32) {
	o.Id = v
}

// GetText returns the Text field value
func (o *InputChecklistTask) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *InputChecklistTask) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *InputChecklistTask) SetText(v string) {
	o.Text = v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *InputChecklistTask) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputChecklistTask) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *InputChecklistTask) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *InputChecklistTask) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *InputChecklistTask) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputChecklistTask) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *InputChecklistTask) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *InputChecklistTask) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


func (o InputChecklistTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputChecklistTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.TextEntities) {
		toSerialize["text_entities"] = o.TextEntities
	}
	return toSerialize, nil
}

func (o *InputChecklistTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"text",
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

	varInputChecklistTask := _InputChecklistTask{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputChecklistTask)

	if err != nil {
		return err
	}

	*o = InputChecklistTask(varInputChecklistTask)

	return err
}

type NullableInputChecklistTask struct {
	value *InputChecklistTask
	isSet bool
}

func (v NullableInputChecklistTask) Get() *InputChecklistTask {
	return v.value
}

func (v *NullableInputChecklistTask) Set(val *InputChecklistTask) {
	v.value = val
	v.isSet = true
}

func (v NullableInputChecklistTask) IsSet() bool {
	return v.isSet
}

func (v *NullableInputChecklistTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputChecklistTask(val *InputChecklistTask) *NullableInputChecklistTask {
	return &NullableInputChecklistTask{value: val, isSet: true}
}

func (v NullableInputChecklistTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputChecklistTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


