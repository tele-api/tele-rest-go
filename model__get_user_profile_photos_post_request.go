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

// checks if the GetUserProfilePhotosPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserProfilePhotosPostRequest{}

// GetUserProfilePhotosPostRequest struct for GetUserProfilePhotosPostRequest
type GetUserProfilePhotosPostRequest struct {
	// Unique identifier of the target user
	UserId int32 `json:"user_id"`
	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset *int32 `json:"offset,omitempty"`
	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int32 `json:"limit,omitempty"`
}

type _GetUserProfilePhotosPostRequest GetUserProfilePhotosPostRequest

// NewGetUserProfilePhotosPostRequest instantiates a new GetUserProfilePhotosPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserProfilePhotosPostRequest(userId int32) *GetUserProfilePhotosPostRequest {
	this := GetUserProfilePhotosPostRequest{}
	this.UserId = userId
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// NewGetUserProfilePhotosPostRequestWithDefaults instantiates a new GetUserProfilePhotosPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserProfilePhotosPostRequestWithDefaults() *GetUserProfilePhotosPostRequest {
	this := GetUserProfilePhotosPostRequest{}
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// GetUserId returns the UserId field value
func (o *GetUserProfilePhotosPostRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetUserProfilePhotosPostRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetUserProfilePhotosPostRequest) SetUserId(v int32) {
	o.UserId = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetUserProfilePhotosPostRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserProfilePhotosPostRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetUserProfilePhotosPostRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetUserProfilePhotosPostRequest) SetOffset(v int32) {
	o.Offset = &v
}


// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetUserProfilePhotosPostRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserProfilePhotosPostRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetUserProfilePhotosPostRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetUserProfilePhotosPostRequest) SetLimit(v int32) {
	o.Limit = &v
}


func (o GetUserProfilePhotosPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserProfilePhotosPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

func (o *GetUserProfilePhotosPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varGetUserProfilePhotosPostRequest := _GetUserProfilePhotosPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUserProfilePhotosPostRequest)

	if err != nil {
		return err
	}

	*o = GetUserProfilePhotosPostRequest(varGetUserProfilePhotosPostRequest)

	return err
}

type NullableGetUserProfilePhotosPostRequest struct {
	value *GetUserProfilePhotosPostRequest
	isSet bool
}

func (v NullableGetUserProfilePhotosPostRequest) Get() *GetUserProfilePhotosPostRequest {
	return v.value
}

func (v *NullableGetUserProfilePhotosPostRequest) Set(val *GetUserProfilePhotosPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserProfilePhotosPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserProfilePhotosPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserProfilePhotosPostRequest(val *GetUserProfilePhotosPostRequest) *NullableGetUserProfilePhotosPostRequest {
	return &NullableGetUserProfilePhotosPostRequest{value: val, isSet: true}
}

func (v NullableGetUserProfilePhotosPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserProfilePhotosPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


