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

// checks if the AnswerShippingQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerShippingQueryRequest{}

// AnswerShippingQueryRequest Request parameters for answerShippingQuery
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered
	ShippingQueryId string `json:"shipping_query_id"`
	// Pass *True* if delivery to the specified address is possible and *False* if there are any problems (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`
	// Required if *ok* is *True*. A JSON-serialized array of available shipping options.
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	// Required if *ok* is *False*. Error message in human readable form that explains why it is impossible to complete the order (e.g. “Sorry, delivery to your desired address is unavailable”). Telegram will display this message to the user.
	ErrorMessage *string `json:"error_message,omitempty"`
}

type _AnswerShippingQueryRequest AnswerShippingQueryRequest

// NewAnswerShippingQueryRequest instantiates a new AnswerShippingQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerShippingQueryRequest(shippingQueryId string, ok bool) *AnswerShippingQueryRequest {
	this := AnswerShippingQueryRequest{}
	this.ShippingQueryId = shippingQueryId
	this.Ok = ok
	return &this
}

// NewAnswerShippingQueryRequestWithDefaults instantiates a new AnswerShippingQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerShippingQueryRequestWithDefaults() *AnswerShippingQueryRequest {
	this := AnswerShippingQueryRequest{}
	return &this
}

// GetShippingQueryId returns the ShippingQueryId field value
func (o *AnswerShippingQueryRequest) GetShippingQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShippingQueryId
}

// GetShippingQueryIdOk returns a tuple with the ShippingQueryId field value
// and a boolean to check if the value has been set.
func (o *AnswerShippingQueryRequest) GetShippingQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingQueryId, true
}

// SetShippingQueryId sets field value
func (o *AnswerShippingQueryRequest) SetShippingQueryId(v string) {
	o.ShippingQueryId = v
}

// GetOk returns the Ok field value
func (o *AnswerShippingQueryRequest) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *AnswerShippingQueryRequest) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *AnswerShippingQueryRequest) SetOk(v bool) {
	o.Ok = v
}

// GetShippingOptions returns the ShippingOptions field value if set, zero value otherwise.
func (o *AnswerShippingQueryRequest) GetShippingOptions() []ShippingOption {
	if o == nil || IsNil(o.ShippingOptions) {
		var ret []ShippingOption
		return ret
	}
	return o.ShippingOptions
}

// GetShippingOptionsOk returns a tuple with the ShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerShippingQueryRequest) GetShippingOptionsOk() ([]ShippingOption, bool) {
	if o == nil || IsNil(o.ShippingOptions) {
		return nil, false
	}
	return o.ShippingOptions, true
}

// HasShippingOptions returns a boolean if a field has been set.
func (o *AnswerShippingQueryRequest) HasShippingOptions() bool {
	if o != nil && !IsNil(o.ShippingOptions) {
		return true
	}

	return false
}

// SetShippingOptions gets a reference to the given []ShippingOption and assigns it to the ShippingOptions field.
func (o *AnswerShippingQueryRequest) SetShippingOptions(v []ShippingOption) {
	o.ShippingOptions = v
}


// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AnswerShippingQueryRequest) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerShippingQueryRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AnswerShippingQueryRequest) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AnswerShippingQueryRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}


func (o AnswerShippingQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerShippingQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipping_query_id"] = o.ShippingQueryId
	toSerialize["ok"] = o.Ok
	if !IsNil(o.ShippingOptions) {
		toSerialize["shipping_options"] = o.ShippingOptions
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

func (o *AnswerShippingQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shipping_query_id",
		"ok",
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

	varAnswerShippingQueryRequest := _AnswerShippingQueryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnswerShippingQueryRequest)

	if err != nil {
		return err
	}

	*o = AnswerShippingQueryRequest(varAnswerShippingQueryRequest)

	return err
}

type NullableAnswerShippingQueryRequest struct {
	value *AnswerShippingQueryRequest
	isSet bool
}

func (v NullableAnswerShippingQueryRequest) Get() *AnswerShippingQueryRequest {
	return v.value
}

func (v *NullableAnswerShippingQueryRequest) Set(val *AnswerShippingQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerShippingQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerShippingQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerShippingQueryRequest(val *AnswerShippingQueryRequest) *NullableAnswerShippingQueryRequest {
	return &NullableAnswerShippingQueryRequest{value: val, isSet: true}
}

func (v NullableAnswerShippingQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerShippingQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


