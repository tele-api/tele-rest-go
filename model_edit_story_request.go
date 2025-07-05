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

// checks if the EditStoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditStoryRequest{}

// EditStoryRequest Request parameters for editStory
type EditStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier of the story to edit
	StoryId int32 `json:"story_id"`
	Content InputStoryContent `json:"content"`
	// Caption of the story, 0-2048 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the story caption. See [formatting options](https://core.telegram.org/bots/api/#formatting-options) for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of *parse\\_mode*
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// A JSON-serialized list of clickable areas to be shown on the story
	Areas []StoryArea `json:"areas,omitempty"`
}

type _EditStoryRequest EditStoryRequest

// NewEditStoryRequest instantiates a new EditStoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditStoryRequest(businessConnectionId string, storyId int32, content InputStoryContent) *EditStoryRequest {
	this := EditStoryRequest{}
	this.BusinessConnectionId = businessConnectionId
	this.StoryId = storyId
	this.Content = content
	return &this
}

// NewEditStoryRequestWithDefaults instantiates a new EditStoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditStoryRequestWithDefaults() *EditStoryRequest {
	this := EditStoryRequest{}
	return &this
}

// GetBusinessConnectionId returns the BusinessConnectionId field value
func (o *EditStoryRequest) GetBusinessConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessConnectionId
}

// GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field value
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetBusinessConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessConnectionId, true
}

// SetBusinessConnectionId sets field value
func (o *EditStoryRequest) SetBusinessConnectionId(v string) {
	o.BusinessConnectionId = v
}

// GetStoryId returns the StoryId field value
func (o *EditStoryRequest) GetStoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StoryId
}

// GetStoryIdOk returns a tuple with the StoryId field value
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetStoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoryId, true
}

// SetStoryId sets field value
func (o *EditStoryRequest) SetStoryId(v int32) {
	o.StoryId = v
}

// GetContent returns the Content field value
func (o *EditStoryRequest) GetContent() InputStoryContent {
	if o == nil {
		var ret InputStoryContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetContentOk() (*InputStoryContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *EditStoryRequest) SetContent(v InputStoryContent) {
	o.Content = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *EditStoryRequest) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *EditStoryRequest) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *EditStoryRequest) SetCaption(v string) {
	o.Caption = &v
}


// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *EditStoryRequest) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *EditStoryRequest) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *EditStoryRequest) SetParseMode(v string) {
	o.ParseMode = &v
}


// GetCaptionEntities returns the CaptionEntities field value if set, zero value otherwise.
func (o *EditStoryRequest) GetCaptionEntities() []MessageEntity {
	if o == nil || IsNil(o.CaptionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.CaptionEntities
}

// GetCaptionEntitiesOk returns a tuple with the CaptionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetCaptionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.CaptionEntities) {
		return nil, false
	}
	return o.CaptionEntities, true
}

// HasCaptionEntities returns a boolean if a field has been set.
func (o *EditStoryRequest) HasCaptionEntities() bool {
	if o != nil && !IsNil(o.CaptionEntities) {
		return true
	}

	return false
}

// SetCaptionEntities gets a reference to the given []MessageEntity and assigns it to the CaptionEntities field.
func (o *EditStoryRequest) SetCaptionEntities(v []MessageEntity) {
	o.CaptionEntities = v
}


// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *EditStoryRequest) GetAreas() []StoryArea {
	if o == nil || IsNil(o.Areas) {
		var ret []StoryArea
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditStoryRequest) GetAreasOk() ([]StoryArea, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *EditStoryRequest) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []StoryArea and assigns it to the Areas field.
func (o *EditStoryRequest) SetAreas(v []StoryArea) {
	o.Areas = v
}


func (o EditStoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditStoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_connection_id"] = o.BusinessConnectionId
	toSerialize["story_id"] = o.StoryId
	toSerialize["content"] = o.Content
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parse_mode"] = o.ParseMode
	}
	if !IsNil(o.CaptionEntities) {
		toSerialize["caption_entities"] = o.CaptionEntities
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

func (o *EditStoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_connection_id",
		"story_id",
		"content",
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

	varEditStoryRequest := _EditStoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEditStoryRequest)

	if err != nil {
		return err
	}

	*o = EditStoryRequest(varEditStoryRequest)

	return err
}

type NullableEditStoryRequest struct {
	value *EditStoryRequest
	isSet bool
}

func (v NullableEditStoryRequest) Get() *EditStoryRequest {
	return v.value
}

func (v *NullableEditStoryRequest) Set(val *EditStoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditStoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditStoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditStoryRequest(val *EditStoryRequest) *NullableEditStoryRequest {
	return &NullableEditStoryRequest{value: val, isSet: true}
}

func (v NullableEditStoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditStoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


