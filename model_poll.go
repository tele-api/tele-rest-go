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

// checks if the Poll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Poll{}

// Poll This object contains information about a poll.
type Poll struct {
	// Unique poll identifier
	Id string `json:"id"`
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// *Optional*. Special entities that appear in the *question*. Currently, only custom emoji entities are allowed in poll questions
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`
	// List of poll options
	Options []PollOption `json:"options"`
	// Total number of users that voted in the poll
	TotalVoterCount int32 `json:"total_voter_count"`
	// *True*, if the poll is closed
	IsClosed bool `json:"is_closed"`
	// *True*, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	// Poll type, currently can be “regular” or “quiz”
	Type string `json:"type"`
	// *True*, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// *Optional*. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionId *int32 `json:"correct_option_id,omitempty"`
	// *Optional*. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation *string `json:"explanation,omitempty"`
	// *Optional*. Special entities like usernames, URLs, bot commands, etc. that appear in the *explanation*
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	// *Optional*. Amount of time in seconds the poll will be active after creation
	OpenPeriod *int32 `json:"open_period,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate *int32 `json:"close_date,omitempty"`
}

type _Poll Poll

// NewPoll instantiates a new Poll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoll(id string, question string, options []PollOption, totalVoterCount int32, isClosed bool, isAnonymous bool, type_ string, allowsMultipleAnswers bool) *Poll {
	this := Poll{}
	this.Id = id
	this.Question = question
	this.Options = options
	this.TotalVoterCount = totalVoterCount
	this.IsClosed = isClosed
	this.IsAnonymous = isAnonymous
	this.Type = type_
	this.AllowsMultipleAnswers = allowsMultipleAnswers
	return &this
}

// NewPollWithDefaults instantiates a new Poll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollWithDefaults() *Poll {
	this := Poll{}
	return &this
}

// GetId returns the Id field value
func (o *Poll) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Poll) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Poll) SetId(v string) {
	o.Id = v
}

// GetQuestion returns the Question field value
func (o *Poll) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *Poll) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *Poll) SetQuestion(v string) {
	o.Question = v
}

// GetQuestionEntities returns the QuestionEntities field value if set, zero value otherwise.
func (o *Poll) GetQuestionEntities() []MessageEntity {
	if o == nil || IsNil(o.QuestionEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.QuestionEntities
}

// GetQuestionEntitiesOk returns a tuple with the QuestionEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetQuestionEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.QuestionEntities) {
		return nil, false
	}
	return o.QuestionEntities, true
}

// HasQuestionEntities returns a boolean if a field has been set.
func (o *Poll) HasQuestionEntities() bool {
	if o != nil && !IsNil(o.QuestionEntities) {
		return true
	}

	return false
}

// SetQuestionEntities gets a reference to the given []MessageEntity and assigns it to the QuestionEntities field.
func (o *Poll) SetQuestionEntities(v []MessageEntity) {
	o.QuestionEntities = v
}


// GetOptions returns the Options field value
func (o *Poll) GetOptions() []PollOption {
	if o == nil {
		var ret []PollOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Poll) GetOptionsOk() ([]PollOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Poll) SetOptions(v []PollOption) {
	o.Options = v
}

// GetTotalVoterCount returns the TotalVoterCount field value
func (o *Poll) GetTotalVoterCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalVoterCount
}

// GetTotalVoterCountOk returns a tuple with the TotalVoterCount field value
// and a boolean to check if the value has been set.
func (o *Poll) GetTotalVoterCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalVoterCount, true
}

// SetTotalVoterCount sets field value
func (o *Poll) SetTotalVoterCount(v int32) {
	o.TotalVoterCount = v
}

// GetIsClosed returns the IsClosed field value
func (o *Poll) GetIsClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *Poll) GetIsClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *Poll) SetIsClosed(v bool) {
	o.IsClosed = v
}

// GetIsAnonymous returns the IsAnonymous field value
func (o *Poll) GetIsAnonymous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value
// and a boolean to check if the value has been set.
func (o *Poll) GetIsAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAnonymous, true
}

// SetIsAnonymous sets field value
func (o *Poll) SetIsAnonymous(v bool) {
	o.IsAnonymous = v
}

// GetType returns the Type field value
func (o *Poll) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Poll) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Poll) SetType(v string) {
	o.Type = v
}

// GetAllowsMultipleAnswers returns the AllowsMultipleAnswers field value
func (o *Poll) GetAllowsMultipleAnswers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowsMultipleAnswers
}

// GetAllowsMultipleAnswersOk returns a tuple with the AllowsMultipleAnswers field value
// and a boolean to check if the value has been set.
func (o *Poll) GetAllowsMultipleAnswersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowsMultipleAnswers, true
}

// SetAllowsMultipleAnswers sets field value
func (o *Poll) SetAllowsMultipleAnswers(v bool) {
	o.AllowsMultipleAnswers = v
}

// GetCorrectOptionId returns the CorrectOptionId field value if set, zero value otherwise.
func (o *Poll) GetCorrectOptionId() int32 {
	if o == nil || IsNil(o.CorrectOptionId) {
		var ret int32
		return ret
	}
	return *o.CorrectOptionId
}

