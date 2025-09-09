# SuggestedPostDeclined

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedPostMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Comment** | Pointer to **string** | *Optional*. Comment with which the post was declined | [optional] 

## Methods

### NewSuggestedPostDeclined

`func NewSuggestedPostDeclined() *SuggestedPostDeclined`

NewSuggestedPostDeclined instantiates a new SuggestedPostDeclined object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostDeclinedWithDefaults

`func NewSuggestedPostDeclinedWithDefaults() *SuggestedPostDeclined`

NewSuggestedPostDeclinedWithDefaults instantiates a new SuggestedPostDeclined object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedPostMessage

`func (o *SuggestedPostDeclined) GetSuggestedPostMessage() Message`

GetSuggestedPostMessage returns the SuggestedPostMessage field if non-nil, zero value otherwise.

### GetSuggestedPostMessageOk

`func (o *SuggestedPostDeclined) GetSuggestedPostMessageOk() (*Message, bool)`

GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPostMessage

`func (o *SuggestedPostDeclined) SetSuggestedPostMessage(v Message)`

SetSuggestedPostMessage sets SuggestedPostMessage field to given value.

### HasSuggestedPostMessage

`func (o *SuggestedPostDeclined) HasSuggestedPostMessage() bool`

HasSuggestedPostMessage returns a boolean if a field has been set.

### GetComment

`func (o *SuggestedPostDeclined) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SuggestedPostDeclined) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SuggestedPostDeclined) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SuggestedPostDeclined) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


