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

// checks if the KeyboardButtonRequestUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyboardButtonRequestUsers{}

// KeyboardButtonRequestUsers This object defines the criteria used to request suitable users. Information about the selected users will be shared with the bot when the corresponding button is pressed. [More about requesting users »](https://core.telegram.org/bots/features#chat-and-user-selection)
type KeyboardButtonRequestUsers struct {
	// Signed 32-bit identifier of the request that will be received back in the [UsersShared](https://core.telegram.org/bots/api/#usersshared) object. Must be unique within the message
	RequestId int32 `json:"request_id"`
	// *Optional*. Pass *True* to request bots, pass *False* to request regular users. If not specified, no additional restrictions are applied.
	UserIsBot *bool `json:"user_is_bot,omitempty"`
	// *Optional*. Pass *True* to request premium users, pass *False* to request non-premium users. If not specified, no additional restrictions are applied.
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
	// *Optional*. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity *int32 `json:"max_quantity,omitempty"`
	// *Optional*. Pass *True* to request the users' first and last names
	RequestName *bool `json:"request_name,omitempty"`
	// *Optional*. Pass *True* to request the users' usernames
	RequestUsername *bool `json:"request_username,omitempty"`
	// *Optional*. Pass *True* to request the users' photos
	RequestPhoto *bool `json:"request_photo,omitempty"`
}

type _KeyboardButtonRequestUsers KeyboardButtonRequestUsers

// NewKeyboardButtonRequestUsers instantiates a new KeyboardButtonRequestUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyboardButtonRequestUsers(requestId int32) *KeyboardButtonRequestUsers {
	this := KeyboardButtonRequestUsers{}
	this.RequestId = requestId
	var maxQuantity int32 = 1
	this.MaxQuantity = &maxQuantity
	return &this
}

// NewKeyboardButtonRequestUsersWithDefaults instantiates a new KeyboardButtonRequestUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyboardButtonRequestUsersWithDefaults() *KeyboardButtonRequestUsers {
	this := KeyboardButtonRequestUsers{}
	var maxQuantity int32 = 1
	this.MaxQuantity = &maxQuantity
	return &this
}

// GetRequestId returns the RequestId field value
func (o *KeyboardButtonRequestUsers) GetRequestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetRequestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *KeyboardButtonRequestUsers) SetRequestId(v int32) {
	o.RequestId = v
}

// GetUserIsBot returns the UserIsBot field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetUserIsBot() bool {
	if o == nil || IsNil(o.UserIsBot) {
		var ret bool
		return ret
	}
	return *o.UserIsBot
}

// GetUserIsBotOk returns a tuple with the UserIsBot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetUserIsBotOk() (*bool, bool) {
	if o == nil || IsNil(o.UserIsBot) {
		return nil, false
	}
	return o.UserIsBot, true
}

// HasUserIsBot returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasUserIsBot() bool {
	if o != nil && !IsNil(o.UserIsBot) {
		return true
	}

	return false
}

// SetUserIsBot gets a reference to the given bool and assigns it to the UserIsBot field.
func (o *KeyboardButtonRequestUsers) SetUserIsBot(v bool) {
	o.UserIsBot = &v
}


// GetUserIsPremium returns the UserIsPremium field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetUserIsPremium() bool {
	if o == nil || IsNil(o.UserIsPremium) {
		var ret bool
		return ret
	}
	return *o.UserIsPremium
}

// GetUserIsPremiumOk returns a tuple with the UserIsPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetUserIsPremiumOk() (*bool, bool) {
	if o == nil || IsNil(o.UserIsPremium) {
		return nil, false
	}
	return o.UserIsPremium, true
}

// HasUserIsPremium returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasUserIsPremium() bool {
	if o != nil && !IsNil(o.UserIsPremium) {
		return true
	}

	return false
}

// SetUserIsPremium gets a reference to the given bool and assigns it to the UserIsPremium field.
func (o *KeyboardButtonRequestUsers) SetUserIsPremium(v bool) {
	o.UserIsPremium = &v
}


// GetMaxQuantity returns the MaxQuantity field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetMaxQuantity() int32 {
	if o == nil || IsNil(o.MaxQuantity) {
		var ret int32
		return ret
	}
	return *o.MaxQuantity
}

// GetMaxQuantityOk returns a tuple with the MaxQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetMaxQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQuantity) {
		return nil, false
	}
	return o.MaxQuantity, true
}

