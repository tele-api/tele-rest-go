# PollAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PollId** | **string** | Unique poll identifier | 
**VoterChat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**OptionIds** | **[]int32** | 0-based identifiers of chosen answer options. May be empty if the vote was retracted. | 

## Methods

### NewPollAnswer

`func NewPollAnswer(pollId string, optionIds []int32, ) *PollAnswer`

NewPollAnswer instantiates a new PollAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollAnswerWithDefaults

`func NewPollAnswerWithDefaults() *PollAnswer`

NewPollAnswerWithDefaults instantiates a new PollAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPollId

`func (o *PollAnswer) GetPollId() string`

GetPollId returns the PollId field if non-nil, zero value otherwise.

### GetPollIdOk

`func (o *PollAnswer) GetPollIdOk() (*string, bool)`

GetPollIdOk returns a tuple with the PollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollId

`func (o *PollAnswer) SetPollId(v string)`

SetPollId sets PollId field to given value.


### GetVoterChat

`func (o *PollAnswer) GetVoterChat() Chat`

GetVoterChat returns the VoterChat field if non-nil, zero value otherwise.

### GetVoterChatOk

`func (o *PollAnswer) GetVoterChatOk() (*Chat, bool)`

GetVoterChatOk returns a tuple with the VoterChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoterChat

`func (o *PollAnswer) SetVoterChat(v Chat)`

SetVoterChat sets VoterChat field to given value.

### HasVoterChat

`func (o *PollAnswer) HasVoterChat() bool`

HasVoterChat returns a boolean if a field has been set.

### GetUser

`func (o *PollAnswer) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PollAnswer) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PollAnswer) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *PollAnswer) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOptionIds

`func (o *PollAnswer) GetOptionIds() []int32`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *PollAnswer) GetOptionIdsOk() (*[]int32, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *PollAnswer) SetOptionIds(v []int32)`

SetOptionIds sets OptionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


