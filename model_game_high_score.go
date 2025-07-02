/** 
 * Telegram Bot API - REST API Client
 * The Bot API is an HTTP-based interface created for developers keen on building bots for Telegram. To learn how to create and set up a bot, please consult our Introduction to Bots and Bot FAQ.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 9.0.0
 *    * - **Modified**: 2025-07-02T07:03:19.642213517Z[Etc/UTC]
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

// checks if the GameHighScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameHighScore{}

// GameHighScore This object represents one row of the high scores table for a game.
type GameHighScore struct {
	// Position in high score table for the game
	Position int32 `json:"position"`
	User User `json:"user"`
	// Score
	Score int32 `json:"score"`
}

type _GameHighScore GameHighScore

// NewGameHighScore instantiates a new GameHighScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameHighScore(position int32, user User, score int32) *GameHighScore {
	this := GameHighScore{}
	this.Position = position
	this.User = user
	this.Score = score
	return &this
}

// NewGameHighScoreWithDefaults instantiates a new GameHighScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameHighScoreWithDefaults() *GameHighScore {
	this := GameHighScore{}
	return &this
}

// GetPosition returns the Position field value
func (o *GameHighScore) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *GameHighScore) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *GameHighScore) SetPosition(v int32) {
	o.Position = v
}

// GetUser returns the User field value
func (o *GameHighScore) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GameHighScore) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GameHighScore) SetUser(v User) {
	o.User = v
}

// GetScore returns the Score field value
func (o *GameHighScore) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *GameHighScore) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *GameHighScore) SetScore(v int32) {
	o.Score = v
}

func (o GameHighScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameHighScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["position"] = o.Position
	toSerialize["user"] = o.User
	toSerialize["score"] = o.Score
	return toSerialize, nil
}

func (o *GameHighScore) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"position",
		"user",
		"score",
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

	varGameHighScore := _GameHighScore{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameHighScore)

	if err != nil {
		return err
	}

	*o = GameHighScore(varGameHighScore)

	return err
}

type NullableGameHighScore struct {
	value *GameHighScore
	isSet bool
}

func (v NullableGameHighScore) Get() *GameHighScore {
	return v.value
}

func (v *NullableGameHighScore) Set(val *GameHighScore) {
	v.value = val
	v.isSet = true
}

func (v NullableGameHighScore) IsSet() bool {
	return v.isSet
}

func (v *NullableGameHighScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameHighScore(val *GameHighScore) *NullableGameHighScore {
	return &NullableGameHighScore{value: val, isSet: true}
}

func (v NullableGameHighScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameHighScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


