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

// checks if the ChosenInlineResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChosenInlineResult{}

// ChosenInlineResult Represents a [result](https://core.telegram.org/bots/api/#inlinequeryresult) of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultId string `json:"result_id"`
	From User `json:"from"`
	Location *Location `json:"location,omitempty"`
	// *Optional*. Identifier of the sent inline message. Available only if there is an [inline keyboard](https://core.telegram.org/bots/api/#inlinekeyboardmarkup) attached to the message. Will be also received in [callback queries](https://core.telegram.org/bots/api/#callbackquery) and can be used to [edit](https://core.telegram.org/bots/api/#updating-messages) the message.
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// The query that was used to obtain the result
	Query string `json:"query"`
}

type _ChosenInlineResult ChosenInlineResult

// NewChosenInlineResult instantiates a new ChosenInlineResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChosenInlineResult(resultId string, from User, query string) *ChosenInlineResult {
	this := ChosenInlineResult{}
	this.ResultId = resultId
	this.From = from
	this.Query = query
	return &this
}

// NewChosenInlineResultWithDefaults instantiates a new ChosenInlineResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChosenInlineResultWithDefaults() *ChosenInlineResult {
	this := ChosenInlineResult{}
	return &this
}

// GetResultId returns the ResultId field value
func (o *ChosenInlineResult) GetResultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value
// and a boolean to check if the value has been set.
func (o *ChosenInlineResult) GetResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultId, true
}

// SetResultId sets field value
func (o *ChosenInlineResult) SetResultId(v string) {
	o.ResultId = v
}

// GetFrom returns the From field value
func (o *ChosenInlineResult) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ChosenInlineResult) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *ChosenInlineResult) SetFrom(v User) {
	o.From = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ChosenInlineResult) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChosenInlineResult) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ChosenInlineResult) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *ChosenInlineResult) SetLocation(v Location) {
	o.Location = &v
}


// GetInlineMessageId returns the InlineMessageId field value if set, zero value otherwise.
func (o *ChosenInlineResult) GetInlineMessageId() string {
	if o == nil || IsNil(o.InlineMessageId) {
		var ret string
		return ret
	}
	return *o.InlineMessageId
}

// GetInlineMessageIdOk returns a tuple with the InlineMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChosenInlineResult) GetInlineMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineMessageId) {
		return nil, false
	}
	return o.InlineMessageId, true
}

// HasInlineMessageId returns a boolean if a field has been set.
func (o *ChosenInlineResult) HasInlineMessageId() bool {
	if o != nil && !IsNil(o.InlineMessageId) {
		return true
	}

	return false
}

// SetInlineMessageId gets a reference to the given string and assigns it to the InlineMessageId field.
func (o *ChosenInlineResult) SetInlineMessageId(v string) {
	o.InlineMessageId = &v
}


// GetQuery returns the Query field value
func (o *ChosenInlineResult) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ChosenInlineResult) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ChosenInlineResult) SetQuery(v string) {
	o.Query = v
}

func (o ChosenInlineResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChosenInlineResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result_id"] = o.ResultId
	toSerialize["from"] = o.From
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.InlineMessageId) {
		toSerialize["inline_message_id"] = o.InlineMessageId
	}
	toSerialize["query"] = o.Query
	return toSerialize, nil
}

func (o *ChosenInlineResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result_id",
		"from",
		"query",
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

	varChosenInlineResult := _ChosenInlineResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChosenInlineResult)

	if err != nil {
		return err
	}

	*o = ChosenInlineResult(varChosenInlineResult)

	return err
}

type NullableChosenInlineResult struct {
	value *ChosenInlineResult
	isSet bool
}

func (v NullableChosenInlineResult) Get() *ChosenInlineResult {
	return v.value
}

func (v *NullableChosenInlineResult) Set(val *ChosenInlineResult) {
	v.value = val
	v.isSet = true
}

func (v NullableChosenInlineResult) IsSet() bool {
	return v.isSet
}

func (v *NullableChosenInlineResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChosenInlineResult(val *ChosenInlineResult) *NullableChosenInlineResult {
	return &NullableChosenInlineResult{value: val, isSet: true}
}

func (v NullableChosenInlineResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChosenInlineResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


