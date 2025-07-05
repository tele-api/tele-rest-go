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

// checks if the GetWebhookInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookInfoResponse{}

// GetWebhookInfoResponse struct for GetWebhookInfoResponse
type GetWebhookInfoResponse struct {
	Ok bool `json:"ok"`
	Result WebhookInfo `json:"result"`
}

type _GetWebhookInfoResponse GetWebhookInfoResponse

// NewGetWebhookInfoResponse instantiates a new GetWebhookInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookInfoResponse(ok bool, result WebhookInfo) *GetWebhookInfoResponse {
	this := GetWebhookInfoResponse{}
	this.Ok = ok
	this.Result = result
	return &this
}

// NewGetWebhookInfoResponseWithDefaults instantiates a new GetWebhookInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookInfoResponseWithDefaults() *GetWebhookInfoResponse {
	this := GetWebhookInfoResponse{}
	var ok bool = true
	this.Ok = ok
	return &this
}

// GetOk returns the Ok field value
func (o *GetWebhookInfoResponse) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *GetWebhookInfoResponse) GetOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *GetWebhookInfoResponse) SetOk(v bool) {
	o.Ok = v
}

// GetResult returns the Result field value
func (o *GetWebhookInfoResponse) GetResult() WebhookInfo {
	if o == nil {
		var ret WebhookInfo
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetWebhookInfoResponse) GetResultOk() (*WebhookInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetWebhookInfoResponse) SetResult(v WebhookInfo) {
	o.Result = v
}

func (o GetWebhookInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ok"] = o.Ok
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *GetWebhookInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ok",
		"result",
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

	varGetWebhookInfoResponse := _GetWebhookInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetWebhookInfoResponse)

	if err != nil {
		return err
	}

	*o = GetWebhookInfoResponse(varGetWebhookInfoResponse)

	return err
}

type NullableGetWebhookInfoResponse struct {
	value *GetWebhookInfoResponse
	isSet bool
}

func (v NullableGetWebhookInfoResponse) Get() *GetWebhookInfoResponse {
	return v.value
}

func (v *NullableGetWebhookInfoResponse) Set(val *GetWebhookInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookInfoResponse(val *GetWebhookInfoResponse) *NullableGetWebhookInfoResponse {
	return &NullableGetWebhookInfoResponse{value: val, isSet: true}
}

func (v NullableGetWebhookInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


