# SetChatStickerSetPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**RestrictChatMemberPostRequestChatId**](RestrictChatMemberPostRequestChatId.md) |  | 
**StickerSetName** | **string** | Name of the sticker set to be set as the group sticker set | 

## Methods

### NewSetChatStickerSetPostRequest

`func NewSetChatStickerSetPostRequest(chatId RestrictChatMemberPostRequestChatId, stickerSetName string, ) *SetChatStickerSetPostRequest`

NewSetChatStickerSetPostRequest instantiates a new SetChatStickerSetPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetChatStickerSetPostRequestWithDefaults

`func NewSetChatStickerSetPostRequestWithDefaults() *SetChatStickerSetPostRequest`

NewSetChatStickerSetPostRequestWithDefaults instantiates a new SetChatStickerSetPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *SetChatStickerSetPostRequest) GetChatId() RestrictChatMemberPostRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SetChatStickerSetPostRequest) GetChatIdOk() (*RestrictChatMemberPostRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SetChatStickerSetPostRequest) SetChatId(v RestrictChatMemberPostRequestChatId)`

SetChatId sets ChatId field to given value.


### GetStickerSetName

`func (o *SetChatStickerSetPostRequest) GetStickerSetName() string`

GetStickerSetName returns the StickerSetName field if non-nil, zero value otherwise.

### GetStickerSetNameOk

`func (o *SetChatStickerSetPostRequest) GetStickerSetNameOk() (*string, bool)`

GetStickerSetNameOk returns a tuple with the StickerSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerSetName

`func (o *SetChatStickerSetPostRequest) SetStickerSetName(v string)`

SetStickerSetName sets StickerSetName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


