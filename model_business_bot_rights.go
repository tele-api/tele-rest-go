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
)

// checks if the BusinessBotRights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessBotRights{}

// BusinessBotRights Represents the rights of a business bot.
type BusinessBotRights struct {
	// *Optional*. True, if the bot can send and edit messages in the private chats that had incoming messages in the last 24 hours
	CanReply *bool `json:"can_reply,omitempty"`
	// *Optional*. True, if the bot can mark incoming private messages as read
	CanReadMessages *bool `json:"can_read_messages,omitempty"`
	// *Optional*. True, if the bot can delete messages sent by the bot
	CanDeleteSentMessages *bool `json:"can_delete_sent_messages,omitempty"`
	// *Optional*. True, if the bot can delete all private messages in managed chats
	CanDeleteAllMessages *bool `json:"can_delete_all_messages,omitempty"`
	// *Optional*. True, if the bot can edit the first and last name of the business account
	CanEditName *bool `json:"can_edit_name,omitempty"`
	// *Optional*. True, if the bot can edit the bio of the business account
	CanEditBio *bool `json:"can_edit_bio,omitempty"`
	// *Optional*. True, if the bot can edit the profile photo of the business account
	CanEditProfilePhoto *bool `json:"can_edit_profile_photo,omitempty"`
	// *Optional*. True, if the bot can edit the username of the business account
	CanEditUsername *bool `json:"can_edit_username,omitempty"`
	// *Optional*. True, if the bot can change the privacy settings pertaining to gifts for the business account
	CanChangeGiftSettings *bool `json:"can_change_gift_settings,omitempty"`
	// *Optional*. True, if the bot can view gifts and the amount of Telegram Stars owned by the business account
	CanViewGiftsAndStars *bool `json:"can_view_gifts_and_stars,omitempty"`
	// *Optional*. True, if the bot can convert regular gifts owned by the business account to Telegram Stars
	CanConvertGiftsToStars *bool `json:"can_convert_gifts_to_stars,omitempty"`
	// *Optional*. True, if the bot can transfer and upgrade gifts owned by the business account
	CanTransferAndUpgradeGifts *bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	// *Optional*. True, if the bot can transfer Telegram Stars received by the business account to its own account, or use them to upgrade and transfer gifts
	CanTransferStars *bool `json:"can_transfer_stars,omitempty"`
	// *Optional*. True, if the bot can post, edit and delete stories on behalf of the business account
	CanManageStories *bool `json:"can_manage_stories,omitempty"`
}

// NewBusinessBotRights instantiates a new BusinessBotRights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessBotRights() *BusinessBotRights {
	this := BusinessBotRights{}
	var canReply bool = true
	this.CanReply = &canReply
	var canReadMessages bool = true
	this.CanReadMessages = &canReadMessages
	var canDeleteSentMessages bool = true
	this.CanDeleteSentMessages = &canDeleteSentMessages
	var canDeleteAllMessages bool = true
	this.CanDeleteAllMessages = &canDeleteAllMessages
	var canEditName bool = true
	this.CanEditName = &canEditName
	var canEditBio bool = true
	this.CanEditBio = &canEditBio
	var canEditProfilePhoto bool = true
	this.CanEditProfilePhoto = &canEditProfilePhoto
	var canEditUsername bool = true
	this.CanEditUsername = &canEditUsername
	var canChangeGiftSettings bool = true
	this.CanChangeGiftSettings = &canChangeGiftSettings
	var canViewGiftsAndStars bool = true
	this.CanViewGiftsAndStars = &canViewGiftsAndStars
	var canConvertGiftsToStars bool = true
	this.CanConvertGiftsToStars = &canConvertGiftsToStars
	var canTransferAndUpgradeGifts bool = true
	this.CanTransferAndUpgradeGifts = &canTransferAndUpgradeGifts
	var canTransferStars bool = true
	this.CanTransferStars = &canTransferStars
	var canManageStories bool = true
	this.CanManageStories = &canManageStories
	return &this
}

