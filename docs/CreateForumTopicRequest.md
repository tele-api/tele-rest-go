# CreateForumTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**Name** | **string** | Topic name, 1-128 characters | 
**IconColor** | Pointer to **int32** | Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F) | [optional] 
**IconCustomEmojiId** | Pointer to **string** | Unique identifier of the custom emoji shown as the topic icon. Use [getForumTopicIconStickers](https://core.telegram.org/bots/api/#getforumtopiciconstickers) to get all allowed custom emoji identifiers. | [optional] 

## Methods

### NewCreateForumTopicRequest

`func NewCreateForumTopicRequest(chatId BotCommandScopeChatChatId, name string, ) *CreateForumTopicRequest`

NewCreateForumTopicRequest instantiates a new CreateForumTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateForumTopicRequestWithDefaults

`func NewCreateForumTopicRequestWithDefaults() *CreateForumTopicRequest`

NewCreateForumTopicRequestWithDefaults instantiates a new CreateForumTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *CreateForumTopicRequest) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CreateForumTopicRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CreateForumTopicRequest) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetName

`func (o *CreateForumTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateForumTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateForumTopicRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIconColor

`func (o *CreateForumTopicRequest) GetIconColor() int32`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *CreateForumTopicRequest) GetIconColorOk() (*int32, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *CreateForumTopicRequest) SetIconColor(v int32)`

SetIconColor sets IconColor field to given value.

### HasIconColor

`func (o *CreateForumTopicRequest) HasIconColor() bool`

HasIconColor returns a boolean if a field has been set.

### GetIconCustomEmojiId

`func (o *CreateForumTopicRequest) GetIconCustomEmojiId() string`

GetIconCustomEmojiId returns the IconCustomEmojiId field if non-nil, zero value otherwise.

### GetIconCustomEmojiIdOk

`func (o *CreateForumTopicRequest) GetIconCustomEmojiIdOk() (*string, bool)`

GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCustomEmojiId

`func (o *CreateForumTopicRequest) SetIconCustomEmojiId(v string)`

SetIconCustomEmojiId sets IconCustomEmojiId field to given value.

### HasIconCustomEmojiId

`func (o *CreateForumTopicRequest) HasIconCustomEmojiId() bool`

HasIconCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


