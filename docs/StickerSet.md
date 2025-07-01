# StickerSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Sticker set name | 
**Title** | **string** | Sticker set title | 
**StickerType** | **string** | Type of stickers in the set, currently one of “regular”, “mask”, “custom\\_emoji” | 
**Stickers** | [**[]Sticker**](Sticker.md) | List of all set stickers | 
**Thumbnail** | Pointer to [**PhotoSize**](PhotoSize.md) |  | [optional] 

## Methods

### NewStickerSet

`func NewStickerSet(name string, title string, stickerType string, stickers []Sticker, ) *StickerSet`

NewStickerSet instantiates a new StickerSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickerSetWithDefaults

`func NewStickerSetWithDefaults() *StickerSet`

NewStickerSetWithDefaults instantiates a new StickerSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StickerSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StickerSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StickerSet) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *StickerSet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StickerSet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StickerSet) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStickerType

`func (o *StickerSet) GetStickerType() string`

GetStickerType returns the StickerType field if non-nil, zero value otherwise.

### GetStickerTypeOk

`func (o *StickerSet) GetStickerTypeOk() (*string, bool)`

GetStickerTypeOk returns a tuple with the StickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerType

`func (o *StickerSet) SetStickerType(v string)`

SetStickerType sets StickerType field to given value.


### GetStickers

`func (o *StickerSet) GetStickers() []Sticker`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *StickerSet) GetStickersOk() (*[]Sticker, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *StickerSet) SetStickers(v []Sticker)`

SetStickers sets Stickers field to given value.


### GetThumbnail

`func (o *StickerSet) GetThumbnail() PhotoSize`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *StickerSet) GetThumbnailOk() (*PhotoSize, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *StickerSet) SetThumbnail(v PhotoSize)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *StickerSet) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


