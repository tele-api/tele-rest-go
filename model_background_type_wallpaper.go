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

// checks if the BackgroundTypeWallpaper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundTypeWallpaper{}

// BackgroundTypeWallpaper The background is a wallpaper in the JPEG format.
type BackgroundTypeWallpaper struct {
	// Type of the background, always “wallpaper”
	Type string `json:"type"`
	Document Document `json:"document"`
	// Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int32 `json:"dark_theme_dimming"`
	// *Optional*. *True*, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsBlurred *bool `json:"is_blurred,omitempty"`
	// *Optional*. *True*, if the background moves slightly when the device is tilted
	IsMoving *bool `json:"is_moving,omitempty"`
}

type _BackgroundTypeWallpaper BackgroundTypeWallpaper

// NewBackgroundTypeWallpaper instantiates a new BackgroundTypeWallpaper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundTypeWallpaper(type_ string, document Document, darkThemeDimming int32) *BackgroundTypeWallpaper {
	this := BackgroundTypeWallpaper{}
	this.Type = type_
	this.Document = document
	this.DarkThemeDimming = darkThemeDimming
	var isBlurred bool = true
	this.IsBlurred = &isBlurred
	var isMoving bool = true
	this.IsMoving = &isMoving
	return &this
}

// NewBackgroundTypeWallpaperWithDefaults instantiates a new BackgroundTypeWallpaper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundTypeWallpaperWithDefaults() *BackgroundTypeWallpaper {
	this := BackgroundTypeWallpaper{}
	var type_ string = "wallpaper"
	this.Type = type_
	var isBlurred bool = true
	this.IsBlurred = &isBlurred
	var isMoving bool = true
	this.IsMoving = &isMoving
	return &this
}

// GetType returns the Type field value
func (o *BackgroundTypeWallpaper) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypeWallpaper) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BackgroundTypeWallpaper) SetType(v string) {
	o.Type = v
}

// GetDocument returns the Document field value
func (o *BackgroundTypeWallpaper) GetDocument() Document {
	if o == nil {
		var ret Document
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypeWallpaper) GetDocumentOk() (*Document, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *BackgroundTypeWallpaper) SetDocument(v Document) {
	o.Document = v
}

// GetDarkThemeDimming returns the DarkThemeDimming field value
func (o *BackgroundTypeWallpaper) GetDarkThemeDimming() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DarkThemeDimming
}

// GetDarkThemeDimmingOk returns a tuple with the DarkThemeDimming field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypeWallpaper) GetDarkThemeDimmingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DarkThemeDimming, true
}

// SetDarkThemeDimming sets field value
func (o *BackgroundTypeWallpaper) SetDarkThemeDimming(v int32) {
	o.DarkThemeDimming = v
}

// GetIsBlurred returns the IsBlurred field value if set, zero value otherwise.
func (o *BackgroundTypeWallpaper) GetIsBlurred() bool {
	if o == nil || IsNil(o.IsBlurred) {
		var ret bool
		return ret
	}
	return *o.IsBlurred
}

// GetIsBlurredOk returns a tuple with the IsBlurred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackgroundTypeWallpaper) GetIsBlurredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBlurred) {
		return nil, false
	}
	return o.IsBlurred, true
}

// HasIsBlurred returns a boolean if a field has been set.
func (o *BackgroundTypeWallpaper) HasIsBlurred() bool {
	if o != nil && !IsNil(o.IsBlurred) {
		return true
	}

	return false
}

// SetIsBlurred gets a reference to the given bool and assigns it to the IsBlurred field.
func (o *BackgroundTypeWallpaper) SetIsBlurred(v bool) {
	o.IsBlurred = &v
}


// GetIsMoving returns the IsMoving field value if set, zero value otherwise.
func (o *BackgroundTypeWallpaper) GetIsMoving() bool {
	if o == nil || IsNil(o.IsMoving) {
		var ret bool
		return ret
	}
	return *o.IsMoving
}

// GetIsMovingOk returns a tuple with the IsMoving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackgroundTypeWallpaper) GetIsMovingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMoving) {
		return nil, false
	}
	return o.IsMoving, true
}

// HasIsMoving returns a boolean if a field has been set.
func (o *BackgroundTypeWallpaper) HasIsMoving() bool {
	if o != nil && !IsNil(o.IsMoving) {
		return true
	}

	return false
}

// SetIsMoving gets a reference to the given bool and assigns it to the IsMoving field.
func (o *BackgroundTypeWallpaper) SetIsMoving(v bool) {
	o.IsMoving = &v
}


func (o BackgroundTypeWallpaper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundTypeWallpaper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["document"] = o.Document
	toSerialize["dark_theme_dimming"] = o.DarkThemeDimming
	if !IsNil(o.IsBlurred) {
		toSerialize["is_blurred"] = o.IsBlurred
	}
	if !IsNil(o.IsMoving) {
		toSerialize["is_moving"] = o.IsMoving
	}
	return toSerialize, nil
}

func (o *BackgroundTypeWallpaper) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"document",
		"dark_theme_dimming",
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

	varBackgroundTypeWallpaper := _BackgroundTypeWallpaper{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackgroundTypeWallpaper)

	if err != nil {
		return err
	}

	*o = BackgroundTypeWallpaper(varBackgroundTypeWallpaper)

	return err
}

type NullableBackgroundTypeWallpaper struct {
	value *BackgroundTypeWallpaper
	isSet bool
}

func (v NullableBackgroundTypeWallpaper) Get() *BackgroundTypeWallpaper {
	return v.value
}

func (v *NullableBackgroundTypeWallpaper) Set(val *BackgroundTypeWallpaper) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTypeWallpaper) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTypeWallpaper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTypeWallpaper(val *BackgroundTypeWallpaper) *NullableBackgroundTypeWallpaper {
	return &NullableBackgroundTypeWallpaper{value: val, isSet: true}
}

func (v NullableBackgroundTypeWallpaper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTypeWallpaper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


