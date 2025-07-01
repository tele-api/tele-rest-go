# Sticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Identifier for this file, which can be used to download or reuse the file | 
**FileUniqueId** | **string** | Unique identifier for this file, which is supposed to be the same over time and for different bots. Can&#39;t be used to download or reuse the file. | 
**Type** | **string** | Type of the sticker, currently one of “regular”, “mask”, “custom\\_emoji”. The type of the sticker is independent from its format, which is determined by the fields *is\\_animated* and *is\\_video*. | 
**Width** | **int32** | Sticker width | 
**Height** | **int32** | Sticker height | 
**IsAnimated** | **bool** | *True*, if the sticker is [animated](https://telegram.org/blog/animated-stickers) | 
**IsVideo** | **bool** | *True*, if the sticker is a [video sticker](https://telegram.org/blog/video-stickers-better-reactions) | 
**Thumbnail** | Pointer to [**PhotoSize**](PhotoSize.md) |  | [optional] 
**Emoji** | Pointer to **string** | *Optional*. Emoji associated with the sticker | [optional] 
**SetName** | Pointer to **string** | *Optional*. Name of the sticker set to which the sticker belongs | [optional] 
**PremiumAnimation** | Pointer to [**File**](File.md) |  | [optional] 
**MaskPosition** | Pointer to [**MaskPosition**](MaskPosition.md) |  | [optional] 
**CustomEmojiId** | Pointer to **string** | *Optional*. For custom emoji stickers, unique identifier of the custom emoji | [optional] 
**NeedsRepainting** | Pointer to **bool** | *Optional*. *True*, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places | [optional] [default to true]
**FileSize** | Pointer to **int32** | *Optional*. File size in bytes | [optional] 

## Methods

### NewSticker

`func NewSticker(fileId string, fileUniqueId string, type_ string, width int32, height int32, isAnimated bool, isVideo bool, ) *Sticker`

NewSticker instantiates a new Sticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickerWithDefaults

`func NewStickerWithDefaults() *Sticker`

NewStickerWithDefaults instantiates a new Sticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *Sticker) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Sticker) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Sticker) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFileUniqueId

`func (o *Sticker) GetFileUniqueId() string`

GetFileUniqueId returns the FileUniqueId field if non-nil, zero value otherwise.

### GetFileUniqueIdOk

`func (o *Sticker) GetFileUniqueIdOk() (*string, bool)`

GetFileUniqueIdOk returns a tuple with the FileUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqueId

`func (o *Sticker) SetFileUniqueId(v string)`

SetFileUniqueId sets FileUniqueId field to given value.


### GetType

`func (o *Sticker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Sticker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Sticker) SetType(v string)`

SetType sets Type field to given value.


### GetWidth

`func (o *Sticker) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Sticker) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Sticker) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Sticker) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Sticker) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Sticker) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetIsAnimated

`func (o *Sticker) GetIsAnimated() bool`

GetIsAnimated returns the IsAnimated field if non-nil, zero value otherwise.

### GetIsAnimatedOk

`func (o *Sticker) GetIsAnimatedOk() (*bool, bool)`

GetIsAnimatedOk returns a tuple with the IsAnimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnimated

`func (o *Sticker) SetIsAnimated(v bool)`

SetIsAnimated sets IsAnimated field to given value.


### GetIsVideo

`func (o *Sticker) GetIsVideo() bool`

GetIsVideo returns the IsVideo field if non-nil, zero value otherwise.

### GetIsVideoOk

`func (o *Sticker) GetIsVideoOk() (*bool, bool)`

GetIsVideoOk returns a tuple with the IsVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideo

`func (o *Sticker) SetIsVideo(v bool)`

SetIsVideo sets IsVideo field to given value.


### GetThumbnail

`func (o *Sticker) GetThumbnail() PhotoSize`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Sticker) GetThumbnailOk() (*PhotoSize, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Sticker) SetThumbnail(v PhotoSize)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Sticker) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetEmoji

`func (o *Sticker) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *Sticker) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *Sticker) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *Sticker) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetSetName

`func (o *Sticker) GetSetName() string`

GetSetName returns the SetName field if non-nil, zero value otherwise.

### GetSetNameOk

`func (o *Sticker) GetSetNameOk() (*string, bool)`

GetSetNameOk returns a tuple with the SetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetName

`func (o *Sticker) SetSetName(v string)`

SetSetName sets SetName field to given value.

### HasSetName

`func (o *Sticker) HasSetName() bool`

HasSetName returns a boolean if a field has been set.

### GetPremiumAnimation

`func (o *Sticker) GetPremiumAnimation() File`

GetPremiumAnimation returns the PremiumAnimation field if non-nil, zero value otherwise.

### GetPremiumAnimationOk

`func (o *Sticker) GetPremiumAnimationOk() (*File, bool)`

GetPremiumAnimationOk returns a tuple with the PremiumAnimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumAnimation

`func (o *Sticker) SetPremiumAnimation(v File)`

SetPremiumAnimation sets PremiumAnimation field to given value.

### HasPremiumAnimation

`func (o *Sticker) HasPremiumAnimation() bool`

HasPremiumAnimation returns a boolean if a field has been set.

### GetMaskPosition

`func (o *Sticker) GetMaskPosition() MaskPosition`

GetMaskPosition returns the MaskPosition field if non-nil, zero value otherwise.

### GetMaskPositionOk

`func (o *Sticker) GetMaskPositionOk() (*MaskPosition, bool)`

GetMaskPositionOk returns a tuple with the MaskPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPosition

`func (o *Sticker) SetMaskPosition(v MaskPosition)`

SetMaskPosition sets MaskPosition field to given value.

### HasMaskPosition

`func (o *Sticker) HasMaskPosition() bool`

HasMaskPosition returns a boolean if a field has been set.

### GetCustomEmojiId

`func (o *Sticker) GetCustomEmojiId() string`

GetCustomEmojiId returns the CustomEmojiId field if non-nil, zero value otherwise.

### GetCustomEmojiIdOk

`func (o *Sticker) GetCustomEmojiIdOk() (*string, bool)`

GetCustomEmojiIdOk returns a tuple with the CustomEmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmojiId

`func (o *Sticker) SetCustomEmojiId(v string)`

SetCustomEmojiId sets CustomEmojiId field to given value.

### HasCustomEmojiId

`func (o *Sticker) HasCustomEmojiId() bool`

HasCustomEmojiId returns a boolean if a field has been set.

### GetNeedsRepainting

`func (o *Sticker) GetNeedsRepainting() bool`

GetNeedsRepainting returns the NeedsRepainting field if non-nil, zero value otherwise.

### GetNeedsRepaintingOk

`func (o *Sticker) GetNeedsRepaintingOk() (*bool, bool)`

GetNeedsRepaintingOk returns a tuple with the NeedsRepainting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRepainting

`func (o *Sticker) SetNeedsRepainting(v bool)`

SetNeedsRepainting sets NeedsRepainting field to given value.

### HasNeedsRepainting

`func (o *Sticker) HasNeedsRepainting() bool`

HasNeedsRepainting returns a boolean if a field has been set.

### GetFileSize

`func (o *Sticker) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Sticker) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Sticker) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Sticker) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