// NewBusinessBotRightsWithDefaults instantiates a new BusinessBotRights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessBotRightsWithDefaults() *BusinessBotRights {
	this := BusinessBotRights{}
	var canReply bool = true
	this.CanReply = &canReply
	var canReadMessages bool = true
	this.CanReadMessages = &canReadMessages
	var canDeleteSentMessages bool = true
	this.CanDeleteSentMessages = &canDeleteSentMessages
	var canDeleteAllMessages bool = true
	this.CanDeleteAllMessages = &canDeleteAllMessages
	var canEditName bool = true
	this.CanEditName = &canEditName
	var canEditBio bool = true
	this.CanEditBio = &canEditBio
	var canEditProfilePhoto bool = true
	this.CanEditProfilePhoto = &canEditProfilePhoto
	var canEditUsername bool = true
	this.CanEditUsername = &canEditUsername
	var canChangeGiftSettings bool = true
	this.CanChangeGiftSettings = &canChangeGiftSettings
	var canViewGiftsAndStars bool = true
	this.CanViewGiftsAndStars = &canViewGiftsAndStars
	var canConvertGiftsToStars bool = true
	this.CanConvertGiftsToStars = &canConvertGiftsToStars
	var canTransferAndUpgradeGifts bool = true
	this.CanTransferAndUpgradeGifts = &canTransferAndUpgradeGifts
	var canTransferStars bool = true
	this.CanTransferStars = &canTransferStars
	var canManageStories bool = true
	this.CanManageStories = &canManageStories
	return &this
}

// GetCanReply returns the CanReply field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanReply() bool {
	if o == nil || IsNil(o.CanReply) {
		var ret bool
		return ret
	}
	return *o.CanReply
}

// GetCanReplyOk returns a tuple with the CanReply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanReplyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanReply) {
		return nil, false
	}
	return o.CanReply, true
}

// HasCanReply returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanReply() bool {
	if o != nil && !IsNil(o.CanReply) {
		return true
	}

	return false
}

// SetCanReply gets a reference to the given bool and assigns it to the CanReply field.
func (o *BusinessBotRights) SetCanReply(v bool) {
	o.CanReply = &v
}


// GetCanReadMessages returns the CanReadMessages field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanReadMessages() bool {
	if o == nil || IsNil(o.CanReadMessages) {
		var ret bool
		return ret
	}
	return *o.CanReadMessages
}

// GetCanReadMessagesOk returns a tuple with the CanReadMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanReadMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanReadMessages) {
		return nil, false
	}
	return o.CanReadMessages, true
}

// HasCanReadMessages returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanReadMessages() bool {
	if o != nil && !IsNil(o.CanReadMessages) {
		return true
	}

	return false
}

// SetCanReadMessages gets a reference to the given bool and assigns it to the CanReadMessages field.
func (o *BusinessBotRights) SetCanReadMessages(v bool) {
	o.CanReadMessages = &v
}


// GetCanDeleteSentMessages returns the CanDeleteSentMessages field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanDeleteSentMessages() bool {
	if o == nil || IsNil(o.CanDeleteSentMessages) {
		var ret bool
		return ret
	}
	return *o.CanDeleteSentMessages
}

// GetCanDeleteSentMessagesOk returns a tuple with the CanDeleteSentMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanDeleteSentMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteSentMessages) {
		return nil, false
	}
	return o.CanDeleteSentMessages, true
}

// HasCanDeleteSentMessages returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanDeleteSentMessages() bool {
	if o != nil && !IsNil(o.CanDeleteSentMessages) {
		return true
	}

	return false
}

// SetCanDeleteSentMessages gets a reference to the given bool and assigns it to the CanDeleteSentMessages field.
func (o *BusinessBotRights) SetCanDeleteSentMessages(v bool) {
	o.CanDeleteSentMessages = &v
}


// GetCanDeleteAllMessages returns the CanDeleteAllMessages field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanDeleteAllMessages() bool {
	if o == nil || IsNil(o.CanDeleteAllMessages) {
		var ret bool
		return ret
	}
	return *o.CanDeleteAllMessages
}

