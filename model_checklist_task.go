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

// checks if the ChecklistTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChecklistTask{}

// ChecklistTask Describes a task in a checklist.
type ChecklistTask struct {
	// Unique identifier of the task
	Id int32 `json:"id"`
	// Text of the task
	Text string `json:"text"`
	// *Optional*. Special entities that appear in the task text
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User `json:"completed_by_user,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the task was completed; 0 if the task wasn't completed
	CompletionDate *int32 `json:"completion_date,omitempty"`
}

type _ChecklistTask ChecklistTask

// NewChecklistTask instantiates a new ChecklistTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecklistTask(id int32, text string) *ChecklistTask {
	this := ChecklistTask{}
	this.Id = id
	this.Text = text
	return &this
}

// NewChecklistTaskWithDefaults instantiates a new ChecklistTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecklistTaskWithDefaults() *ChecklistTask {
	this := ChecklistTask{}
	return &this
}

// GetId returns the Id field value
func (o *ChecklistTask) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChecklistTask) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChecklistTask) SetId(v int32) {
	o.Id = v
}

// GetText returns the Text field value
func (o *ChecklistTask) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ChecklistTask) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ChecklistTask) SetText(v string) {
	o.Text = v
}

// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *ChecklistTask) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTask) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *ChecklistTask) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *ChecklistTask) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


// GetCompletedByUser returns the CompletedByUser field value if set, zero value otherwise.
func (o *ChecklistTask) GetCompletedByUser() User {
	if o == nil || IsNil(o.CompletedByUser) {
		var ret User
		return ret
	}
	return *o.CompletedByUser
}

// GetCompletedByUserOk returns a tuple with the CompletedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTask) GetCompletedByUserOk() (*User, bool) {
	if o == nil || IsNil(o.CompletedByUser) {
		return nil, false
	}
	return o.CompletedByUser, true
}

// HasCompletedByUser returns a boolean if a field has been set.
func (o *ChecklistTask) HasCompletedByUser() bool {
	if o != nil && !IsNil(o.CompletedByUser) {
		return true
	}

	return false
}

// SetCompletedByUser gets a reference to the given User and assigns it to the CompletedByUser field.
func (o *ChecklistTask) SetCompletedByUser(v User) {
	o.CompletedByUser = &v
}


// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *ChecklistTask) GetCompletionDate() int32 {
	if o == nil || IsNil(o.CompletionDate) {
		var ret int32
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTask) GetCompletionDateOk() (*int32, bool) {
	if o == nil || IsNil(o.CompletionDate) {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *ChecklistTask) HasCompletionDate() bool {
	if o != nil && !IsNil(o.CompletionDate) {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given int32 and assigns it to the CompletionDate field.
func (o *ChecklistTask) SetCompletionDate(v int32) {
	o.CompletionDate = &v
}


func (o ChecklistTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChecklistTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	if !IsNil(o.TextEntities) {
		toSerialize["text_entities"] = o.TextEntities
	}
	if !IsNil(o.CompletedByUser) {
		toSerialize["completed_by_user"] = o.CompletedByUser
	}
	if !IsNil(o.CompletionDate) {
		toSerialize["completion_date"] = o.CompletionDate
	}
	return toSerialize, nil
}

func (o *ChecklistTask) UnmarshalJSON(data []byte) (err error) {
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

	varChecklistTask := _ChecklistTask{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChecklistTask)

	if err != nil {
		return err
	}

	*o = ChecklistTask(varChecklistTask)

	return err
}

type NullableChecklistTask struct {
	value *ChecklistTask
	isSet bool
}

func (v NullableChecklistTask) Get() *ChecklistTask {
	return v.value
}

func (v *NullableChecklistTask) Set(val *ChecklistTask) {
	v.value = val
	v.isSet = true
}

func (v NullableChecklistTask) IsSet() bool {
	return v.isSet
}

func (v *NullableChecklistTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecklistTask(val *ChecklistTask) *NullableChecklistTask {
	return &NullableChecklistTask{value: val, isSet: true}
}

func (v NullableChecklistTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecklistTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


