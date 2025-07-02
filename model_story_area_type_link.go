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

// checks if the StoryAreaTypeLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryAreaTypeLink{}

// StoryAreaTypeLink Describes a story area pointing to an HTTP or tg:// link. Currently, a story can have up to 3 link areas.
type StoryAreaTypeLink struct {
	// Type of the area, always “link”
	Type string `json:"type"`
	// HTTP or tg:// URL to be opened when the area is clicked
	Url string `json:"url"`
}

type _StoryAreaTypeLink StoryAreaTypeLink

// NewStoryAreaTypeLink instantiates a new StoryAreaTypeLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryAreaTypeLink(type_ string, url string) *StoryAreaTypeLink {
	this := StoryAreaTypeLink{}
	this.Type = type_
	this.Url = url
	return &this
}

// NewStoryAreaTypeLinkWithDefaults instantiates a new StoryAreaTypeLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryAreaTypeLinkWithDefaults() *StoryAreaTypeLink {
	this := StoryAreaTypeLink{}
	var type_ string = "link"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StoryAreaTypeLink) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLink) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StoryAreaTypeLink) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *StoryAreaTypeLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *StoryAreaTypeLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *StoryAreaTypeLink) SetUrl(v string) {
	o.Url = v
}

func (o StoryAreaTypeLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryAreaTypeLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *StoryAreaTypeLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"url",
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

	varStoryAreaTypeLink := _StoryAreaTypeLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStoryAreaTypeLink)

	if err != nil {
		return err
	}

	*o = StoryAreaTypeLink(varStoryAreaTypeLink)

	return err
}

type NullableStoryAreaTypeLink struct {
	value *StoryAreaTypeLink
	isSet bool
}

func (v NullableStoryAreaTypeLink) Get() *StoryAreaTypeLink {
	return v.value
}

func (v *NullableStoryAreaTypeLink) Set(val *StoryAreaTypeLink) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryAreaTypeLink) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryAreaTypeLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryAreaTypeLink(val *StoryAreaTypeLink) *NullableStoryAreaTypeLink {
	return &NullableStoryAreaTypeLink{value: val, isSet: true}
}

func (v NullableStoryAreaTypeLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryAreaTypeLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


