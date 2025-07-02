# AddStickerToSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier of sticker set owner | 
**Name** | **string** | Sticker set name | 
**Sticker** | [**InputSticker**](InputSticker.md) |  | 

## Methods

### NewAddStickerToSetRequest

`func NewAddStickerToSetRequest(userId int32, name string, sticker InputSticker, ) *AddStickerToSetRequest`

NewAddStickerToSetRequest instantiates a new AddStickerToSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStickerToSetRequestWithDefaults

`func NewAddStickerToSetRequestWithDefaults() *AddStickerToSetRequest`

NewAddStickerToSetRequestWithDefaults instantiates a new AddStickerToSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AddStickerToSetRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AddStickerToSetRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AddStickerToSetRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetName

`func (o *AddStickerToSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStickerToSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStickerToSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSticker

`func (o *AddStickerToSetRequest) GetSticker() InputSticker`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *AddStickerToSetRequest) GetStickerOk() (*InputSticker, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *AddStickerToSetRequest) SetSticker(v InputSticker)`

SetSticker sets Sticker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


