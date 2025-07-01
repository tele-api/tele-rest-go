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

// checks if the OwnedGifts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnedGifts{}

// OwnedGifts Contains the list of gifts received and owned by a user or a chat.
type OwnedGifts struct {
	// The total number of gifts owned by the user or the chat
	TotalCount int32 `json:"total_count"`
	// The list of gifts
	Gifts []OwnedGift `json:"gifts"`
	// *Optional*. Offset for the next request. If empty, then there are no more results
	NextOffset *string `json:"next_offset,omitempty"`
}

type _OwnedGifts OwnedGifts

// NewOwnedGifts instantiates a new OwnedGifts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnedGifts(totalCount int32, gifts []OwnedGift) *OwnedGifts {
	this := OwnedGifts{}
	this.TotalCount = totalCount
	this.Gifts = gifts
	return &this
}

// NewOwnedGiftsWithDefaults instantiates a new OwnedGifts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnedGiftsWithDefaults() *OwnedGifts {
	this := OwnedGifts{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *OwnedGifts) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *OwnedGifts) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *OwnedGifts) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetGifts returns the Gifts field value
func (o *OwnedGifts) GetGifts() []OwnedGift {
	if o == nil {
		var ret []OwnedGift
		return ret
	}

	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value
// and a boolean to check if the value has been set.
func (o *OwnedGifts) GetGiftsOk() ([]OwnedGift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gifts, true
}

// SetGifts sets field value
func (o *OwnedGifts) SetGifts(v []OwnedGift) {
	o.Gifts = v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *OwnedGifts) GetNextOffset() string {
	if o == nil || IsNil(o.NextOffset) {
		var ret string
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnedGifts) GetNextOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.NextOffset) {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *OwnedGifts) HasNextOffset() bool {
	if o != nil && !IsNil(o.NextOffset) {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given string and assigns it to the NextOffset field.
func (o *OwnedGifts) SetNextOffset(v string) {
	o.NextOffset = &v
}


func (o OwnedGifts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnedGifts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_count"] = o.TotalCount
	toSerialize["gifts"] = o.Gifts
	if !IsNil(o.NextOffset) {
		toSerialize["next_offset"] = o.NextOffset
	}
	return toSerialize, nil
}

func (o *OwnedGifts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_count",
		"gifts",
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

	varOwnedGifts := _OwnedGifts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOwnedGifts)

	if err != nil {
		return err
	}

	*o = OwnedGifts(varOwnedGifts)

	return err
}

type NullableOwnedGifts struct {
	value *OwnedGifts
	isSet bool
}

func (v NullableOwnedGifts) Get() *OwnedGifts {
	return v.value
}

func (v *NullableOwnedGifts) Set(val *OwnedGifts) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnedGifts) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnedGifts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnedGifts(val *OwnedGifts) *NullableOwnedGifts {
	return &NullableOwnedGifts{value: val, isSet: true}
}

func (v NullableOwnedGifts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnedGifts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


