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

// checks if the PostAnswerCallbackQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAnswerCallbackQueryRequest{}

// PostAnswerCallbackQueryRequest struct for PostAnswerCallbackQueryRequest
type PostAnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryId string `json:"callback_query_id"`
	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text *string `json:"text,omitempty"`
	// If *True*, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to *false*.
	ShowAlert *bool `json:"show_alert,omitempty"`
	// URL that will be opened by the user's client. If you have created a [Game](https://core.telegram.org/bots/api/#game) and accepted the conditions via [@BotFather](https://t.me/botfather), specify the URL that opens your game - note that this will only work if the query comes from a [*callback\\_game*](https://core.telegram.org/bots/api/#inlinekeyboardbutton) button.    Otherwise, you may use links like `t.me/your_bot?start=XXXX` that open your bot with a parameter.
	Url *string `json:"url,omitempty"`
	// The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime *int32 `json:"cache_time,omitempty"`
}

type _PostAnswerCallbackQueryRequest PostAnswerCallbackQueryRequest

// NewPostAnswerCallbackQueryRequest instantiates a new PostAnswerCallbackQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAnswerCallbackQueryRequest(callbackQueryId string) *PostAnswerCallbackQueryRequest {
	this := PostAnswerCallbackQueryRequest{}
	this.CallbackQueryId = callbackQueryId
	var showAlert bool = false
	this.ShowAlert = &showAlert
	var cacheTime int32 = 0
	this.CacheTime = &cacheTime
	return &this
}

// NewPostAnswerCallbackQueryRequestWithDefaults instantiates a new PostAnswerCallbackQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAnswerCallbackQueryRequestWithDefaults() *PostAnswerCallbackQueryRequest {
	this := PostAnswerCallbackQueryRequest{}
	var showAlert bool = false
	this.ShowAlert = &showAlert
	var cacheTime int32 = 0
	this.CacheTime = &cacheTime
	return &this
}

// GetCallbackQueryId returns the CallbackQueryId field value
func (o *PostAnswerCallbackQueryRequest) GetCallbackQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackQueryId
}

// GetCallbackQueryIdOk returns a tuple with the CallbackQueryId field value
// and a boolean to check if the value has been set.
func (o *PostAnswerCallbackQueryRequest) GetCallbackQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackQueryId, true
}

// SetCallbackQueryId sets field value
func (o *PostAnswerCallbackQueryRequest) SetCallbackQueryId(v string) {
	o.CallbackQueryId = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PostAnswerCallbackQueryRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnswerCallbackQueryRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PostAnswerCallbackQueryRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PostAnswerCallbackQueryRequest) SetText(v string) {
	o.Text = &v
}


// GetShowAlert returns the ShowAlert field value if set, zero value otherwise.
func (o *PostAnswerCallbackQueryRequest) GetShowAlert() bool {
	if o == nil || IsNil(o.ShowAlert) {
		var ret bool
		return ret
	}
	return *o.ShowAlert
}

// GetShowAlertOk returns a tuple with the ShowAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnswerCallbackQueryRequest) GetShowAlertOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAlert) {
		return nil, false
	}
	return o.ShowAlert, true
}

// HasShowAlert returns a boolean if a field has been set.
func (o *PostAnswerCallbackQueryRequest) HasShowAlert() bool {
	if o != nil && !IsNil(o.ShowAlert) {
		return true
	}

	return false
}

// SetShowAlert gets a reference to the given bool and assigns it to the ShowAlert field.
func (o *PostAnswerCallbackQueryRequest) SetShowAlert(v bool) {
	o.ShowAlert = &v
}


// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PostAnswerCallbackQueryRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnswerCallbackQueryRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PostAnswerCallbackQueryRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PostAnswerCallbackQueryRequest) SetUrl(v string) {
	o.Url = &v
}


// GetCacheTime returns the CacheTime field value if set, zero value otherwise.
func (o *PostAnswerCallbackQueryRequest) GetCacheTime() int32 {
	if o == nil || IsNil(o.CacheTime) {
		var ret int32
		return ret
	}
	return *o.CacheTime
}

// GetCacheTimeOk returns a tuple with the CacheTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnswerCallbackQueryRequest) GetCacheTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTime) {
		return nil, false
	}
	return o.CacheTime, true
}

// HasCacheTime returns a boolean if a field has been set.
func (o *PostAnswerCallbackQueryRequest) HasCacheTime() bool {
	if o != nil && !IsNil(o.CacheTime) {
		return true
	}

	return false
}

// SetCacheTime gets a reference to the given int32 and assigns it to the CacheTime field.
func (o *PostAnswerCallbackQueryRequest) SetCacheTime(v int32) {
	o.CacheTime = &v
}


func (o PostAnswerCallbackQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAnswerCallbackQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["callback_query_id"] = o.CallbackQueryId
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.ShowAlert) {
		toSerialize["show_alert"] = o.ShowAlert
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CacheTime) {
		toSerialize["cache_time"] = o.CacheTime
	}
	return toSerialize, nil
}

func (o *PostAnswerCallbackQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"callback_query_id",
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

	varPostAnswerCallbackQueryRequest := _PostAnswerCallbackQueryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAnswerCallbackQueryRequest)

	if err != nil {
		return err
	}

	*o = PostAnswerCallbackQueryRequest(varPostAnswerCallbackQueryRequest)

	return err
}

type NullablePostAnswerCallbackQueryRequest struct {
	value *PostAnswerCallbackQueryRequest
	isSet bool
}

func (v NullablePostAnswerCallbackQueryRequest) Get() *PostAnswerCallbackQueryRequest {
	return v.value
}

func (v *NullablePostAnswerCallbackQueryRequest) Set(val *PostAnswerCallbackQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAnswerCallbackQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAnswerCallbackQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAnswerCallbackQueryRequest(val *PostAnswerCallbackQueryRequest) *NullablePostAnswerCallbackQueryRequest {
	return &NullablePostAnswerCallbackQueryRequest{value: val, isSet: true}
}

func (v NullablePostAnswerCallbackQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAnswerCallbackQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


