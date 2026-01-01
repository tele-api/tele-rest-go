/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.3.0
 *    * - **Modified**: 2026-01-01T02:06:16.405655294Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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

// checks if the GetChatGiftsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChatGiftsRequest{}

// GetChatGiftsRequest Request parameters for getChatGifts
type GetChatGiftsRequest struct {
	ChatId SendMessageRequestChatId `json:"chat_id"`
	// Pass *True* to exclude gifts that aren't saved to the chat's profile page. Always *True*, unless the bot has the *can\\_post\\_messages* administrator right in the channel.
	ExcludeUnsaved *bool `json:"exclude_unsaved,omitempty"`
	// Pass *True* to exclude gifts that are saved to the chat's profile page. Always *False*, unless the bot has the *can\\_post\\_messages* administrator right in the channel.
	ExcludeSaved *bool `json:"exclude_saved,omitempty"`
	// Pass *True* to exclude gifts that can be purchased an unlimited number of times
	ExcludeUnlimited *bool `json:"exclude_unlimited,omitempty"`
	// Pass *True* to exclude gifts that can be purchased a limited number of times and can be upgraded to unique
	ExcludeLimitedUpgradable *bool `json:"exclude_limited_upgradable,omitempty"`
	// Pass *True* to exclude gifts that can be purchased a limited number of times and can't be upgraded to unique
	ExcludeLimitedNonUpgradable *bool `json:"exclude_limited_non_upgradable,omitempty"`
	// Pass *True* to exclude gifts that were assigned from the TON blockchain and can't be resold or transferred in Telegram
	ExcludeFromBlockchain *bool `json:"exclude_from_blockchain,omitempty"`
	// Pass *True* to exclude unique gifts
	ExcludeUnique *bool `json:"exclude_unique,omitempty"`
	// Pass *True* to sort results by gift price instead of send date. Sorting is applied before pagination.
	SortByPrice *bool `json:"sort_by_price,omitempty"`
	// Offset of the first entry to return as received from the previous request; use an empty string to get the first chunk of results
	Offset *string `json:"offset,omitempty"`
	// The maximum number of gifts to be returned; 1-100. Defaults to 100
	Limit *int32 `json:"limit,omitempty"`
}

type _GetChatGiftsRequest GetChatGiftsRequest

// NewGetChatGiftsRequest instantiates a new GetChatGiftsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChatGiftsRequest(chatId SendMessageRequestChatId) *GetChatGiftsRequest {
	this := GetChatGiftsRequest{}
	this.ChatId = chatId
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// NewGetChatGiftsRequestWithDefaults instantiates a new GetChatGiftsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChatGiftsRequestWithDefaults() *GetChatGiftsRequest {
	this := GetChatGiftsRequest{}
	var limit int32 = 100
	this.Limit = &limit
	return &this
}

// GetChatId returns the ChatId field value
func (o *GetChatGiftsRequest) GetChatId() SendMessageRequestChatId {
	if o == nil {
		var ret SendMessageRequestChatId
		return ret
	}

	return o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetChatIdOk() (*SendMessageRequestChatId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatId, true
}

// SetChatId sets field value
func (o *GetChatGiftsRequest) SetChatId(v SendMessageRequestChatId) {
	o.ChatId = v
}

// GetExcludeUnsaved returns the ExcludeUnsaved field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeUnsaved() bool {
	if o == nil || IsNil(o.ExcludeUnsaved) {
		var ret bool
		return ret
	}
	return *o.ExcludeUnsaved
}

// GetExcludeUnsavedOk returns a tuple with the ExcludeUnsaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeUnsavedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUnsaved) {
		return nil, false
	}
	return o.ExcludeUnsaved, true
}

// HasExcludeUnsaved returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeUnsaved() bool {
	if o != nil && !IsNil(o.ExcludeUnsaved) {
		return true
	}

	return false
}

// SetExcludeUnsaved gets a reference to the given bool and assigns it to the ExcludeUnsaved field.
func (o *GetChatGiftsRequest) SetExcludeUnsaved(v bool) {
	o.ExcludeUnsaved = &v
}


// GetExcludeSaved returns the ExcludeSaved field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeSaved() bool {
	if o == nil || IsNil(o.ExcludeSaved) {
		var ret bool
		return ret
	}
	return *o.ExcludeSaved
}

