# UploadStickerFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | User identifier of sticker file owner | 
**Sticker** | **interface{}** |  | 
**StickerFormat** | **string** | Format of the sticker, must be one of “static”, “animated”, “video” | 

## Methods

### NewUploadStickerFileRequest

`func NewUploadStickerFileRequest(userId int32, sticker interface{}, stickerFormat string, ) *UploadStickerFileRequest`

NewUploadStickerFileRequest instantiates a new UploadStickerFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadStickerFileRequestWithDefaults

`func NewUploadStickerFileRequestWithDefaults() *UploadStickerFileRequest`

NewUploadStickerFileRequestWithDefaults instantiates a new UploadStickerFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UploadStickerFileRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UploadStickerFileRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UploadStickerFileRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetSticker

`func (o *UploadStickerFileRequest) GetSticker() interface{}`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *UploadStickerFileRequest) GetStickerOk() (*interface{}, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *UploadStickerFileRequest) SetSticker(v interface{})`

SetSticker sets Sticker field to given value.


### SetStickerNil

`func (o *UploadStickerFileRequest) SetStickerNil(b bool)`

 SetStickerNil sets the value for Sticker to be an explicit nil

### UnsetSticker
`func (o *UploadStickerFileRequest) UnsetSticker()`

UnsetSticker ensures that no value is present for Sticker, not even an explicit nil
### GetStickerFormat

`func (o *UploadStickerFileRequest) GetStickerFormat() string`

GetStickerFormat returns the StickerFormat field if non-nil, zero value otherwise.

### GetStickerFormatOk

`func (o *UploadStickerFileRequest) GetStickerFormatOk() (*string, bool)`

GetStickerFormatOk returns a tuple with the StickerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerFormat

`func (o *UploadStickerFileRequest) SetStickerFormat(v string)`

SetStickerFormat sets StickerFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