// HasMaxQuantity returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasMaxQuantity() bool {
	if o != nil && !IsNil(o.MaxQuantity) {
		return true
	}

	return false
}

// SetMaxQuantity gets a reference to the given int32 and assigns it to the MaxQuantity field.
func (o *KeyboardButtonRequestUsers) SetMaxQuantity(v int32) {
	o.MaxQuantity = &v
}


// GetRequestName returns the RequestName field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetRequestName() bool {
	if o == nil || IsNil(o.RequestName) {
		var ret bool
		return ret
	}
	return *o.RequestName
}

// GetRequestNameOk returns a tuple with the RequestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetRequestNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestName) {
		return nil, false
	}
	return o.RequestName, true
}

// HasRequestName returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasRequestName() bool {
	if o != nil && !IsNil(o.RequestName) {
		return true
	}

	return false
}

// SetRequestName gets a reference to the given bool and assigns it to the RequestName field.
func (o *KeyboardButtonRequestUsers) SetRequestName(v bool) {
	o.RequestName = &v
}


// GetRequestUsername returns the RequestUsername field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetRequestUsername() bool {
	if o == nil || IsNil(o.RequestUsername) {
		var ret bool
		return ret
	}
	return *o.RequestUsername
}

// GetRequestUsernameOk returns a tuple with the RequestUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetRequestUsernameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestUsername) {
		return nil, false
	}
	return o.RequestUsername, true
}

// HasRequestUsername returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasRequestUsername() bool {
	if o != nil && !IsNil(o.RequestUsername) {
		return true
	}

	return false
}

// SetRequestUsername gets a reference to the given bool and assigns it to the RequestUsername field.
func (o *KeyboardButtonRequestUsers) SetRequestUsername(v bool) {
	o.RequestUsername = &v
}


// GetRequestPhoto returns the RequestPhoto field value if set, zero value otherwise.
func (o *KeyboardButtonRequestUsers) GetRequestPhoto() bool {
	if o == nil || IsNil(o.RequestPhoto) {
		var ret bool
		return ret
	}
	return *o.RequestPhoto
}

// GetRequestPhotoOk returns a tuple with the RequestPhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestUsers) GetRequestPhotoOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestPhoto) {
		return nil, false
	}
	return o.RequestPhoto, true
}

// HasRequestPhoto returns a boolean if a field has been set.
func (o *KeyboardButtonRequestUsers) HasRequestPhoto() bool {
	if o != nil && !IsNil(o.RequestPhoto) {
		return true
	}

	return false
}

// SetRequestPhoto gets a reference to the given bool and assigns it to the RequestPhoto field.
func (o *KeyboardButtonRequestUsers) SetRequestPhoto(v bool) {
	o.RequestPhoto = &v
}


func (o KeyboardButtonRequestUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyboardButtonRequestUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	if !IsNil(o.UserIsBot) {
		toSerialize["user_is_bot"] = o.UserIsBot
	}
	if !IsNil(o.UserIsPremium) {
		toSerialize["user_is_premium"] = o.UserIsPremium
	}
	if !IsNil(o.MaxQuantity) {
		toSerialize["max_quantity"] = o.MaxQuantity
	}
	if !IsNil(o.RequestName) {
		toSerialize["request_name"] = o.RequestName
	}
	if !IsNil(o.RequestUsername) {
		toSerialize["request_username"] = o.RequestUsername
	}
	if !IsNil(o.RequestPhoto) {
		toSerialize["request_photo"] = o.RequestPhoto
	}
	return toSerialize, nil
}

func (o *KeyboardButtonRequestUsers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
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

	varKeyboardButtonRequestUsers := _KeyboardButtonRequestUsers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyboardButtonRequestUsers)

	if err != nil {
		return err
	}

	*o = KeyboardButtonRequestUsers(varKeyboardButtonRequestUsers)

	return err
}

type NullableKeyboardButtonRequestUsers struct {
	value *KeyboardButtonRequestUsers
	isSet bool
}

func (v NullableKeyboardButtonRequestUsers) Get() *KeyboardButtonRequestUsers {
	return v.value
}

func (v *NullableKeyboardButtonRequestUsers) Set(val *KeyboardButtonRequestUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyboardButtonRequestUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyboardButtonRequestUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyboardButtonRequestUsers(val *KeyboardButtonRequestUsers) *NullableKeyboardButtonRequestUsers {
	return &NullableKeyboardButtonRequestUsers{value: val, isSet: true}
}

func (v NullableKeyboardButtonRequestUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyboardButtonRequestUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


