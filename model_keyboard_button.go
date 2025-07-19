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

// checks if the KeyboardButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyboardButton{}

// KeyboardButton This object represents one button of the reply keyboard. At most one of the optional fields must be used to specify type of the button. For simple text buttons, *String* can be used instead of this object to specify the button text.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	Text string `json:"text"`
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	// *Optional*. If *True*, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestContact *bool `json:"request_contact,omitempty"`
	// *Optional*. If *True*, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestLocation *bool `json:"request_location,omitempty"`
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

type _KeyboardButton KeyboardButton

// NewKeyboardButton instantiates a new KeyboardButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyboardButton(text string) *KeyboardButton {
	this := KeyboardButton{}
	this.Text = text
	return &this
}

// NewKeyboardButtonWithDefaults instantiates a new KeyboardButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyboardButtonWithDefaults() *KeyboardButton {
	this := KeyboardButton{}
	return &this
}

// GetText returns the Text field value
func (o *KeyboardButton) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *KeyboardButton) SetText(v string) {
	o.Text = v
}

// GetRequestUsers returns the RequestUsers field value if set, zero value otherwise.
func (o *KeyboardButton) GetRequestUsers() KeyboardButtonRequestUsers {
	if o == nil || IsNil(o.RequestUsers) {
		var ret KeyboardButtonRequestUsers
		return ret
	}
	return *o.RequestUsers
}

// GetRequestUsersOk returns a tuple with the RequestUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetRequestUsersOk() (*KeyboardButtonRequestUsers, bool) {
	if o == nil || IsNil(o.RequestUsers) {
		return nil, false
	}
	return o.RequestUsers, true
}

// HasRequestUsers returns a boolean if a field has been set.
func (o *KeyboardButton) HasRequestUsers() bool {
	if o != nil && !IsNil(o.RequestUsers) {
		return true
	}

	return false
}

// SetRequestUsers gets a reference to the given KeyboardButtonRequestUsers and assigns it to the RequestUsers field.
func (o *KeyboardButton) SetRequestUsers(v KeyboardButtonRequestUsers) {
	o.RequestUsers = &v
}


// GetRequestChat returns the RequestChat field value if set, zero value otherwise.
func (o *KeyboardButton) GetRequestChat() KeyboardButtonRequestChat {
	if o == nil || IsNil(o.RequestChat) {
		var ret KeyboardButtonRequestChat
		return ret
	}
	return *o.RequestChat
}

// GetRequestChatOk returns a tuple with the RequestChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetRequestChatOk() (*KeyboardButtonRequestChat, bool) {
	if o == nil || IsNil(o.RequestChat) {
		return nil, false
	}
	return o.RequestChat, true
}

// HasRequestChat returns a boolean if a field has been set.
func (o *KeyboardButton) HasRequestChat() bool {
	if o != nil && !IsNil(o.RequestChat) {
		return true
	}

	return false
}

// SetRequestChat gets a reference to the given KeyboardButtonRequestChat and assigns it to the RequestChat field.
func (o *KeyboardButton) SetRequestChat(v KeyboardButtonRequestChat) {
	o.RequestChat = &v
}


// GetRequestContact returns the RequestContact field value if set, zero value otherwise.
func (o *KeyboardButton) GetRequestContact() bool {
	if o == nil || IsNil(o.RequestContact) {
		var ret bool
		return ret
	}
	return *o.RequestContact
}

// GetRequestContactOk returns a tuple with the RequestContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetRequestContactOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestContact) {
		return nil, false
	}
	return o.RequestContact, true
}

// HasRequestContact returns a boolean if a field has been set.
func (o *KeyboardButton) HasRequestContact() bool {
	if o != nil && !IsNil(o.RequestContact) {
		return true
	}

	return false
}

// SetRequestContact gets a reference to the given bool and assigns it to the RequestContact field.
func (o *KeyboardButton) SetRequestContact(v bool) {
	o.RequestContact = &v
}


