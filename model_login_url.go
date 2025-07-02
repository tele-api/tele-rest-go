/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the LoginUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginUrl{}

// LoginUrl This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the [Telegram Login Widget](https://core.telegram.org/widgets/login) when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:  Telegram apps support these buttons as of [version 5.7](https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots).  Sample bot: [@discussbot](https://t.me/discussbot)
type LoginUrl struct {
	// An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in [Receiving authorization data](https://core.telegram.org/widgets/login#receiving-authorization-data).    **NOTE:** You **must** always check the hash of the received data to verify the authentication and the integrity of the data as described in [Checking authorization](https://core.telegram.org/widgets/login#checking-authorization).
	Url string `json:"url"`
	// *Optional*. New text of the button in forwarded messages.
	ForwardText *string `json:"forward_text,omitempty"`
	// *Optional*. Username of a bot, which will be used for user authorization. See [Setting up a bot](https://core.telegram.org/widgets/login#setting-up-a-bot) for more details. If not specified, the current bot's username will be assumed. The *url*'s domain must be the same as the domain linked with the bot. See [Linking your domain to the bot](https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot) for more details.
	BotUsername *string `json:"bot_username,omitempty"`
	// *Optional*. Pass *True* to request the permission for your bot to send messages to the user.
	RequestWriteAccess *bool `json:"request_write_access,omitempty"`
}

type _LoginUrl LoginUrl

// NewLoginUrl instantiates a new LoginUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginUrl(url string) *LoginUrl {
	this := LoginUrl{}
	this.Url = url
	return &this
}

// NewLoginUrlWithDefaults instantiates a new LoginUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginUrlWithDefaults() *LoginUrl {
	this := LoginUrl{}
	return &this
}

// GetUrl returns the Url field value
func (o *LoginUrl) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LoginUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LoginUrl) SetUrl(v string) {
	o.Url = v
}

// GetForwardText returns the ForwardText field value if set, zero value otherwise.
func (o *LoginUrl) GetForwardText() string {
	if o == nil || IsNil(o.ForwardText) {
		var ret string
		return ret
	}
	return *o.ForwardText
}

// GetForwardTextOk returns a tuple with the ForwardText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginUrl) GetForwardTextOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardText) {
		return nil, false
	}
	return o.ForwardText, true
}

// HasForwardText returns a boolean if a field has been set.
func (o *LoginUrl) HasForwardText() bool {
	if o != nil && !IsNil(o.ForwardText) {
		return true
	}

	return false
}

// SetForwardText gets a reference to the given string and assigns it to the ForwardText field.
func (o *LoginUrl) SetForwardText(v string) {
	o.ForwardText = &v
}


// GetBotUsername returns the BotUsername field value if set, zero value otherwise.
func (o *LoginUrl) GetBotUsername() string {
	if o == nil || IsNil(o.BotUsername) {
		var ret string
		return ret
	}
	return *o.BotUsername
}

// GetBotUsernameOk returns a tuple with the BotUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginUrl) GetBotUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BotUsername) {
		return nil, false
	}
	return o.BotUsername, true
}

// HasBotUsername returns a boolean if a field has been set.
func (o *LoginUrl) HasBotUsername() bool {
	if o != nil && !IsNil(o.BotUsername) {
		return true
	}

	return false
}

// SetBotUsername gets a reference to the given string and assigns it to the BotUsername field.
func (o *LoginUrl) SetBotUsername(v string) {
	o.BotUsername = &v
}


// GetRequestWriteAccess returns the RequestWriteAccess field value if set, zero value otherwise.
func (o *LoginUrl) GetRequestWriteAccess() bool {
	if o == nil || IsNil(o.RequestWriteAccess) {
		var ret bool
		return ret
	}
	return *o.RequestWriteAccess
}

// GetRequestWriteAccessOk returns a tuple with the RequestWriteAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginUrl) GetRequestWriteAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestWriteAccess) {
		return nil, false
	}
	return o.RequestWriteAccess, true
}

// HasRequestWriteAccess returns a boolean if a field has been set.
func (o *LoginUrl) HasRequestWriteAccess() bool {
	if o != nil && !IsNil(o.RequestWriteAccess) {
		return true
	}

	return false
}

// SetRequestWriteAccess gets a reference to the given bool and assigns it to the RequestWriteAccess field.
func (o *LoginUrl) SetRequestWriteAccess(v bool) {
	o.RequestWriteAccess = &v
}


func (o LoginUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.ForwardText) {
		toSerialize["forward_text"] = o.ForwardText
	}
	if !IsNil(o.BotUsername) {
		toSerialize["bot_username"] = o.BotUsername
	}
	if !IsNil(o.RequestWriteAccess) {
		toSerialize["request_write_access"] = o.RequestWriteAccess
	}
	return toSerialize, nil
}

func (o *LoginUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varLoginUrl := _LoginUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginUrl)

	if err != nil {
		return err
	}

	*o = LoginUrl(varLoginUrl)

	return err
}

type NullableLoginUrl struct {
	value *LoginUrl
	isSet bool
}

func (v NullableLoginUrl) Get() *LoginUrl {
	return v.value
}

func (v *NullableLoginUrl) Set(val *LoginUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginUrl(val *LoginUrl) *NullableLoginUrl {
	return &NullableLoginUrl{value: val, isSet: true}
}

func (v NullableLoginUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


