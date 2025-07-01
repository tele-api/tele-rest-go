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

// checks if the Game type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Game{}

// Game This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`
	// *Optional*. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls [setGameScore](https://core.telegram.org/bots/api/#setgamescore), or manually edited using [editMessageText](https://core.telegram.org/bots/api/#editmessagetext). 0-4096 characters.
	Text *string `json:"text,omitempty"`
	// *Optional*. Special entities that appear in *text*, such as usernames, URLs, bot commands, etc.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation *Animation `json:"animation,omitempty"`
}

type _Game Game

// NewGame instantiates a new Game object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGame(title string, description string, photo []PhotoSize) *Game {
	this := Game{}
	this.Title = title
	this.Description = description
	this.Photo = photo
	return &this
}

// NewGameWithDefaults instantiates a new Game object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameWithDefaults() *Game {
	this := Game{}
	return &this
}

// GetTitle returns the Title field value
func (o *Game) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Game) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Game) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *Game) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Game) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Game) SetDescription(v string) {
	o.Description = v
}

// GetPhoto returns the Photo field value
func (o *Game) GetPhoto() []PhotoSize {
	if o == nil {
		var ret []PhotoSize
		return ret
	}

	return o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value
// and a boolean to check if the value has been set.
func (o *Game) GetPhotoOk() ([]PhotoSize, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photo, true
}

// SetPhoto sets field value
func (o *Game) SetPhoto(v []PhotoSize) {
	o.Photo = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Game) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Game) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Game) SetText(v string) {
	o.Text = &v
}


// GetTextEntities returns the TextEntities field value if set, zero value otherwise.
func (o *Game) GetTextEntities() []MessageEntity {
	if o == nil || IsNil(o.TextEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.TextEntities
}

// GetTextEntitiesOk returns a tuple with the TextEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetTextEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.TextEntities) {
		return nil, false
	}
	return o.TextEntities, true
}

// HasTextEntities returns a boolean if a field has been set.
func (o *Game) HasTextEntities() bool {
	if o != nil && !IsNil(o.TextEntities) {
		return true
	}

	return false
}

// SetTextEntities gets a reference to the given []MessageEntity and assigns it to the TextEntities field.
func (o *Game) SetTextEntities(v []MessageEntity) {
	o.TextEntities = v
}


// GetAnimation returns the Animation field value if set, zero value otherwise.
func (o *Game) GetAnimation() Animation {
	if o == nil || IsNil(o.Animation) {
		var ret Animation
		return ret
	}
	return *o.Animation
}

// GetAnimationOk returns a tuple with the Animation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetAnimationOk() (*Animation, bool) {
	if o == nil || IsNil(o.Animation) {
		return nil, false
	}
	return o.Animation, true
}

// HasAnimation returns a boolean if a field has been set.
func (o *Game) HasAnimation() bool {
	if o != nil && !IsNil(o.Animation) {
		return true
	}

	return false
}

// SetAnimation gets a reference to the given Animation and assigns it to the Animation field.
func (o *Game) SetAnimation(v Animation) {
	o.Animation = &v
}


func (o Game) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Game) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["photo"] = o.Photo
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.TextEntities) {
		toSerialize["text_entities"] = o.TextEntities
	}
	if !IsNil(o.Animation) {
		toSerialize["animation"] = o.Animation
	}
	return toSerialize, nil
}

func (o *Game) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"description",
		"photo",
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

	varGame := _Game{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGame)

	if err != nil {
		return err
	}

	*o = Game(varGame)

	return err
}

type NullableGame struct {
	value *Game
	isSet bool
}

func (v NullableGame) Get() *Game {
	return v.value
}

func (v *NullableGame) Set(val *Game) {
	v.value = val
	v.isSet = true
}

func (v NullableGame) IsSet() bool {
	return v.isSet
}

func (v *NullableGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGame(val *Game) *NullableGame {
	return &NullableGame{value: val, isSet: true}
}

func (v NullableGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


