# MessageEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the entity. Currently, can be “mention” (&#x60;@username&#x60;), “hashtag” (&#x60;#hashtag&#x60; or &#x60;#hashtag@chatusername&#x60;), “cashtag” (&#x60;$USD&#x60; or &#x60;$USD@chatusername&#x60;), “bot\\_command” (&#x60;/start@jobs_bot&#x60;), “url” (&#x60;https://telegram.org&#x60;), “email” (&#x60;do-not-reply@telegram.org&#x60;), “phone\\_number” (&#x60;+1-212-555-0123&#x60;), “bold” (**bold text**), “italic” (*italic text*), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “expandable\\_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block), “text\\_link” (for clickable text URLs), “text\\_mention” (for users [without usernames](https://telegram.org/blog/edit#new-mentions)), “custom\\_emoji” (for inline custom emoji stickers) | 
**Offset** | **int32** | Offset in [UTF-16 code units](https://core.telegram.org/api/entities#entity-length) to the start of the entity | 
**Length** | **int32** | Length of the entity in [UTF-16 code units](https://core.telegram.org/api/entities#entity-length) | 
**Url** | Pointer to **string** | *Optional*. For “text\\_link” only, URL that will be opened after user taps on the text | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Language** | Pointer to **string** | *Optional*. For “pre” only, the programming language of the entity text | [optional] 
**CustomEmojiId** | Pointer to **string** | *Optional*. For “custom\\_emoji” only, unique identifier of the custom emoji. Use [getCustomEmojiStickers](https://core.telegram.org/bots/api/#getcustomemojistickers) to get full information about the sticker | [optional] 

## Methods

### NewMessageEntity

`func NewMessageEntity(type_ string, offset int32, length int32, ) *MessageEntity`

NewMessageEntity instantiates a new MessageEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageEntityWithDefaults

`func NewMessageEntityWithDefaults() *MessageEntity`

NewMessageEntityWithDefaults instantiates a new MessageEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageEntity) SetType(v string)`

SetType sets Type field to given value.


### GetOffset

`func (o *MessageEntity) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MessageEntity) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MessageEntity) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLength

`func (o *MessageEntity) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MessageEntity) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MessageEntity) SetLength(v int32)`

SetLength sets Length field to given value.


### GetUrl

`func (o *MessageEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUser

`func (o *MessageEntity) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MessageEntity) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MessageEntity) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *MessageEntity) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLanguage

`func (o *MessageEntity) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MessageEntity) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MessageEntity) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MessageEntity) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCustomEmojiId

`func (o *MessageEntity) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *MessageEntity) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *MessageEntity) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.

### HasCustomEmojiId

`func (o *MessageEntity) HasCustomEmojiId() bool`

HasCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


