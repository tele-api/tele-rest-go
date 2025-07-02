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

// checks if the TransferGiftRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferGiftRequest{}

// TransferGiftRequest Request parameters for transferGift
type TransferGiftRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier of the regular gift that should be transferred
	OwnedGiftId string `json:"owned_gift_id"`
	// Unique identifier of the chat which will own the gift. The chat must be active in the last 24 hours.
	NewOwnerChatId int32 `json:"new_owner_chat_id"`
	// The amount of Telegram Stars that will be paid for the transfer from the business account balance. If positive, then the *can\\_transfer\\_stars* business bot right is required.
	StarCount *int32 `json:"star_count,omitempty"`
}

type _TransferGiftRequest TransferGiftRequest

// NewTransferGiftRequest instantiates a new TransferGiftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferGiftRequest(businessConnectionId string, ownedGiftId string, newOwnerChatId int32) *TransferGiftRequest {
	this := TransferGiftRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.OwnedGiftId = ownedGiftId
	this.NewOwnerChatId = newOwnerChatId
	return &this
}

// NewTransferGiftRequestWithDefaults instantiates a new TransferGiftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferGiftRequestWithDefaults() *TransferGiftRequest {
	this := TransferGiftRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *TransferGiftRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *TransferGiftRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *TransferGiftRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetOwnedGiftId returns the OwnedGiftId field value
func (o *TransferGiftRequest) GetOwnedGiftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnedGiftId
}

// GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field value
// and a boolean to check if the value has been set.
func (o *TransferGiftRequest) GetOwnedGiftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedGiftId, true
}

// SetOwnedGiftId sets field value
func (o *TransferGiftRequest) SetOwnedGiftId(v string) {
	o.OwnedGiftId = v
}

// GetNewOwnerChatId returns the NewOwnerChatId field value
func (o *TransferGiftRequest) GetNewOwnerChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NewOwnerChatId
}

// GetNewOwnerChatIdOk returns a tuple with the NewOwnerChatId field value
// and a boolean to check if the value has been set.
func (o *TransferGiftRequest) GetNewOwnerChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewOwnerChatId, true
}

// SetNewOwnerChatId sets field value
func (o *TransferGiftRequest) SetNewOwnerChatId(v int32) {
	o.NewOwnerChatId = v
}

// GetStarCount returns the StarCount field value if set, zero value otherwise.
func (o *TransferGiftRequest) GetStarCount() int32 {
	if o == nil || IsNil(o.StarCount) {
		var ret int32
		return ret
	}
	return *o.StarCount
}

// GetStarCountOk returns a tuple with the StarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferGiftRequest) GetStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StarCount) {
		return nil, false
	}
	return o.StarCount, true
}

// HasStarCount returns a boolean if a field has been set.
func (o *TransferGiftRequest) HasStarCount() bool {
	if o != nil && !IsNil(o.StarCount) {
		return true
	}

	return false
}

// SetStarCount gets a reference to the given int32 and assigns it to the StarCount field.
func (o *TransferGiftRequest) SetStarCount(v int32) {
	o.StarCount = &v
}


func (o TransferGiftRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferGiftRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["owned_gift_id"] = o.OwnedGiftId
	toSerialize["new_owner_chat_id"] = o.NewOwnerChatId
	if !IsNil(o.StarCount) {
		toSerialize["star_count"] = o.StarCount
	}
	return toSerialize, nil
}

func (o *TransferGiftRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"owned_gift_id",
		"new_owner_chat_id",
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

	varTransferGiftRequest := _TransferGiftRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferGiftRequest)

	if err != nil {
		return err
	}

	*o = TransferGiftRequest(varTransferGiftRequest)

	return err
}

type NullableTransferGiftRequest struct {
	value *TransferGiftRequest
	isSet bool
}

func (v NullableTransferGiftRequest) Get() *TransferGiftRequest {
	return v.value
}

func (v *NullableTransferGiftRequest) Set(val *TransferGiftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferGiftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferGiftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferGiftRequest(val *TransferGiftRequest) *NullableTransferGiftRequest {
	return &NullableTransferGiftRequest{value: val, isSet: true}
}

func (v NullableTransferGiftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferGiftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


