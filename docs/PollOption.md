# PollOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Option text, 1-100 characters | 
**TextEntities** | Pointer to [**[]MessageEntity**](MessageEntity.md) | *Optional*. Special entities that appear in the option *text*. Currently, only custom emoji entities are allowed in poll option texts | [optional] 
**VoterCount** | **int32** | Number of users that voted for this option | 

## Methods

### NewPollOption

`func NewPollOption(text string, voterCount int32, ) *PollOption`

NewPollOption instantiates a new PollOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollOptionWithDefaults

`func NewPollOptionWithDefaults() *PollOption`

NewPollOptionWithDefaults instantiates a new PollOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *PollOption) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PollOption) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PollOption) SetText(v string)`

SetText sets Text field to given value.


### GetTextEntities

`func (o *PollOption) GetTextEntities() []MessageEntity`

GetTextEntities returns the TextEntities field if non-nil, zero value otherwise.

### GetTextEntitiesOk

`func (o *PollOption) GetTextEntitiesOk() (*[]MessageEntity, bool)`

GetTextEntitiesOk returns a tuple with the TextEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextEntities

`func (o *PollOption) SetTextEntities(v []MessageEntity)`

SetTextEntities sets TextEntities field to given value.

### HasTextEntities

`func (o *PollOption) HasTextEntities() bool`

HasTextEntities returns a boolean if a field has been set.

### GetVoterCount

`func (o *PollOption) GetVoterCount() int32`

GetVoterCount returns the VoterCount field if non-nil, zero value otherwise.

### GetVoterCountOk

`func (o *PollOption) GetVoterCountOk() (*int32, bool)`

GetVoterCountOk returns a tuple with the VoterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoterCount

`func (o *PollOption) SetVoterCount(v int32)`

SetVoterCount sets VoterCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