// GetRequestLocation returns the RequestLocation field value if set, zero value otherwise.
func (o *KeyboardButton) GetRequestLocation() bool {
	if o == nil || IsNil(o.RequestLocation) {
		var ret bool
		return ret
	}
	return *o.RequestLocation
}

// GetRequestLocationOk returns a tuple with the RequestLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetRequestLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestLocation) {
		return nil, false
	}
	return o.RequestLocation, true
}

// HasRequestLocation returns a boolean if a field has been set.
func (o *KeyboardButton) HasRequestLocation() bool {
	if o != nil && !IsNil(o.RequestLocation) {
		return true
	}

	return false
}

// SetRequestLocation gets a reference to the given bool and assigns it to the RequestLocation field.
func (o *KeyboardButton) SetRequestLocation(v bool) {
	o.RequestLocation = &v
}


// GetRequestPoll returns the RequestPoll field value if set, zero value otherwise.
func (o *KeyboardButton) GetRequestPoll() KeyboardButtonPollType {
	if o == nil || IsNil(o.RequestPoll) {
		var ret KeyboardButtonPollType
		return ret
	}
	return *o.RequestPoll
}

// GetRequestPollOk returns a tuple with the RequestPoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetRequestPollOk() (*KeyboardButtonPollType, bool) {
	if o == nil || IsNil(o.RequestPoll) {
		return nil, false
	}
	return o.RequestPoll, true
}

// HasRequestPoll returns a boolean if a field has been set.
func (o *KeyboardButton) HasRequestPoll() bool {
	if o != nil && !IsNil(o.RequestPoll) {
		return true
	}

	return false
}

// SetRequestPoll gets a reference to the given KeyboardButtonPollType and assigns it to the RequestPoll field.
func (o *KeyboardButton) SetRequestPoll(v KeyboardButtonPollType) {
	o.RequestPoll = &v
}


// GetWebApp returns the WebApp field value if set, zero value otherwise.
func (o *KeyboardButton) GetWebApp() WebAppInfo {
	if o == nil || IsNil(o.WebApp) {
		var ret WebAppInfo
		return ret
	}
	return *o.WebApp
}

// GetWebAppOk returns a tuple with the WebApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButton) GetWebAppOk() (*WebAppInfo, bool) {
	if o == nil || IsNil(o.WebApp) {
		return nil, false
	}
	return o.WebApp, true
}

// HasWebApp returns a boolean if a field has been set.
func (o *KeyboardButton) HasWebApp() bool {
	if o != nil && !IsNil(o.WebApp) {
		return true
	}

	return false
}

// SetWebApp gets a reference to the given WebAppInfo and assigns it to the WebApp field.
func (o *KeyboardButton) SetWebApp(v WebAppInfo) {
	o.WebApp = &v
}


func (o KeyboardButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyboardButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.RequestUsers) {
		toSerialize["request_users"] = o.RequestUsers
	}
	if !IsNil(o.RequestChat) {
		toSerialize["request_chat"] = o.RequestChat
	}
	if !IsNil(o.RequestContact) {
		toSerialize["request_contact"] = o.RequestContact
	}
	if !IsNil(o.RequestLocation) {
		toSerialize["request_location"] = o.RequestLocation
	}
	if !IsNil(o.RequestPoll) {
		toSerialize["request_poll"] = o.RequestPoll
	}
	if !IsNil(o.WebApp) {
		toSerialize["web_app"] = o.WebApp
	}
	return toSerialize, nil
}

func (o *KeyboardButton) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varKeyboardButton := _KeyboardButton{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyboardButton)

	if err != nil {
		return err
	}

	*o = KeyboardButton(varKeyboardButton)

	return err
}

type NullableKeyboardButton struct {
	value *KeyboardButton
	isSet bool
}

func (v NullableKeyboardButton) Get() *KeyboardButton {
	return v.value
}

func (v *NullableKeyboardButton) Set(val *KeyboardButton) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyboardButton) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyboardButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyboardButton(val *KeyboardButton) *NullableKeyboardButton {
	return &NullableKeyboardButton{value: val, isSet: true}
}

func (v NullableKeyboardButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyboardButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


