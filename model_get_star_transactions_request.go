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
)

// checks if the GetStarTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStarTransactionsRequest{}

// GetStarTransactionsRequest Request parameters for getStarTransactions
type GetStarTransactionsRequest struct {
	// Number of transactions to skip in the response
	Offset *int32 `json:"offset,omitempty"`
	// The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int32 `json:"limit,omitempty"`
}

// NewGetStarTransactionsRequest instantiates a new GetStarTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStarTransactionsRequest() *GetStarTransactionsRequest {
	this := GetStarTransactionsRequest{}
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// NewGetStarTransactionsRequestWithDefaults instantiates a new GetStarTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStarTransactionsRequestWithDefaults() *GetStarTransactionsRequest {
	this := GetStarTransactionsRequest{}
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetStarTransactionsRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStarTransactionsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetStarTransactionsRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetStarTransactionsRequest) SetOffset(v int32) {
	o.Offset = &v
}


// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetStarTransactionsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStarTransactionsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetStarTransactionsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetStarTransactionsRequest) SetLimit(v int32) {
	o.Limit = &v
}


func (o GetStarTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStarTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableGetStarTransactionsRequest struct {
	value *GetStarTransactionsRequest
	isSet bool
}

func (v NullableGetStarTransactionsRequest) Get() *GetStarTransactionsRequest {
	return v.value
}

func (v *NullableGetStarTransactionsRequest) Set(val *GetStarTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStarTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStarTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStarTransactionsRequest(val *GetStarTransactionsRequest) *NullableGetStarTransactionsRequest {
	return &NullableGetStarTransactionsRequest{value: val, isSet: true}
}

func (v NullableGetStarTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStarTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


