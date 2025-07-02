# ReplaceStickerInSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier of the sticker set owner | 
**Name** | **string** | Sticker set name | 
**OldSticker** | **string** | File identifier of the replaced sticker | 
**Sticker** | [**InputSticker**](InputSticker.md) |  | 

## Methods

### NewReplaceStickerInSetRequest

`func NewReplaceStickerInSetRequest(userId int32, name string, oldSticker string, sticker InputSticker, ) *ReplaceStickerInSetRequest`

NewReplaceStickerInSetRequest instantiates a new ReplaceStickerInSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceStickerInSetRequestWithDefaults

`func NewReplaceStickerInSetRequestWithDefaults() *ReplaceStickerInSetRequest`

NewReplaceStickerInSetRequestWithDefaults instantiates a new ReplaceStickerInSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ReplaceStickerInSetRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReplaceStickerInSetRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReplaceStickerInSetRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *ReplaceStickerInSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplaceStickerInSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplaceStickerInSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOldSticker

`func (o *ReplaceStickerInSetRequest) GetOldSticker() string`

GetOldSticker returns the OldSticker field if non-nil, zero value otherwise.

### GetOldStickerOk

`func (o *ReplaceStickerInSetRequest) GetOldStickerOk() (*string, bool)`

GetOldStickerOk returns a tuple with the OldSticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSticker

`func (o *ReplaceStickerInSetRequest) SetOldSticker(v string)`

SetOldSticker sets OldSticker field to given value.


### GetSticker

`func (o *ReplaceStickerInSetRequest) GetSticker() InputSticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *ReplaceStickerInSetRequest) GetStickerOk() (*InputSticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *ReplaceStickerInSetRequest) SetSticker(v InputSticker)`

SetSticker sets Sticker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


