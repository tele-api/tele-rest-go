# ApproveChatJoinRequestPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessagePostRequestChatId**](SendMessagePostRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 

## Methods

### NewApproveChatJoinRequestPostRequest

`func NewApproveChatJoinRequestPostRequest(chatId SendMessagePostRequestChatId, userId int32, ) *ApproveChatJoinRequestPostRequest`

NewApproveChatJoinRequestPostRequest instantiates a new ApproveChatJoinRequestPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveChatJoinRequestPostRequestWithDefaults

`func NewApproveChatJoinRequestPostRequestWithDefaults() *ApproveChatJoinRequestPostRequest`

NewApproveChatJoinRequestPostRequestWithDefaults instantiates a new ApproveChatJoinRequestPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *ApproveChatJoinRequestPostRequest) GetChatId() SendMessagePostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ApproveChatJoinRequestPostRequest) GetChatIdOk() (*SendMessagePostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ApproveChatJoinRequestPostRequest) SetChatId(v SendMessagePostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *ApproveChatJoinRequestPostRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApproveChatJoinRequestPostRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApproveChatJoinRequestPostRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


