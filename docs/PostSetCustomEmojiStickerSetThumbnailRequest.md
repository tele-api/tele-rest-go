# PostSetCustomEmojiStickerSetThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Sticker set name | 
**CustomEmojiId** | Pointer to **string** | Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail. | [optional] 

## Methods

### NewPostSetCustomEmojiStickerSetThumbnailRequest

`func NewPostSetCustomEmojiStickerSetThumbnailRequest(name string, ) *PostSetCustomEmojiStickerSetThumbnailRequest`

NewPostSetCustomEmojiStickerSetThumbnailRequest instantiates a new PostSetCustomEmojiStickerSetThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetCustomEmojiStickerSetThumbnailRequestWithDefaults

`func NewPostSetCustomEmojiStickerSetThumbnailRequestWithDefaults() *PostSetCustomEmojiStickerSetThumbnailRequest`

NewPostSetCustomEmojiStickerSetThumbnailRequestWithDefaults instantiates a new PostSetCustomEmojiStickerSetThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCustomEmojiId

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.

### HasCustomEmojiId

`func (o *PostSetCustomEmojiStickerSetThumbnailRequest) HasCustomEmojiId() bool`

HasCustomEmojiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


