# Story

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**Id** | **int32** | Unique identifier for the story in the chat | 

## Methods

### NewStory

`func NewStory(chat Chat, id int32, ) *Story`

NewStory instantiates a new Story object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryWithDefaults

`func NewStoryWithDefaults() *Story`

NewStoryWithDefaults instantiates a new Story object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *Story) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *Story) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *Story) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetId

`func (o *Story) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Story) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Story) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


