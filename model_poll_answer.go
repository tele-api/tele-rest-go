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

// checks if the PollAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollAnswer{}

// PollAnswer This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// Unique poll identifier
	PollId string `json:"poll_id"`
	VoterChat *Chat `json:"voter_chat,omitempty"`
	User *User `json:"user,omitempty"`
	// 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
	OptionIds []int32 `json:"option_ids"`
}

type _PollAnswer PollAnswer

// NewPollAnswer instantiates a new PollAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollAnswer(pollId string, optionIds []int32) *PollAnswer {
	this := PollAnswer{}
	this.PollId = pollId
	this.OptionIds = optionIds
	return &this
}

// NewPollAnswerWithDefaults instantiates a new PollAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollAnswerWithDefaults() *PollAnswer {
	this := PollAnswer{}
	return &this
}

// GetPollId returns the PollId field value
func (o *PollAnswer) GetPollId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PollId
}

// GetPollIdOk returns a tuple with the PollId field value
// and a boolean to check if the value has been set.
func (o *PollAnswer) GetPollIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollId, true
}

// SetPollId sets field value
func (o *PollAnswer) SetPollId(v string) {
	o.PollId = v
}

// GetVoterChat returns the VoterChat field value if set, zero value otherwise.
func (o *PollAnswer) GetVoterChat() Chat {
	if o == nil || IsNil(o.VoterChat) {
		var ret Chat
		return ret
	}
	return *o.VoterChat
}

// GetVoterChatOk returns a tuple with the VoterChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollAnswer) GetVoterChatOk() (*Chat, bool) {
	if o == nil || IsNil(o.VoterChat) {
		return nil, false
	}
	return o.VoterChat, true
}

// HasVoterChat returns a boolean if a field has been set.
func (o *PollAnswer) HasVoterChat() bool {
	if o != nil && !IsNil(o.VoterChat) {
		return true
	}

	return false
}

// SetVoterChat gets a reference to the given Chat and assigns it to the VoterChat field.
func (o *PollAnswer) SetVoterChat(v Chat) {
	o.VoterChat = &v
}


// GetUser returns the User field value if set, zero value otherwise.
func (o *PollAnswer) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollAnswer) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PollAnswer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *PollAnswer) SetUser(v User) {
	o.User = &v
}


// GetOptionIds returns the OptionIds field value
func (o *PollAnswer) GetOptionIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value
// and a boolean to check if the value has been set.
func (o *PollAnswer) GetOptionIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptionIds, true
}

// SetOptionIds sets field value
func (o *PollAnswer) SetOptionIds(v []int32) {
	o.OptionIds = v
}

func (o PollAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollAnswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["poll_id"] = o.PollId
	if !IsNil(o.VoterChat) {
		toSerialize["voter_chat"] = o.VoterChat
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["option_ids"] = o.OptionIds
	return toSerialize, nil
}

func (o *PollAnswer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"poll_id",
		"option_ids",
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

	varPollAnswer := _PollAnswer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollAnswer)

	if err != nil {
		return err
	}

	*o = PollAnswer(varPollAnswer)

	return err
}

type NullablePollAnswer struct {
	value *PollAnswer
	isSet bool
}

func (v NullablePollAnswer) Get() *PollAnswer {
	return v.value
}

func (v *NullablePollAnswer) Set(val *PollAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullablePollAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullablePollAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollAnswer(val *PollAnswer) *NullablePollAnswer {
	return &NullablePollAnswer{value: val, isSet: true}
}

func (v NullablePollAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


