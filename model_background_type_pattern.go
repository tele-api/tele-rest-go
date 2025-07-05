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
	"bytes"
	"fmt"
)

// checks if the BackgroundTypePattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundTypePattern{}

// BackgroundTypePattern The background is a .PNG or .TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”) pattern to be combined with the background fill chosen by the user.
type BackgroundTypePattern struct {
	// Type of the background, always “pattern”
	Type string `json:"type"`
	Document Document `json:"document"`
	Fill BackgroundFill `json:"fill"`
	// Intensity of the pattern when it is shown above the filled background; 0-100
	Intensity int32 `json:"intensity"`
	// *Optional*. *True*, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	IsInverted *bool `json:"is_inverted,omitempty"`
	// *Optional*. *True*, if the background moves slightly when the device is tilted
	IsMoving *bool `json:"is_moving,omitempty"`
}

type _BackgroundTypePattern BackgroundTypePattern

// NewBackgroundTypePattern instantiates a new BackgroundTypePattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundTypePattern(type_ string, document Document, fill BackgroundFill, intensity int32) *BackgroundTypePattern {
	this := BackgroundTypePattern{}
	this.Type = type_
	this.Document = document
	this.Fill = fill
	this.Intensity = intensity
	var isInverted bool = true
	this.IsInverted = &isInverted
	var isMoving bool = true
	this.IsMoving = &isMoving
	return &this
}

// NewBackgroundTypePatternWithDefaults instantiates a new BackgroundTypePattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundTypePatternWithDefaults() *BackgroundTypePattern {
	this := BackgroundTypePattern{}
	var type_ string = "pattern"
	this.Type = type_
	var isInverted bool = true
	this.IsInverted = &isInverted
	var isMoving bool = true
	this.IsMoving = &isMoving
	return &this
}

// GetType returns the Type field value
func (o *BackgroundTypePattern) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BackgroundTypePattern) SetType(v string) {
	o.Type = v
}

// GetDocument returns the Document field value
func (o *BackgroundTypePattern) GetDocument() Document {
	if o == nil {
		var ret Document
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetDocumentOk() (*Document, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *BackgroundTypePattern) SetDocument(v Document) {
	o.Document = v
}

// GetFill returns the Fill field value
func (o *BackgroundTypePattern) GetFill() BackgroundFill {
	if o == nil {
		var ret BackgroundFill
		return ret
	}

	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetFillOk() (*BackgroundFill, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fill, true
}

// SetFill sets field value
func (o *BackgroundTypePattern) SetFill(v BackgroundFill) {
	o.Fill = v
}

// GetIntensity returns the Intensity field value
func (o *BackgroundTypePattern) GetIntensity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Intensity
}

// GetIntensityOk returns a tuple with the Intensity field value
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetIntensityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intensity, true
}

// SetIntensity sets field value
func (o *BackgroundTypePattern) SetIntensity(v int32) {
	o.Intensity = v
}

// GetIsInverted returns the IsInverted field value if set, zero value otherwise.
func (o *BackgroundTypePattern) GetIsInverted() bool {
	if o == nil || IsNil(o.IsInverted) {
		var ret bool
		return ret
	}
	return *o.IsInverted
}

// GetIsInvertedOk returns a tuple with the IsInverted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetIsInvertedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInverted) {
		return nil, false
	}
	return o.IsInverted, true
}

// HasIsInverted returns a boolean if a field has been set.
func (o *BackgroundTypePattern) HasIsInverted() bool {
	if o != nil && !IsNil(o.IsInverted) {
		return true
	}

	return false
}

// SetIsInverted gets a reference to the given bool and assigns it to the IsInverted field.
func (o *BackgroundTypePattern) SetIsInverted(v bool) {
	o.IsInverted = &v
}


// GetIsMoving returns the IsMoving field value if set, zero value otherwise.
func (o *BackgroundTypePattern) GetIsMoving() bool {
	if o == nil || IsNil(o.IsMoving) {
		var ret bool
		return ret
	}
	return *o.IsMoving
}

// GetIsMovingOk returns a tuple with the IsMoving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackgroundTypePattern) GetIsMovingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMoving) {
		return nil, false
	}
	return o.IsMoving, true
}

// HasIsMoving returns a boolean if a field has been set.
func (o *BackgroundTypePattern) HasIsMoving() bool {
	if o != nil && !IsNil(o.IsMoving) {
		return true
	}

	return false
}

// SetIsMoving gets a reference to the given bool and assigns it to the IsMoving field.
func (o *BackgroundTypePattern) SetIsMoving(v bool) {
	o.IsMoving = &v
}


func (o BackgroundTypePattern) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundTypePattern) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["document"] = o.Document
	toSerialize["fill"] = o.Fill
	toSerialize["intensity"] = o.Intensity
	if !IsNil(o.IsInverted) {
		toSerialize["is_inverted"] = o.IsInverted
	}
	if !IsNil(o.IsMoving) {
		toSerialize["is_moving"] = o.IsMoving
	}
	return toSerialize, nil
}

func (o *BackgroundTypePattern) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"document",
		"fill",
		"intensity",
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

	varBackgroundTypePattern := _BackgroundTypePattern{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackgroundTypePattern)

	if err != nil {
		return err
	}

	*o = BackgroundTypePattern(varBackgroundTypePattern)

	return err
}

type NullableBackgroundTypePattern struct {
	value *BackgroundTypePattern
	isSet bool
}

func (v NullableBackgroundTypePattern) Get() *BackgroundTypePattern {
	return v.value
}

func (v *NullableBackgroundTypePattern) Set(val *BackgroundTypePattern) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTypePattern) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTypePattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTypePattern(val *BackgroundTypePattern) *NullableBackgroundTypePattern {
	return &NullableBackgroundTypePattern{value: val, isSet: true}
}

func (v NullableBackgroundTypePattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTypePattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


