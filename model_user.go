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
	"bytes"
	"fmt"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User This object represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int32 `json:"id"`
	// *True*, if this user is a bot
	IsBot bool `json:"is_bot"`
	// User's or bot's first name
	FirstName string `json:"first_name"`
	// *Optional*. User's or bot's last name
	LastName *string `json:"last_name,omitempty"`
	// *Optional*. User's or bot's username
	Username *string `json:"username,omitempty"`
	// *Optional*. [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag) of the user's language
	LanguageCode *string `json:"language_code,omitempty"`
	// *Optional*. *True*, if this user is a Telegram Premium user
	IsPremium *bool `json:"is_premium,omitempty"`
	// *Optional*. *True*, if this user added the bot to the attachment menu
	AddedToAttachmentMenu *bool `json:"added_to_attachment_menu,omitempty"`
	// *Optional*. *True*, if the bot can be invited to groups. Returned only in [getMe](https://core.telegram.org/bots/api/#getme).
	CanJoinGroups *bool `json:"can_join_groups,omitempty"`
	// *Optional*. *True*, if [privacy mode](https://core.telegram.org/bots/features#privacy-mode) is disabled for the bot. Returned only in [getMe](https://core.telegram.org/bots/api/#getme).
	CanReadAllGroupMessages *bool `json:"can_read_all_group_messages,omitempty"`
	// *Optional*. *True*, if the bot supports inline queries. Returned only in [getMe](https://core.telegram.org/bots/api/#getme).
	SupportsInlineQueries *bool `json:"supports_inline_queries,omitempty"`
	// *Optional*. *True*, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in [getMe](https://core.telegram.org/bots/api/#getme).
	CanConnectToBusiness *bool `json:"can_connect_to_business,omitempty"`
	// *Optional*. *True*, if the bot has a main Web App. Returned only in [getMe](https://core.telegram.org/bots/api/#getme).
	HasMainWebApp *bool `json:"has_main_web_app,omitempty"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id int32, isBot bool, firstName string) *User {
	this := User{}
	this.Id = id
	this.IsBot = isBot
	this.FirstName = firstName
	var isPremium bool = true
	this.IsPremium = &isPremium
	var addedToAttachmentMenu bool = true
	this.AddedToAttachmentMenu = &addedToAttachmentMenu
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var isPremium bool = true
	this.IsPremium = &isPremium
	var addedToAttachmentMenu bool = true
	this.AddedToAttachmentMenu = &addedToAttachmentMenu
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v int32) {
	o.Id = v
}

// GetIsBot returns the IsBot field value
func (o *User) GetIsBot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBot
}

// GetIsBotOk returns a tuple with the IsBot field value
// and a boolean to check if the value has been set.
func (o *User) GetIsBotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBot, true
}

// SetIsBot sets field value
func (o *User) SetIsBot(v bool) {
	o.IsBot = v
}

// GetFirstName returns the FirstName field value
func (o *User) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *User) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}


// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}


// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *User) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *User) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *User) SetLanguageCode(v string) {
	o.LanguageCode = &v
}


// GetIsPremium returns the IsPremium field value if set, zero value otherwise.
func (o *User) GetIsPremium() bool {
	if o == nil || IsNil(o.IsPremium) {
		var ret bool
		return ret
	}
	return *o.IsPremium
}

// GetIsPremiumOk returns a tuple with the IsPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsPremiumOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPremium) {
		return nil, false
	}
	return o.IsPremium, true
}

// HasIsPremium returns a boolean if a field has been set.
func (o *User) HasIsPremium() bool {
	if o != nil && !IsNil(o.IsPremium) {
		return true
	}

	return false
}

// SetIsPremium gets a reference to the given bool and assigns it to the IsPremium field.
func (o *User) SetIsPremium(v bool) {
	o.IsPremium = &v
}


// GetAddedToAttachmentMenu returns the AddedToAttachmentMenu field value if set, zero value otherwise.
func (o *User) GetAddedToAttachmentMenu() bool {
	if o == nil || IsNil(o.AddedToAttachmentMenu) {
		var ret bool
		return ret
	}
	return *o.AddedToAttachmentMenu
}

// GetAddedToAttachmentMenuOk returns a tuple with the AddedToAttachmentMenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAddedToAttachmentMenuOk() (*bool, bool) {
	if o == nil || IsNil(o.AddedToAttachmentMenu) {
		return nil, false
	}
	return o.AddedToAttachmentMenu, true
}

// HasAddedToAttachmentMenu returns a boolean if a field has been set.
func (o *User) HasAddedToAttachmentMenu() bool {
	if o != nil && !IsNil(o.AddedToAttachmentMenu) {
		return true
	}

	return false
}

// SetAddedToAttachmentMenu gets a reference to the given bool and assigns it to the AddedToAttachmentMenu field.
func (o *User) SetAddedToAttachmentMenu(v bool) {
	o.AddedToAttachmentMenu = &v
}


// GetCanJoinGroups returns the CanJoinGroups field value if set, zero value otherwise.
func (o *User) GetCanJoinGroups() bool {
	if o == nil || IsNil(o.CanJoinGroups) {
		var ret bool
		return ret
	}
	return *o.CanJoinGroups
}

// GetCanJoinGroupsOk returns a tuple with the CanJoinGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCanJoinGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanJoinGroups) {
		return nil, false
	}
	return o.CanJoinGroups, true
}

// HasCanJoinGroups returns a boolean if a field has been set.
func (o *User) HasCanJoinGroups() bool {
	if o != nil && !IsNil(o.CanJoinGroups) {
		return true
	}

	return false
}

// SetCanJoinGroups gets a reference to the given bool and assigns it to the CanJoinGroups field.
func (o *User) SetCanJoinGroups(v bool) {
	o.CanJoinGroups = &v
}


// GetCanReadAllGroupMessages returns the CanReadAllGroupMessages field value if set, zero value otherwise.
func (o *User) GetCanReadAllGroupMessages() bool {
	if o == nil || IsNil(o.CanReadAllGroupMessages) {
		var ret bool
		return ret
	}
	return *o.CanReadAllGroupMessages
}

// GetCanReadAllGroupMessagesOk returns a tuple with the CanReadAllGroupMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCanReadAllGroupMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanReadAllGroupMessages) {
		return nil, false
	}
	return o.CanReadAllGroupMessages, true
}

// HasCanReadAllGroupMessages returns a boolean if a field has been set.
func (o *User) HasCanReadAllGroupMessages() bool {
	if o != nil && !IsNil(o.CanReadAllGroupMessages) {
		return true
	}

	return false
}

// SetCanReadAllGroupMessages gets a reference to the given bool and assigns it to the CanReadAllGroupMessages field.
func (o *User) SetCanReadAllGroupMessages(v bool) {
	o.CanReadAllGroupMessages = &v
}


// GetSupportsInlineQueries returns the SupportsInlineQueries field value if set, zero value otherwise.
func (o *User) GetSupportsInlineQueries() bool {
	if o == nil || IsNil(o.SupportsInlineQueries) {
		var ret bool
		return ret
	}
	return *o.SupportsInlineQueries
}

// GetSupportsInlineQueriesOk returns a tuple with the SupportsInlineQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSupportsInlineQueriesOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsInlineQueries) {
		return nil, false
	}
	return o.SupportsInlineQueries, true
}

// HasSupportsInlineQueries returns a boolean if a field has been set.
func (o *User) HasSupportsInlineQueries() bool {
	if o != nil && !IsNil(o.SupportsInlineQueries) {
		return true
	}

	return false
}

// SetSupportsInlineQueries gets a reference to the given bool and assigns it to the SupportsInlineQueries field.
func (o *User) SetSupportsInlineQueries(v bool) {
	o.SupportsInlineQueries = &v
}


// GetCanConnectToBusiness returns the CanConnectToBusiness field value if set, zero value otherwise.
func (o *User) GetCanConnectToBusiness() bool {
	if o == nil || IsNil(o.CanConnectToBusiness) {
		var ret bool
		return ret
	}
	return *o.CanConnectToBusiness
}

// GetCanConnectToBusinessOk returns a tuple with the CanConnectToBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCanConnectToBusinessOk() (*bool, bool) {
	if o == nil || IsNil(o.CanConnectToBusiness) {
		return nil, false
	}
	return o.CanConnectToBusiness, true
}

// HasCanConnectToBusiness returns a boolean if a field has been set.
func (o *User) HasCanConnectToBusiness() bool {
	if o != nil && !IsNil(o.CanConnectToBusiness) {
		return true
	}

	return false
}

// SetCanConnectToBusiness gets a reference to the given bool and assigns it to the CanConnectToBusiness field.
func (o *User) SetCanConnectToBusiness(v bool) {
	o.CanConnectToBusiness = &v
}


// GetHasMainWebApp returns the HasMainWebApp field value if set, zero value otherwise.
func (o *User) GetHasMainWebApp() bool {
	if o == nil || IsNil(o.HasMainWebApp) {
		var ret bool
		return ret
	}
	return *o.HasMainWebApp
}

// GetHasMainWebAppOk returns a tuple with the HasMainWebApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHasMainWebAppOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMainWebApp) {
		return nil, false
	}
	return o.HasMainWebApp, true
}

// HasHasMainWebApp returns a boolean if a field has been set.
func (o *User) HasHasMainWebApp() bool {
	if o != nil && !IsNil(o.HasMainWebApp) {
		return true
	}

	return false
}

// SetHasMainWebApp gets a reference to the given bool and assigns it to the HasMainWebApp field.
func (o *User) SetHasMainWebApp(v bool) {
	o.HasMainWebApp = &v
}


func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["is_bot"] = o.IsBot
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["language_code"] = o.LanguageCode
	}
	if !IsNil(o.IsPremium) {
		toSerialize["is_premium"] = o.IsPremium
	}
	if !IsNil(o.AddedToAttachmentMenu) {
		toSerialize["added_to_attachment_menu"] = o.AddedToAttachmentMenu
	}
	if !IsNil(o.CanJoinGroups) {
		toSerialize["can_join_groups"] = o.CanJoinGroups
	}
	if !IsNil(o.CanReadAllGroupMessages) {
		toSerialize["can_read_all_group_messages"] = o.CanReadAllGroupMessages
	}
	if !IsNil(o.SupportsInlineQueries) {
		toSerialize["supports_inline_queries"] = o.SupportsInlineQueries
	}
	if !IsNil(o.CanConnectToBusiness) {
		toSerialize["can_connect_to_business"] = o.CanConnectToBusiness
	}
	if !IsNil(o.HasMainWebApp) {
		toSerialize["has_main_web_app"] = o.HasMainWebApp
	}
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"is_bot",
		"first_name",
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

	varUser := _User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