// GetCorrectOptionIdOk returns a tuple with the CorrectOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetCorrectOptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CorrectOptionId) {
		return nil, false
	}
	return o.CorrectOptionId, true
}

// HasCorrectOptionId returns a boolean if a field has been set.
func (o *Poll) HasCorrectOptionId() bool {
	if o != nil && !IsNil(o.CorrectOptionId) {
		return true
	}

	return false
}

// SetCorrectOptionId gets a reference to the given int32 and assigns it to the CorrectOptionId field.
func (o *Poll) SetCorrectOptionId(v int32) {
	o.CorrectOptionId = &v
}


// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *Poll) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *Poll) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *Poll) SetExplanation(v string) {
	o.Explanation = &v
}


// GetExplanationEntities returns the ExplanationEntities field value if set, zero value otherwise.
func (o *Poll) GetExplanationEntities() []MessageEntity {
	if o == nil || IsNil(o.ExplanationEntities) {
		var ret []MessageEntity
		return ret
	}
	return o.ExplanationEntities
}

// GetExplanationEntitiesOk returns a tuple with the ExplanationEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetExplanationEntitiesOk() ([]MessageEntity, bool) {
	if o == nil || IsNil(o.ExplanationEntities) {
		return nil, false
	}
	return o.ExplanationEntities, true
}

// HasExplanationEntities returns a boolean if a field has been set.
func (o *Poll) HasExplanationEntities() bool {
	if o != nil && !IsNil(o.ExplanationEntities) {
		return true
	}

	return false
}

// SetExplanationEntities gets a reference to the given []MessageEntity and assigns it to the ExplanationEntities field.
func (o *Poll) SetExplanationEntities(v []MessageEntity) {
	o.ExplanationEntities = v
}


// GetOpenPeriod returns the OpenPeriod field value if set, zero value otherwise.
func (o *Poll) GetOpenPeriod() int32 {
	if o == nil || IsNil(o.OpenPeriod) {
		var ret int32
		return ret
	}
	return *o.OpenPeriod
}

// GetOpenPeriodOk returns a tuple with the OpenPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetOpenPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenPeriod) {
		return nil, false
	}
	return o.OpenPeriod, true
}

// HasOpenPeriod returns a boolean if a field has been set.
func (o *Poll) HasOpenPeriod() bool {
	if o != nil && !IsNil(o.OpenPeriod) {
		return true
	}

	return false
}

// SetOpenPeriod gets a reference to the given int32 and assigns it to the OpenPeriod field.
func (o *Poll) SetOpenPeriod(v int32) {
	o.OpenPeriod = &v
}


// GetCloseDate returns the CloseDate field value if set, zero value otherwise.
func (o *Poll) GetCloseDate() int32 {
	if o == nil || IsNil(o.CloseDate) {
		var ret int32
		return ret
	}
	return *o.CloseDate
}

// GetCloseDateOk returns a tuple with the CloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Poll) GetCloseDateOk() (*int32, bool) {
	if o == nil || IsNil(o.CloseDate) {
		return nil, false
	}
	return o.CloseDate, true
}

// HasCloseDate returns a boolean if a field has been set.
func (o *Poll) HasCloseDate() bool {
	if o != nil && !IsNil(o.CloseDate) {
		return true
	}

	return false
}

// SetCloseDate gets a reference to the given int32 and assigns it to the CloseDate field.
func (o *Poll) SetCloseDate(v int32) {
	o.CloseDate = &v
}


func (o Poll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Poll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["question"] = o.Question
	if !IsNil(o.QuestionEntities) {
		toSerialize["question_entities"] = o.QuestionEntities
	}
	toSerialize["options"] = o.Options
	toSerialize["total_voter_count"] = o.TotalVoterCount
	toSerialize["is_closed"] = o.IsClosed
	toSerialize["is_anonymous"] = o.IsAnonymous
	toSerialize["type"] = o.Type
	toSerialize["allows_multiple_answers"] = o.AllowsMultipleAnswers
	if !IsNil(o.CorrectOptionId) {
		toSerialize["correct_option_id"] = o.CorrectOptionId
	}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	if !IsNil(o.ExplanationEntities) {
		toSerialize["explanation_entities"] = o.ExplanationEntities
	}
	if !IsNil(o.OpenPeriod) {
		toSerialize["open_period"] = o.OpenPeriod
	}
	if !IsNil(o.CloseDate) {
		toSerialize["close_date"] = o.CloseDate
	}
	return toSerialize, nil
}

func (o *Poll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"question",
		"options",
		"total_voter_count",
		"is_closed",
		"is_anonymous",
		"type",
		"allows_multiple_answers",
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

	varPoll := _Poll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoll)

	if err != nil {
		return err
	}

	*o = Poll(varPoll)

	return err
}

type NullablePoll struct {
	value *Poll
	isSet bool
}

func (v NullablePoll) Get() *Poll {
	return v.value
}

func (v *NullablePoll) Set(val *Poll) {
	v.value = val
	v.isSet = true
}

func (v NullablePoll) IsSet() bool {
	return v.isSet
}

func (v *NullablePoll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoll(val *Poll) *NullablePoll {
	return &NullablePoll{value: val, isSet: true}
}

func (v NullablePoll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


