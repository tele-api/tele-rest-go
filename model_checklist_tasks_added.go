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

// checks if the ChecklistTasksAdded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChecklistTasksAdded{}

// ChecklistTasksAdded Describes a service message about tasks added to a checklist.
type ChecklistTasksAdded struct {
	ChecklistMessage *Message `json:"checklist_message,omitempty"`
	// List of tasks added to the checklist
	Tasks []ChecklistTask `json:"tasks"`
}

type _ChecklistTasksAdded ChecklistTasksAdded

// NewChecklistTasksAdded instantiates a new ChecklistTasksAdded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecklistTasksAdded(tasks []ChecklistTask) *ChecklistTasksAdded {
	this := ChecklistTasksAdded{}
	this.Tasks = tasks
	return &this
}

// NewChecklistTasksAddedWithDefaults instantiates a new ChecklistTasksAdded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecklistTasksAddedWithDefaults() *ChecklistTasksAdded {
	this := ChecklistTasksAdded{}
	return &this
}

// GetChecklistMessage returns the ChecklistMessage field value if set, zero value otherwise.
func (o *ChecklistTasksAdded) GetChecklistMessage() Message {
	if o == nil || IsNil(o.ChecklistMessage) {
		var ret Message
		return ret
	}
	return *o.ChecklistMessage
}

// GetChecklistMessageOk returns a tuple with the ChecklistMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTasksAdded) GetChecklistMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.ChecklistMessage) {
		return nil, false
	}
	return o.ChecklistMessage, true
}

// HasChecklistMessage returns a boolean if a field has been set.
func (o *ChecklistTasksAdded) HasChecklistMessage() bool {
	if o != nil && !IsNil(o.ChecklistMessage) {
		return true
	}

	return false
}

// SetChecklistMessage gets a reference to the given Message and assigns it to the ChecklistMessage field.
func (o *ChecklistTasksAdded) SetChecklistMessage(v Message) {
	o.ChecklistMessage = &v
}


// GetTasks returns the Tasks field value
func (o *ChecklistTasksAdded) GetTasks() []ChecklistTask {
	if o == nil {
		var ret []ChecklistTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *ChecklistTasksAdded) GetTasksOk() ([]ChecklistTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *ChecklistTasksAdded) SetTasks(v []ChecklistTask) {
	o.Tasks = v
}

func (o ChecklistTasksAdded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChecklistTasksAdded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChecklistMessage) {
		toSerialize["checklist_message"] = o.ChecklistMessage
	}
	toSerialize["tasks"] = o.Tasks
	return toSerialize, nil
}

func (o *ChecklistTasksAdded) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tasks",
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

	varChecklistTasksAdded := _ChecklistTasksAdded{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChecklistTasksAdded)

	if err != nil {
		return err
	}

	*o = ChecklistTasksAdded(varChecklistTasksAdded)

	return err
}

type NullableChecklistTasksAdded struct {
	value *ChecklistTasksAdded
	isSet bool
}

func (v NullableChecklistTasksAdded) Get() *ChecklistTasksAdded {
	return v.value
}

func (v *NullableChecklistTasksAdded) Set(val *ChecklistTasksAdded) {
	v.value = val
	v.isSet = true
}

func (v NullableChecklistTasksAdded) IsSet() bool {
	return v.isSet
}

func (v *NullableChecklistTasksAdded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecklistTasksAdded(val *ChecklistTasksAdded) *NullableChecklistTasksAdded {
	return &NullableChecklistTasksAdded{value: val, isSet: true}
}

func (v NullableChecklistTasksAdded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecklistTasksAdded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


