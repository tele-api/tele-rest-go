# ChatBoostRemoved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | [**Chat**](Chat.md) |  | 
**BoostId** | **string** | Unique identifier of the boost | 
**RemoveDate** | **int32** | Point in time (Unix timestamp) when the boost was removed | 
**Source** | [**ChatBoostSource**](ChatBoostSource.md) |  | 

## Methods

### NewChatBoostRemoved

`func NewChatBoostRemoved(chat Chat, boostId string, removeDate int32, source ChatBoostSource, ) *ChatBoostRemoved`

NewChatBoostRemoved instantiates a new ChatBoostRemoved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatBoostRemovedWithDefaults

`func NewChatBoostRemovedWithDefaults() *ChatBoostRemoved`

NewChatBoostRemovedWithDefaults instantiates a new ChatBoostRemoved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatBoostRemoved) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatBoostRemoved) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatBoostRemoved) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetBoostId

`func (o *ChatBoostRemoved) GetBoostId() string`

GetBoostId returns the BoostId field if non-nil, zero value otherwise.

### GetBoostIdOk

`func (o *ChatBoostRemoved) GetBoostIdOk() (*string, bool)`

GetBoostIdOk returns a tuple with the BoostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostId

`func (o *ChatBoostRemoved) SetBoostId(v string)`

SetBoostId sets BoostId field to given value.


### GetRemoveDate

`func (o *ChatBoostRemoved) GetRemoveDate() int32`

GetRemoveDate returns the RemoveDate field if non-nil, zero value otherwise.

### GetRemoveDateOk

`func (o *ChatBoostRemoved) GetRemoveDateOk() (*int32, bool)`

GetRemoveDateOk returns a tuple with the RemoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDate

`func (o *ChatBoostRemoved) SetRemoveDate(v int32)`

SetRemoveDate sets RemoveDate field to given value.


### GetSource

`func (o *ChatBoostRemoved) GetSource() ChatBoostSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatBoostRemoved) GetSourceOk() (*ChatBoostSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatBoostRemoved) SetSource(v ChatBoostSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


