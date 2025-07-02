# GetChatMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**LeaveChatRequestChatId**](LeaveChatRequestChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 

## Methods

### NewGetChatMemberRequest

`func NewGetChatMemberRequest(chatId LeaveChatRequestChatId, userId int32, ) *GetChatMemberRequest`

NewGetChatMemberRequest instantiates a new GetChatMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChatMemberRequestWithDefaults

`func NewGetChatMemberRequestWithDefaults() *GetChatMemberRequest`

NewGetChatMemberRequestWithDefaults instantiates a new GetChatMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *GetChatMemberRequest) GetChatId() LeaveChatRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *GetChatMemberRequest) GetChatIdOk() (*LeaveChatRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *GetChatMemberRequest) SetChatId(v LeaveChatRequestChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *GetChatMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetChatMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetChatMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