// GetExcludeSavedOk returns a tuple with the ExcludeSaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeSavedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeSaved) {
		return nil, false
	}
	return o.ExcludeSaved, true
}

// HasExcludeSaved returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeSaved() bool {
	if o != nil && !IsNil(o.ExcludeSaved) {
		return true
	}

	return false
}

// SetExcludeSaved gets a reference to the given bool and assigns it to the ExcludeSaved field.
func (o *GetChatGiftsRequest) SetExcludeSaved(v bool) {
	o.ExcludeSaved = &v
}


// GetExcludeUnlimited returns the ExcludeUnlimited field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeUnlimited() bool {
	if o == nil || IsNil(o.ExcludeUnlimited) {
		var ret bool
		return ret
	}
	return *o.ExcludeUnlimited
}

// GetExcludeUnlimitedOk returns a tuple with the ExcludeUnlimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeUnlimitedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUnlimited) {
		return nil, false
	}
	return o.ExcludeUnlimited, true
}

// HasExcludeUnlimited returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeUnlimited() bool {
	if o != nil && !IsNil(o.ExcludeUnlimited) {
		return true
	}

	return false
}

// SetExcludeUnlimited gets a reference to the given bool and assigns it to the ExcludeUnlimited field.
func (o *GetChatGiftsRequest) SetExcludeUnlimited(v bool) {
	o.ExcludeUnlimited = &v
}


// GetExcludeLimitedUpgradable returns the ExcludeLimitedUpgradable field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeLimitedUpgradable() bool {
	if o == nil || IsNil(o.ExcludeLimitedUpgradable) {
		var ret bool
		return ret
	}
	return *o.ExcludeLimitedUpgradable
}

// GetExcludeLimitedUpgradableOk returns a tuple with the ExcludeLimitedUpgradable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeLimitedUpgradableOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeLimitedUpgradable) {
		return nil, false
	}
	return o.ExcludeLimitedUpgradable, true
}

// HasExcludeLimitedUpgradable returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeLimitedUpgradable() bool {
	if o != nil && !IsNil(o.ExcludeLimitedUpgradable) {
		return true
	}

	return false
}

// SetExcludeLimitedUpgradable gets a reference to the given bool and assigns it to the ExcludeLimitedUpgradable field.
func (o *GetChatGiftsRequest) SetExcludeLimitedUpgradable(v bool) {
	o.ExcludeLimitedUpgradable = &v
}


// GetExcludeLimitedNonUpgradable returns the ExcludeLimitedNonUpgradable field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeLimitedNonUpgradable() bool {
	if o == nil || IsNil(o.ExcludeLimitedNonUpgradable) {
		var ret bool
		return ret
	}
	return *o.ExcludeLimitedNonUpgradable
}

// GetExcludeLimitedNonUpgradableOk returns a tuple with the ExcludeLimitedNonUpgradable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeLimitedNonUpgradableOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeLimitedNonUpgradable) {
		return nil, false
	}
	return o.ExcludeLimitedNonUpgradable, true
}

// HasExcludeLimitedNonUpgradable returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeLimitedNonUpgradable() bool {
	if o != nil && !IsNil(o.ExcludeLimitedNonUpgradable) {
		return true
	}

	return false
}

// SetExcludeLimitedNonUpgradable gets a reference to the given bool and assigns it to the ExcludeLimitedNonUpgradable field.
func (o *GetChatGiftsRequest) SetExcludeLimitedNonUpgradable(v bool) {
	o.ExcludeLimitedNonUpgradable = &v
}


// GetExcludeFromBlockchain returns the ExcludeFromBlockchain field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeFromBlockchain() bool {
	if o == nil || IsNil(o.ExcludeFromBlockchain) {
		var ret bool
		return ret
	}
	return *o.ExcludeFromBlockchain
}

// GetExcludeFromBlockchainOk returns a tuple with the ExcludeFromBlockchain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeFromBlockchainOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeFromBlockchain) {
		return nil, false
	}
	return o.ExcludeFromBlockchain, true
}

// HasExcludeFromBlockchain returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeFromBlockchain() bool {
	if o != nil && !IsNil(o.ExcludeFromBlockchain) {
		return true
	}

	return false
}

