/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the StoryAreaTypeUniqueGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryAreaTypeUniqueGift{}

// StoryAreaTypeUniqueGift Describes a story area pointing to a unique gift. Currently, a story can have at most 1 unique gift area.
type StoryAreaTypeUniqueGift struct {
	// Type of the area, always “unique\\_gift”
	Type string `json:"type"`
	// Unique name of the gift
	Name string `json:"name"`
}

type _StoryAreaTypeUniqueGift StoryAreaTypeUniqueGift

// NewStoryAreaTypeUniqueGift instantiates a new StoryAreaTypeUniqueGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryAreaTypeUniqueGift(type_ string, name string) *StoryAreaTypeUniqueGift {
	this := StoryAreaTypeUniqueGift{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewStoryAreaTypeUniqueGiftWithDefaults instantiates a new StoryAreaTypeUniqueGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryAreaTypeUniqueGiftWithDefaults() *StoryAreaTypeUniqueGift {
	this := StoryAreaTypeUniqueGift{}
	var type_ string = "unique_gift"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StoryAreaTypeUniqueGift) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeUniqueGift) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StoryAreaTypeUniqueGift) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *StoryAreaTypeUniqueGift) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeUniqueGift) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StoryAreaTypeUniqueGift) SetName(v string) {
	o.Name = v
}

func (o StoryAreaTypeUniqueGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryAreaTypeUniqueGift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *StoryAreaTypeUniqueGift) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
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

	varStoryAreaTypeUniqueGift := _StoryAreaTypeUniqueGift{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStoryAreaTypeUniqueGift)

	if err != nil {
		return err
	}

	*o = StoryAreaTypeUniqueGift(varStoryAreaTypeUniqueGift)

	return err
}

type NullableStoryAreaTypeUniqueGift struct {
	value *StoryAreaTypeUniqueGift
	isSet bool
}

func (v NullableStoryAreaTypeUniqueGift) Get() *StoryAreaTypeUniqueGift {
	return v.value
}

func (v *NullableStoryAreaTypeUniqueGift) Set(val *StoryAreaTypeUniqueGift) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaTypeUniqueGift) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaTypeUniqueGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaTypeUniqueGift(val *StoryAreaTypeUniqueGift) *NullableStoryAreaTypeUniqueGift {
	return &NullableStoryAreaTypeUniqueGift{value: val, isSet: true}
}

func (v NullableStoryAreaTypeUniqueGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaTypeUniqueGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


