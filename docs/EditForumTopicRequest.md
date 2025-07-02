# EditForumTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**BotCommandScopeChatChatId**](BotCommandScopeChatChatId.md) |  | 
**MessageThreadId** | **int32** | Unique identifier for the target message thread of the forum topic | 
**Name** | Pointer to **string** | New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept | [optional] 
**IconCustomEmojiId** | Pointer to **string** | New unique identifier of the custom emoji shown as the topic icon. Use [getForumTopicIconStickers](https://core.telegram.org/bots/api/#getforumtopiciconstickers) to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept | [optional] 

## Methods

### NewEditForumTopicRequest

`func NewEditForumTopicRequest(chatId BotCommandScopeChatChatId, messageThreadId int32, ) *EditForumTopicRequest`

NewEditForumTopicRequest instantiates a new EditForumTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditForumTopicRequestWithDefaults

`func NewEditForumTopicRequestWithDefaults() *EditForumTopicRequest`

NewEditForumTopicRequestWithDefaults instantiates a new EditForumTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *EditForumTopicRequest) GetChatId() BotCommandScopeChatChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditForumTopicRequest) GetChatIdOk() (*BotCommandScopeChatChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditForumTopicRequest) SetChatId(v BotCommandScopeChatChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *EditForumTopicRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *EditForumTopicRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *EditForumTopicRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.


### GetName

`func (o *EditForumTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditForumTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditForumTopicRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditForumTopicRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIconCustomEmojiId

`func (o *EditForumTopicRequest) GetIconCustomEmojiId() string`

GetIconCustomEmojiId returns the IconCustomEmojiId field if non-nil, zero value otherwise.

### GetIconCustomEmojiIdOk

`func (o *EditForumTopicRequest) GetIconCustomEmojiIdOk() (*string, bool)`

GetIconCustomEmojiIdOk returns a tuple with the IconCustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconCustomEmojiId

`func (o *EditForumTopicRequest) SetIconCustomEmojiId(v string)`

SetIconCustomEmojiId sets IconCustomEmojiId field to given value.

### HasIconCustomEmojiId

`func (o *EditForumTopicRequest) HasIconCustomEmojiId() bool`

HasIconCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


