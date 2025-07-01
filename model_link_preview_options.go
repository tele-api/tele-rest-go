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
)

// checks if the LinkPreviewOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkPreviewOptions{}

// LinkPreviewOptions Describes the options used for link preview generation.
type LinkPreviewOptions struct {
	// *Optional*. *True*, if the link preview is disabled
	IsDisabled *bool `json:"is_disabled,omitempty"`
	// *Optional*. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	Url *string `json:"url,omitempty"`
	// *Optional*. *True*, if the media in the link preview is supposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferSmallMedia *bool `json:"prefer_small_media,omitempty"`
	// *Optional*. *True*, if the media in the link preview is supposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia *bool `json:"prefer_large_media,omitempty"`
	// *Optional*. *True*, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
	ShowAboveText *bool `json:"show_above_text,omitempty"`
}

// NewLinkPreviewOptions instantiates a new LinkPreviewOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkPreviewOptions() *LinkPreviewOptions {
	this := LinkPreviewOptions{}
	return &this
}

// NewLinkPreviewOptionsWithDefaults instantiates a new LinkPreviewOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkPreviewOptionsWithDefaults() *LinkPreviewOptions {
	this := LinkPreviewOptions{}
	return &this
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *LinkPreviewOptions) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPreviewOptions) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *LinkPreviewOptions) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *LinkPreviewOptions) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}


// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LinkPreviewOptions) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPreviewOptions) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LinkPreviewOptions) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LinkPreviewOptions) SetUrl(v string) {
	o.Url = &v
}


// GetPreferSmallMedia returns the PreferSmallMedia field value if set, zero value otherwise.
func (o *LinkPreviewOptions) GetPreferSmallMedia() bool {
	if o == nil || IsNil(o.PreferSmallMedia) {
		var ret bool
		return ret
	}
	return *o.PreferSmallMedia
}

// GetPreferSmallMediaOk returns a tuple with the PreferSmallMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPreviewOptions) GetPreferSmallMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferSmallMedia) {
		return nil, false
	}
	return o.PreferSmallMedia, true
}

// HasPreferSmallMedia returns a boolean if a field has been set.
func (o *LinkPreviewOptions) HasPreferSmallMedia() bool {
	if o != nil && !IsNil(o.PreferSmallMedia) {
		return true
	}

	return false
}

// SetPreferSmallMedia gets a reference to the given bool and assigns it to the PreferSmallMedia field.
func (o *LinkPreviewOptions) SetPreferSmallMedia(v bool) {
	o.PreferSmallMedia = &v
}


// GetPreferLargeMedia returns the PreferLargeMedia field value if set, zero value otherwise.
func (o *LinkPreviewOptions) GetPreferLargeMedia() bool {
	if o == nil || IsNil(o.PreferLargeMedia) {
		var ret bool
		return ret
	}
	return *o.PreferLargeMedia
}

// GetPreferLargeMediaOk returns a tuple with the PreferLargeMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPreviewOptions) GetPreferLargeMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferLargeMedia) {
		return nil, false
	}
	return o.PreferLargeMedia, true
}

// HasPreferLargeMedia returns a boolean if a field has been set.
func (o *LinkPreviewOptions) HasPreferLargeMedia() bool {
	if o != nil && !IsNil(o.PreferLargeMedia) {
		return true
	}

	return false
}

// SetPreferLargeMedia gets a reference to the given bool and assigns it to the PreferLargeMedia field.
func (o *LinkPreviewOptions) SetPreferLargeMedia(v bool) {
	o.PreferLargeMedia = &v
}


// GetShowAboveText returns the ShowAboveText field value if set, zero value otherwise.
func (o *LinkPreviewOptions) GetShowAboveText() bool {
	if o == nil || IsNil(o.ShowAboveText) {
		var ret bool
		return ret
	}
	return *o.ShowAboveText
}

// GetShowAboveTextOk returns a tuple with the ShowAboveText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkPreviewOptions) GetShowAboveTextOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAboveText) {
		return nil, false
	}
	return o.ShowAboveText, true
}

// HasShowAboveText returns a boolean if a field has been set.
func (o *LinkPreviewOptions) HasShowAboveText() bool {
	if o != nil && !IsNil(o.ShowAboveText) {
		return true
	}

	return false
}

// SetShowAboveText gets a reference to the given bool and assigns it to the ShowAboveText field.
func (o *LinkPreviewOptions) SetShowAboveText(v bool) {
	o.ShowAboveText = &v
}


func (o LinkPreviewOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkPreviewOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDisabled) {
		toSerialize["is_disabled"] = o.IsDisabled
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.PreferSmallMedia) {
		toSerialize["prefer_small_media"] = o.PreferSmallMedia
	}
	if !IsNil(o.PreferLargeMedia) {
		toSerialize["prefer_large_media"] = o.PreferLargeMedia
	}
	if !IsNil(o.ShowAboveText) {
		toSerialize["show_above_text"] = o.ShowAboveText
	}
	return toSerialize, nil
}

type NullableLinkPreviewOptions struct {
	value *LinkPreviewOptions
	isSet bool
}

func (v NullableLinkPreviewOptions) Get() *LinkPreviewOptions {
	return v.value
}

func (v *NullableLinkPreviewOptions) Set(val *LinkPreviewOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkPreviewOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkPreviewOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkPreviewOptions(val *LinkPreviewOptions) *NullableLinkPreviewOptions {
	return &NullableLinkPreviewOptions{value: val, isSet: true}
}

func (v NullableLinkPreviewOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkPreviewOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


