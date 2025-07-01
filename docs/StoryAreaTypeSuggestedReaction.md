# StoryAreaTypeSuggestedReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the area, always “suggested\\_reaction” | [default to "suggested_reaction"]
**ReactionType** | [**ReactionType**](ReactionType.md) |  | 
**IsDark** | Pointer to **bool** | *Optional*. Pass *True* if the reaction area has a dark background | [optional] 
**IsFlipped** | Pointer to **bool** | *Optional*. Pass *True* if reaction area corner is flipped | [optional] 

## Methods

### NewStoryAreaTypeSuggestedReaction

`func NewStoryAreaTypeSuggestedReaction(type_ string, reactionType ReactionType, ) *StoryAreaTypeSuggestedReaction`

NewStoryAreaTypeSuggestedReaction instantiates a new StoryAreaTypeSuggestedReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryAreaTypeSuggestedReactionWithDefaults

`func NewStoryAreaTypeSuggestedReactionWithDefaults() *StoryAreaTypeSuggestedReaction`

NewStoryAreaTypeSuggestedReactionWithDefaults instantiates a new StoryAreaTypeSuggestedReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoryAreaTypeSuggestedReaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoryAreaTypeSuggestedReaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoryAreaTypeSuggestedReaction) SetType(v string)`

SetType sets Type field to given value.


### GetReactionType

`func (o *StoryAreaTypeSuggestedReaction) GetReactionType() ReactionType`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *StoryAreaTypeSuggestedReaction) GetReactionTypeOk() (*ReactionType, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *StoryAreaTypeSuggestedReaction) SetReactionType(v ReactionType)`

SetReactionType sets ReactionType field to given value.


### GetIsDark

`func (o *StoryAreaTypeSuggestedReaction) GetIsDark() bool`

GetIsDark returns the IsDark field if non-nil, zero value otherwise.

### GetIsDarkOk

`func (o *StoryAreaTypeSuggestedReaction) GetIsDarkOk() (*bool, bool)`

GetIsDarkOk returns a tuple with the IsDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDark

`func (o *StoryAreaTypeSuggestedReaction) SetIsDark(v bool)`

SetIsDark sets IsDark field to given value.

### HasIsDark

`func (o *StoryAreaTypeSuggestedReaction) HasIsDark() bool`

HasIsDark returns a boolean if a field has been set.

### GetIsFlipped

`func (o *StoryAreaTypeSuggestedReaction) GetIsFlipped() bool`

GetIsFlipped returns the IsFlipped field if non-nil, zero value otherwise.

### GetIsFlippedOk

`func (o *StoryAreaTypeSuggestedReaction) GetIsFlippedOk() (*bool, bool)`

GetIsFlippedOk returns a tuple with the IsFlipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlipped

`func (o *StoryAreaTypeSuggestedReaction) SetIsFlipped(v bool)`

SetIsFlipped sets IsFlipped field to given value.

### HasIsFlipped

`func (o *StoryAreaTypeSuggestedReaction) HasIsFlipped() bool`

HasIsFlipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