// GetCanDeleteAllMessagesOk returns a tuple with the CanDeleteAllMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanDeleteAllMessagesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDeleteAllMessages) {
		return nil, false
	}
	return o.CanDeleteAllMessages, true
}

// HasCanDeleteAllMessages returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanDeleteAllMessages() bool {
	if o != nil && !IsNil(o.CanDeleteAllMessages) {
		return true
	}

	return false
}

// SetCanDeleteAllMessages gets a reference to the given bool and assigns it to the CanDeleteAllMessages field.
func (o *BusinessBotRights) SetCanDeleteAllMessages(v bool) {
	o.CanDeleteAllMessages = &v
}


// GetCanEditName returns the CanEditName field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanEditName() bool {
	if o == nil || IsNil(o.CanEditName) {
		var ret bool
		return ret
	}
	return *o.CanEditName
}

// GetCanEditNameOk returns a tuple with the CanEditName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanEditNameOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditName) {
		return nil, false
	}
	return o.CanEditName, true
}

// HasCanEditName returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanEditName() bool {
	if o != nil && !IsNil(o.CanEditName) {
		return true
	}

	return false
}

// SetCanEditName gets a reference to the given bool and assigns it to the CanEditName field.
func (o *BusinessBotRights) SetCanEditName(v bool) {
	o.CanEditName = &v
}


// GetCanEditBio returns the CanEditBio field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanEditBio() bool {
	if o == nil || IsNil(o.CanEditBio) {
		var ret bool
		return ret
	}
	return *o.CanEditBio
}

// GetCanEditBioOk returns a tuple with the CanEditBio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanEditBioOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditBio) {
		return nil, false
	}
	return o.CanEditBio, true
}

// HasCanEditBio returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanEditBio() bool {
	if o != nil && !IsNil(o.CanEditBio) {
		return true
	}

	return false
}

// SetCanEditBio gets a reference to the given bool and assigns it to the CanEditBio field.
func (o *BusinessBotRights) SetCanEditBio(v bool) {
	o.CanEditBio = &v
}


// GetCanEditProfilePhoto returns the CanEditProfilePhoto field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanEditProfilePhoto() bool {
	if o == nil || IsNil(o.CanEditProfilePhoto) {
		var ret bool
		return ret
	}
	return *o.CanEditProfilePhoto
}

// GetCanEditProfilePhotoOk returns a tuple with the CanEditProfilePhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanEditProfilePhotoOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditProfilePhoto) {
		return nil, false
	}
	return o.CanEditProfilePhoto, true
}

// HasCanEditProfilePhoto returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanEditProfilePhoto() bool {
	if o != nil && !IsNil(o.CanEditProfilePhoto) {
		return true
	}

	return false
}

// SetCanEditProfilePhoto gets a reference to the given bool and assigns it to the CanEditProfilePhoto field.
func (o *BusinessBotRights) SetCanEditProfilePhoto(v bool) {
	o.CanEditProfilePhoto = &v
}


// GetCanEditUsername returns the CanEditUsername field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanEditUsername() bool {
	if o == nil || IsNil(o.CanEditUsername) {
		var ret bool
		return ret
	}
	return *o.CanEditUsername
}

// GetCanEditUsernameOk returns a tuple with the CanEditUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanEditUsernameOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEditUsername) {
		return nil, false
	}
	return o.CanEditUsername, true
}

// HasCanEditUsername returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanEditUsername() bool {
	if o != nil && !IsNil(o.CanEditUsername) {
		return true
	}

	return false
}

// SetCanEditUsername gets a reference to the given bool and assigns it to the CanEditUsername field.
func (o *BusinessBotRights) SetCanEditUsername(v bool) {
	o.CanEditUsername = &v
}


// GetCanChangeGiftSettings returns the CanChangeGiftSettings field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanChangeGiftSettings() bool {
	if o == nil || IsNil(o.CanChangeGiftSettings) {
		var ret bool
		return ret
	}
	return *o.CanChangeGiftSettings
}

