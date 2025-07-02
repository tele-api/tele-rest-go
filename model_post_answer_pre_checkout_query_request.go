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

// checks if the PostAnswerPreCheckoutQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAnswerPreCheckoutQueryRequest{}

// PostAnswerPreCheckoutQueryRequest struct for PostAnswerPreCheckoutQueryRequest
type PostAnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	// Specify *True* if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use *False* if there are any problems.
	Ok bool `json:"ok"`
	// Required if *ok* is *False*. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. \"Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!\"). Telegram will display this message to the user.
	ErrorMessage *string `json:"error_message,omitempty"`
}

type _PostAnswerPreCheckoutQueryRequest PostAnswerPreCheckoutQueryRequest

// NewPostAnswerPreCheckoutQueryRequest instantiates a new PostAnswerPreCheckoutQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAnswerPreCheckoutQueryRequest(preCheckoutQueryId string, ok bool) *PostAnswerPreCheckoutQueryRequest {
	this := PostAnswerPreCheckoutQueryRequest{}
	this.PreCheckoutQueryId = preCheckoutQueryId
	this.Ok = ok
	return &this
}

// NewPostAnswerPreCheckoutQueryRequestWithDefaults instantiates a new PostAnswerPreCheckoutQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAnswerPreCheckoutQueryRequestWithDefaults() *PostAnswerPreCheckoutQueryRequest {
	this := PostAnswerPreCheckoutQueryRequest{}
	return &this
}

// GetPreCheckoutQueryId returns the PreCheckoutQueryId field value
func (o *PostAnswerPreCheckoutQueryRequest) GetPreCheckoutQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreCheckoutQueryId
}

// GetPreCheckoutQueryIdOk returns a tuple with the PreCheckoutQueryId field value
// and a boolean to check if the value has been set.
func (o *PostAnswerPreCheckoutQueryRequest) GetPreCheckoutQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreCheckoutQueryId, true
}

// SetPreCheckoutQueryId sets field value
func (o *PostAnswerPreCheckoutQueryRequest) SetPreCheckoutQueryId(v string) {
	o.PreCheckoutQueryId = v
}

// GetOk returns the Ok field value
func (o *PostAnswerPreCheckoutQueryRequest) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *PostAnswerPreCheckoutQueryRequest) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *PostAnswerPreCheckoutQueryRequest) SetOk(v bool) {
	o.Ok = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PostAnswerPreCheckoutQueryRequest) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnswerPreCheckoutQueryRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PostAnswerPreCheckoutQueryRequest) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PostAnswerPreCheckoutQueryRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}


func (o PostAnswerPreCheckoutQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAnswerPreCheckoutQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pre_checkout_query_id"] = o.PreCheckoutQueryId
	toSerialize["ok"] = o.Ok
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

func (o *PostAnswerPreCheckoutQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pre_checkout_query_id",
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

	varPostAnswerPreCheckoutQueryRequest := _PostAnswerPreCheckoutQueryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAnswerPreCheckoutQueryRequest)

	if err != nil {
		return err
	}

	*o = PostAnswerPreCheckoutQueryRequest(varPostAnswerPreCheckoutQueryRequest)

	return err
}

type NullablePostAnswerPreCheckoutQueryRequest struct {
	value *PostAnswerPreCheckoutQueryRequest
	isSet bool
}

func (v NullablePostAnswerPreCheckoutQueryRequest) Get() *PostAnswerPreCheckoutQueryRequest {
	return v.value
}

func (v *NullablePostAnswerPreCheckoutQueryRequest) Set(val *PostAnswerPreCheckoutQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAnswerPreCheckoutQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAnswerPreCheckoutQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAnswerPreCheckoutQueryRequest(val *PostAnswerPreCheckoutQueryRequest) *NullablePostAnswerPreCheckoutQueryRequest {
	return &NullablePostAnswerPreCheckoutQueryRequest{value: val, isSet: true}
}

func (v NullablePostAnswerPreCheckoutQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAnswerPreCheckoutQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


