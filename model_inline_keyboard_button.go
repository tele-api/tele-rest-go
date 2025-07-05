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

// checks if the InlineKeyboardButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineKeyboardButton{}

// InlineKeyboardButton This object represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify type of the button.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`
	// *Optional*. HTTP or tg:// URL to be opened when the button is pressed. Links `tg://user?id=<user_id>` can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	Url *string `json:"url,omitempty"`
	// *Optional*. Data to be sent in a [callback query](https://core.telegram.org/bots/api/#callbackquery) to the bot when the button is pressed, 1-64 bytes
	CallbackData *string `json:"callback_data,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	LoginUrl *LoginUrl `json:"login_url,omitempty"`
	// *Optional*. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
	// *Optional*. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.    This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText *CopyTextButton `json:"copy_text,omitempty"`
	CallbackGame interface{} `json:"callback_game,omitempty"`
	// *Optional*. Specify *True*, to send a [Pay button](https://core.telegram.org/bots/api/#payments). Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.    **NOTE:** This type of button **must** always be the first button in the first row and can only be used in invoice messages.
	Pay *bool `json:"pay,omitempty"`
}

type _InlineKeyboardButton InlineKeyboardButton

// NewInlineKeyboardButton instantiates a new InlineKeyboardButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineKeyboardButton(text string) *InlineKeyboardButton {
	this := InlineKeyboardButton{}
	this.Text = text
	return &this
}

// NewInlineKeyboardButtonWithDefaults instantiates a new InlineKeyboardButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineKeyboardButtonWithDefaults() *InlineKeyboardButton {
	this := InlineKeyboardButton{}
	return &this
}

// GetText returns the Text field value
func (o *InlineKeyboardButton) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *InlineKeyboardButton) SetText(v string) {
	o.Text = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineKeyboardButton) SetUrl(v string) {
	o.Url = &v
}


// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *InlineKeyboardButton) SetCallbackData(v string) {
	o.CallbackData = &v
}


// GetWebApp returns the WebApp field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetWebApp() WebAppInfo {
	if o == nil || IsNil(o.WebApp) {
		var ret WebAppInfo
		return ret
	}
	return *o.WebApp
}

// GetWebAppOk returns a tuple with the WebApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetWebAppOk() (*WebAppInfo, bool) {
	if o == nil || IsNil(o.WebApp) {
		return nil, false
	}
	return o.WebApp, true
}

// HasWebApp returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasWebApp() bool {
	if o != nil && !IsNil(o.WebApp) {
		return true
	}

	return false
}

// SetWebApp gets a reference to the given WebAppInfo and assigns it to the WebApp field.
func (o *InlineKeyboardButton) SetWebApp(v WebAppInfo) {
	o.WebApp = &v
}


// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetLoginUrl() LoginUrl {
	if o == nil || IsNil(o.LoginUrl) {
		var ret LoginUrl
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetLoginUrlOk() (*LoginUrl, bool) {
	if o == nil || IsNil(o.LoginUrl) {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasLoginUrl() bool {
	if o != nil && !IsNil(o.LoginUrl) {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given LoginUrl and assigns it to the LoginUrl field.
func (o *InlineKeyboardButton) SetLoginUrl(v LoginUrl) {
	o.LoginUrl = &v
}


// GetSwitchInlineQuery returns the SwitchInlineQuery field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetSwitchInlineQuery() string {
	if o == nil || IsNil(o.SwitchInlineQuery) {
		var ret string
		return ret
	}
	return *o.SwitchInlineQuery
}

// GetSwitchInlineQueryOk returns a tuple with the SwitchInlineQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetSwitchInlineQueryOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchInlineQuery) {
		return nil, false
	}
	return o.SwitchInlineQuery, true
}

// HasSwitchInlineQuery returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasSwitchInlineQuery() bool {
	if o != nil && !IsNil(o.SwitchInlineQuery) {
		return true
	}

	return false
}

// SetSwitchInlineQuery gets a reference to the given string and assigns it to the SwitchInlineQuery field.
func (o *InlineKeyboardButton) SetSwitchInlineQuery(v string) {
	o.SwitchInlineQuery = &v
}


// GetSwitchInlineQueryCurrentChat returns the SwitchInlineQueryCurrentChat field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetSwitchInlineQueryCurrentChat() string {
	if o == nil || IsNil(o.SwitchInlineQueryCurrentChat) {
		var ret string
		return ret
	}
	return *o.SwitchInlineQueryCurrentChat
}

// GetSwitchInlineQueryCurrentChatOk returns a tuple with the SwitchInlineQueryCurrentChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetSwitchInlineQueryCurrentChatOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchInlineQueryCurrentChat) {
		return nil, false
	}
	return o.SwitchInlineQueryCurrentChat, true
}

// HasSwitchInlineQueryCurrentChat returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasSwitchInlineQueryCurrentChat() bool {
	if o != nil && !IsNil(o.SwitchInlineQueryCurrentChat) {
		return true
	}

	return false
}

// SetSwitchInlineQueryCurrentChat gets a reference to the given string and assigns it to the SwitchInlineQueryCurrentChat field.
func (o *InlineKeyboardButton) SetSwitchInlineQueryCurrentChat(v string) {
	o.SwitchInlineQueryCurrentChat = &v
}


// GetSwitchInlineQueryChosenChat returns the SwitchInlineQueryChosenChat field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetSwitchInlineQueryChosenChat() SwitchInlineQueryChosenChat {
	if o == nil || IsNil(o.SwitchInlineQueryChosenChat) {
		var ret SwitchInlineQueryChosenChat
		return ret
	}
	return *o.SwitchInlineQueryChosenChat
}

// GetSwitchInlineQueryChosenChatOk returns a tuple with the SwitchInlineQueryChosenChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetSwitchInlineQueryChosenChatOk() (*SwitchInlineQueryChosenChat, bool) {
	if o == nil || IsNil(o.SwitchInlineQueryChosenChat) {
		return nil, false
	}
	return o.SwitchInlineQueryChosenChat, true
}

// HasSwitchInlineQueryChosenChat returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasSwitchInlineQueryChosenChat() bool {
	if o != nil && !IsNil(o.SwitchInlineQueryChosenChat) {
		return true
	}

	return false
}

// SetSwitchInlineQueryChosenChat gets a reference to the given SwitchInlineQueryChosenChat and assigns it to the SwitchInlineQueryChosenChat field.
func (o *InlineKeyboardButton) SetSwitchInlineQueryChosenChat(v SwitchInlineQueryChosenChat) {
	o.SwitchInlineQueryChosenChat = &v
}


// GetCopyText returns the CopyText field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetCopyText() CopyTextButton {
	if o == nil || IsNil(o.CopyText) {
		var ret CopyTextButton
		return ret
	}
	return *o.CopyText
}

// GetCopyTextOk returns a tuple with the CopyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetCopyTextOk() (*CopyTextButton, bool) {
	if o == nil || IsNil(o.CopyText) {
		return nil, false
	}
	return o.CopyText, true
}

// HasCopyText returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasCopyText() bool {
	if o != nil && !IsNil(o.CopyText) {
		return true
	}

	return false
}

// SetCopyText gets a reference to the given CopyTextButton and assigns it to the CopyText field.
func (o *InlineKeyboardButton) SetCopyText(v CopyTextButton) {
	o.CopyText = &v
}


// GetCallbackGame returns the CallbackGame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineKeyboardButton) GetCallbackGame() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CallbackGame
}

// GetCallbackGameOk returns a tuple with the CallbackGame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineKeyboardButton) GetCallbackGameOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackGame, true
}

// HasCallbackGame returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasCallbackGame() bool {
	if o != nil && !IsNil(o.CallbackGame) {
		return true
	}

	return false
}

// SetCallbackGame gets a reference to the given interface{} and assigns it to the CallbackGame field.
func (o *InlineKeyboardButton) SetCallbackGame(v interface{}) {
	o.CallbackGame = v
}


// GetPay returns the Pay field value if set, zero value otherwise.
func (o *InlineKeyboardButton) GetPay() bool {
	if o == nil || IsNil(o.Pay) {
		var ret bool
		return ret
	}
	return *o.Pay
}

// GetPayOk returns a tuple with the Pay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineKeyboardButton) GetPayOk() (*bool, bool) {
	if o == nil || IsNil(o.Pay) {
		return nil, false
	}
	return o.Pay, true
}

// HasPay returns a boolean if a field has been set.
func (o *InlineKeyboardButton) HasPay() bool {
	if o != nil && !IsNil(o.Pay) {
		return true
	}

	return false
}

// SetPay gets a reference to the given bool and assigns it to the Pay field.
func (o *InlineKeyboardButton) SetPay(v bool) {
	o.Pay = &v
}


func (o InlineKeyboardButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineKeyboardButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callback_data"] = o.CallbackData
	}
	if !IsNil(o.WebApp) {
		toSerialize["web_app"] = o.WebApp
	}
	if !IsNil(o.LoginUrl) {
		toSerialize["login_url"] = o.LoginUrl
	}
	if !IsNil(o.SwitchInlineQuery) {
		toSerialize["switch_inline_query"] = o.SwitchInlineQuery
	}
	if !IsNil(o.SwitchInlineQueryCurrentChat) {
		toSerialize["switch_inline_query_current_chat"] = o.SwitchInlineQueryCurrentChat
	}
	if !IsNil(o.SwitchInlineQueryChosenChat) {
		toSerialize["switch_inline_query_chosen_chat"] = o.SwitchInlineQueryChosenChat
	}
	if !IsNil(o.CopyText) {
		toSerialize["copy_text"] = o.CopyText
	}
	if o.CallbackGame != nil {
		toSerialize["callback_game"] = o.CallbackGame
	}
	if !IsNil(o.Pay) {
		toSerialize["pay"] = o.Pay
	}
	return toSerialize, nil
}

func (o *InlineKeyboardButton) UnmarshalJSON(data []byte) (err error) {
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

	varInlineKeyboardButton := _InlineKeyboardButton{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineKeyboardButton)

	if err != nil {
		return err
	}

	*o = InlineKeyboardButton(varInlineKeyboardButton)

	return err
}

type NullableInlineKeyboardButton struct {
	value *InlineKeyboardButton
	isSet bool
}

func (v NullableInlineKeyboardButton) Get() *InlineKeyboardButton {
	return v.value
}

func (v *NullableInlineKeyboardButton) Set(val *InlineKeyboardButton) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineKeyboardButton) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineKeyboardButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineKeyboardButton(val *InlineKeyboardButton) *NullableInlineKeyboardButton {
	return &NullableInlineKeyboardButton{value: val, isSet: true}
}

func (v NullableInlineKeyboardButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineKeyboardButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