// GetCanChangeGiftSettingsOk returns a tuple with the CanChangeGiftSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanChangeGiftSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanChangeGiftSettings) {
		return nil, false
	}
	return o.CanChangeGiftSettings, true
}

// HasCanChangeGiftSettings returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanChangeGiftSettings() bool {
	if o != nil && !IsNil(o.CanChangeGiftSettings) {
		return true
	}

	return false
}

// SetCanChangeGiftSettings gets a reference to the given bool and assigns it to the CanChangeGiftSettings field.
func (o *BusinessBotRights) SetCanChangeGiftSettings(v bool) {
	o.CanChangeGiftSettings = &v
}


// GetCanViewGiftsAndStars returns the CanViewGiftsAndStars field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanViewGiftsAndStars() bool {
	if o == nil || IsNil(o.CanViewGiftsAndStars) {
		var ret bool
		return ret
	}
	return *o.CanViewGiftsAndStars
}

// GetCanViewGiftsAndStarsOk returns a tuple with the CanViewGiftsAndStars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanViewGiftsAndStarsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanViewGiftsAndStars) {
		return nil, false
	}
	return o.CanViewGiftsAndStars, true
}

// HasCanViewGiftsAndStars returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanViewGiftsAndStars() bool {
	if o != nil && !IsNil(o.CanViewGiftsAndStars) {
		return true
	}

	return false
}

// SetCanViewGiftsAndStars gets a reference to the given bool and assigns it to the CanViewGiftsAndStars field.
func (o *BusinessBotRights) SetCanViewGiftsAndStars(v bool) {
	o.CanViewGiftsAndStars = &v
}


// GetCanConvertGiftsToStars returns the CanConvertGiftsToStars field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanConvertGiftsToStars() bool {
	if o == nil || IsNil(o.CanConvertGiftsToStars) {
		var ret bool
		return ret
	}
	return *o.CanConvertGiftsToStars
}

// GetCanConvertGiftsToStarsOk returns a tuple with the CanConvertGiftsToStars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanConvertGiftsToStarsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanConvertGiftsToStars) {
		return nil, false
	}
	return o.CanConvertGiftsToStars, true
}

// HasCanConvertGiftsToStars returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanConvertGiftsToStars() bool {
	if o != nil && !IsNil(o.CanConvertGiftsToStars) {
		return true
	}

	return false
}

// SetCanConvertGiftsToStars gets a reference to the given bool and assigns it to the CanConvertGiftsToStars field.
func (o *BusinessBotRights) SetCanConvertGiftsToStars(v bool) {
	o.CanConvertGiftsToStars = &v
}


// GetCanTransferAndUpgradeGifts returns the CanTransferAndUpgradeGifts field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanTransferAndUpgradeGifts() bool {
	if o == nil || IsNil(o.CanTransferAndUpgradeGifts) {
		var ret bool
		return ret
	}
	return *o.CanTransferAndUpgradeGifts
}

// GetCanTransferAndUpgradeGiftsOk returns a tuple with the CanTransferAndUpgradeGifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanTransferAndUpgradeGiftsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTransferAndUpgradeGifts) {
		return nil, false
	}
	return o.CanTransferAndUpgradeGifts, true
}

// HasCanTransferAndUpgradeGifts returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanTransferAndUpgradeGifts() bool {
	if o != nil && !IsNil(o.CanTransferAndUpgradeGifts) {
		return true
	}

	return false
}

// SetCanTransferAndUpgradeGifts gets a reference to the given bool and assigns it to the CanTransferAndUpgradeGifts field.
func (o *BusinessBotRights) SetCanTransferAndUpgradeGifts(v bool) {
	o.CanTransferAndUpgradeGifts = &v
}


// GetCanTransferStars returns the CanTransferStars field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanTransferStars() bool {
	if o == nil || IsNil(o.CanTransferStars) {
		var ret bool
		return ret
	}
	return *o.CanTransferStars
}

