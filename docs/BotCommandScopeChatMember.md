# BotCommandScopeChatMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Scope type, must be *chat\\_member* | [default to "chat_member"]
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**UserId** | **int32** | Unique identifier of the target user | 

## Methods

### NewBotCommandScopeChatMember

`func NewBotCommandScopeChatMember(type_ string, chatId BotCommandScopeChatChatId, userId int32, ) *BotCommandScopeChatMember`

NewBotCommandScopeChatMember instantiates a new BotCommandScopeChatMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotCommandScopeChatMemberWithDefaults

`func NewBotCommandScopeChatMemberWithDefaults() *BotCommandScopeChatMember`

NewBotCommandScopeChatMemberWithDefaults instantiates a new BotCommandScopeChatMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BotCommandScopeChatMember) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BotCommandScopeChatMember) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BotCommandScopeChatMember) SetType(v string)`

SetType sets Type field to given value.


### GetChatId

`func (o *BotCommandScopeChatMember) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *BotCommandScopeChatMember) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *BotCommandScopeChatMember) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetUserId

`func (o *BotCommandScopeChatMember) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BotCommandScopeChatMember) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BotCommandScopeChatMember) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


