/** 
 * Telegram Bot API - REST API Client
 * Auto-generated OpenAPI schema
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-01T14:14:20.091913680Z[Etc/UTC]
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

// checks if the KeyboardButtonRequestChat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyboardButtonRequestChat{}

// KeyboardButtonRequestChat This object defines the criteria used to request a suitable chat. Information about the selected chat will be shared with the bot when the corresponding button is pressed. The bot will be granted requested rights in the chat if appropriate. [More about requesting chats »](https://core.telegram.org/bots/features#chat-and-user-selection).
type KeyboardButtonRequestChat struct {
	// Signed 32-bit identifier of the request, which will be received back in the [ChatShared](https://core.telegram.org/bots/api/#chatshared) object. Must be unique within the message
	RequestId int32 `json:"request_id"`
	// Pass *True* to request a channel chat, pass *False* to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`
	// *Optional*. Pass *True* to request a forum supergroup, pass *False* to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatIsForum *bool `json:"chat_is_forum,omitempty"`
	// *Optional*. Pass *True* to request a supergroup or a channel with a username, pass *False* to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatHasUsername *bool `json:"chat_has_username,omitempty"`
	// *Optional*. Pass *True* to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	ChatIsCreated *bool `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	// *Optional*. Pass *True* to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	BotIsMember *bool `json:"bot_is_member,omitempty"`
	// *Optional*. Pass *True* to request the chat's title
	RequestTitle *bool `json:"request_title,omitempty"`
	// *Optional*. Pass *True* to request the chat's username
	RequestUsername *bool `json:"request_username,omitempty"`
	// *Optional*. Pass *True* to request the chat's photo
	RequestPhoto *bool `json:"request_photo,omitempty"`
}

type _KeyboardButtonRequestChat KeyboardButtonRequestChat

// NewKeyboardButtonRequestChat instantiates a new KeyboardButtonRequestChat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyboardButtonRequestChat(requestId int32, chatIsChannel bool) *KeyboardButtonRequestChat {
	this := KeyboardButtonRequestChat{}
	this.RequestId = requestId
	this.ChatIsChannel = chatIsChannel
	return &this
}

// NewKeyboardButtonRequestChatWithDefaults instantiates a new KeyboardButtonRequestChat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyboardButtonRequestChatWithDefaults() *KeyboardButtonRequestChat {
	this := KeyboardButtonRequestChat{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *KeyboardButtonRequestChat) GetRequestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetRequestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *KeyboardButtonRequestChat) SetRequestId(v int32) {
	o.RequestId = v
}

// GetChatIsChannel returns the ChatIsChannel field value
func (o *KeyboardButtonRequestChat) GetChatIsChannel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ChatIsChannel
}

// GetChatIsChannelOk returns a tuple with the ChatIsChannel field value
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetChatIsChannelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatIsChannel, true
}

// SetChatIsChannel sets field value
func (o *KeyboardButtonRequestChat) SetChatIsChannel(v bool) {
	o.ChatIsChannel = v
}

// GetChatIsForum returns the ChatIsForum field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetChatIsForum() bool {
	if o == nil || IsNil(o.ChatIsForum) {
		var ret bool
		return ret
	}
	return *o.ChatIsForum
}

// GetChatIsForumOk returns a tuple with the ChatIsForum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetChatIsForumOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatIsForum) {
		return nil, false
	}
	return o.ChatIsForum, true
}

// HasChatIsForum returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasChatIsForum() bool {
	if o != nil && !IsNil(o.ChatIsForum) {
		return true
	}

	return false
}

// SetChatIsForum gets a reference to the given bool and assigns it to the ChatIsForum field.
func (o *KeyboardButtonRequestChat) SetChatIsForum(v bool) {
	o.ChatIsForum = &v
}


// GetChatHasUsername returns the ChatHasUsername field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetChatHasUsername() bool {
	if o == nil || IsNil(o.ChatHasUsername) {
		var ret bool
		return ret
	}
	return *o.ChatHasUsername
}

// GetChatHasUsernameOk returns a tuple with the ChatHasUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetChatHasUsernameOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatHasUsername) {
		return nil, false
	}
	return o.ChatHasUsername, true
}

// HasChatHasUsername returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasChatHasUsername() bool {
	if o != nil && !IsNil(o.ChatHasUsername) {
		return true
	}

	return false
}

// SetChatHasUsername gets a reference to the given bool and assigns it to the ChatHasUsername field.
func (o *KeyboardButtonRequestChat) SetChatHasUsername(v bool) {
	o.ChatHasUsername = &v
}


// GetChatIsCreated returns the ChatIsCreated field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetChatIsCreated() bool {
	if o == nil || IsNil(o.ChatIsCreated) {
		var ret bool
		return ret
	}
	return *o.ChatIsCreated
}

// GetChatIsCreatedOk returns a tuple with the ChatIsCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetChatIsCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.ChatIsCreated) {
		return nil, false
	}
	return o.ChatIsCreated, true
}

// HasChatIsCreated returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasChatIsCreated() bool {
	if o != nil && !IsNil(o.ChatIsCreated) {
		return true
	}

	return false
}

// SetChatIsCreated gets a reference to the given bool and assigns it to the ChatIsCreated field.
func (o *KeyboardButtonRequestChat) SetChatIsCreated(v bool) {
	o.ChatIsCreated = &v
}


// GetUserAdministratorRights returns the UserAdministratorRights field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetUserAdministratorRights() ChatAdministratorRights {
	if o == nil || IsNil(o.UserAdministratorRights) {
		var ret ChatAdministratorRights
		return ret
	}
	return *o.UserAdministratorRights
}

// GetUserAdministratorRightsOk returns a tuple with the UserAdministratorRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetUserAdministratorRightsOk() (*ChatAdministratorRights, bool) {
	if o == nil || IsNil(o.UserAdministratorRights) {
		return nil, false
	}
	return o.UserAdministratorRights, true
}

// HasUserAdministratorRights returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasUserAdministratorRights() bool {
	if o != nil && !IsNil(o.UserAdministratorRights) {
		return true
	}

	return false
}

// SetUserAdministratorRights gets a reference to the given ChatAdministratorRights and assigns it to the UserAdministratorRights field.
func (o *KeyboardButtonRequestChat) SetUserAdministratorRights(v ChatAdministratorRights) {
	o.UserAdministratorRights = &v
}


// GetBotAdministratorRights returns the BotAdministratorRights field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetBotAdministratorRights() ChatAdministratorRights {
	if o == nil || IsNil(o.BotAdministratorRights) {
		var ret ChatAdministratorRights
		return ret
	}
	return *o.BotAdministratorRights
}

// GetBotAdministratorRightsOk returns a tuple with the BotAdministratorRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetBotAdministratorRightsOk() (*ChatAdministratorRights, bool) {
	if o == nil || IsNil(o.BotAdministratorRights) {
		return nil, false
	}
	return o.BotAdministratorRights, true
}

// HasBotAdministratorRights returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasBotAdministratorRights() bool {
	if o != nil && !IsNil(o.BotAdministratorRights) {
		return true
	}

	return false
}

// SetBotAdministratorRights gets a reference to the given ChatAdministratorRights and assigns it to the BotAdministratorRights field.
func (o *KeyboardButtonRequestChat) SetBotAdministratorRights(v ChatAdministratorRights) {
	o.BotAdministratorRights = &v
}


// GetBotIsMember returns the BotIsMember field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetBotIsMember() bool {
	if o == nil || IsNil(o.BotIsMember) {
		var ret bool
		return ret
	}
	return *o.BotIsMember
}

// GetBotIsMemberOk returns a tuple with the BotIsMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetBotIsMemberOk() (*bool, bool) {
	if o == nil || IsNil(o.BotIsMember) {
		return nil, false
	}
	return o.BotIsMember, true
}

// HasBotIsMember returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasBotIsMember() bool {
	if o != nil && !IsNil(o.BotIsMember) {
		return true
	}

	return false
}

// SetBotIsMember gets a reference to the given bool and assigns it to the BotIsMember field.
func (o *KeyboardButtonRequestChat) SetBotIsMember(v bool) {
	o.BotIsMember = &v
}


// GetRequestTitle returns the RequestTitle field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetRequestTitle() bool {
	if o == nil || IsNil(o.RequestTitle) {
		var ret bool
		return ret
	}
	return *o.RequestTitle
}

// GetRequestTitleOk returns a tuple with the RequestTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetRequestTitleOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTitle) {
		return nil, false
	}
	return o.RequestTitle, true
}

// HasRequestTitle returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasRequestTitle() bool {
	if o != nil && !IsNil(o.RequestTitle) {
		return true
	}

	return false
}

// SetRequestTitle gets a reference to the given bool and assigns it to the RequestTitle field.
func (o *KeyboardButtonRequestChat) SetRequestTitle(v bool) {
	o.RequestTitle = &v
}


// GetRequestUsername returns the RequestUsername field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetRequestUsername() bool {
	if o == nil || IsNil(o.RequestUsername) {
		var ret bool
		return ret
	}
	return *o.RequestUsername
}

// GetRequestUsernameOk returns a tuple with the RequestUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetRequestUsernameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestUsername) {
		return nil, false
	}
	return o.RequestUsername, true
}

// HasRequestUsername returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasRequestUsername() bool {
	if o != nil && !IsNil(o.RequestUsername) {
		return true
	}

	return false
}

// SetRequestUsername gets a reference to the given bool and assigns it to the RequestUsername field.
func (o *KeyboardButtonRequestChat) SetRequestUsername(v bool) {
	o.RequestUsername = &v
}


// GetRequestPhoto returns the RequestPhoto field value if set, zero value otherwise.
func (o *KeyboardButtonRequestChat) GetRequestPhoto() bool {
	if o == nil || IsNil(o.RequestPhoto) {
		var ret bool
		return ret
	}
	return *o.RequestPhoto
}

// GetRequestPhotoOk returns a tuple with the RequestPhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyboardButtonRequestChat) GetRequestPhotoOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestPhoto) {
		return nil, false
	}
	return o.RequestPhoto, true
}

// HasRequestPhoto returns a boolean if a field has been set.
func (o *KeyboardButtonRequestChat) HasRequestPhoto() bool {
	if o != nil && !IsNil(o.RequestPhoto) {
		return true
	}

	return false
}

// SetRequestPhoto gets a reference to the given bool and assigns it to the RequestPhoto field.
func (o *KeyboardButtonRequestChat) SetRequestPhoto(v bool) {
	o.RequestPhoto = &v
}


func (o KeyboardButtonRequestChat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyboardButtonRequestChat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request_id"] = o.RequestId
	toSerialize["chat_is_channel"] = o.ChatIsChannel
	if !IsNil(o.ChatIsForum) {
		toSerialize["chat_is_forum"] = o.ChatIsForum
	}
	if !IsNil(o.ChatHasUsername) {
		toSerialize["chat_has_username"] = o.ChatHasUsername
	}
	if !IsNil(o.ChatIsCreated) {
		toSerialize["chat_is_created"] = o.ChatIsCreated
	}
	if !IsNil(o.UserAdministratorRights) {
		toSerialize["user_administrator_rights"] = o.UserAdministratorRights
	}
	if !IsNil(o.BotAdministratorRights) {
		toSerialize["bot_administrator_rights"] = o.BotAdministratorRights
	}
	if !IsNil(o.BotIsMember) {
		toSerialize["bot_is_member"] = o.BotIsMember
	}
	if !IsNil(o.RequestTitle) {
		toSerialize["request_title"] = o.RequestTitle
	}
	if !IsNil(o.RequestUsername) {
		toSerialize["request_username"] = o.RequestUsername
	}
	if !IsNil(o.RequestPhoto) {
		toSerialize["request_photo"] = o.RequestPhoto
	}
	return toSerialize, nil
}

func (o *KeyboardButtonRequestChat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"request_id",
		"chat_is_channel",
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

	varKeyboardButtonRequestChat := _KeyboardButtonRequestChat{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyboardButtonRequestChat)

	if err != nil {
		return err
	}

	*o = KeyboardButtonRequestChat(varKeyboardButtonRequestChat)

	return err
}

type NullableKeyboardButtonRequestChat struct {
	value *KeyboardButtonRequestChat
	isSet bool
}

func (v NullableKeyboardButtonRequestChat) Get() *KeyboardButtonRequestChat {
	return v.value
}

func (v *NullableKeyboardButtonRequestChat) Set(val *KeyboardButtonRequestChat) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyboardButtonRequestChat) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyboardButtonRequestChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyboardButtonRequestChat(val *KeyboardButtonRequestChat) *NullableKeyboardButtonRequestChat {
	return &NullableKeyboardButtonRequestChat{value: val, isSet: true}
}

func (v NullableKeyboardButtonRequestChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyboardButtonRequestChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


