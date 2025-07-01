# BusinessMessagesDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**Chat** | [**Chat**](Chat.md) |  | 
**MessageIds** | **[]int32** | The list of identifiers of deleted messages in the chat of the business account | 

## Methods

### NewBusinessMessagesDeleted

`func NewBusinessMessagesDeleted(businessConnectionId string, chat Chat, messageIds []int32, ) *BusinessMessagesDeleted`

NewBusinessMessagesDeleted instantiates a new BusinessMessagesDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessMessagesDeletedWithDefaults

`func NewBusinessMessagesDeletedWithDefaults() *BusinessMessagesDeleted`

NewBusinessMessagesDeletedWithDefaults instantiates a new BusinessMessagesDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *BusinessMessagesDeleted) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *BusinessMessagesDeleted) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *BusinessMessagesDeleted) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetChat

`func (o *BusinessMessagesDeleted) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *BusinessMessagesDeleted) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *BusinessMessagesDeleted) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetMessageIds

`func (o *BusinessMessagesDeleted) GetMessageIds() []int32`

GetMessageIds returns the MessageIds field if non-nil, zero value otherwise.

### GetMessageIdsOk

`func (o *BusinessMessagesDeleted) GetMessageIdsOk() (*[]int32, bool)`

GetMessageIdsOk returns a tuple with the MessageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIds

`func (o *BusinessMessagesDeleted) SetMessageIds(v []int32)`

SetMessageIds sets MessageIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


