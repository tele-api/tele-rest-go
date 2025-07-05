/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.1.0
 *    * - **Modified**: 2025-07-05T02:41:44.515216840Z[Etc/UTC]
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

// checks if the WriteAccessAllowed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteAccessAllowed{}

// WriteAccessAllowed This object represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method [requestWriteAccess](https://core.telegram.org/bots/webapps#initializing-mini-apps).
type WriteAccessAllowed struct {
	// *Optional*. *True*, if the access was granted after the user accepted an explicit request from a Web App sent by the method [requestWriteAccess](https://core.telegram.org/bots/webapps#initializing-mini-apps)
	FromRequest *bool `json:"from_request,omitempty"`
	// *Optional*. Name of the Web App, if the access was granted when the Web App was launched from a link
	WebAppName *string `json:"web_app_name,omitempty"`
	// *Optional*. *True*, if the access was granted when the bot was added to the attachment or side menu
	FromAttachmentMenu *bool `json:"from_attachment_menu,omitempty"`
}

// NewWriteAccessAllowed instantiates a new WriteAccessAllowed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteAccessAllowed() *WriteAccessAllowed {
	this := WriteAccessAllowed{}
	return &this
}

// NewWriteAccessAllowedWithDefaults instantiates a new WriteAccessAllowed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteAccessAllowedWithDefaults() *WriteAccessAllowed {
	this := WriteAccessAllowed{}
	return &this
}

// GetFromRequest returns the FromRequest field value if set, zero value otherwise.
func (o *WriteAccessAllowed) GetFromRequest() bool {
	if o == nil || IsNil(o.FromRequest) {
		var ret bool
		return ret
	}
	return *o.FromRequest
}

// GetFromRequestOk returns a tuple with the FromRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteAccessAllowed) GetFromRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.FromRequest) {
		return nil, false
	}
	return o.FromRequest, true
}

// HasFromRequest returns a boolean if a field has been set.
func (o *WriteAccessAllowed) HasFromRequest() bool {
	if o != nil && !IsNil(o.FromRequest) {
		return true
	}

	return false
}

// SetFromRequest gets a reference to the given bool and assigns it to the FromRequest field.
func (o *WriteAccessAllowed) SetFromRequest(v bool) {
	o.FromRequest = &v
}


// GetWebAppName returns the WebAppName field value if set, zero value otherwise.
func (o *WriteAccessAllowed) GetWebAppName() string {
	if o == nil || IsNil(o.WebAppName) {
		var ret string
		return ret
	}
	return *o.WebAppName
}

// GetWebAppNameOk returns a tuple with the WebAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteAccessAllowed) GetWebAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.WebAppName) {
		return nil, false
	}
	return o.WebAppName, true
}

// HasWebAppName returns a boolean if a field has been set.
func (o *WriteAccessAllowed) HasWebAppName() bool {
	if o != nil && !IsNil(o.WebAppName) {
		return true
	}

	return false
}

// SetWebAppName gets a reference to the given string and assigns it to the WebAppName field.
func (o *WriteAccessAllowed) SetWebAppName(v string) {
	o.WebAppName = &v
}


// GetFromAttachmentMenu returns the FromAttachmentMenu field value if set, zero value otherwise.
func (o *WriteAccessAllowed) GetFromAttachmentMenu() bool {
	if o == nil || IsNil(o.FromAttachmentMenu) {
		var ret bool
		return ret
	}
	return *o.FromAttachmentMenu
}

// GetFromAttachmentMenuOk returns a tuple with the FromAttachmentMenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteAccessAllowed) GetFromAttachmentMenuOk() (*bool, bool) {
	if o == nil || IsNil(o.FromAttachmentMenu) {
		return nil, false
	}
	return o.FromAttachmentMenu, true
}

// HasFromAttachmentMenu returns a boolean if a field has been set.
func (o *WriteAccessAllowed) HasFromAttachmentMenu() bool {
	if o != nil && !IsNil(o.FromAttachmentMenu) {
		return true
	}

	return false
}

// SetFromAttachmentMenu gets a reference to the given bool and assigns it to the FromAttachmentMenu field.
func (o *WriteAccessAllowed) SetFromAttachmentMenu(v bool) {
	o.FromAttachmentMenu = &v
}


func (o WriteAccessAllowed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteAccessAllowed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromRequest) {
		toSerialize["from_request"] = o.FromRequest
	}
	if !IsNil(o.WebAppName) {
		toSerialize["web_app_name"] = o.WebAppName
	}
	if !IsNil(o.FromAttachmentMenu) {
		toSerialize["from_attachment_menu"] = o.FromAttachmentMenu
	}
	return toSerialize, nil
}

type NullableWriteAccessAllowed struct {
	value *WriteAccessAllowed
	isSet bool
}

func (v NullableWriteAccessAllowed) Get() *WriteAccessAllowed {
	return v.value
}

func (v *NullableWriteAccessAllowed) Set(val *WriteAccessAllowed) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteAccessAllowed) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteAccessAllowed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteAccessAllowed(val *WriteAccessAllowed) *NullableWriteAccessAllowed {
	return &NullableWriteAccessAllowed{value: val, isSet: true}
}

func (v NullableWriteAccessAllowed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteAccessAllowed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


