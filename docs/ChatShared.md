# ChatShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int32** | Identifier of the request | 
**ChatId** | **int32** | Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means. | 
**Title** | Pointer to **string** | *Optional*. Title of the chat, if the title was requested by the bot. | [optional] 
**Username** | Pointer to **string** | *Optional*. Username of the chat, if the username was requested by the bot and available. | [optional] 
**Photo** | Pointer to [**[]PhotoSize**](PhotoSize.md) | *Optional*. Available sizes of the chat photo, if the photo was requested by the bot | [optional] 

## Methods

### NewChatShared

`func NewChatShared(requestId int32, chatId int32, ) *ChatShared`

NewChatShared instantiates a new ChatShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatSharedWithDefaults

`func NewChatSharedWithDefaults() *ChatShared`

NewChatSharedWithDefaults instantiates a new ChatShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ChatShared) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChatShared) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChatShared) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.


### GetChatId

`func (o *ChatShared) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ChatShared) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ChatShared) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetTitle

`func (o *ChatShared) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChatShared) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChatShared) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChatShared) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsername

`func (o *ChatShared) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChatShared) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChatShared) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ChatShared) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhoto

`func (o *ChatShared) GetPhoto() []PhotoSize`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ChatShared) GetPhotoOk() (*[]PhotoSize, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ChatShared) SetPhoto(v []PhotoSize)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ChatShared) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


