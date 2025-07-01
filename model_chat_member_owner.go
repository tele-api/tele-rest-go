/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the ChatMemberOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatMemberOwner{}

// ChatMemberOwner Represents a [chat member](https://core.telegram.org/bots/api/#chatmember) that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	// The member's status in the chat, always “creator”
	Status string `json:"status"`
	User User `json:"user"`
	// *True*, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous"`
	// *Optional*. Custom title for this user
	CustomTitle *string `json:"custom_title,omitempty"`
}

type _ChatMemberOwner ChatMemberOwner

// NewChatMemberOwner instantiates a new ChatMemberOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatMemberOwner(status string, user User, isAnonymous bool) *ChatMemberOwner {
	this := ChatMemberOwner{}
	this.Status = status
	this.User = user
	this.IsAnonymous = isAnonymous
	return &this
}

// NewChatMemberOwnerWithDefaults instantiates a new ChatMemberOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatMemberOwnerWithDefaults() *ChatMemberOwner {
	this := ChatMemberOwner{}
	var status string = "creator"
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *ChatMemberOwner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChatMemberOwner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChatMemberOwner) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value
func (o *ChatMemberOwner) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ChatMemberOwner) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ChatMemberOwner) SetUser(v User) {
	o.User = v
}

// GetIsAnonymous returns the IsAnonymous field value
func (o *ChatMemberOwner) GetIsAnonymous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value
// and a boolean to check if the value has been set.
func (o *ChatMemberOwner) GetIsAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnonymous, true
}

// SetIsAnonymous sets field value
func (o *ChatMemberOwner) SetIsAnonymous(v bool) {
	o.IsAnonymous = v
}

// GetCustomTitle returns the CustomTitle field value if set, zero value otherwise.
func (o *ChatMemberOwner) GetCustomTitle() string {
	if o == nil || IsNil(o.CustomTitle) {
		var ret string
		return ret
	}
	return *o.CustomTitle
}

// GetCustomTitleOk returns a tuple with the CustomTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatMemberOwner) GetCustomTitleOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTitle) {
		return nil, false
	}
	return o.CustomTitle, true
}

// HasCustomTitle returns a boolean if a field has been set.
func (o *ChatMemberOwner) HasCustomTitle() bool {
	if o != nil && !IsNil(o.CustomTitle) {
		return true
	}

	return false
}

// SetCustomTitle gets a reference to the given string and assigns it to the CustomTitle field.
func (o *ChatMemberOwner) SetCustomTitle(v string) {
	o.CustomTitle = &v
}


func (o ChatMemberOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatMemberOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["user"] = o.User
	toSerialize["is_anonymous"] = o.IsAnonymous
	if !IsNil(o.CustomTitle) {
		toSerialize["custom_title"] = o.CustomTitle
	}
	return toSerialize, nil
}

func (o *ChatMemberOwner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"user",
		"is_anonymous",
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

	varChatMemberOwner := _ChatMemberOwner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatMemberOwner)

	if err != nil {
		return err
	}

	*o = ChatMemberOwner(varChatMemberOwner)

	return err
}

type NullableChatMemberOwner struct {
	value *ChatMemberOwner
	isSet bool
}

func (v NullableChatMemberOwner) Get() *ChatMemberOwner {
	return v.value
}

func (v *NullableChatMemberOwner) Set(val *ChatMemberOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableChatMemberOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableChatMemberOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatMemberOwner(val *ChatMemberOwner) *NullableChatMemberOwner {
	return &NullableChatMemberOwner{value: val, isSet: true}
}

func (v NullableChatMemberOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatMemberOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


