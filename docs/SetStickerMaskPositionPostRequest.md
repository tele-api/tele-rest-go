# SetStickerMaskPositionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**MaskPosition** | Pointer to [**MaskPosition**](MaskPosition.md) |  | [optional] 

## Methods

### NewSetStickerMaskPositionPostRequest

`func NewSetStickerMaskPositionPostRequest(sticker string, ) *SetStickerMaskPositionPostRequest`

NewSetStickerMaskPositionPostRequest instantiates a new SetStickerMaskPositionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerMaskPositionPostRequestWithDefaults

`func NewSetStickerMaskPositionPostRequestWithDefaults() *SetStickerMaskPositionPostRequest`

NewSetStickerMaskPositionPostRequestWithDefaults instantiates a new SetStickerMaskPositionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerMaskPositionPostRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerMaskPositionPostRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerMaskPositionPostRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetMaskPosition

`func (o *SetStickerMaskPositionPostRequest) GetMaskPosition() MaskPosition`

GetMaskPosition returns the MaskPosition field if non-nil, zero value otherwise.

### GetMaskPositionOk

`func (o *SetStickerMaskPositionPostRequest) GetMaskPositionOk() (*MaskPosition, bool)`

GetMaskPositionOk returns a tuple with the MaskPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPosition

`func (o *SetStickerMaskPositionPostRequest) SetMaskPosition(v MaskPosition)`

SetMaskPosition sets MaskPosition field to given value.

### HasMaskPosition

`func (o *SetStickerMaskPositionPostRequest) HasMaskPosition() bool`

HasMaskPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


