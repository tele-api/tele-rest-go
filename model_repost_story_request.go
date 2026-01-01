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

// checks if the RepostStoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepostStoryRequest{}

// RepostStoryRequest Request parameters for repostStory
type RepostStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier of the chat which posted the story that should be reposted
	FromChatId int32 `json:"from_chat_id"`
	// Unique identifier of the story that should be reposted
	FromStoryId int32 `json:"from_story_id"`
	// Period after which the story is moved to the archive, in seconds; must be one of `6 * 3600`, `12 * 3600`, `86400`, or `2 * 86400`
	ActivePeriod int32 `json:"active_period"`
	// Pass *True* to keep the story accessible after it expires
	PostToChatPage *bool `json:"post_to_chat_page,omitempty"`
	// Pass *True* if the content of the story must be protected from forwarding and screenshotting
	ProtectContent *bool `json:"protect_content,omitempty"`
}

type _RepostStoryRequest RepostStoryRequest

// NewRepostStoryRequest instantiates a new RepostStoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepostStoryRequest(businessConnectionId string, fromChatId int32, fromStoryId int32, activePeriod int32) *RepostStoryRequest {
	this := RepostStoryRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.FromChatId = fromChatId
	this.FromStoryId = fromStoryId
	this.ActivePeriod = activePeriod
	return &this
}

// NewRepostStoryRequestWithDefaults instantiates a new RepostStoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepostStoryRequestWithDefaults() *RepostStoryRequest {
	this := RepostStoryRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *RepostStoryRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *RepostStoryRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetFromChatId returns the FromChatId field value
func (o *RepostStoryRequest) GetFromChatId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromChatId
}

// GetFromChatIdOk returns a tuple with the FromChatId field value
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetFromChatIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromChatId, true
}

// SetFromChatId sets field value
func (o *RepostStoryRequest) SetFromChatId(v int32) {
	o.FromChatId = v
}

// GetFromStoryId returns the FromStoryId field value
func (o *RepostStoryRequest) GetFromStoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromStoryId
}

// GetFromStoryIdOk returns a tuple with the FromStoryId field value
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetFromStoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromStoryId, true
}

// SetFromStoryId sets field value
func (o *RepostStoryRequest) SetFromStoryId(v int32) {
	o.FromStoryId = v
}

// GetActivePeriod returns the ActivePeriod field value
func (o *RepostStoryRequest) GetActivePeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActivePeriod
}

// GetActivePeriodOk returns a tuple with the ActivePeriod field value
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetActivePeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivePeriod, true
}

// SetActivePeriod sets field value
func (o *RepostStoryRequest) SetActivePeriod(v int32) {
	o.ActivePeriod = v
}

// GetPostToChatPage returns the PostToChatPage field value if set, zero value otherwise.
func (o *RepostStoryRequest) GetPostToChatPage() bool {
	if o == nil || IsNil(o.PostToChatPage) {
		var ret bool
		return ret
	}
	return *o.PostToChatPage
}

// GetPostToChatPageOk returns a tuple with the PostToChatPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetPostToChatPageOk() (*bool, bool) {
	if o == nil || IsNil(o.PostToChatPage) {
		return nil, false
	}
	return o.PostToChatPage, true
}

// HasPostToChatPage returns a boolean if a field has been set.
func (o *RepostStoryRequest) HasPostToChatPage() bool {
	if o != nil && !IsNil(o.PostToChatPage) {
		return true
	}

	return false
}

// SetPostToChatPage gets a reference to the given bool and assigns it to the PostToChatPage field.
func (o *RepostStoryRequest) SetPostToChatPage(v bool) {
	o.PostToChatPage = &v
}


// GetProtectContent returns the ProtectContent field value if set, zero value otherwise.
func (o *RepostStoryRequest) GetProtectContent() bool {
	if o == nil || IsNil(o.ProtectContent) {
		var ret bool
		return ret
	}
	return *o.ProtectContent
}

// GetProtectContentOk returns a tuple with the ProtectContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepostStoryRequest) GetProtectContentOk() (*bool, bool) {
	if o == nil || IsNil(o.ProtectContent) {
		return nil, false
	}
	return o.ProtectContent, true
}

// HasProtectContent returns a boolean if a field has been set.
func (o *RepostStoryRequest) HasProtectContent() bool {
	if o != nil && !IsNil(o.ProtectContent) {
		return true
	}

	return false
}

// SetProtectContent gets a reference to the given bool and assigns it to the ProtectContent field.
func (o *RepostStoryRequest) SetProtectContent(v bool) {
	o.ProtectContent = &v
}


func (o RepostStoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepostStoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["from_chat_id"] = o.FromChatId
	toSerialize["from_story_id"] = o.FromStoryId
	toSerialize["active_period"] = o.ActivePeriod
	if !IsNil(o.PostToChatPage) {
		toSerialize["post_to_chat_page"] = o.PostToChatPage
	}
	if !IsNil(o.ProtectContent) {
		toSerialize["protect_content"] = o.ProtectContent
	}
	return toSerialize, nil
}

func (o *RepostStoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"from_chat_id",
		"from_story_id",
		"active_period",
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

	varRepostStoryRequest := _RepostStoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRepostStoryRequest)

	if err != nil {
		return err
	}

	*o = RepostStoryRequest(varRepostStoryRequest)

	return err
}

type NullableRepostStoryRequest struct {
	value *RepostStoryRequest
	isSet bool
}

func (v NullableRepostStoryRequest) Get() *RepostStoryRequest {
	return v.value
}

func (v *NullableRepostStoryRequest) Set(val *RepostStoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRepostStoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRepostStoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepostStoryRequest(val *RepostStoryRequest) *NullableRepostStoryRequest {
	return &NullableRepostStoryRequest{value: val, isSet: true}
}

func (v NullableRepostStoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepostStoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

