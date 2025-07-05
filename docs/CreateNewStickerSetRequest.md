# CreateNewStickerSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier of created sticker set owner | 
**Name** | **string** | Short name of sticker set, to be used in &#x60;t.me/addstickers/&#x60; URLs (e.g., *animals*). Can contain only English letters, digits and underscores. Must begin with a letter, can&#39;t contain consecutive underscores and must end in &#x60;\&quot;_by_&lt;bot_username&gt;\&quot;&#x60;. &#x60;&lt;bot_username&gt;&#x60; is case insensitive. 1-64 characters. | 
**Title** | **string** | Sticker set title, 1-64 characters | 
**Stickers** | [**[]InputSticker**](InputSticker.md) | A JSON-serialized list of 1-50 initial stickers to be added to the sticker set | 
**StickerType** | Pointer to **string** | Type of stickers in the set, pass “regular”, “mask”, or “custom\\_emoji”. By default, a regular sticker set is created. | [optional] 
**NeedsRepainting** | Pointer to **bool** | Pass *True* if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only | [optional] 

## Methods

### NewCreateNewStickerSetRequest

`func NewCreateNewStickerSetRequest(userId int32, name string, title string, stickers []InputSticker, ) *CreateNewStickerSetRequest`

NewCreateNewStickerSetRequest instantiates a new CreateNewStickerSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewStickerSetRequestWithDefaults

`func NewCreateNewStickerSetRequestWithDefaults() *CreateNewStickerSetRequest`

NewCreateNewStickerSetRequestWithDefaults instantiates a new CreateNewStickerSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateNewStickerSetRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateNewStickerSetRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateNewStickerSetRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *CreateNewStickerSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNewStickerSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNewStickerSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateNewStickerSetRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateNewStickerSetRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateNewStickerSetRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStickers

`func (o *CreateNewStickerSetRequest) GetStickers() []InputSticker`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *CreateNewStickerSetRequest) GetStickersOk() (*[]InputSticker, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *CreateNewStickerSetRequest) SetStickers(v []InputSticker)`

SetStickers sets Stickers field to given value.


### GetStickerType

`func (o *CreateNewStickerSetRequest) GetStickerType() string`

GetStickerType returns the StickerType field if non-nil, zero value otherwise.

### GetStickerTypeOk

`func (o *CreateNewStickerSetRequest) GetStickerTypeOk() (*string, bool)`

GetStickerTypeOk returns a tuple with the StickerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerType

`func (o *CreateNewStickerSetRequest) SetStickerType(v string)`

SetStickerType sets StickerType field to given value.

### HasStickerType

`func (o *CreateNewStickerSetRequest) HasStickerType() bool`

HasStickerType returns a boolean if a field has been set.

### GetNeedsRepainting

`func (o *CreateNewStickerSetRequest) GetNeedsRepainting() bool`

GetNeedsRepainting returns the NeedsRepainting field if non-nil, zero value otherwise.

### GetNeedsRepaintingOk

`func (o *CreateNewStickerSetRequest) GetNeedsRepaintingOk() (*bool, bool)`

GetNeedsRepaintingOk returns a tuple with the NeedsRepainting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRepainting

`func (o *CreateNewStickerSetRequest) SetNeedsRepainting(v bool)`

SetNeedsRepainting sets NeedsRepainting field to given value.

### HasNeedsRepainting

`func (o *CreateNewStickerSetRequest) HasNeedsRepainting() bool`

HasNeedsRepainting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


