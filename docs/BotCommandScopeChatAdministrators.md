# BotCommandScopeChatAdministrators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Scope type, must be *chat\\_administrators* | [default to "chat_administrators"]
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 

## Methods

### NewBotCommandScopeChatAdministrators

`func NewBotCommandScopeChatAdministrators(type_ string, chatId BotCommandScopeChatChatId, ) *BotCommandScopeChatAdministrators`

NewBotCommandScopeChatAdministrators instantiates a new BotCommandScopeChatAdministrators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotCommandScopeChatAdministratorsWithDefaults

`func NewBotCommandScopeChatAdministratorsWithDefaults() *BotCommandScopeChatAdministrators`

NewBotCommandScopeChatAdministratorsWithDefaults instantiates a new BotCommandScopeChatAdministrators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BotCommandScopeChatAdministrators) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BotCommandScopeChatAdministrators) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BotCommandScopeChatAdministrators) SetType(v string)`

SetType sets Type field to given value.


### GetChatId

`func (o *BotCommandScopeChatAdministrators) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *BotCommandScopeChatAdministrators) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *BotCommandScopeChatAdministrators) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


