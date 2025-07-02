# EditGeneralForumTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**Name** | **string** | New topic name, 1-128 characters | 

## Methods

### NewEditGeneralForumTopicRequest

`func NewEditGeneralForumTopicRequest(chatId BotCommandScopeChatChatId, name string, ) *EditGeneralForumTopicRequest`

NewEditGeneralForumTopicRequest instantiates a new EditGeneralForumTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditGeneralForumTopicRequestWithDefaults

`func NewEditGeneralForumTopicRequestWithDefaults() *EditGeneralForumTopicRequest`

NewEditGeneralForumTopicRequestWithDefaults instantiates a new EditGeneralForumTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *EditGeneralForumTopicRequest) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditGeneralForumTopicRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditGeneralForumTopicRequest) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *EditGeneralForumTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditGeneralForumTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditGeneralForumTopicRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


