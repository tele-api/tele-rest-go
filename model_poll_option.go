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

// checks if the PollOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollOption{}

// PollOption This object contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// *Optional*. Special entities that appear in the option *text*. Currently, only custom emoji entities are allowed in poll option texts
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	// Number of users that voted for this option
	VoterCount int32 `json:"voter_count"`
}

type _PollOption PollOption

// NewPollOption instantiates a new PollOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollOption(text string, voterCount int32) *PollOption {
	this := PollOption{}
	this.Text = text
	this.VoterCount = voterCount
	return &this
}

// NewPollOptionWithDefaults instantiates a new PollOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollOptionWithDefaults() *PollOption {
	this := PollOption{}
	return &this
}

// GetText returns the Text field value
func (o *PollOption) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *PollOption) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *PollOption) SetText(v string) {
	o.Text = v
}

// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *PollOption) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollOption) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *PollOption) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *PollOption) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


// GetVoterCount returns the VoterCount field value
func (o *PollOption) GetVoterCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VoterCount
}

// GetVoterCountOk returns a tuple with the VoterCount field value
// and a boolean to check if the value has been set.
func (o *PollOption) GetVoterCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoterCount, true
}

// SetVoterCount sets field value
func (o *PollOption) SetVoterCount(v int32) {
	o.VoterCount = v
}

func (o PollOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.TextEntities) {
		toSerialize["text_entities"] = o.TextEntities
	}
	toSerialize["voter_count"] = o.VoterCount
	return toSerialize, nil
}

func (o *PollOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
		"voter_count",
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

	varPollOption := _PollOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollOption)

	if err != nil {
		return err
	}

	*o = PollOption(varPollOption)

	return err
}

type NullablePollOption struct {
	value *PollOption
	isSet bool
}

func (v NullablePollOption) Get() *PollOption {
	return v.value
}

func (v *NullablePollOption) Set(val *PollOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePollOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePollOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollOption(val *PollOption) *NullablePollOption {
	return &NullablePollOption{value: val, isSet: true}
}

func (v NullablePollOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


