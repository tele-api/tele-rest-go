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
)

// checks if the ChecklistTasksDone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChecklistTasksDone{}

// ChecklistTasksDone Describes a service message about checklist tasks marked as done or not done.
type ChecklistTasksDone struct {
	ChecklistMessage *Message `json:"checklist_message,omitempty"`
	// *Optional*. Identifiers of the tasks that were marked as done
	MarkedAsDoneTaskIds []int32 `json:"marked_as_done_task_ids,omitempty"`
	// *Optional*. Identifiers of the tasks that were marked as not done
	MarkedAsNotDoneTaskIds []int32 `json:"marked_as_not_done_task_ids,omitempty"`
}

// NewChecklistTasksDone instantiates a new ChecklistTasksDone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecklistTasksDone() *ChecklistTasksDone {
	this := ChecklistTasksDone{}
	return &this
}

// NewChecklistTasksDoneWithDefaults instantiates a new ChecklistTasksDone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecklistTasksDoneWithDefaults() *ChecklistTasksDone {
	this := ChecklistTasksDone{}
	return &this
}

// GetChecklistMessage returns the ChecklistMessage field value if set, zero value otherwise.
func (o *ChecklistTasksDone) GetChecklistMessage() Message {
	if o == nil || IsNil(o.ChecklistMessage) {
		var ret Message
		return ret
	}
	return *o.ChecklistMessage
}

// GetChecklistMessageOk returns a tuple with the ChecklistMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTasksDone) GetChecklistMessageOk() (*Message, bool) {
	if o == nil || IsNil(o.ChecklistMessage) {
		return nil, false
	}
	return o.ChecklistMessage, true
}

// HasChecklistMessage returns a boolean if a field has been set.
func (o *ChecklistTasksDone) HasChecklistMessage() bool {
	if o != nil && !IsNil(o.ChecklistMessage) {
		return true
	}

	return false
}

// SetChecklistMessage gets a reference to the given Message and assigns it to the ChecklistMessage field.
func (o *ChecklistTasksDone) SetChecklistMessage(v Message) {
	o.ChecklistMessage = &v
}


// GetMarkedAsDoneTaskIds returns the MarkedAsDoneTaskIds field value if set, zero value otherwise.
func (o *ChecklistTasksDone) GetMarkedAsDoneTaskIds() []int32 {
	if o == nil || IsNil(o.MarkedAsDoneTaskIds) {
		var ret []int32
		return ret
	}
	return o.MarkedAsDoneTaskIds
}

// GetMarkedAsDoneTaskIdsOk returns a tuple with the MarkedAsDoneTaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTasksDone) GetMarkedAsDoneTaskIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MarkedAsDoneTaskIds) {
		return nil, false
	}
	return o.MarkedAsDoneTaskIds, true
}

// HasMarkedAsDoneTaskIds returns a boolean if a field has been set.
func (o *ChecklistTasksDone) HasMarkedAsDoneTaskIds() bool {
	if o != nil && !IsNil(o.MarkedAsDoneTaskIds) {
		return true
	}

	return false
}

// SetMarkedAsDoneTaskIds gets a reference to the given []int32 and assigns it to the MarkedAsDoneTaskIds field.
func (o *ChecklistTasksDone) SetMarkedAsDoneTaskIds(v []int32) {
	o.MarkedAsDoneTaskIds = v
}


// GetMarkedAsNotDoneTaskIds returns the MarkedAsNotDoneTaskIds field value if set, zero value otherwise.
func (o *ChecklistTasksDone) GetMarkedAsNotDoneTaskIds() []int32 {
	if o == nil || IsNil(o.MarkedAsNotDoneTaskIds) {
		var ret []int32
		return ret
	}
	return o.MarkedAsNotDoneTaskIds
}

// GetMarkedAsNotDoneTaskIdsOk returns a tuple with the MarkedAsNotDoneTaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChecklistTasksDone) GetMarkedAsNotDoneTaskIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MarkedAsNotDoneTaskIds) {
		return nil, false
	}
	return o.MarkedAsNotDoneTaskIds, true
}

// HasMarkedAsNotDoneTaskIds returns a boolean if a field has been set.
func (o *ChecklistTasksDone) HasMarkedAsNotDoneTaskIds() bool {
	if o != nil && !IsNil(o.MarkedAsNotDoneTaskIds) {
		return true
	}

	return false
}

// SetMarkedAsNotDoneTaskIds gets a reference to the given []int32 and assigns it to the MarkedAsNotDoneTaskIds field.
func (o *ChecklistTasksDone) SetMarkedAsNotDoneTaskIds(v []int32) {
	o.MarkedAsNotDoneTaskIds = v
}


func (o ChecklistTasksDone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChecklistTasksDone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChecklistMessage) {
		toSerialize["checklist_message"] = o.ChecklistMessage
	}
	if !IsNil(o.MarkedAsDoneTaskIds) {
		toSerialize["marked_as_done_task_ids"] = o.MarkedAsDoneTaskIds
	}
	if !IsNil(o.MarkedAsNotDoneTaskIds) {
		toSerialize["marked_as_not_done_task_ids"] = o.MarkedAsNotDoneTaskIds
	}
	return toSerialize, nil
}

type NullableChecklistTasksDone struct {
	value *ChecklistTasksDone
	isSet bool
}

func (v NullableChecklistTasksDone) Get() *ChecklistTasksDone {
	return v.value
}

func (v *NullableChecklistTasksDone) Set(val *ChecklistTasksDone) {
	v.value = val
	v.isSet = true
}

func (v NullableChecklistTasksDone) IsSet() bool {
	return v.isSet
}

func (v *NullableChecklistTasksDone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecklistTasksDone(val *ChecklistTasksDone) *NullableChecklistTasksDone {
	return &NullableChecklistTasksDone{value: val, isSet: true}
}

func (v NullableChecklistTasksDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecklistTasksDone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


