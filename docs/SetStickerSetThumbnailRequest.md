# SetStickerSetThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Sticker set name | 
**UserId** | **int32** | User identifier of the sticker set owner | 
**Thumbnail** | Pointer to **NullableString** |  | [optional] 
**Format** | **string** | Format of the thumbnail, must be one of “static” for a **.WEBP** or **.PNG** image, “animated” for a **.TGS** animation, or “video” for a **.WEBM** video | 

## Methods

### NewSetStickerSetThumbnailRequest

`func NewSetStickerSetThumbnailRequest(name string, userId int32, format string, ) *SetStickerSetThumbnailRequest`

NewSetStickerSetThumbnailRequest instantiates a new SetStickerSetThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerSetThumbnailRequestWithDefaults

`func NewSetStickerSetThumbnailRequestWithDefaults() *SetStickerSetThumbnailRequest`

NewSetStickerSetThumbnailRequestWithDefaults instantiates a new SetStickerSetThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetStickerSetThumbnailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetStickerSetThumbnailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetStickerSetThumbnailRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *SetStickerSetThumbnailRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SetStickerSetThumbnailRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SetStickerSetThumbnailRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetThumbnail

`func (o *SetStickerSetThumbnailRequest) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *SetStickerSetThumbnailRequest) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *SetStickerSetThumbnailRequest) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *SetStickerSetThumbnailRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *SetStickerSetThumbnailRequest) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *SetStickerSetThumbnailRequest) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetFormat

`func (o *SetStickerSetThumbnailRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SetStickerSetThumbnailRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SetStickerSetThumbnailRequest) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


