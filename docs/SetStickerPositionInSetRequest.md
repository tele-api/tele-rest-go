# SetStickerPositionInSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**Position** | **int32** | New sticker position in the set, zero-based | 

## Methods

### NewSetStickerPositionInSetRequest

`func NewSetStickerPositionInSetRequest(sticker string, position int32, ) *SetStickerPositionInSetRequest`

NewSetStickerPositionInSetRequest instantiates a new SetStickerPositionInSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerPositionInSetRequestWithDefaults

`func NewSetStickerPositionInSetRequestWithDefaults() *SetStickerPositionInSetRequest`

NewSetStickerPositionInSetRequestWithDefaults instantiates a new SetStickerPositionInSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerPositionInSetRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerPositionInSetRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerPositionInSetRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetPosition

`func (o *SetStickerPositionInSetRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetStickerPositionInSetRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SetStickerPositionInSetRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


