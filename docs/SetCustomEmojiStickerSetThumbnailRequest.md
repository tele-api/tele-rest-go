# SetCustomEmojiStickerSetThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Sticker set name | 
**CustomEmojiId** | Pointer to **string** | Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail. | [optional] 

## Methods

### NewSetCustomEmojiStickerSetThumbnailRequest

`func NewSetCustomEmojiStickerSetThumbnailRequest(name string, ) *SetCustomEmojiStickerSetThumbnailRequest`

NewSetCustomEmojiStickerSetThumbnailRequest instantiates a new SetCustomEmojiStickerSetThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCustomEmojiStickerSetThumbnailRequestWithDefaults

`func NewSetCustomEmojiStickerSetThumbnailRequestWithDefaults() *SetCustomEmojiStickerSetThumbnailRequest`

NewSetCustomEmojiStickerSetThumbnailRequestWithDefaults instantiates a new SetCustomEmojiStickerSetThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetCustomEmojiStickerSetThumbnailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetCustomEmojiStickerSetThumbnailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetCustomEmojiStickerSetThumbnailRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCustomEmojiId

`func (o *SetCustomEmojiStickerSetThumbnailRequest) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *SetCustomEmojiStickerSetThumbnailRequest) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *SetCustomEmojiStickerSetThumbnailRequest) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.

### HasCustomEmojiId

`func (o *SetCustomEmojiStickerSetThumbnailRequest) HasCustomEmojiId() bool`

HasCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


