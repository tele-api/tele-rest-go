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

// checks if the PaidMediaPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaidMediaPreview{}

// PaidMediaPreview The paid media isn't available before the payment.
type PaidMediaPreview struct {
	// Type of the paid media, always “preview”
	Type string `json:"type"`
	// *Optional*. Media width as defined by the sender
	Width *int32 `json:"width,omitempty"`
	// *Optional*. Media height as defined by the sender
	Height *int32 `json:"height,omitempty"`
	// *Optional*. Duration of the media in seconds as defined by the sender
	Duration *int32 `json:"duration,omitempty"`
}

type _PaidMediaPreview PaidMediaPreview

// NewPaidMediaPreview instantiates a new PaidMediaPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaidMediaPreview(type_ string) *PaidMediaPreview {
	this := PaidMediaPreview{}
	this.Type = type_
	return &this
}

// NewPaidMediaPreviewWithDefaults instantiates a new PaidMediaPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaidMediaPreviewWithDefaults() *PaidMediaPreview {
	this := PaidMediaPreview{}
	var type_ string = "preview"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PaidMediaPreview) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaidMediaPreview) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaidMediaPreview) SetType(v string) {
	o.Type = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *PaidMediaPreview) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidMediaPreview) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PaidMediaPreview) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *PaidMediaPreview) SetWidth(v int32) {
	o.Width = &v
}


// GetHeight returns the Height field value if set, zero value otherwise.
func (o *PaidMediaPreview) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidMediaPreview) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *PaidMediaPreview) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *PaidMediaPreview) SetHeight(v int32) {
	o.Height = &v
}


// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PaidMediaPreview) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaidMediaPreview) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PaidMediaPreview) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *PaidMediaPreview) SetDuration(v int32) {
	o.Duration = &v
}


func (o PaidMediaPreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaidMediaPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

func (o *PaidMediaPreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varPaidMediaPreview := _PaidMediaPreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaidMediaPreview)

	if err != nil {
		return err
	}

	*o = PaidMediaPreview(varPaidMediaPreview)

	return err
}

type NullablePaidMediaPreview struct {
	value *PaidMediaPreview
	isSet bool
}

func (v NullablePaidMediaPreview) Get() *PaidMediaPreview {
	return v.value
}

func (v *NullablePaidMediaPreview) Set(val *PaidMediaPreview) {
	v.value = val
	v.isSet = true
}

func (v NullablePaidMediaPreview) IsSet() bool {
	return v.isSet
}

func (v *NullablePaidMediaPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaidMediaPreview(val *PaidMediaPreview) *NullablePaidMediaPreview {
	return &NullablePaidMediaPreview{value: val, isSet: true}
}

func (v NullablePaidMediaPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaidMediaPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


