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

// checks if the Checklist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Checklist{}

// Checklist Describes a checklist.
type Checklist struct {
	// Title of the checklist
	Title string `json:"title"`
	// *Optional*. Special entities that appear in the checklist title
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`
	// List of tasks in the checklist
	Tasks []ChecklistTask `json:"tasks"`
	// *Optional*. *True*, if users other than the creator of the list can add tasks to the list
	OthersCanAddTasks *bool `json:"others_can_add_tasks,omitempty"`
	// *Optional*. *True*, if users other than the creator of the list can mark tasks as done or not done
	OthersCanMarkTasksAsDone *bool `json:"others_can_mark_tasks_as_done,omitempty"`
}

type _Checklist Checklist

// NewChecklist instantiates a new Checklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecklist(title string, tasks []ChecklistTask) *Checklist {
	this := Checklist{}
	this.Title = title
	this.Tasks = tasks
	var othersCanAddTasks bool = true
	this.OthersCanAddTasks = &othersCanAddTasks
	var othersCanMarkTasksAsDone bool = true
	this.OthersCanMarkTasksAsDone = &othersCanMarkTasksAsDone
	return &this
}

// NewChecklistWithDefaults instantiates a new Checklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecklistWithDefaults() *Checklist {
	this := Checklist{}
	var othersCanAddTasks bool = true
	this.OthersCanAddTasks = &othersCanAddTasks
	var othersCanMarkTasksAsDone bool = true
	this.OthersCanMarkTasksAsDone = &othersCanMarkTasksAsDone
	return &this
}

// GetTitle returns the Title field value
func (o *Checklist) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Checklist) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Checklist) SetTitle(v string) {
	o.Title = v
}

// GetTitleEntities returns the TitleEntities field value if set, zero value otherwise.
func (o *Checklist) GetTitleEntities() []MessageEntity {
	if o == nil || IsNil(o.TitleEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TitleEntities
}

// GetTitleEntitiesOk returns a tuple with the TitleEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checklist) GetTitleEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TitleEntities) {
		return nil, false
	}
	return o.TitleEntities, true
}

// HasTitleEntities returns a boolean if a field has been set.
func (o *Checklist) HasTitleEntities() bool {
	if o != nil && !IsNil(o.TitleEntities) {
		return true
	}

	return false
}

// SetTitleEntities gets a reference to the given []MessageEntity and assigns it to the TitleEntities field.
func (o *Checklist) SetTitleEntities(v []MessageEntity) {
	o.TitleEntities = v
}


// GetTasks returns the Tasks field value
func (o *Checklist) GetTasks() []ChecklistTask {
	if o == nil {
		var ret []ChecklistTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *Checklist) GetTasksOk() ([]ChecklistTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *Checklist) SetTasks(v []ChecklistTask) {
	o.Tasks = v
}

// GetOthersCanAddTasks returns the OthersCanAddTasks field value if set, zero value otherwise.
func (o *Checklist) GetOthersCanAddTasks() bool {
	if o == nil || IsNil(o.OthersCanAddTasks) {
		var ret bool
		return ret
	}
	return *o.OthersCanAddTasks
}

// GetOthersCanAddTasksOk returns a tuple with the OthersCanAddTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checklist) GetOthersCanAddTasksOk() (*bool, bool) {
	if o == nil || IsNil(o.OthersCanAddTasks) {
		return nil, false
	}
	return o.OthersCanAddTasks, true
}

// HasOthersCanAddTasks returns a boolean if a field has been set.
func (o *Checklist) HasOthersCanAddTasks() bool {
	if o != nil && !IsNil(o.OthersCanAddTasks) {
		return true
	}

	return false
}

// SetOthersCanAddTasks gets a reference to the given bool and assigns it to the OthersCanAddTasks field.
func (o *Checklist) SetOthersCanAddTasks(v bool) {
	o.OthersCanAddTasks = &v
}


// GetOthersCanMarkTasksAsDone returns the OthersCanMarkTasksAsDone field value if set, zero value otherwise.
func (o *Checklist) GetOthersCanMarkTasksAsDone() bool {
	if o == nil || IsNil(o.OthersCanMarkTasksAsDone) {
		var ret bool
		return ret
	}
	return *o.OthersCanMarkTasksAsDone
}

// GetOthersCanMarkTasksAsDoneOk returns a tuple with the OthersCanMarkTasksAsDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checklist) GetOthersCanMarkTasksAsDoneOk() (*bool, bool) {
	if o == nil || IsNil(o.OthersCanMarkTasksAsDone) {
		return nil, false
	}
	return o.OthersCanMarkTasksAsDone, true
}

// HasOthersCanMarkTasksAsDone returns a boolean if a field has been set.
func (o *Checklist) HasOthersCanMarkTasksAsDone() bool {
	if o != nil && !IsNil(o.OthersCanMarkTasksAsDone) {
		return true
	}

	return false
}

// SetOthersCanMarkTasksAsDone gets a reference to the given bool and assigns it to the OthersCanMarkTasksAsDone field.
func (o *Checklist) SetOthersCanMarkTasksAsDone(v bool) {
	o.OthersCanMarkTasksAsDone = &v
}


func (o Checklist) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Checklist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.TitleEntities) {
		toSerialize["title_entities"] = o.TitleEntities
	}
	toSerialize["tasks"] = o.Tasks
	if !IsNil(o.OthersCanAddTasks) {
		toSerialize["others_can_add_tasks"] = o.OthersCanAddTasks
	}
	if !IsNil(o.OthersCanMarkTasksAsDone) {
		toSerialize["others_can_mark_tasks_as_done"] = o.OthersCanMarkTasksAsDone
	}
	return toSerialize, nil
}

func (o *Checklist) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varChecklist := _Checklist{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChecklist)

	if err != nil {
		return err
	}

	*o = Checklist(varChecklist)

	return err
}

type NullableChecklist struct {
	value *Checklist
	isSet bool
}

func (v NullableChecklist) Get() *Checklist {
	return v.value
}

func (v *NullableChecklist) Set(val *Checklist) {
	v.value = val
	v.isSet = true
}

func (v NullableChecklist) IsSet() bool {
	return v.isSet
}

func (v *NullableChecklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecklist(val *Checklist) *NullableChecklist {
	return &NullableChecklist{value: val, isSet: true}
}

func (v NullableChecklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


