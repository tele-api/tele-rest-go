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

// checks if the InlineQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineQuery{}

// InlineQuery This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	From User `json:"from"`
	// Text of the query (up to 256 characters)
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
	// *Optional*. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	ChatType *string `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type _InlineQuery InlineQuery

// NewInlineQuery instantiates a new InlineQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineQuery(id string, from User, query string, offset string) *InlineQuery {
	this := InlineQuery{}
	this.Id = id
	this.From = from
	this.Query = query
	this.Offset = offset
	return &this
}

// NewInlineQueryWithDefaults instantiates a new InlineQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineQueryWithDefaults() *InlineQuery {
	this := InlineQuery{}
	return &this
}

// GetId returns the Id field value
func (o *InlineQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineQuery) SetId(v string) {
	o.Id = v
}

// GetFrom returns the From field value
func (o *InlineQuery) GetFrom() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetFromOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *InlineQuery) SetFrom(v User) {
	o.From = v
}

// GetQuery returns the Query field value
func (o *InlineQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *InlineQuery) SetQuery(v string) {
	o.Query = v
}

// GetOffset returns the Offset field value
func (o *InlineQuery) GetOffset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetOffsetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *InlineQuery) SetOffset(v string) {
	o.Offset = v
}

// GetChatType returns the ChatType field value if set, zero value otherwise.
func (o *InlineQuery) GetChatType() string {
	if o == nil || IsNil(o.ChatType) {
		var ret string
		return ret
	}
	return *o.ChatType
}

// GetChatTypeOk returns a tuple with the ChatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetChatTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChatType) {
		return nil, false
	}
	return o.ChatType, true
}

// HasChatType returns a boolean if a field has been set.
func (o *InlineQuery) HasChatType() bool {
	if o != nil && !IsNil(o.ChatType) {
		return true
	}

	return false
}

// SetChatType gets a reference to the given string and assigns it to the ChatType field.
func (o *InlineQuery) SetChatType(v string) {
	o.ChatType = &v
}


// GetLocation returns the Location field value if set, zero value otherwise.
func (o *InlineQuery) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineQuery) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *InlineQuery) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *InlineQuery) SetLocation(v Location) {
	o.Location = &v
}


func (o InlineQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["offset"] = o.Offset
	if !IsNil(o.ChatType) {
		toSerialize["chat_type"] = o.ChatType
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

func (o *InlineQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"from",
		"query",
		"offset",
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

	varInlineQuery := _InlineQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineQuery)

	if err != nil {
		return err
	}

	*o = InlineQuery(varInlineQuery)

	return err
}

type NullableInlineQuery struct {
	value *InlineQuery
	isSet bool
}

func (v NullableInlineQuery) Get() *InlineQuery {
	return v.value
}

func (v *NullableInlineQuery) Set(val *InlineQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineQuery(val *InlineQuery) *NullableInlineQuery {
	return &NullableInlineQuery{value: val, isSet: true}
}

func (v NullableInlineQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


