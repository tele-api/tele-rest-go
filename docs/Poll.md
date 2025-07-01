# Poll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique poll identifier | 
**Question** | **string** | Poll question, 1-300 characters | 
**QuestionEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the *question*. Currently, only custom emoji entities are allowed in poll questions | [optional] 
**Options** | [**[]PollOption**](PollOption.md) | List of poll options | 
**TotalVoterCount** | **int32** | Total number of users that voted in the poll | 
**IsClosed** | **bool** | *True*, if the poll is closed | 
**IsAnonymous** | **bool** | *True*, if the poll is anonymous | 
**Type** | **string** | Poll type, currently can be “regular” or “quiz” | 
**AllowsMultipleAnswers** | **bool** | *True*, if the poll allows multiple answers | 
**CorrectOptionId** | Pointer to **int32** | *Optional*. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot. | [optional] 
**Explanation** | Pointer to **string** | *Optional*. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters | [optional] 
**ExplanationEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities like usernames, URLs, bot commands, etc. that appear in the *explanation* | [optional] 
**OpenPeriod** | Pointer to **int32** | *Optional*. Amount of time in seconds the poll will be active after creation | [optional] 
**CloseDate** | Pointer to **int32** | *Optional*. Point in time (Unix timestamp) when the poll will be automatically closed | [optional] 

## Methods

### NewPoll

`func NewPoll(id string, question string, options []PollOption, totalVoterCount int32, isClosed bool, isAnonymous bool, type_ string, allowsMultipleAnswers bool, ) *Poll`

NewPoll instantiates a new Poll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollWithDefaults

`func NewPollWithDefaults() *Poll`

NewPollWithDefaults instantiates a new Poll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Poll) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Poll) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Poll) SetId(v string)`

SetId sets Id field to given value.


### GetQuestion

`func (o *Poll) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *Poll) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *Poll) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetQuestionEntities

`func (o *Poll) GetQuestionEntities() []MessageEntity`

GetQuestionEntities returns the QuestionEntities field if non-nil, zero value otherwise.

### GetQuestionEntitiesOk

`func (o *Poll) GetQuestionEntitiesOk() (*[]MessageEntity, bool)`

GetQuestionEntitiesOk returns a tuple with the QuestionEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionEntities

`func (o *Poll) SetQuestionEntities(v []MessageEntity)`

SetQuestionEntities sets QuestionEntities field to given value.

### HasQuestionEntities

`func (o *Poll) HasQuestionEntities() bool`

HasQuestionEntities returns a boolean if a field has been set.

### GetOptions

`func (o *Poll) GetOptions() []PollOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Poll) GetOptionsOk() (*[]PollOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Poll) SetOptions(v []PollOption)`

SetOptions sets Options field to given value.


### GetTotalVoterCount

`func (o *Poll) GetTotalVoterCount() int32`

GetTotalVoterCount returns the TotalVoterCount field if non-nil, zero value otherwise.

### GetTotalVoterCountOk

`func (o *Poll) GetTotalVoterCountOk() (*int32, bool)`

GetTotalVoterCountOk returns a tuple with the TotalVoterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVoterCount

`func (o *Poll) SetTotalVoterCount(v int32)`

SetTotalVoterCount sets TotalVoterCount field to given value.


### GetIsClosed

`func (o *Poll) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *Poll) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *Poll) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.


### GetIsAnonymous

`func (o *Poll) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *Poll) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *Poll) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.


### GetType

`func (o *Poll) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Poll) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Poll) SetType(v string)`

SetType sets Type field to given value.


### GetAllowsMultipleAnswers

`func (o *Poll) GetAllowsMultipleAnswers() bool`

GetAllowsMultipleAnswers returns the AllowsMultipleAnswers field if non-nil, zero value otherwise.

### GetAllowsMultipleAnswersOk

`func (o *Poll) GetAllowsMultipleAnswersOk() (*bool, bool)`

GetAllowsMultipleAnswersOk returns a tuple with the AllowsMultipleAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsMultipleAnswers

`func (o *Poll) SetAllowsMultipleAnswers(v bool)`

SetAllowsMultipleAnswers sets AllowsMultipleAnswers field to given value.


### GetCorrectOptionId

`func (o *Poll) GetCorrectOptionId() int32`

GetCorrectOptionId returns the CorrectOptionId field if non-nil, zero value otherwise.

### GetCorrectOptionIdOk

`func (o *Poll) GetCorrectOptionIdOk() (*int32, bool)`

GetCorrectOptionIdOk returns a tuple with the CorrectOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectOptionId

`func (o *Poll) SetCorrectOptionId(v int32)`

SetCorrectOptionId sets CorrectOptionId field to given value.

### HasCorrectOptionId

`func (o *Poll) HasCorrectOptionId() bool`

HasCorrectOptionId returns a boolean if a field has been set.

### GetExplanation

`func (o *Poll) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *Poll) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *Poll) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *Poll) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetExplanationEntities

`func (o *Poll) GetExplanationEntities() []MessageEntity`

GetExplanationEntities returns the ExplanationEntities field if non-nil, zero value otherwise.

### GetExplanationEntitiesOk

`func (o *Poll) GetExplanationEntitiesOk() (*[]MessageEntity, bool)`

GetExplanationEntitiesOk returns a tuple with the ExplanationEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationEntities

`func (o *Poll) SetExplanationEntities(v []MessageEntity)`

SetExplanationEntities sets ExplanationEntities field to given value.

### HasExplanationEntities

`func (o *Poll) HasExplanationEntities() bool`

HasExplanationEntities returns a boolean if a field has been set.

### GetOpenPeriod

`func (o *Poll) GetOpenPeriod() int32`

GetOpenPeriod returns the OpenPeriod field if non-nil, zero value otherwise.

### GetOpenPeriodOk

`func (o *Poll) GetOpenPeriodOk() (*int32, bool)`

GetOpenPeriodOk returns a tuple with the OpenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriod

`func (o *Poll) SetOpenPeriod(v int32)`

SetOpenPeriod sets OpenPeriod field to given value.

### HasOpenPeriod

`func (o *Poll) HasOpenPeriod() bool`

HasOpenPeriod returns a boolean if a field has been set.

### GetCloseDate

`func (o *Poll) GetCloseDate() int32`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *Poll) GetCloseDateOk() (*int32, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *Poll) SetCloseDate(v int32)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *Poll) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