// GetCanTransferStarsOk returns a tuple with the CanTransferStars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanTransferStarsOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTransferStars) {
		return nil, false
	}
	return o.CanTransferStars, true
}

// HasCanTransferStars returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanTransferStars() bool {
	if o != nil && !IsNil(o.CanTransferStars) {
		return true
	}

	return false
}

// SetCanTransferStars gets a reference to the given bool and assigns it to the CanTransferStars field.
func (o *BusinessBotRights) SetCanTransferStars(v bool) {
	o.CanTransferStars = &v
}


// GetCanManageStories returns the CanManageStories field value if set, zero value otherwise.
func (o *BusinessBotRights) GetCanManageStories() bool {
	if o == nil || IsNil(o.CanManageStories) {
		var ret bool
		return ret
	}
	return *o.CanManageStories
}

// GetCanManageStoriesOk returns a tuple with the CanManageStories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessBotRights) GetCanManageStoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManageStories) {
		return nil, false
	}
	return o.CanManageStories, true
}

// HasCanManageStories returns a boolean if a field has been set.
func (o *BusinessBotRights) HasCanManageStories() bool {
	if o != nil && !IsNil(o.CanManageStories) {
		return true
	}

	return false
}

// SetCanManageStories gets a reference to the given bool and assigns it to the CanManageStories field.
func (o *BusinessBotRights) SetCanManageStories(v bool) {
	o.CanManageStories = &v
}


func (o BusinessBotRights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessBotRights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanReply) {
		toSerialize["can_reply"] = o.CanReply
	}
	if !IsNil(o.CanReadMessages) {
		toSerialize["can_read_messages"] = o.CanReadMessages
	}
	if !IsNil(o.CanDeleteSentMessages) {
		toSerialize["can_delete_sent_messages"] = o.CanDeleteSentMessages
	}
	if !IsNil(o.CanDeleteAllMessages) {
		toSerialize["can_delete_all_messages"] = o.CanDeleteAllMessages
	}
	if !IsNil(o.CanEditName) {
		toSerialize["can_edit_name"] = o.CanEditName
	}
	if !IsNil(o.CanEditBio) {
		toSerialize["can_edit_bio"] = o.CanEditBio
	}
	if !IsNil(o.CanEditProfilePhoto) {
		toSerialize["can_edit_profile_photo"] = o.CanEditProfilePhoto
	}
	if !IsNil(o.CanEditUsername) {
		toSerialize["can_edit_username"] = o.CanEditUsername
	}
	if !IsNil(o.CanChangeGiftSettings) {
		toSerialize["can_change_gift_settings"] = o.CanChangeGiftSettings
	}
	if !IsNil(o.CanViewGiftsAndStars) {
		toSerialize["can_view_gifts_and_stars"] = o.CanViewGiftsAndStars
	}
	if !IsNil(o.CanConvertGiftsToStars) {
		toSerialize["can_convert_gifts_to_stars"] = o.CanConvertGiftsToStars
	}
	if !IsNil(o.CanTransferAndUpgradeGifts) {
		toSerialize["can_transfer_and_upgrade_gifts"] = o.CanTransferAndUpgradeGifts
	}
	if !IsNil(o.CanTransferStars) {
		toSerialize["can_transfer_stars"] = o.CanTransferStars
	}
	if !IsNil(o.CanManageStories) {
		toSerialize["can_manage_stories"] = o.CanManageStories
	}
	return toSerialize, nil
}

type NullableBusinessBotRights struct {
	value *BusinessBotRights
	isSet bool
}

func (v NullableBusinessBotRights) Get() *BusinessBotRights {
	return v.value
}

func (v *NullableBusinessBotRights) Set(val *BusinessBotRights) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessBotRights) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessBotRights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessBotRights(val *BusinessBotRights) *NullableBusinessBotRights {
	return &NullableBusinessBotRights{value: val, isSet: true}
}

func (v NullableBusinessBotRights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessBotRights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


