/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:36:13.209453861Z[Etc/UTC]
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

// checks if the AnswerInlineQueryPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerInlineQueryPostRequest{}

// AnswerInlineQueryPostRequest struct for AnswerInlineQueryPostRequest
type AnswerInlineQueryPostRequest struct {
	// Unique identifier for the answered query
	InlineQueryId string `json:"inline_query_id"`
	// A JSON-serialized array of results for the inline query
	Results []InlineQueryResult `json:"results"`
	// The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	CacheTime *int32 `json:"cache_time,omitempty"`
	// Pass *True* if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	IsPersonal *bool `json:"is_personal,omitempty"`
	// Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	NextOffset *string `json:"next_offset,omitempty"`
	Button *InlineQueryResultsButton `json:"button,omitempty"`
}

type _AnswerInlineQueryPostRequest AnswerInlineQueryPostRequest

// NewAnswerInlineQueryPostRequest instantiates a new AnswerInlineQueryPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerInlineQueryPostRequest(inlineQueryId string, results []InlineQueryResult) *AnswerInlineQueryPostRequest {
	this := AnswerInlineQueryPostRequest{}
	this.InlineQueryId = inlineQueryId
	this.Results = results
	var cacheTime int32 = 300
	this.CacheTime = &cacheTime
	return &this
}

// NewAnswerInlineQueryPostRequestWithDefaults instantiates a new AnswerInlineQueryPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerInlineQueryPostRequestWithDefaults() *AnswerInlineQueryPostRequest {
	this := AnswerInlineQueryPostRequest{}
	var cacheTime int32 = 300
	this.CacheTime = &cacheTime
	return &this
}

// GetInlineQueryId returns the InlineQueryId field value
func (o *AnswerInlineQueryPostRequest) GetInlineQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InlineQueryId
}

// GetInlineQueryIdOk returns a tuple with the InlineQueryId field value
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetInlineQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InlineQueryId, true
}

// SetInlineQueryId sets field value
func (o *AnswerInlineQueryPostRequest) SetInlineQueryId(v string) {
	o.InlineQueryId = v
}

// GetResults returns the Results field value
func (o *AnswerInlineQueryPostRequest) GetResults() []InlineQueryResult {
	if o == nil {
		var ret []InlineQueryResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetResultsOk() ([]InlineQueryResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *AnswerInlineQueryPostRequest) SetResults(v []InlineQueryResult) {
	o.Results = v
}

// GetCacheTime returns the CacheTime field value if set, zero value otherwise.
func (o *AnswerInlineQueryPostRequest) GetCacheTime() int32 {
	if o == nil || IsNil(o.CacheTime) {
		var ret int32
		return ret
	}
	return *o.CacheTime
}

// GetCacheTimeOk returns a tuple with the CacheTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetCacheTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheTime) {
		return nil, false
	}
	return o.CacheTime, true
}

// HasCacheTime returns a boolean if a field has been set.
func (o *AnswerInlineQueryPostRequest) HasCacheTime() bool {
	if o != nil && !IsNil(o.CacheTime) {
		return true
	}

	return false
}

// SetCacheTime gets a reference to the given int32 and assigns it to the CacheTime field.
func (o *AnswerInlineQueryPostRequest) SetCacheTime(v int32) {
	o.CacheTime = &v
}


// GetIsPersonal returns the IsPersonal field value if set, zero value otherwise.
func (o *AnswerInlineQueryPostRequest) GetIsPersonal() bool {
	if o == nil || IsNil(o.IsPersonal) {
		var ret bool
		return ret
	}
	return *o.IsPersonal
}

// GetIsPersonalOk returns a tuple with the IsPersonal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetIsPersonalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPersonal) {
		return nil, false
	}
	return o.IsPersonal, true
}

// HasIsPersonal returns a boolean if a field has been set.
func (o *AnswerInlineQueryPostRequest) HasIsPersonal() bool {
	if o != nil && !IsNil(o.IsPersonal) {
		return true
	}

	return false
}

// SetIsPersonal gets a reference to the given bool and assigns it to the IsPersonal field.
func (o *AnswerInlineQueryPostRequest) SetIsPersonal(v bool) {
	o.IsPersonal = &v
}


// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *AnswerInlineQueryPostRequest) GetNextOffset() string {
	if o == nil || IsNil(o.NextOffset) {
		var ret string
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetNextOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.NextOffset) {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *AnswerInlineQueryPostRequest) HasNextOffset() bool {
	if o != nil && !IsNil(o.NextOffset) {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given string and assigns it to the NextOffset field.
func (o *AnswerInlineQueryPostRequest) SetNextOffset(v string) {
	o.NextOffset = &v
}


// GetButton returns the Button field value if set, zero value otherwise.
func (o *AnswerInlineQueryPostRequest) GetButton() InlineQueryResultsButton {
	if o == nil || IsNil(o.Button) {
		var ret InlineQueryResultsButton
		return ret
	}
	return *o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerInlineQueryPostRequest) GetButtonOk() (*InlineQueryResultsButton, bool) {
	if o == nil || IsNil(o.Button) {
		return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *AnswerInlineQueryPostRequest) HasButton() bool {
	if o != nil && !IsNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given InlineQueryResultsButton and assigns it to the Button field.
func (o *AnswerInlineQueryPostRequest) SetButton(v InlineQueryResultsButton) {
	o.Button = &v
}


func (o AnswerInlineQueryPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerInlineQueryPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inline_query_id"] = o.InlineQueryId
	toSerialize["results"] = o.Results
	if !IsNil(o.CacheTime) {
		toSerialize["cache_time"] = o.CacheTime
	}
	if !IsNil(o.IsPersonal) {
		toSerialize["is_personal"] = o.IsPersonal
	}
	if !IsNil(o.NextOffset) {
		toSerialize["next_offset"] = o.NextOffset
	}
	if !IsNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	return toSerialize, nil
}

func (o *AnswerInlineQueryPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inline_query_id",
		"results",
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

	varAnswerInlineQueryPostRequest := _AnswerInlineQueryPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnswerInlineQueryPostRequest)

	if err != nil {
		return err
	}

	*o = AnswerInlineQueryPostRequest(varAnswerInlineQueryPostRequest)

	return err
}

type NullableAnswerInlineQueryPostRequest struct {
	value *AnswerInlineQueryPostRequest
	isSet bool
}

func (v NullableAnswerInlineQueryPostRequest) Get() *AnswerInlineQueryPostRequest {
	return v.value
}

func (v *NullableAnswerInlineQueryPostRequest) Set(val *AnswerInlineQueryPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerInlineQueryPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerInlineQueryPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerInlineQueryPostRequest(val *AnswerInlineQueryPostRequest) *NullableAnswerInlineQueryPostRequest {
	return &NullableAnswerInlineQueryPostRequest{value: val, isSet: true}
}

func (v NullableAnswerInlineQueryPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerInlineQueryPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


