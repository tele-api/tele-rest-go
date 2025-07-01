# SetStickerPositionInSetPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sticker** | **string** | File identifier of the sticker | 
**Position** | **int32** | New sticker position in the set, zero-based | 

## Methods

### NewSetStickerPositionInSetPostRequest

`func NewSetStickerPositionInSetPostRequest(sticker string, position int32, ) *SetStickerPositionInSetPostRequest`

NewSetStickerPositionInSetPostRequest instantiates a new SetStickerPositionInSetPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetStickerPositionInSetPostRequestWithDefaults

`func NewSetStickerPositionInSetPostRequestWithDefaults() *SetStickerPositionInSetPostRequest`

NewSetStickerPositionInSetPostRequestWithDefaults instantiates a new SetStickerPositionInSetPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSticker

`func (o *SetStickerPositionInSetPostRequest) GetSticker() string`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *SetStickerPositionInSetPostRequest) GetStickerOk() (*string, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *SetStickerPositionInSetPostRequest) SetSticker(v string)`

SetSticker sets Sticker field to given value.


### GetPosition

`func (o *SetStickerPositionInSetPostRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetStickerPositionInSetPostRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SetStickerPositionInSetPostRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


