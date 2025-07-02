# DeclineChatJoinRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 

## Methods

### NewDeclineChatJoinRequestRequest

`func NewDeclineChatJoinRequestRequest(chatId SendMessageRequestChatId, userId int32, ) *DeclineChatJoinRequestRequest`

NewDeclineChatJoinRequestRequest instantiates a new DeclineChatJoinRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclineChatJoinRequestRequestWithDefaults

`func NewDeclineChatJoinRequestRequestWithDefaults() *DeclineChatJoinRequestRequest`

NewDeclineChatJoinRequestRequestWithDefaults instantiates a new DeclineChatJoinRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *DeclineChatJoinRequestRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *DeclineChatJoinRequestRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *DeclineChatJoinRequestRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *DeclineChatJoinRequestRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeclineChatJoinRequestRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeclineChatJoinRequestRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