// SetExcludeFromBlockchain gets a reference to the given bool and assigns it to the ExcludeFromBlockchain field.
func (o *GetChatGiftsRequest) SetExcludeFromBlockchain(v bool) {
	o.ExcludeFromBlockchain = &v
}


// GetExcludeUnique returns the ExcludeUnique field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetExcludeUnique() bool {
	if o == nil || IsNil(o.ExcludeUnique) {
		var ret bool
		return ret
	}
	return *o.ExcludeUnique
}

// GetExcludeUniqueOk returns a tuple with the ExcludeUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetExcludeUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUnique) {
		return nil, false
	}
	return o.ExcludeUnique, true
}

// HasExcludeUnique returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasExcludeUnique() bool {
	if o != nil && !IsNil(o.ExcludeUnique) {
		return true
	}

	return false
}

// SetExcludeUnique gets a reference to the given bool and assigns it to the ExcludeUnique field.
func (o *GetChatGiftsRequest) SetExcludeUnique(v bool) {
	o.ExcludeUnique = &v
}


// GetSortByPrice returns the SortByPrice field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetSortByPrice() bool {
	if o == nil || IsNil(o.SortByPrice) {
		var ret bool
		return ret
	}
	return *o.SortByPrice
}

// GetSortByPriceOk returns a tuple with the SortByPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetSortByPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.SortByPrice) {
		return nil, false
	}
	return o.SortByPrice, true
}

// HasSortByPrice returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasSortByPrice() bool {
	if o != nil && !IsNil(o.SortByPrice) {
		return true
	}

	return false
}

// SetSortByPrice gets a reference to the given bool and assigns it to the SortByPrice field.
func (o *GetChatGiftsRequest) SetSortByPrice(v bool) {
	o.SortByPrice = &v
}


// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetOffset() string {
	if o == nil || IsNil(o.Offset) {
		var ret string
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given string and assigns it to the Offset field.
func (o *GetChatGiftsRequest) SetOffset(v string) {
	o.Offset = &v
}


// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetChatGiftsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatGiftsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetChatGiftsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetChatGiftsRequest) SetLimit(v int32) {
	o.Limit = &v
}


func (o GetChatGiftsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChatGiftsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chat_id"] = o.ChatId
	if !IsNil(o.ExcludeUnsaved) {
		toSerialize["exclude_unsaved"] = o.ExcludeUnsaved
	}
	if !IsNil(o.ExcludeSaved) {
		toSerialize["exclude_saved"] = o.ExcludeSaved
	}
	if !IsNil(o.ExcludeUnlimited) {
		toSerialize["exclude_unlimited"] = o.ExcludeUnlimited
	}
	if !IsNil(o.ExcludeLimitedUpgradable) {
		toSerialize["exclude_limited_upgradable"] = o.ExcludeLimitedUpgradable
	}
	if !IsNil(o.ExcludeLimitedNonUpgradable) {
		toSerialize["exclude_limited_non_upgradable"] = o.ExcludeLimitedNonUpgradable
	}
	if !IsNil(o.ExcludeFromBlockchain) {
		toSerialize["exclude_from_blockchain"] = o.ExcludeFromBlockchain
	}
	if !IsNil(o.ExcludeUnique) {
		toSerialize["exclude_unique"] = o.ExcludeUnique
	}
	if !IsNil(o.SortByPrice) {
		toSerialize["sort_by_price"] = o.SortByPrice
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

func (o *GetChatGiftsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chat_id",
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

	varGetChatGiftsRequest := _GetChatGiftsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetChatGiftsRequest)

	if err != nil {
		return err
	}

	*o = GetChatGiftsRequest(varGetChatGiftsRequest)

	return err
}

type NullableGetChatGiftsRequest struct {
	value *GetChatGiftsRequest
	isSet bool
}

func (v NullableGetChatGiftsRequest) Get() *GetChatGiftsRequest {
	return v.value
}

func (v *NullableGetChatGiftsRequest) Set(val *GetChatGiftsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChatGiftsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChatGiftsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChatGiftsRequest(val *GetChatGiftsRequest) *NullableGetChatGiftsRequest {
	return &NullableGetChatGiftsRequest{value: val, isSet: true}
}

func (v NullableGetChatGiftsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChatGiftsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

