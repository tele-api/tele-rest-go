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

// checks if the InlineQueryResultsButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQueryResultsButton{}

// InlineQueryResultsButton This object represents a button to be shown above inline query results. You **must** use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	// Label text on the button
	Text string `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// *Optional*. [Deep-linking](https://core.telegram.org/bots/features#deep-linking) parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only `A-Z`, `a-z`, `0-9`, `_` and `-` are allowed.    *Example:* An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a [*switch\\_inline*](https://core.telegram.org/bots/api/#inlinekeyboardmarkup) button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
	StartParameter *string `json:"start_parameter,omitempty"`
}

type _InlineQueryResultsButton InlineQueryResultsButton

// NewInlineQueryResultsButton instantiates a new InlineQueryResultsButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQueryResultsButton(text string) *InlineQueryResultsButton {
	this := InlineQueryResultsButton{}
	this.Text = text
	return &this
}

// NewInlineQueryResultsButtonWithDefaults instantiates a new InlineQueryResultsButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryResultsButtonWithDefaults() *InlineQueryResultsButton {
	this := InlineQueryResultsButton{}
	return &this
}

// GetText returns the Text field value
func (o *InlineQueryResultsButton) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *InlineQueryResultsButton) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *InlineQueryResultsButton) SetText(v string) {
	o.Text = v
}

// GetWebApp returns the WebApp field value if set, zero value otherwise.
func (o *InlineQueryResultsButton) GetWebApp() WebAppInfo {
	if o == nil || IsNil(o.WebApp) {
		var ret WebAppInfo
		return ret
	}
	return *o.WebApp
}

// GetWebAppOk returns a tuple with the WebApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultsButton) GetWebAppOk() (*WebAppInfo, bool) {
	if o == nil || IsNil(o.WebApp) {
		return nil, false
	}
	return o.WebApp, true
}

// HasWebApp returns a boolean if a field has been set.
func (o *InlineQueryResultsButton) HasWebApp() bool {
	if o != nil && !IsNil(o.WebApp) {
		return true
	}

	return false
}

// SetWebApp gets a reference to the given WebAppInfo and assigns it to the WebApp field.
func (o *InlineQueryResultsButton) SetWebApp(v WebAppInfo) {
	o.WebApp = &v
}


// GetStartParameter returns the StartParameter field value if set, zero value otherwise.
func (o *InlineQueryResultsButton) GetStartParameter() string {
	if o == nil || IsNil(o.StartParameter) {
		var ret string
		return ret
	}
	return *o.StartParameter
}

// GetStartParameterOk returns a tuple with the StartParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQueryResultsButton) GetStartParameterOk() (*string, bool) {
	if o == nil || IsNil(o.StartParameter) {
		return nil, false
	}
	return o.StartParameter, true
}

// HasStartParameter returns a boolean if a field has been set.
func (o *InlineQueryResultsButton) HasStartParameter() bool {
	if o != nil && !IsNil(o.StartParameter) {
		return true
	}

	return false
}

// SetStartParameter gets a reference to the given string and assigns it to the StartParameter field.
func (o *InlineQueryResultsButton) SetStartParameter(v string) {
	o.StartParameter = &v
}


func (o InlineQueryResultsButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQueryResultsButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.WebApp) {
		toSerialize["web_app"] = o.WebApp
	}
	if !IsNil(o.StartParameter) {
		toSerialize["start_parameter"] = o.StartParameter
	}
	return toSerialize, nil
}

func (o *InlineQueryResultsButton) UnmarshalJSON(data []byte) (err error) {
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

	varInlineQueryResultsButton := _InlineQueryResultsButton{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQueryResultsButton)

	if err != nil {
		return err
	}

	*o = InlineQueryResultsButton(varInlineQueryResultsButton)

	return err
}

type NullableInlineQueryResultsButton struct {
	value *InlineQueryResultsButton
	isSet bool
}

func (v NullableInlineQueryResultsButton) Get() *InlineQueryResultsButton {
	return v.value
}

func (v *NullableInlineQueryResultsButton) Set(val *InlineQueryResultsButton) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQueryResultsButton) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQueryResultsButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQueryResultsButton(val *InlineQueryResultsButton) *NullableInlineQueryResultsButton {
	return &NullableInlineQueryResultsButton{value: val, isSet: true}
}

func (v NullableInlineQueryResultsButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQueryResultsButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


